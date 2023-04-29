package act

import (
	"github.com/qunv/eql/src/core/antlr"
	"math"
)

type abs struct {
	ctx antlr.IActSpecContext
}

func newAbs(ctx antlr.IActSpecContext) Act {
	return abs{ctx}
}

func (a abs) Calculate(input [][]float64) float64 {
	params := a.ctx.AllParam()
	if len(params) != 1 {
		panic("error: len param must be 1")
	}

	p := a.ctx.Param(0)
	result := newParam(p, nil).evaluate(input)
	return math.Abs(result)
}
