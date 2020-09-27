package main

import (
	"net"
	"net/rpc"
	"time"
)

type ServerTCP struct {
	listener  *net.Listener
	serverRPC *rpc.Server
}

func NewServerTCP(address string) (*ServerTCP, error) {
	sqrt := new(SqrtRPC)

	server := rpc.NewServer()
	server.RegisterName("Sqrt", sqrt)

	ln, err := net.Listen("tcp", address)

	return &ServerTCP{
		listener:  &ln,
		serverRPC: server,
	}, err
}

func (s *ServerTCP) ListenTCP(exit NotifChan, exited NotifChan) {
	listener := (*s.listener).(*net.TCPListener)
	for {
		listener.SetDeadline(time.Now().Add(1 * time.Second))

		s.serverRPC.Accept(listener)
	}
}

func (s *ServerTCP) Close() {
	(*s.listener).Close()
}
