package main

import (
	"math"
	"log"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) InvokeSqrt(ctx context.Context, args *Args) (*Reply, error) {
	log.Printf("Receive new message")
	result := []float64{}
	var a = float64(args.A)
	var b = float64(args.B)
	var c = float64(args.C)

	deltaValue := CalculateDelta(a, b, c)

	if deltaValue < 0 {
		return &Reply{
			Result: result,
		}, nil
	}

	if deltaValue == 0 {
		return &Reply{
			Result: append(result, (b*(-1))/(2*a)),
		}, nil
	}

	return &Reply{
		Result: append(result, (math.Sqrt(deltaValue)-b)/2*a, ((-1)*math.Sqrt(deltaValue)-b)/2*a),
	}, nil
}

func CalculateDelta(a, b, c float64) float64 {
	return (b * b) - (4 * a * c)
}
