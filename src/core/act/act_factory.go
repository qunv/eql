package act

import (
	"github.com/qunv/eql/src/core/antlr"
)

func GetActSpec(ctx antlr.IActSpecContext) Act {
	if ctx.Act().Math().SUM() != nil {
		return newSum(ctx)
	}
	if ctx.Act().Math().AVG() != nil {
		return newAvg(ctx)
	}
	if ctx.Act().Math().ABS() != nil {
		return newAbs(ctx)
	}
	panic("Act not support!")
}
