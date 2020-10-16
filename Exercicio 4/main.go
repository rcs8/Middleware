package main

import (
	"fmt"
	"math"
	"time"
)

type NotifChan chan interface{}

type BenchResult struct {
	mean float64
	sd   float64
}

const address = "0.0.0.0:6666"

func initServer(protocol string, exit NotifChan, exited NotifChan) {
	switch protocol {
	case "TCP":
		server, err := NewServerTCP(address)
		if err != nil {
			panic(err)
		}
		defer server.Close()
		server.ListenTCP(exit, exited)
	case "UDP":
		server, err := NewServerUDP(address)
		if err != nil {
			panic(err)
		}
		defer server.Close()
		server.ListenUDP(exit, exited)
	case "RPC":
		server, err := NewServerRPC(address)
		if err != nil {
			panic(err)
		}
		defer server.Close()
		server.ListenRPC(exit, exited)
	case "RabbitMQ":
		server, err := NewServerRabbitMQ(address)
		if err != nil {
			panic(err)
		}
		defer server.Close()
		server.ListenRabbitMQ(exit, exited)
	}
}

func initClient(protocol string) Client {
	var client Client
	var err error
	switch protocol {
	case "TCP":
		client, err = NewClientTCP(address)
		if err != nil {
			panic(err)
		}
	case "UDP":
		client, err = NewClientUDP(address)
		if err != nil {
			panic(err)
		}
	case "RPC":
		client, err = NewClientRPC(address)
		if err != nil {
			panic(err)
		}
	case "RabbitMQ":
		client, err = NewClientRabbitMQ(address)
		if err != nil {
			panic(err)
		}
	default:
		panic("panico")
	}
	return client
}

const iterations = 10

func simpleClient(protocol string, client Client) {
	defer client.Close()

	for i := 0; i < iterations; i++ {
		_, _ = client.MakeRequest()
	}
}

func benchmarkClient(protocol string, result chan BenchResult, client Client) {
	defer client.Close()

	var sum int64 = 0
	iterationTime := make([]int64, iterations)
	for i := 0; i < iterations; i++ {
		_, time, _ := client.MakeRequestBenchmark()
		sum += time
		iterationTime[i] = time
	}

	var variation float64 = 0
	mean := float64(sum) / float64(iterations)
	for _, time := range iterationTime {
		diff := float64(time) - mean
		variation += diff * diff
	}
	variation /= float64(iterations)
	sd := math.Sqrt(variation)

	result <- BenchResult{mean, sd}
}

func benchmarkProtocolClients(protocol string, nClients int) BenchResult {
	result := make(chan BenchResult)
	clients := make([]Client, nClients)

	for i := 0; i < nClients; i++ {
		clients[i] = initClient(protocol)
	}

	go benchmarkClient(protocol, result, clients[0])
	for i := 1; i < nClients; i++ {
		go simpleClient(protocol, clients[i])
	}

	return <-result
}

var ClientAmounts = []int{1}

func suite() (map[string][]BenchResult, float64, float64, float64) {
	results := make(map[string][]BenchResult)
	var maxMean float64 = 0
	var minMeanSD float64 = 0
	var maxMeanSD float64 = 0

	exit := make(NotifChan)
	exited := make(NotifChan)

	for _, protocol := range []string{"RabbitMQ"} {
		results[protocol] = make([]BenchResult, 0)
		for _, clientAmount := range ClientAmounts {
			go initServer(protocol, exit, exited)
			time.Sleep(100 * time.Millisecond)

			result := benchmarkProtocolClients(protocol, clientAmount)
			results[protocol] = append(results[protocol], result)
			maxMean = math.Max(maxMean, result.mean)
			minMeanSD = math.Min(minMeanSD, result.mean-result.sd)
			maxMeanSD = math.Max(maxMeanSD, result.mean+result.sd)
			fmt.Printf("%s with %d clients avg: %f\n", protocol, clientAmount, result.mean)

			exit <- true
			<-exited
		}
	}

	return results, maxMean, minMeanSD, maxMeanSD
}

func main() {
	protocolResults, maxMean, minMeanSD, maxMeanSD := suite()

	Plot(protocolResults, maxMean, minMeanSD, maxMeanSD)
}
