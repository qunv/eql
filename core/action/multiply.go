package action

import (
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/val"
)

type _multiply struct {
	ctx   antlr.IActionSpecContext
	store mem.Mem
}

func multiply(ctx antlr.IActionSpecContext, store mem.Mem) Action {
	return _multiply{ctx, store}
}

func (d _multiply) Evaluate(input EqlInput) (val.EqlValue, error) {
	if len(d.ctx.AllParam()) != 2 {
		panic("Len params just accept 2")
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
	return val.NewEqlValue(f1 * f2), nil
}
