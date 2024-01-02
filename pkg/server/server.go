package server

import (
	"io"
	"log"
	"net"

	"github.com/suyash0x/piqo/pkg/helpers"
)

type Server struct{}

func (s *Server) handleConn(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			helpers.LogError("closing connection", err)
			return
		}
		log.Println("connection closed successfully", conn.RemoteAddr().String())

	}()

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)

		helpers.LogError("Reading data from connection", err)

		if err == io.EOF {
			break
		}

		request := buffer[:n]

		_, err = conn.Write(request)

		helpers.LogError("Writing data to connection", err)
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
	return &Server{}
}
