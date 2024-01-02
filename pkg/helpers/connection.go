package helpers

import (
	"net"
)

func ReadConn(conn net.Conn) (request []byte, err error) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)

	if err != nil {
		return request, err
	}

	request = buffer[:n]
	return request, err

}

func WriteConn(conn net.Conn, msg string) (err error) {
	_, err = conn.Write([]byte(msg + "\n"))
	return err
}
