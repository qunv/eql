package action

import "github.com/qunv/eql/src/core/antlr"

type eq struct {
	ctx antlr.IActionSpecContext
}

func newEq(ctx antlr.IActionSpecContext) Action {
	return eq{ctx}
}

func (e eq) Calculate(input [][]interface{}) (EqlValue, error) {
	if len(e.ctx.AllParam()) != 2 {
		panic("Len params just accept 2")
	}
	val1, err := newParam(e.ctx.Param(0), nil).evaluate(input)
	if err != nil {
		return nil, err
	}
	val2, err := newParam(e.ctx.Param(1), nil).evaluate(input)
	if err != nil {
		return nil, err
	}

	return NewEqlValue(val1.Value() == val2.Value()), nil
}
