package action

import "github.com/qunv/eql/src/core/antlr"

type add struct {
	ctx antlr.IActionSpecContext
}

func newAdd(ctx antlr.IActionSpecContext) Action {
	return add{ctx}
}

func (a add) Calculate(input [][]interface{}) (EqlValue, error) {
	if len(a.ctx.AllParam()) != 2 {
		panic("Len params just accept 2")
	}
	return newSum(a.ctx).Calculate(input)
}
