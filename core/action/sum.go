package action

import (
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/val"
)

type _sum struct {
	ctx   antlr.IActionSpecContext
	store mem.Mem
}

func sum(ctx antlr.IActionSpecContext, store mem.Mem) Action {
	return _sum{
		ctx:   ctx,
		store: store,
	}
}

func (s _sum) Evaluate(input EqlInput) (val.EqlValue, error) {
	params := s.ctx.AllParam()
	f := func(values []val.EqlValue) (val.EqlValue, error) {
		r := val.NewEqlValue(0.0)
		for _, value := range values {
			if err := r.Add(value); err != nil {
				return nil, err
			}
		}
		return r, nil
	}
	result := val.NewEqlValue(0.0)
	for _, p := range params {
		par, err := param(p, f, s.store).evaluate(input)
		if err != nil {
			return nil, err
		}
		if err = result.Add(par); err != nil {
			return nil, err
		}
	}
	return result, nil
}
