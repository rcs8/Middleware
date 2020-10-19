package main

import (
	"encoding/json"
	"math/rand"
	"net"
	"time"
)

type ClientSocket struct {
	conn    *net.Conn
	encoder *json.Encoder
	decoder *json.Decoder
}

func NewClientTCP(address string) (*ClientSocket, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return CreateEncoders(conn), err
}

func NewClientUDP(address string) (*ClientSocket, error) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}

	return CreateEncoders(conn), err
}

func CreateEncoders(conn net.Conn) *ClientSocket {
	jsonEncoder := json.NewEncoder(conn)
	jsonDecoder := json.NewDecoder(conn)

	return &ClientSocket{
		conn:    &conn,
		encoder: jsonEncoder,
		decoder: jsonDecoder,
	}
}

func (c *ClientSocket) MakeRequest() ([]float64, error) {
	var response Reply
	var err error

	message := c.prepareArgs()

	err = c.encoder.Encode(message)

	if err != nil {
		return nil, err
	}

	err = c.decoder.Decode(&response)

	if err != nil {
		return nil, err
	}

	return response.Result, err
}

func (c *ClientSocket) MakeRequestBenchmark() ([]float64, int64, error) {
	var response Reply
	var err error

	message := c.prepareArgs()

	startTime := time.Now()
	err = c.encoder.Encode(message)

	if err != nil {
		return nil, 0, err
	}

	err = c.decoder.Decode(&response)
	totalTime := time.Now().Sub(startTime).Microseconds()

	if err != nil {
		return nil, 0, err
	}

	return response.Result, totalTime, err
}

func (c *ClientSocket) prepareArgs() Args {
	rand.Seed(time.Now().UnixNano())
	return Args{
		A: rand.Intn(10) + 1,
		B: rand.Intn(10) + 1,
		C: rand.Intn(10) + 1,
	}
}

func (c *ClientSocket) Close() {
	(*c.conn).Close()
}
