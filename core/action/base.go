package action

import (
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/val"
	"sync"
)

type baseConcurrent struct {
}

func (b baseConcurrent) evaluateParams(
	store mem.Mem,
	params []antlr.IParamContext,
	done <-chan struct{},
	input EqlInput,
	inputRangeFunc func(values []val.EqlValue) (val.EqlValue, error),
) <-chan result {
	c := make(chan result)
	go func() {
		var wg sync.WaitGroup
		for _, p := range params {
			wg.Add(1)
			go func(ctx antlr.IParamContext) {
				defer wg.Done()
				par, err := param(ctx, inputRangeFunc, store).evaluate(input)
				select {
				case c <- result{par, err}:
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
