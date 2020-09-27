package main

import (
	"net/rpc"
	"time"
)

type Client struct {
	client *rpc.Client
}

func NewClientRPC(address string) (*Client, error) {
	client, err := rpc.Dial("tcp", address)

	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, err
}

func (c *Client) MakeRequest() ([]float64, error) {
	var response Reply
	var err error

	message := PrepareArgs()

	err = c.client.Call("Sqrt.Sqrt", message, &response)

	if err != nil {
		return nil, err
	}

	return response.Result, err
}

func (c *Client) MakeRequestBenchmark() ([]float64, int64, error) {
	var response Reply
	var err error

	message := PrepareArgs()

	startTime := time.Now()

	err = c.client.Call("Sqrt.Sqrt", message, &response)

	totalTime := time.Now().Sub(startTime).Microseconds()

	if err != nil {
		return nil, 0, err
	}

	return response.Result, totalTime, err
}

func PrepareArgs() Args {
	return Args{
		A: 4,
		B: 3,
		C: -5,
	}
}

func (c *Client) Close() {
	(*c.client).Close()
}
