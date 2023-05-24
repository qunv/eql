package action

import (
	"errors"
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/val"
)

type _divide struct {
	ctx   antlr.IActionSpecContext
	store mem.Mem
}

func divide(ctx antlr.IActionSpecContext, store mem.Mem) Action {
	return _divide{ctx, store}
}

func (d _divide) Evaluate(input EqlInput) (val.EqlValue, error) {
	if len(d.ctx.AllParam()) != 2 {
		return nil, errors.New("len params just accept 2")
	}

	val1, err := param(d.ctx.Param(0), nil, d.store).evaluate(input)
	if err != nil {
		return nil, err
	}
	f1, err := val1.Float64()
	if err != nil {
		return nil, err
	}
	val2, err := param(d.ctx.Param(1), nil, d.store).evaluate(input)
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

	return val.NewEqlValue(f1 / f2), nil
}
