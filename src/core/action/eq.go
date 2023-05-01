package action

import "github.com/qunv/eql/src/core/antlr"

type _eq struct {
	ctx antlr.IActionSpecContext
}

func eq(ctx antlr.IActionSpecContext) Action {
	return _eq{ctx}
}

func (e _eq) Calculate(input [][]interface{}) (EqlValue, error) {
	if len(e.ctx.AllParam()) != 2 {
		panic("Len params just accept 2")
	}
	val1, err := param(e.ctx.Param(0), nil).evaluate(input)
	if err != nil {
		return nil, err
	}
	val2, err := param(e.ctx.Param(1), nil).evaluate(input)
	if err != nil {
		return nil, err
	}

	return NewEqlValue(val1.Value() == val2.Value()), nil
}
