package action

import (
	"errors"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/val"

	"github.com/qunv/eql/core/antlr"
)

type _gt struct {
	ctx   antlr.IActionSpecContext
	store mem.Mem
}

func gt(ctx antlr.IActionSpecContext, store mem.Mem) Action {
	return _gt{ctx, store}
}

func (g _gt) Evaluate(input EqlInput) (val.EqlValue, error) {
	if len(g.ctx.AllParam()) != 2 {
		return nil, errors.New("len params just accept 2")
	}
	val1, err := param(g.ctx.Param(0), nil, g.store).evaluate(input)
	if err != nil {
		return nil, err
	}
	val2, err := param(g.ctx.Param(1), nil, g.store).evaluate(input)
	if err != nil {
		return nil, err
	}

	return val.NewEqlValue(val1.String() > val2.String()), nil
}
