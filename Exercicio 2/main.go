package main

import (
	"fmt"
	"math"
	"time"
)

type NotifChan chan interface{}

const address = "0.0.0.0:6666"
const iterations = 10000

var ClientAmounts = []int{1, 2, 5, 10}

func initServer(protocol string, exit NotifChan, exited NotifChan) {
	if protocol == "TCP" {
		server, err := NewServerTCP(address)
		if err != nil {
			panic(err)
		}
		defer server.Close()
		server.ListenTCP(exit, exited)
	} else {
		server, err := NewServerUDP(address)
		if err != nil {
			panic(err)
		}
		defer server.Close()
		server.ListenUDP(exit, exited)
	}
}

func initClient(protocol string) *Client {
	if protocol == "TCP" {
		client, err := NewClientTCP(address)
		if err != nil {
			panic(err)
		}
		return client
	} else {
		client, err := NewClientUDP(address)
		if err != nil {
			panic(err)
		}
		return client
	}
}

func simpleClient(protocol string) {
	client := initClient(protocol)
	defer client.Close()

	for i := 0; i < iterations; i++ {
		_, _ = client.MakeRequest()
	}
}

func benchmarkClient(protocol string, result chan float64) {
	client := initClient(protocol)
	defer client.Close()

	var totalTime int64 = 0
	for i := 0; i < iterations; i++ {
		_, time, _ := client.MakeRequestBenchmark()
		totalTime += time
	}
	result <- float64(totalTime) / float64(iterations)
}

func benchmarkProtocolClients(protocol string, clients int) float64 {
	result := make(chan float64)

	go benchmarkClient(protocol, result)
	for i := 1; i < clients; i++ {
		go simpleClient(protocol)
	}

	return <-result
}

func suite() (map[string][]float64, float64) {
	results := make(map[string][]float64)
	var maxAvg float64 = 0

	exit := make(NotifChan)
	exited := make(NotifChan)

	for _, protocol := range []string{"TCP", "UDP"} {
		results[protocol] = make([]float64, 0)
		for _, clientAmount := range ClientAmounts {
			go initServer(protocol, exit, exited)
			time.Sleep(100 * time.Millisecond)

			avg := benchmarkProtocolClients(protocol, clientAmount)
			results[protocol] = append(results[protocol], avg)
			maxAvg = math.Max(maxAvg, avg)
			fmt.Printf("%s with %d clients avg: %f\n", protocol, clientAmount, avg)

			exit <- true
			<-exited
		}
	}

	return results, maxAvg
}

func main() {
	protocolResults, maxAvg := suite()

	Plot(protocolResults, maxAvg)
}
