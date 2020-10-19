package main

import "math"

func InvokeSqrt(args Args) Reply {
	var a = float64(args.A)
	var b = float64(args.B)
	var c = float64(args.C)

	deltaValue := CalculateDelta(a, b, c)

	if deltaValue < 0 {
		return Reply{
			Result: []float64{},
		}
	}

	if deltaValue == 0 {
		return Reply{
			Result: []float64{(b * (-1)) / (2 * a)},
		}
	}

	return Reply{
		Result: []float64{(math.Sqrt(deltaValue) - b) / 2 * a, ((-1)*math.Sqrt(deltaValue) - b) / 2 * a},
	}
}

func CalculateDelta(a, b, c float64) float64 {
	return (b * b) - (4 * a * c)
}
