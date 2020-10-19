package main

import (
	"os"
)

const address = "0.0.0.0:6666"

func initServer(protocol string) {
	switch protocol {
	case "TCP":
		server, err := NewServerTCP(address)
		if err != nil {
			panic(err)
		}
		defer server.Close()
		server.ListenTCP()
	case "UDP":
		server, err := NewServerUDP(address)
		if err != nil {
			panic(err)
		}
		defer server.Close()
		server.ListenUDP()
	case "RPC":
		server, err := NewServerRPC(address)
		if err != nil {
			panic(err)
		}
		defer server.Close()
		server.ListenRPC()
	case "RabbitMQ":
		server, err := NewServerRabbitMQ(address)
		if err != nil {
			panic(err)
		}
		defer server.Close()
		server.ListenRabbitMQ()
	}
}

func main() {
	initServer(os.Args[1])
}
