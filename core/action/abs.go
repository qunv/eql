package action

import (
	"errors"
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/val"
	"math"
)

type _abs struct {
	ctx   antlr.IActionSpecContext
	store mem.Mem
}

func abs(ctx antlr.IActionSpecContext, store mem.Mem) Action {
	return _abs{ctx, store}
}

func (a _abs) Evaluate(input EqlInput) (val.EqlValue, error) {
	params := a.ctx.AllParam()
	if len(params) != 1 {
		return nil, errors.New("error: len _param must be 1")
	}

	p := a.ctx.Param(0)
	result, err := param(p, nil, a.store).evaluate(input)

	if err != nil {
		return nil, err
	}
	value, err := result.Float64()
	if err != nil {
		return nil, err
	}
	return val.NewEqlValue(math.Abs(value)), nil
}
