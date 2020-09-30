package main

type Client interface {
	MakeRequest() ([]float64, error)
	MakeRequestBenchmark() ([]float64, int64, error)
	Close()
}
