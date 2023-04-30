package action

import (
	antlr2 "github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/qunv/eql/src/core/antlr"
	"github.com/qunv/eql/src/core/utils"
)

type Expression struct {
	ctx   antlr.IExpressionContext
	input [][]interface{}
}

func NewExpression(ctx antlr.IExpressionContext, input [][]interface{}) *Expression {
	return &Expression{
		ctx:   ctx,
		input: input,
	}
}

func (e *Expression) Evaluate() (EqlValue, error) {
	ctx := e.ctx
	result, err := e.evaluateTerm(ctx.Term(0))
	if err != nil {
		return nil, err
	}
	for i := 1; i < len(ctx.AllTerm()); i++ {
		op := ctx.GetChild(2*i - 1).(*antlr2.TerminalNodeImpl).GetSymbol().GetTokenType()
		operand, err := e.evaluateTerm(ctx.Term(i))
		if err != nil {
			return nil, err
		}

		if op == antlr.EqlParserPLUS {
			err = result.Add(operand)
		} else {
			err = result.Minus(operand)
		}
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (e *Expression) evaluateActSpec(ctx antlr.IActionSpecContext) (EqlValue, error) {
	return GetActSpec(ctx).Calculate(e.input)
}

func (e *Expression) evaluateFactor(ctx antlr.IFactorContext) (EqlValue, error) {
	if ctx.Number() != nil {
		return NewEqlValue(utils.GetNumber(ctx.Number())), nil
	}

	if ctx.ActionSpec() != nil {
		return e.evaluateActSpec(ctx.ActionSpec())
	}

	if ctx.Def() != nil {
		row, col := utils.GetRowAndColum(ctx.Def())
		return NewEqlValue(e.input[row][col]), nil
	}

	return NewExpression(ctx.Expression(), e.input).Evaluate()
}

func (e *Expression) evaluateTerm(ctx antlr.ITermContext) (EqlValue, error) {
	factors := ctx.AllFactor()

	result, err := e.evaluateFactor(factors[0])
	if err != nil {
		return nil, err
	}
	for i := 1; i < len(factors); i++ {
		op := ctx.GetChild(2*i - 1).(*antlr2.TerminalNodeImpl).GetSymbol().GetTokenType()
		operand, err := e.evaluateFactor(factors[i].(*antlr.FactorContext))
		if err != nil {
			return nil, err
		}
		if op == antlr.EqlParserMULT {
			err = result.Mul(operand)
		} else {
			err = result.Div(operand)
		}
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
