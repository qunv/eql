package act

import (
	"github.com/qunv/eql/src/core/antlr"
	"github.com/qunv/eql/src/core/utils"
)

type param struct {
	ctx antlr.IParamContext
	f   func([]float64) float64
}

func newParam(ctx antlr.IParamContext, f func([]float64) float64) param {
	return param{
		ctx: ctx,
		f:   f,
	}
}

func (e param) evaluate(input [][]float64) float64 {
	if e.ctx.ActSpec() != nil {
		return GetActSpec(e.ctx.ActSpec()).Calculate(input)
	}

	if e.ctx.COLON() != nil {
		magic1 := e.ctx.Magic(0)
		row1, col1 := utils.GetRowAndColum(magic1)
		magic2 := e.ctx.Magic(1)
		row2, col2 := utils.GetRowAndColum(magic2)
		var values []float64
		for i := row1; i <= row2; i++ {
			values = append(values, input[i][col1], input[i][col2])
		}
		return e.f(values)
	}
	row, col := utils.GetRowAndColum(e.ctx.Magic(0))
	return input[row][col]
}
