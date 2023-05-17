package action

import (
	"fmt"
	"github.com/qunv/eql/core/antlr"
)

type _concat struct {
	ctx antlr.IActionSpecContext
}

func concat(ctx antlr.IActionSpecContext) Action {
	return _concat{ctx}
}

func (c _concat) Evaluate(input EqlInput) (EqlValue, error) {
	if len(c.ctx.AllParam()) != 2 {
		panic("Len params just accept 2")
	}

	val1, err := param(c.ctx.Param(0), nil).evaluate(input)
	if err != nil {
		return nil, err
	}
	val2, err := param(c.ctx.Param(1), nil).evaluate(input)
	if err != nil {
		return nil, err
	}
	return NewEqlValue(fmt.Sprintf("%v%v", val1.Value(), val2.Value())), nil
}
