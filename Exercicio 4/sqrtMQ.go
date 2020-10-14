package main

type SqrtRabbitMQ struct{}

func (s *SqrtRabbitMQ) Sqrt(req *Request) Reply {
	var reply Reply

	op := req.Op

	args := Args{
		A: req.A,
		B: req.B,
		C: req.C,
	}

	switch op {
	case "sqrt":
		reply = InvokeSqrt(args)
	}
	return reply
}
