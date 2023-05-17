package action

import (
	"github.com/qunv/eql/core/antlr"
)

func GetActSpec(ctx antlr.IActionSpecContext) Action {
	type_ := ctx.Type_()
	if type_ == nil {
		panic("type is nil")
	}
	if math := type_.Math(); math != nil {
		if math.SUM() != nil {
			return sum(ctx)
		}
		if math.ABS() != nil {
			return abs(ctx)
		}
	}

	if operator := type_.Operator(); operator != nil {
		if operator.AVG() != nil {
			return avg(ctx)
		}
		if operator.ADD() != nil {
			return add(ctx)
		}
		if operator.EQ() != nil {
			return eq(ctx)
		}
		if operator.CONCAT() != nil {
			return concat(ctx)
		}
		if operator.MULTIPLY() != nil {
			return multiply(ctx)
		}
		if operator.DIVIDE() != nil {
			return divide(ctx)
		}
		if operator.GT() != nil {
			return gt(ctx)
		}
		if operator.GTE() != nil {
			return gte(ctx)
		}
	}

	if logical := type_.Logical(); logical != nil {
		if logical.IF() != nil {
			return if_(ctx)
		}
	}
	panic("type not support!")
}
