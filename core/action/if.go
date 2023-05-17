package action

import (
	"errors"
	"github.com/qunv/eql/core/antlr"
	"reflect"
)

type _if struct {
	ctx antlr.IActionSpecContext
}

func if_(ctx antlr.IActionSpecContext) Action {
	return _if{ctx}
}

func (i _if) Evaluate(input EqlInput) (EqlValue, error) {
	if len(i.ctx.AllParam()) != 3 {
		return nil, errors.New("len params just accept 2")
	}
	logicalExpr, err := param(i.ctx.Param(0), nil).evaluate(input)
	if err != nil {
		return nil, err
	}
	trueValue, err := param(i.ctx.Param(1), nil).evaluate(input)
	if err != nil {
		return nil, err
	}

	falseValue, err := param(i.ctx.Param(2), nil).evaluate(input)
	if err != nil {
		return nil, err
	}

	if !logicalExpr.IsBool() {
		return nil, errors.New("parameter 1 expects boolean values")
	}

	check := reflect.ValueOf(logicalExpr.Value()).Bool()
	if check {
		return trueValue, nil
	}
	return falseValue, nil
}
