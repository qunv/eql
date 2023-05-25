package action

import (
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
)

func GetActSpec(ctx antlr.IActionSpecContext, store mem.Mem) Action {
	type_ := ctx.Type_()
	if type_ == nil {
		panic("type is nil")
	}
	if math := type_.Math(); math != nil {
		if math.SUM() != nil {
			return sum(ctx, store)
		}
		if math.ABS() != nil {
			return abs(ctx, store)
		}
	}

	if operator := type_.Operator(); operator != nil {
		if operator.AVG() != nil {
			return avg(ctx, store)
		}
		if operator.ADD() != nil {
			return add(ctx, store)
		}
		if operator.EQ() != nil {
			return eq(ctx, store)
		}
		if operator.CONCAT() != nil {
			return concat(ctx, store)
		}
		if operator.MULTIPLY() != nil {
			return multiply(ctx, store)
		}
		if operator.DIVIDE() != nil {
			return divide(ctx, store)
		}
		if operator.GT() != nil {
			return gt(ctx, store)
		}
		if operator.GTE() != nil {
			return gte(ctx, store)
		}
	}

	if logical := type_.Logical(); logical != nil {
		if logical.IF() != nil {
			return if_(ctx, store)
		}
	}
	panic("type not support!")
}
