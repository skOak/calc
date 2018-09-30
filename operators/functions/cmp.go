package functions

import (
	"math"
)

import (
	"github.com/skOak/calc/operators"
)

var (
	max = &operators.Operator{
		Name:          ">",
		Precedence:    0,
		Associativity: operators.R,
		Args:          2,
		Operation: func(args []float64) float64 {
			return math.Max(args[0], args[1])
		},
	}
	min = &operators.Operator{
		Name:          "<",
		Precedence:    0,
		Associativity: operators.R,
		Args:          2,
		Operation: func(args []float64) float64 {
			return math.Min(args[0], args[1])
		},
	}
)

func init() {
	Register(max)
	Register(min)
}
