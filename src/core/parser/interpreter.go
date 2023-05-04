package parser

import (
	"github.com/qunv/eql/src/core/action"
	eqlantlr "github.com/qunv/eql/src/core/antlr"
)

type EqlInterpreter struct {
	*eqlantlr.BaseEqlParserListener
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

func (e *EqlInterpreter) ExitStatement(ctx *eqlantlr.StatementContext) {
	expr := action.NewExpression(ctx.Expression(), e.input)
	result, err := expr.Evaluate()
	if err != nil {
		e.err = err
	}
	e.result = result
}
