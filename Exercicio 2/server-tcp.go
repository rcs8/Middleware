package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type ServerTCP struct {
	listener *net.Listener
}

func NewServerTCP(address string) (*ServerTCP, error) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return &ServerTCP{
		listener: &ln,
	}, err
}

func (s *ServerTCP) ListenTCP(exit NotifChan, exited NotifChan) {
	listener := (*s.listener).(*net.TCPListener)
	for {
		listener.SetDeadline(time.Now().Add(1 * time.Second))
		conn, err := listener.Accept()
		if err != nil {
			_, stop := <-exit
			if stop {
				listener.Close()
				exited <- true
				return
			}
			continue
		}

		go HandleTCP(conn)
	}
}

func (s *ServerTCP) Close() {
	(*s.listener).Close()
}

func HandleTCP(conn net.Conn) {
	var messageFromClient Args

	defer conn.Close()

	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	for {
		err := jsonDecoder.Decode(&messageFromClient)
		if err != nil && err.Error() == "EOF" {
			conn.Close()
			break
		}

		msgToClient := InvokeSqrt(messageFromClient)

		err = jsonEncoder.Encode(msgToClient)
		if err != nil {
			fmt.Println(err)
		}
	}
}
