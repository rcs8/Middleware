package main

import (
	"net"
	"net/rpc"
	"time"
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

func (s *ServerRPC) ListenRPC(exit NotifChan, exited NotifChan) {
	listener := (*s.listener).(*net.TCPListener)
	for {
		listener.SetDeadline(time.Now().Add(1 * time.Second))

		s.serverRPC.Accept(listener)

		_, stop := <-exit
		if stop {
			listener.Close()
			exited <- true
			return
		}
	}
}

func (s *ServerRPC) Close() {
	(*s.listener).Close()
}
