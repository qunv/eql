package action

import (
	"github.com/qunv/eql/src/core/antlr"
	"github.com/qunv/eql/src/core/utils"
)

type fn func([]EqlValue) (EqlValue, error)

type _param struct {
	ctx antlr.IParamContext
	f   fn
}

func param(ctx antlr.IParamContext, f fn) _param {
	return _param{
		ctx: ctx,
		f:   f,
	}
}

func (e _param) evaluate(input [][]interface{}) (EqlValue, error) {
	if e.ctx.ActionSpec() != nil {
		return GetActSpec(e.ctx.ActionSpec()).Calculate(input)
	}

	if e.ctx.Number() != nil {
		return NewEqlValue(utils.GetNumber(e.ctx.Number())), nil
	}

	if e.ctx.Expression() != nil {
		return NewExpression(e.ctx.Expression(), input).Evaluate()
	}

	if e.ctx.InputRange() != nil {
		magic1 := e.ctx.InputRange().Def(0)
		row1, col1 := utils.GetRowAndColum(magic1)
		magic2 := e.ctx.InputRange().Def(1)
		row2, col2 := utils.GetRowAndColum(magic2)
		var values []EqlValue
		for i := row1; i <= row2; i++ {
			values = append(values, NewEqlValue(input[i][col1]), NewEqlValue(input[i][col2]))
		}
		return e.f(values)
	}
	row, col := utils.GetRowAndColum(e.ctx.Def())
	return NewEqlValue(input[row][col]), nil
}
