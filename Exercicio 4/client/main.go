package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"time"
)

type NotifChan chan interface{}

type BenchResult struct {
	mean float64
	sd   float64
}

func initClient(protocol string) Client {
	var client Client
	var err error
	switch protocol {
	case "TCP":
		client, err = NewClientTCP("server-tcp:6666")
		if err != nil {
			panic(err)
		}
	case "UDP":
		client, err = NewClientUDP("server-udp:6666")
		if err != nil {
			panic(err)
		}
	case "RPC":
		client, err = NewClientRPC("server-rpc:6666")
		if err != nil {
			panic(err)
		}
	case "RabbitMQ":
		client, err = NewClientRabbitMQ("server-rabbitmq:6666")
		if err != nil {
			panic(err)
		}
	default:
		panic("panico")
	}
	return client
}

func simpleClient(protocol string, client Client, wg *sync.WaitGroup) {
	defer wg.Done()
	defer client.Close()

	for i := 0; i < iterations; i++ {
		_, _ = client.MakeRequest()
	}
}

func benchmarkClient(protocol string, result *BenchResult, client Client, wg *sync.WaitGroup) {
	defer wg.Done()
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

	result.mean = mean
	result.sd = sd
}

func benchmarkProtocolClients(protocol string, nClients int) BenchResult {
	result := BenchResult{}
	clients := make([]Client, nClients)

	wg := sync.WaitGroup{}

	for i := 0; i < nClients; i++ {
		clients[i] = initClient(protocol)
	}

	time.Sleep(1 * time.Second)

	wg.Add(1)
	go benchmarkClient(protocol, &result, clients[0], &wg)
	for i := 1; i < nClients; i++ {
		wg.Add(1)
		go simpleClient(protocol, clients[i], &wg)
	}

	wg.Wait()

	return result
}

const iterations = 10000

var ClientAmounts = []int{1, 2, 5, 10}

func suite(protocol string) (map[string][]BenchResult, float64, float64, float64) {
	results := make(map[string][]BenchResult)
	var maxMean float64 = 0
	var minMeanSD float64 = 0
	var maxMeanSD float64 = 0

	results[protocol] = make([]BenchResult, 0)
	for _, clientAmount := range ClientAmounts {
		result := benchmarkProtocolClients(protocol, clientAmount)
		results[protocol] = append(results[protocol], result)
		maxMean = math.Max(maxMean, result.mean)
		minMeanSD = math.Min(minMeanSD, result.mean-result.sd)
		maxMeanSD = math.Max(maxMeanSD, result.mean+result.sd)
		fmt.Printf("%s with %d clients avg: %f\n", protocol, clientAmount, result.mean)
	}

	return results, maxMean, minMeanSD, maxMeanSD
}

func main() {
	time.Sleep(1 * time.Second)
	protocolResults, maxMean, minMeanSD, maxMeanSD := suite(os.Args[1])

	Plot(protocolResults, maxMean, minMeanSD, maxMeanSD)
}
