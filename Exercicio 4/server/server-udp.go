package main

import (
	"encoding/json"
	"fmt"
	"net"
	"unsafe"
)

type ServerUDP struct {
	conn *net.UDPConn
}

func NewServerUDP(port string) (*ServerUDP, error) {
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	return &ServerUDP{
		conn: conn,
	}, err
}

func (s *ServerUDP) ListenUDP() {
	var args Args
	for {
		msgFromClient := make([]byte, unsafe.Sizeof(args))
		n, addr, err := (*s.conn).ReadFromUDP(msgFromClient)
		if err != nil {
			panic(err)
		}

		go HandleUDP(s.conn, msgFromClient, n, addr)
	}
}

func (s *ServerUDP) Close() {
	(*s.conn).Close()
}

func HandleUDP(conn *net.UDPConn, msgFromClient []byte, n int, addr *net.UDPAddr) {
	var msgToClient []byte
	var args Args

	err := json.Unmarshal(msgFromClient[:n], &args)
	if err != nil {
		fmt.Println(string(msgFromClient[:n]), err)
	}

	result := InvokeSqrt(args)

	msgToClient, err = json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	_, err = conn.WriteTo(msgToClient, addr)
	if err != nil {
		fmt.Println(err)
	}
}
