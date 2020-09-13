package socket

import (
	"encoding/json"
	"fmt"
	"net"

	"../shared"
)

type ServerUDP struct {
	conn *net.UDPConn
}

func NewServerUDP(port string) (*ServerUDP, error) {
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		return nil, err
	}

	// Listen on udp port
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	return &ServerUDP{
		conn: conn,
	}, err
}

func (s *ServerUDP) ListenUDP() {
	msgFromClient := make([]byte, 1024)
	for {
		// Receive request
		n, addr, err := (*s.conn).ReadFromUDP(msgFromClient)
		if err != nil {
			fmt.Println(err)
		}

		// Handle request
		go HandleUDP(s.conn, msgFromClient, n, addr)
	}
}

func HandleUDP(conn *net.UDPConn, msgFromClient []byte, n int, addr *net.UDPAddr) {
	var msgToClient []byte
	var args shared.Args

	err := json.Unmarshal(msgFromClient[:n], &args)
	if err != nil {
		fmt.Println(err)
	}

	result := shared.InvokeSqrt(args)

	msgToClient, err = json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	_, err = conn.WriteTo(msgToClient, addr)
	if err != nil {
		fmt.Println(err)
	}
}
