package action

import (
	"errors"
	antlr2 "github.com/antlr4-go/antlr/v4"
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/utils"
)

type Expression struct {
	ctx   antlr.IExpressionContext
	input EqlInput
}

func NewExpression(ctx antlr.IExpressionContext, input EqlInput) *Expression {
	return &Expression{
		ctx:   ctx,
		input: input,
	}
}

func (e *Expression) Evaluate() (EqlValue, error) {
	ctx := e.ctx
	if ctx.Compair() != nil {
		return e.evaluateLogicalExpression(ctx)
	}
	return e.evaluateExpression(ctx)
}

func (e *Expression) evaluateLogicalExpression(ctx antlr.IExpressionContext) (EqlValue, error) {
	if len(ctx.AllTerm()) != 1 && len(ctx.AllTerm()) != 2 {
		return nil, errors.New("logical expression must contain 1 or 2 param")
	}

	if len(ctx.AllTerm()) == 1 {
		result, err := e.evaluateTerm(ctx.Term(0))
		if err != nil {
			return nil, err
		}
		if result.IsBool() {
			return result, nil
		}
		return nil, errors.New("parameter 1 expects boolean values")
	}

	param1, err := e.evaluateTerm(ctx.Term(0))
	if err != nil {
		return nil, err
	}
	param2, err := e.evaluateTerm(ctx.Term(1))
	if err != nil {
		return nil, err
	}
	if ctx.Compair().GREATER_THAN() != nil {
		if ctx.Compair().EQUAL() != nil {
			return NewEqlValue(param1.String() >= param2.String()), nil
		}
		return NewEqlValue(param1.String() > param2.String()), nil
	}

	if ctx.Compair().LESS_THAN() != nil {
		if ctx.Compair().EQUAL() != nil {
			return NewEqlValue(param1.String() <= param2.String()), nil
		}
		return NewEqlValue(param1.String() < param2.String()), nil
	}

	return NewEqlValue(param1.String() == param2.String()), nil
}

func (e *Expression) evaluateExpression(ctx antlr.IExpressionContext) (EqlValue, error) {
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
	return GetActSpec(ctx).Evaluate(e.input)
}

func (e *Expression) evaluateFactor(ctx antlr.IFactorContext) (EqlValue, error) {
	if ctx.Number() != nil {
		return NewEqlValue(utils.GetNumber(ctx.Number())), nil
	}

	if ctx.TRUE() != nil {
		return NewEqlValue(true), nil
	}

	if ctx.FALSE() != nil {
		return NewEqlValue(false), nil
	}

	if ctx.STRING() != nil {
		str := ctx.STRING().GetText()
		if len(str) <= 2 {
			return NewEqlValue(""), nil
		}
		return NewEqlValue(str[1 : len(str)-1]), nil
	}

	if ctx.ActionSpec() != nil {
		return e.evaluateActSpec(ctx.ActionSpec())
	}

	if ctx.Def() != nil {
		row, col, err := utils.GetRowAndColum(ctx.Def())
		if err != nil {
			return nil, err
		}
		return NewEqlValue(e.input.Get(row, col)), nil
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
