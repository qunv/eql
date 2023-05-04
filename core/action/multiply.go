package action

import (
	"github.com/qunv/eql/core/antlr"
)

type _multiply struct {
	ctx antlr.IActionSpecContext
}

func multiply(ctx antlr.IActionSpecContext) Action {
	return _multiply{ctx}
}

func (d _multiply) Calculate(input EqlInput) (EqlValue, error) {
	if len(d.ctx.AllParam()) != 2 {
		panic("Len params just accept 2")
	}

	val1, err := param(d.ctx.Param(0), nil).evaluate(input)
	if err != nil {
		return nil, err
	}
	f1, err := val1.Float64()
	if err != nil {
		return nil, err
	}
	val2, err := param(d.ctx.Param(1), nil).evaluate(input)
	if err != nil {
		return nil, err
	}
	f2, err := val2.Float64()
	if err != nil {
		return nil, err
	}
	return NewEqlValue(f1 * f2), nil
}
