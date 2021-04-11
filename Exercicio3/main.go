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

	var maxN int64 = math.MinInt64
	var minN int64 = math.MaxInt64

	var sum int64 = 0
	iterationTime := make([]int64, iterations)
	for i := 0; i < iterations; i++ {
		_, time, _ := client.MakeRequestBenchmark()
		maxN = int64(math.Max(float64(maxN), float64(time)))
		minN = int64(math.Min(float64(minN), float64(time)))
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

func benchmarkProtocolClients(protocol string,
	clients int) BenchResult {
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

	for _, protocol := range []string{"TCP", "UDP"} {
		results[protocol] = make([]BenchResult, 0)
		for _, clientAmount := range ClientAmounts {
			go initServer(protocol, exit, exited)
			time.Sleep(100 * time.Millisecond)

			result := benchmarkProtocolClients(protocol, clientAmount)
			results[protocol] = append(results[protocol], result)
			maxMean = math.Max(maxMean, result.mean)
			minMeanSD = math.Min(minMeanSD, result.mean-result.sd)
			maxMeanSD = math.Max(maxMeanSD, result.mean+result.sd)
			fmt.Printf("%s with %d clients avg: %f\n",
			protocol, clientAmount, result.mean)

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