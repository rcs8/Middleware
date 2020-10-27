package main

import (
	"fmt"
	"net"
)

type crhUDP struct {
	serverHost string
	serverPort int
	conn       *net.UDPConn
}

func newcrhUDP() (*crhUDP, error) {
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:8080")
	if err != nil {
		return nil, err
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}
	return &crhUDP{
		conn:       conn,
		serverHost: "0.0.0.0",
		serverPort: 8080,
	}, nil
}

func (crh crhUDP) SendReceive(msgToServer string) []byte {
	var conn net.UDPConn

	conn = *crh.conn

	conn.Write([]byte(msgToServer))
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)

	if err != nil {
		fmt.Println(err)
	}

	conn.Close()

	return buffer[0:n]
}
