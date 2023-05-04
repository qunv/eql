package core

import (
	"github.com/qunv/eql/core/action"
	"github.com/qunv/eql/core/antlr"
)

type EqlInterpreter struct {
	*antlr.BaseEqlParserListener
	input  action.EqlInput
	result action.EqlValue
	err    error
}

func NewEqlInterpreter(input action.EqlInput) *EqlInterpreter {
	return &EqlInterpreter{
		input: input,
	}
}

func (e *EqlInterpreter) Result() (action.EqlValue, error) {
	return e.result, e.err
}

func (e *EqlInterpreter) ExitStatement(ctx *antlr.StatementContext) {
	expr := action.NewExpression(ctx.Expression(), e.input)
	result, err := expr.Evaluate()
	if err != nil {
		e.err = err
	}
	e.result = result
}
