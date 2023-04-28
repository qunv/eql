package act

import (
	antlr2 "github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/qunv/eql/src/core/antlr"
	"github.com/qunv/eql/src/core/utils"
)

type Expression struct {
	ctx   antlr.IExpressionContext
	input [][]float64
}

func NewExpression(ctx antlr.IExpressionContext, input [][]float64) *Expression {
	return &Expression{
		ctx:   ctx,
		input: input,
	}
}

func (e *Expression) Evaluate() float64 {
	ctx := e.ctx
	result := e.evaluateTerm(ctx.Term(0))
	for i := 1; i < len(ctx.AllTerm()); i++ {
		op := ctx.GetChild(2*i - 1).(*antlr2.TerminalNodeImpl).GetSymbol().GetTokenType()
		operand := e.evaluateTerm(ctx.Term(i))
		if op == antlr.EqlParserPLUS {
			result += operand
		} else {
			result -= operand
		}
	}
	return result
}

func (e *Expression) evaluateActSpec(ctx antlr.IActSpecContext) float64 {
	return GetActSpec(ctx).Calculate(e.input)
}

func (e *Expression) evaluateFactor(ctx antlr.IFactorContext) float64 {
	if ctx.Number() != nil {
		return utils.GetNumber(ctx.Number())
	}

	if ctx.ActSpec() != nil {
		return e.evaluateActSpec(ctx.ActSpec())
	}

	if ctx.Def() != nil {
		row, col := utils.GetRowAndColum(ctx.Def())
		return e.input[row][col]
	}

	return NewExpression(ctx.Expression(), e.input).Evaluate()
}

func (e *Expression) evaluateTerm(ctx antlr.ITermContext) float64 {
	factors := ctx.AllFactor()

	result := e.evaluateFactor(factors[0])
	for i := 1; i < len(factors); i++ {
		op := ctx.GetChild(2*i - 1).(*antlr2.TerminalNodeImpl).GetSymbol().GetTokenType()
		operand := e.evaluateFactor(factors[i].(*antlr.FactorContext))
		if op == antlr.EqlParserMULT {
			result *= operand
		} else {
			result /= operand
		}
	}
	return result
}
