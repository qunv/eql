package action

import (
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/val"
)

type _sum struct {
	baseConcurrent
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
	doneChan := make(chan struct{})
	defer close(doneChan)
	params := s.ctx.AllParam()
	response := val.NewEqlValue(0.0)
	resultChan := s.evaluateParams(params, doneChan, input)
	for r := range resultChan {
		if r.err != nil {
			return nil, r.err
		}
		if err := response.Add(r.value); err != nil {
			return nil, err
		}
	}
	return response, nil
}

func (s _sum) evaluateParams(params []antlr.IParamContext, done <-chan struct{}, input EqlInput) <-chan result {
	f := func(values []val.EqlValue) (val.EqlValue, error) {
		r := val.NewEqlValue(0.0)
		for _, value := range values {
			if err := r.Add(value); err != nil {
				return nil, err
			}
		}
		return r, nil
	}
	return s.baseConcurrent.evaluateParams(s.store, params, done, input, f)
}
