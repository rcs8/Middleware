package main

type SqrtRPC struct{}

func (s *SqrtRPC) Sqrt(args *Args, reply *Reply) error {
	*reply = InvokeSqrt(*args)
	return nil
}
