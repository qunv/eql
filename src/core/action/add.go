package action

import (
	"errors"
	"github.com/qunv/eql/src/core/antlr"
)

type _add struct {
	ctx antlr.IActionSpecContext
}

func add(ctx antlr.IActionSpecContext) Action {
	return _add{ctx}
}

func (a _add) Calculate(input EqlInput) (EqlValue, error) {
	if len(a.ctx.AllParam()) != 2 {
		return nil, errors.New("len params just accept 2")
	}
	return sum(a.ctx).Calculate(input)
}
