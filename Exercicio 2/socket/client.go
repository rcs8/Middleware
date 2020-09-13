package socket

import (
	"encoding/json"
	"math/rand"
	"net"
	"time"

	"../shared"
)

type Client struct {
	conn    *net.Conn
	encoder *json.Encoder
	decoder *json.Decoder
}

func NewClientTCP(port string) (*Client, error) {
	conn, err := net.Dial("tcp", port)
	if err != nil {
		return nil, err
	}

	return CreateEncoders(conn), err
}

func NewClientUDP(port string) (*Client, error) {
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}

	return CreateEncoders(conn), err
}

func CreateEncoders(conn net.Conn) *Client {
	jsonEncoder := json.NewEncoder(conn)
	jsonDecoder := json.NewDecoder(conn)

	return &Client{
		conn:    &conn,
		encoder: jsonEncoder,
		decoder: jsonDecoder,
	}
}

func (c *Client) MakeRequest() (string, error) {
	var response shared.Reply
	var err error

	message := PrepareArgs()

	err = c.encoder.Encode(message)

	if err != nil {
		return "", err
	}

	err = c.decoder.Decode(&response)

	if err != nil {
		return "", err
	}

	return response.Result, err
}

func PrepareArgs() shared.Args {
	rand.Seed(time.Now().UnixNano())
	return shared.Args{
		A: rand.Intn(10) + 1,
		B: rand.Intn(10) + 1,
		C: rand.Intn(10) + 1,
	}
}

func (c *Client) CloseConnection() {
	(*c.conn).Close()
}
