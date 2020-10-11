package main

type SqrtRabbitMQ struct{}

func (s *SqrtRabbitMQ) Sqrt(req *Request) Reply {
	var reply Reply

	op := req.Op
	args := req.args

	switch op {
	case "sqrt":
		reply = InvokeSqrt(args)
	}
	return reply
}
