package action

import (
	"errors"
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/val"
)

type _eq struct {
	ctx antlr.IActionSpecContext
	store mem.Mem
}

func eq(ctx antlr.IActionSpecContext, store mem.Mem) Action {
	return _eq{ctx, store}
}

func (e _eq) Evaluate(input EqlInput) (val.EqlValue, error) {
	if len(e.ctx.AllParam()) != 2 {
		return nil, errors.New("len params just accept 2")
	}
	val1, err := param(e.ctx.Param(0), nil, e.store).evaluate(input)
	if err != nil {
		return nil, err
	}
	val2, err := param(e.ctx.Param(1), nil, e.store).evaluate(input)
	if err != nil {
		return nil, err
	}

	return val.NewEqlValue(val1.Value() == val2.Value()), nil
}
