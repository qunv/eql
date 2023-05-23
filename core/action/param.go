package action

import (
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/utils"
	"github.com/qunv/eql/core/val"
)

type fn func([]val.EqlValue) (val.EqlValue, error)

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

func (e _param) evaluate(input EqlInput) (val.EqlValue, error) {
	if e.ctx.InputRange() != nil {
		magic1 := e.ctx.InputRange().Cell(0)
		row1, col1, err := utils.GetRowAndColum(magic1)
		if err != nil {
			return nil, err
		}
		magic2 := e.ctx.InputRange().Cell(1)
		row2, col2, err := utils.GetRowAndColum(magic2)
		if err != nil {
			return nil, err
		}
		var values []val.EqlValue
		for i := row1; i <= row2; i++ {
			if col1 != col2 {
				values = append(values, val.NewEqlValue(input.Get(i, col1)), val.NewEqlValue(input.Get(i, col2)))
			} else {
				values = append(values, val.NewEqlValue(input.Get(i, col1)))
			}
		}
		return e.f(values)
	}
	return NewExpression(e.ctx.Expression(), input).Evaluate()
}
