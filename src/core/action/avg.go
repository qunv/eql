package action

import (
	"github.com/qunv/eql/src/core/antlr"
)

type _avg struct {
	ctx antlr.IActionSpecContext
}

func avg(ctx antlr.IActionSpecContext) Action {
	return _avg{
		ctx: ctx,
	}
}

func (a _avg) Calculate(input [][]interface{}) (EqlValue, error) {
	params := a.ctx.AllParam()
	result := NewEqlValue(0.0)
	f := func(values []EqlValue) (EqlValue, error) {
		r := NewEqlValue(0.0)
		lenValue := NewEqlValue(float64(len(values)))
		for _, val := range values {
			err := val.Div(lenValue)
			if err != nil {
				return nil, err
			}
			err = r.Add(val)
			if err != nil {
				return nil, err
			}
		}
		return r, nil
	}
	lenValue := NewEqlValue(float64(len(params)))
	for _, p := range params {
		par, err := param(p, f).evaluate(input)
		if err != nil {
			return nil, err
		}
		err = par.Div(lenValue)
		if err != nil {
			return nil, err
		}
		err = result.Add(par)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
