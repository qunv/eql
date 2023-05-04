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

func (e _param) evaluate(input EqlInput) (EqlValue, error) {
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
		row1, col1, err := utils.GetRowAndColum(magic1)
		if err != nil {
			return nil, err
		}
		magic2 := e.ctx.InputRange().Def(1)
		row2, col2, err := utils.GetRowAndColum(magic2)
		if err != nil {
			return nil, err
		}
		var values []EqlValue
		for i := row1; i <= row2; i++ {
			values = append(values, NewEqlValue(input.Get(i, col1)), NewEqlValue(input.Get(i, col2)))
		}
		return e.f(values)
	}
	row, col, err := utils.GetRowAndColum(e.ctx.Def())
	if err != nil {
		return nil, err
	}
	return NewEqlValue(input.Get(row, col)), nil
}
