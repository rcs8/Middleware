package main

import (
	"net"
	"net/rpc"
)

type ServerRPC struct {
	listener  *net.Listener
	serverRPC *rpc.Server
}

func NewServerRPC(address string) (*ServerRPC, error) {
	sqrt := new(SqrtRPC)

	server := rpc.NewServer()
	server.RegisterName("Sqrt", sqrt)

	ln, err := net.Listen("tcp", address)

	return &ServerRPC{
		listener:  &ln,
		serverRPC: server,
	}, err
}

func (s *ServerRPC) ListenRPC() {
	for {
		s.serverRPC.Accept(*s.listener)
	}
}

func (s *ServerRPC) Close() {
	(*s.listener).Close()
}
