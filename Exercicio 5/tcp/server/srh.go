package main

import (
	"fmt"
	"net"
	"time"
)

type srhTCP struct {
	listener *net.Listener
}

func newSRHTCP() (*srhTCP, error) {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		return nil, err
	}
	return &srhTCP{
		listener: &ln,
	}, err
}

func (srh *srhTCP) Receive() {
	request := make([]byte, 64)
	fmt.Println("listening tcp server")
	listener := (*srh.listener).(*net.TCPListener)
	defer listener.Close()
	for {
		listener.SetDeadline(time.Now().Add(10 * time.Second))
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		conn.Read(request)

		go HandleTCP(*srh, conn, request)
	}
}

func (srh *srhTCP) SendTCP(msgToClient []byte, conn net.Conn) {
	var err error
	conn.Write(msgToClient)
	if err != nil {
		fmt.Println(err)
	}
}

func (srh *srhTCP) Close() {
	(*srh.listener).Close()
}
