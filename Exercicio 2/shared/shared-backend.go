import (
	"fmt"
	"math"
)

func InvokeSqrt(args shared.Args) shared.Reply {
	var a = float64(args.A)
	var b = float64(args.B)
	var c = float64(args.C)

	deltaValue := CalculateDelta(a, b, c)

	if deltaValue < 0 {
		return shared.Reply{
			Result: "Nenhuma raiz real\n",
		}
	}

	if deltaValue == 0 {
		return shared.Reply{
			Result: fmt.Sprintf("%f\n", (b*(-1))/(2*a)),
		}
	}

	return shared.Reply{
		Result: fmt.Sprintf("%f e %f\n", (math.Sqrt(deltaValue)-b)/2*a, ((-1)*math.Sqrt(deltaValue)-b)/2*a),
	}
}

func CalculateDelta(a, b, c float64) float64 {
	return (b * b) - (4 * a * c)
}
