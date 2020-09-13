package shared

const Iterations = 1
const Port = 8080

type Args struct {
	A int
	B int
	C int
}

type Reply struct {
	Result []float64
}
