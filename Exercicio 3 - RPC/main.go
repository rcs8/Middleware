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
	if protocol == "RPC" {
		server, err := NewServerRPC(address)
		if err != nil {
			panic(err)
		}
		defer server.Close()
		server.ListenRPC(exit, exited)
	}
}

func initClient(protocol string) *Client {
	client, err := NewClientRPC(address)
	if err != nil {
		panic(err)
	}
	return client
}

const iterations = 10000

func simpleClient(protocol string) {
	client := initClient(protocol)
	defer client.Close()

	for i := 0; i < iterations; i++ {
		_, _ = client.MakeRequest()
	}
}

func benchmarkClient(protocol string, result chan BenchResult) {
	client := initClient(protocol)
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

func benchmarkProtocolClients(protocol string, clients int) BenchResult {
	result := make(chan BenchResult)

	go benchmarkClient(protocol, result)
	for i := 1; i < clients; i++ {
		go simpleClient(protocol)
	}

	return <-result
}

var ClientAmounts = []int{1, 2, 5, 10}

func suite() (map[string][]BenchResult, float64, float64, float64) {
	results := make(map[string][]BenchResult)
	var maxMean float64 = 0
	var minMeanSD float64 = 0
	var maxMeanSD float64 = 0

	exit := make(NotifChan)
	exited := make(NotifChan)

	for _, protocol := range []string{"RPC"} {
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
