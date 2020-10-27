package main

import (
	"fmt"
	"net"
)

type crhTCP struct {
	serverHost string
	serverPort int
	conn       net.Conn
}

func newCRHTCP() (*crhTCP, error) {
	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	if err != nil {
		return nil, err
	}
	return &crhTCP{
		conn:       conn,
		serverHost: "0.0.0.0",
		serverPort: 6666,
	}, nil
}

func (crh crhTCP) SendReceive(msgToServer string) []byte {
	var conn net.Conn

	conn = crh.conn

	conn.Write([]byte(msgToServer))
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)

	if err != nil {
		fmt.Println(err)
	}

	conn.Close()

	return buffer
}

// func NewClientUDP(address string) (*crhTCP, error) {
// 	addr, err := net.ResolveUDPAddr("udp", address)
// 	if err != nil {
// 		return nil, err
// 	}

// 	conn, err := net.DialUDP("udp", nil, addr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return CreateEncoders(conn, "udp"), err
// }

// func CreateEncoders(conn net.Conn, protocol string) *crhTCP {
// 	jsonEncoder := json.NewEncoder(conn)
// 	jsonDecoder := json.NewDecoder(conn)

// 	return &crhTCP{
// 		conn:       &conn,
// 		encoder:    jsonEncoder,
// 		decoder:    jsonDecoder,
// 		serverHost: "0.0.0.0",
// 		serverPort: 6666,
// 		protocol:   protocol,
// 	}
// }
