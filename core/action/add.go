package action

import (
	"errors"
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/val"
)

type _add struct {
	ctx   antlr.IActionSpecContext
	store mem.Mem
}

func add(ctx antlr.IActionSpecContext, store mem.Mem) Action {
	return _add{ctx, store}
}

func (a _add) Evaluate(input EqlInput) (val.EqlValue, error) {
	if len(a.ctx.AllParam()) != 2 {
		return nil, errors.New("len params just accept 2")
	}
	return sum(a.ctx, a.store).Evaluate(input)
}
