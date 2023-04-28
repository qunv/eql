package act

import (
	"github.com/qunv/eql/src/core/antlr"
)

type sum struct {
	ctx antlr.IActSpecContext
}

func newSum(ctx antlr.IActSpecContext) Act {
	return sum{
		ctx: ctx,
	}
}

func (s sum) Calculate(input [][]float64) float64 {
	params := s.ctx.AllParam()
	result := 0.0
	f := func(values []float64) float64 {
		r := 0.0
		for _, value := range values {
			r += value
		}
		return r
	}
	for _, p := range params {
		result += newParam(p, f).evaluate(input)
	}
	return result
}
