package action

import (
	"errors"
	"github.com/qunv/eql/src/core/antlr"
)

type abs struct {
	ctx antlr.IActionSpecContext
}

func newAbs(ctx antlr.IActionSpecContext) Action {
	return abs{ctx}
}

func (a abs) Calculate(input [][]interface{}) (EqlValue, error) {
	params := a.ctx.AllParam()
	if len(params) != 1 {
		return nil, errors.New("error: len param must be 1")
	}

	p := a.ctx.Param(0)
	result, err := newParam(p, nil).evaluate(input)

	if err != nil {
		return nil, err
	}

	return NewEqlValue(result), nil
}
