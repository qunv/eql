package action

import (
	"fmt"
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/val"
)

type _concat struct {
	ctx   antlr.IActionSpecContext
	store mem.Mem
}

func concat(ctx antlr.IActionSpecContext, store mem.Mem) Action {
	return _concat{ctx, store}
}

func (c _concat) Evaluate(input EqlInput) (val.EqlValue, error) {
	if len(c.ctx.AllParam()) != 2 {
		panic("Len params just accept 2")
	}

	val1, err := param(c.ctx.Param(0), nil, c.store).evaluate(input)
	if err != nil {
		return nil, err
	}
	val2, err := param(c.ctx.Param(1), nil, c.store).evaluate(input)
	if err != nil {
		return nil, err
	}
	return val.NewEqlValue(fmt.Sprintf("%v%v", val1.Value(), val2.Value())), nil
}
