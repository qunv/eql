package parser

import (
	"github.com/qunv/eql/src/core/act"
	eqlantlr "github.com/qunv/eql/src/core/antlr"
)

type EqlInterpreter struct {
	*eqlantlr.BaseEqlParserListener
	input  [][]float64
	result float64
}

func NewEqlInterpreter(input [][]float64) *EqlInterpreter {
	return &EqlInterpreter{
		input: input,
	}
}

func (e *EqlInterpreter) Result() float64 {
	return e.result
}

func (e *EqlInterpreter) ExitStatement(ctx *eqlantlr.StatementContext) {
	expr := act.NewExpression(ctx.Expression(), e.input)
	e.result = expr.Evaluate()
}

func (e *EqlInterpreter) evaluateActSpec(ctx eqlantlr.IActSpecContext) float64 {
	return act.GetActSpec(ctx).Calculate(e.input)
}
