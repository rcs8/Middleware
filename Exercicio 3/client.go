package main

import (
	"net/rpc"
	"time"
)

type Client struct {
	client *rpc.Client
}

func NewClient(address string) (*Client, error) {
	client, err := rpc.Dial("tcp", address)

	return &Client{
		client: client,
	}, err
}

func (c *Client) MakeRequest() []float64 {
	var response Reply

	message := PrepareArgs()

	c.client.Call("Sqrt.Sqrt", message, &response)

	return response.Result
}

func (c *Client) MakeRequestBenchmark() ([]float64, int64) {
	var response Reply

	message := PrepareArgs()

	startTime := time.Now()

	c.client.Call("Sqrt.Sqrt", message, &response)

	totalTime := time.Now().Sub(startTime).Microseconds()

	return response.Result, totalTime
}

func PrepareArgs() Args {
	return Args{
		A: 5,
		B: 7,
		C: 8,
	}
}

func (c *Client) Close() {
	(*c.client).Close()
}
