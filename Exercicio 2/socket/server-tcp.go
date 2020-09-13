package socket

import (
	"encoding/json"
	"fmt"
	"net"

	"../shared"
)

type ServerTCP struct {
	listener *net.Listener
}

func NewServerTCP(port string) (*ServerTCP, error) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}

	return &ServerTCP{
		listener: &ln,
	}, err
}

func (s *ServerTCP) ListenTCP() {
	for {
		conn, err := (*s.listener).Accept()
		if err != nil {
			fmt.Println(err)
		}

		go HandleTCP(conn)
	}
}

func HandleTCP(conn net.Conn) {
	var messageFromClient shared.Args

	defer conn.Close()

	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	for {
		err := jsonDecoder.Decode(&messageFromClient)
		if err != nil && err.Error() == "EOF" {
			conn.Close()
			break
		}

		msgToClient := shared.InvokeSqrt(messageFromClient)

		err = jsonEncoder.Encode(msgToClient)
		if err != nil {
			fmt.Println(err)
		}
	}
}
