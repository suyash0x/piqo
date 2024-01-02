package server

import (
	"encoding/json"
	"io"
	"log"
	"net"

	"github.com/suyash0x/piqo/pkg/broker"
	"github.com/suyash0x/piqo/pkg/helpers"
)

type brokerModule interface {
	AddTopic(net.Conn)
	GetMessage(net.Conn)
	PostMessage(net.Conn)
}

type Server struct {
	brokerModule
}

type Request struct {
	Action  string `json:"action"`
	Topic   string `json:"topic"`
	Message any    `json:"message"`
}

func (s *Server) handleConn(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			helpers.LogError("closing connection", err)
			return
		}
	}()

	for {
		var request Request
		data, err := helpers.ReadConn(conn)

		if err != nil {
			if err == io.EOF {
				break
			}
			helpers.LogError("Reading connection", err)
			continue
		}

		err = json.Unmarshal(data, &request)

		if err != nil {
			helpers.WriteConn(conn, "Error converting request JSON")
			return
		}

		switch request.Action {
		case "produce":
			s.PostMessage(conn)
		case "consume":
			s.GetMessage(conn)
		case "add-topic":
			s.AddTopic(conn)
		default:
			helpers.WriteConn(conn, "Invalid action")
			continue
		}

		helpers.WriteConn(conn, "OK")
	}

}

func (s *Server) StartServer() {
	listener, err := net.Listen("tcp", ":8080")
	helpers.FatalOutError("Starting piqo server", err)

	log.Println("piqo sever listing on", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		helpers.FatalOutError("Accepting connection", err)
		log.Println(conn.RemoteAddr().String(), "connected to server")

		go s.handleConn(conn)
	}
}

func NewServer() *Server {
	return &Server{
		brokerModule: broker.NewBroker(),
	}
}
