package action

import (
	"errors"
	"github.com/qunv/eql/core/antlr"
)

type _divide struct {
	ctx antlr.IActionSpecContext
}

func divide(ctx antlr.IActionSpecContext) Action {
	return _divide{ctx}
}

func (d _divide) Evaluate(input EqlInput) (EqlValue, error) {
	if len(d.ctx.AllParam()) != 2 {
		return nil, errors.New("len params just accept 2")
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

	if f2 == 0 {
		return nil, errors.New("2nd param must be not zero")
	}

	return NewEqlValue(f1 / f2), nil
}
