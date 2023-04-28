package act

import (
	"github.com/qunv/eql/src/core/antlr"
)

type avg struct {
	ctx antlr.IActSpecContext
}

func newAvg(ctx antlr.IActSpecContext) Act {
	return avg{
		ctx: ctx,
	}
}

func (a avg) Calculate(input [][]float64) float64 {
	params := a.ctx.AllParam()
	result := 0.0
	f := func(values []float64) float64 {
		r := 0.0
		for _, value := range values {
			r += value / float64(len(values))
		}
		return r
	}
	for _, p := range params {
		result += newParam(p, f).evaluate(input) / float64(len(params))
	}
	return result
}
