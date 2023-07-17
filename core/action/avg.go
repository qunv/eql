package action

import (
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/val"
)

type _avg struct {
	baseConcurrent
	ctx   antlr.IActionSpecContext
	store mem.Mem
}

func avg(ctx antlr.IActionSpecContext, store mem.Mem) Action {
	return _avg{
		ctx:   ctx,
		store: store,
	}
}

func (a _avg) Evaluate(input EqlInput) (val.EqlValue, error) {
	doneChan := make(chan struct{})
	defer close(doneChan)
	params := a.ctx.AllParam()
	response := val.NewEqlValue(0.0)
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
	resultChan := a.baseConcurrent.evaluateParams(a.store, params, doneChan, input, f)
	for r := range resultChan {
		if r.err != nil {
			return nil, r.err
		}
		err := r.value.Div(lenValue)
		if err != nil {
			return nil, err
		}
		err = response.Add(r.value)
		if err != nil {
			return nil, err
		}
	}
	return response, nil
}
