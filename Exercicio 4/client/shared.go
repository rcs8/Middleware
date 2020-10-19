package main

type Args struct {
	A int
	B int
	C int
}

type Reply struct {
	Result []float64
}

type Request struct {
	Op string
	A  int
	B  int
	C  int
}
