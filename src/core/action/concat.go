package action

import (
	"fmt"
	"github.com/qunv/eql/src/core/antlr"
)

type concat struct {
	ctx antlr.IActionSpecContext
}

func newConcat(ctx antlr.IActionSpecContext) Action {
	return concat{ctx}
}

func (c concat) Calculate(input [][]interface{}) (EqlValue, error) {
	if len(c.ctx.AllParam()) != 2 {
		panic("Len params just accept 2")
	}

	val1, err := newParam(c.ctx.Param(0), nil).evaluate(input)
	if err != nil {
		return nil, err
	}
	val2, err := newParam(c.ctx.Param(1), nil).evaluate(input)
	if err != nil {
		return nil, err
	}
	return NewEqlValue(fmt.Sprintf("%v%v", val1.Value(), val2.Value())), nil
}
