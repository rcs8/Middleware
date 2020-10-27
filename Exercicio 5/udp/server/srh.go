package main

import (
	"fmt"
	"net"
	"time"
)

type srhUDP struct {
	conn *net.UDPConn
}

func newSRHUDP() (*srhUDP, error) {
	s, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	conn, err := net.ListenUDP("udp", s)
	if err != nil {
		return nil, err
	}
	return &srhUDP{
		conn: conn,
	}, err
}

func (srh *srhUDP) Receive() {
	var conn net.UDPConn = *srh.conn
	request := make([]byte, 64)
	fmt.Println("listening UDP server")
	defer conn.Close()
	for {
		conn.SetDeadline(time.Now().Add(10 * time.Second))
		n, addr, err := conn.ReadFromUDP(request)
		if err != nil {
			panic(err)
		}

		go HandleUDP(*srh, conn, request[0:n], addr)
	}
}

func (srh *srhUDP) SendUDP(msgToClient []byte, conn net.UDPConn, addr *net.UDPAddr) {
	var err error
	conn.WriteToUDP(msgToClient, addr)
	if err != nil {
		fmt.Println(err)
	}
}

func (srh *srhUDP) Close() {
	(*srh.conn).Close()
}
