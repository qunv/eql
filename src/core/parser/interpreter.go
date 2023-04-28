package parser

import (
	"github.com/qunv/eql/src/core/act"
	eqlantlr "github.com/qunv/eql/src/core/antlr"
	"github.com/qunv/eql/src/core/utils"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
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

func (e *EqlInterpreter) ExitExpression(ctx *eqlantlr.ExpressionContext) {
	terms := ctx.AllTerm()
	result := e.evaluateTerm(terms[0])
	for i := 1; i < len(terms); i++ {
		op := ctx.GetChild(2*i - 1).(*antlr.TerminalNodeImpl).GetSymbol().GetTokenType()
		operand := e.evaluateTerm(terms[i].(*eqlantlr.TermContext))
		if op == eqlantlr.EqlParserPLUS {
			result += operand
		} else {
			result -= operand
		}
	}
	e.result = result
}

func (e *EqlInterpreter) evaluateActSpec(ctx eqlantlr.IActSpecContext) float64 {
	return act.GetActSpec(ctx).Calculate(e.input)
}

func (e *EqlInterpreter) evaluateFactor(ctx eqlantlr.IFactorContext) float64 {
	if ctx.NUMBER() != nil {
		f, _ := strconv.ParseFloat(ctx.NUMBER().GetText(), 64)
		return f
	}

	if ctx.ActSpec() != nil {
		return e.evaluateActSpec(ctx.ActSpec())
	}

	if ctx.Magic() != nil {
		row, col := utils.GetRowAndColum(ctx.Magic())
		return e.input[row][col]
	}

	return e.evaluateExpr(ctx.Expression())
}

func (e *EqlInterpreter) evaluateTerm(ctx eqlantlr.ITermContext) float64 {
	factors := ctx.AllFactor()

	result := e.evaluateFactor(factors[0])
	for i := 1; i < len(factors); i++ {
		op := ctx.GetChild(2*i - 1).(*antlr.TerminalNodeImpl).GetSymbol().GetTokenType()
		operand := e.evaluateFactor(factors[i].(*eqlantlr.FactorContext))
		if op == eqlantlr.EqlParserMULT {
			result *= operand
		} else {
			result /= operand
		}
	}
	return result
}

func (e *EqlInterpreter) evaluateExpr(ctx eqlantlr.IExpressionContext) float64 {
	result := e.evaluateTerm(ctx.Term(0))
	for i := 1; i < len(ctx.AllTerm()); i++ {
		op := ctx.GetChild(2*i - 1).(*antlr.TerminalNodeImpl).GetSymbol().GetTokenType()
		operand := e.evaluateTerm(ctx.Term(i))
		if op == eqlantlr.EqlParserPLUS {
			result += operand
		} else {
			result -= operand
		}
	}
	return result
}
