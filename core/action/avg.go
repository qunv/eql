package action

import (
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/val"
)

type _avg struct {
	ctx antlr.IActionSpecContext
}

func avg(ctx antlr.IActionSpecContext) Action {
	return _avg{
		ctx: ctx,
	}
}

func (a _avg) Evaluate(input EqlInput) (val.EqlValue, error) {
	params := a.ctx.AllParam()
	result := val.NewEqlValue(0.0)
	f := func(values []val.EqlValue) (val.EqlValue, error) {
		r := val.NewEqlValue(0.0)
		lenValue := val.NewEqlValue(float64(len(values)))
		for _, value := range values {
			err := value.Div(lenValue)
			if err != nil {
				return nil, err
			}
			err = r.Add(value)
			if err != nil {
				return nil, err
			}
		}
		return r, nil
	}
	lenValue := val.NewEqlValue(float64(len(params)))
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
