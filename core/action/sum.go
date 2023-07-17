package action

import (
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/val"
	"sync"
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
	donec := make(chan struct{})
	defer close(donec)
	params := s.ctx.AllParam()
	result := val.NewEqlValue(0.0)
	resultChan := s.evaluateParams(params, donec, input)
	for r := range resultChan {
		if r.err != nil {
			return nil, r.err
		}
		if err := result.Add(r.value); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (s _sum) evaluateParams(params []antlr.IParamContext, done <-chan struct{}, input EqlInput) <-chan sumResult {
	c := make(chan sumResult)
	f := func(values []val.EqlValue) (val.EqlValue, error) {
		r := val.NewEqlValue(0.0)
		for _, value := range values {
			if err := r.Add(value); err != nil {
				return nil, err
			}
		}
		return r, nil
	}
	go func() {
		var wg sync.WaitGroup
		for _, p := range params {
			wg.Add(1)
			go func(ctx antlr.IParamContext) {
				defer wg.Done()
				par, err := param(ctx, f, s.store).evaluate(input)
				select {
				case c <- sumResult{par, err}:
				case <-done:
					return
				}
			}(p)
		}
		go func() {
			wg.Wait()
			close(c)
		}()
	}()
	return c
}

type sumResult struct {
	value val.EqlValue
	err   error
}
