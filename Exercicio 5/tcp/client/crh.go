package main

import (
	"fmt"
	"net"
)

type crhTCP struct {
	serverHost string
	serverPort int
	conn       net.Conn
}

func newCRHTCP() (*crhTCP, error) {
	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	if err != nil {
		return nil, err
	}
	return &crhTCP{
		conn:       conn,
		serverHost: "0.0.0.0",
		serverPort: 8080,
	}, nil
}

func (crh crhTCP) SendReceive(msgToServer string) []byte {
	var conn net.Conn

	conn = crh.conn

	conn.Write([]byte(msgToServer))
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)

	if err != nil {
		fmt.Println(err)
	}

	conn.Close()

	return buffer
}
