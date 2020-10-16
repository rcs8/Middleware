package main

import (
	"net/rpc"
	"time"
)

type ClientRPC struct {
	client *rpc.Client
}

func (c *ClientRPC) MakeRequest() ([]float64, error) {
	var response Reply
	var err error

	message := c.prepareArgs()

	err = c.client.Call("Sqrt.Sqrt", message, &response)

	if err != nil {
		return nil, err
	}

	return response.Result, err
}

func (c *ClientRPC) MakeRequestBenchmark() ([]float64, int64, error) {
	var response Reply
	var err error

	message := c.prepareArgs()

	startTime := time.Now()

	err = c.client.Call("Sqrt.Sqrt", message, &response)

	totalTime := time.Now().Sub(startTime).Microseconds()

	if err != nil {
		return nil, 0, err
	}

	return response.Result, totalTime, err
}

func (c *ClientRPC) Close() {
	(*c.client).Close()
}

func NewClientRPC(address string) (*ClientRPC, error) {
	client, err := rpc.Dial("tcp", address)

	if err != nil {
		return nil, err
	}

	return &ClientRPC{
		client: client,
	}, err
}

func (c *ClientRPC) prepareArgs() Args {
	return Args{
		A: 4,
		B: 3,
		C: -5,
	}
}
