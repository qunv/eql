package action

import (
	"github.com/qunv/eql/src/core/antlr"
)

type sum struct {
	ctx antlr.IActionSpecContext
}

func newSum(ctx antlr.IActionSpecContext) Action {
	return sum{
		ctx: ctx,
	}
}

func (s sum) Calculate(input [][]interface{}) (EqlValue, error) {
	params := s.ctx.AllParam()
	f := func(values []EqlValue) (EqlValue, error) {
		r := NewEqlValue(0.0)
		for _, val := range values {
			if err := r.Add(val); err != nil {
				return nil, err
			}
		}
		return r, nil
	}
	result := NewEqlValue(0.0)
	for _, p := range params {
		par, err := newParam(p, f).evaluate(input)
		if err != nil {
			return nil, err
		}
		if err = result.Add(par); err != nil {
			return nil, err
		}
	}
	return result, nil
}
