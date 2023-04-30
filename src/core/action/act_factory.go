package action

import (
	"github.com/qunv/eql/src/core/antlr"
)

func GetActSpec(ctx antlr.IActionSpecContext) Action {
	type_ := ctx.Type_()
	if type_ == nil {
		panic("type is nil")
	}
	if math := type_.Math(); math != nil {
		if math.SUM() != nil {
			return newSum(ctx)
		}
		if math.ABS() != nil {
			return newAbs(ctx)
		}
	}

	if operator := type_.Operator(); operator != nil {
		if operator.AVG() != nil {
			return newAvg(ctx)
		}
		if operator.ADD() != nil {
			return newAdd(ctx)
		}
	}
	panic("type not support!")
}
