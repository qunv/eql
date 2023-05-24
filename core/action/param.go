package action

import (
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/utils"
	"github.com/qunv/eql/core/val"
)

type fn func([]val.EqlValue) (val.EqlValue, error)

type _param struct {
	ctx   antlr.IParamContext
	store mem.Mem
	f     fn
}

func param(ctx antlr.IParamContext, f fn, store mem.Mem) _param {
	return _param{
		ctx:   ctx,
		f:     f,
		store: store,
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
				for j := col1; j <= col2; j++ {
					values = append(values, val.NewEqlValue(input.Get(i, j)))
				}
			} else {
				values = append(values, val.NewEqlValue(input.Get(i, col1)))
			}
		}
		return e.f(values)
	}
	return newExpression(e.ctx.Expression(), input, e.store).Evaluate()
}
