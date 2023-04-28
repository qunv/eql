package act

import (
	"github.com/qunv/eql/src/core/antlr"
)

func GetActSpec(ctx antlr.IActSpecContext) Act {
	if ctx.Act().SUM() != nil {
		return newSum(ctx)
	}
	if ctx.Act().AVG() != nil {
		return newAvg(ctx)
	}
	panic("Act not support!")
}
