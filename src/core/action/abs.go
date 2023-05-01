package action

import (
	"errors"
	"github.com/qunv/eql/src/core/antlr"
)

type _abs struct {
	ctx antlr.IActionSpecContext
}

func abs(ctx antlr.IActionSpecContext) Action {
	return _abs{ctx}
}

func (a _abs) Calculate(input [][]interface{}) (EqlValue, error) {
	params := a.ctx.AllParam()
	if len(params) != 1 {
		return nil, errors.New("error: len _param must be 1")
	}

	p := a.ctx.Param(0)
	result, err := param(p, nil).evaluate(input)

	if err != nil {
		return nil, err
	}

	return NewEqlValue(result), nil
}
