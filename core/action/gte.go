package action

import (
	"errors"
	"github.com/qunv/eql/core/antlr"
)

type _gte struct {
	ctx antlr.IActionSpecContext
}

func gte(ctx antlr.IActionSpecContext) Action {
	return _gte{ctx}
}

func (g _gte) Calculate(input EqlInput) (EqlValue, error) {
	if len(g.ctx.AllParam()) != 2 {
		return nil, errors.New("len params just accept 2")
	}
	val1, err := param(g.ctx.Param(0), nil).evaluate(input)
	if err != nil {
		return nil, err
	}
	val2, err := param(g.ctx.Param(1), nil).evaluate(input)
	if err != nil {
		return nil, err
	}

	return NewEqlValue(val1.String() >= val2.String()), nil
}
