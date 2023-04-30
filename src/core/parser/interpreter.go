package parser

import (
	"github.com/qunv/eql/src/core/action"
	eqlantlr "github.com/qunv/eql/src/core/antlr"
)

type EqlInterpreter struct {
	*eqlantlr.BaseEqlParserListener
	input  [][]interface{}
	result action.EqlValue
}

func NewEqlInterpreter(input [][]interface{}) *EqlInterpreter {
	return &EqlInterpreter{
		input: input,
	}
}

func (e *EqlInterpreter) Result() action.EqlValue {
	return e.result
}

func (e *EqlInterpreter) ExitStatement(ctx *eqlantlr.StatementContext) {
	expr := action.NewExpression(ctx.Expression(), e.input)
	result, err := expr.Evaluate()
	if err != nil {
		panic(err)
	}
	e.result = result
}
