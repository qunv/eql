package action

import (
	"errors"
	antlr2 "github.com/antlr4-go/antlr/v4"
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/utils"
	"github.com/qunv/eql/core/val"
)

type expression struct {
	ctx   antlr.IExpressionContext
	input EqlInput
	store mem.Mem
}

func newExpression(ctx antlr.IExpressionContext, input EqlInput, store mem.Mem) *expression {
	return &expression{
		ctx:   ctx,
		input: input,
		store: store,
	}
}

func (e *expression) Evaluate() (val.EqlValue, error) {
	ctx := e.ctx
	if ctx.Compair() != nil {
		return e.evaluateLogicalExpression(ctx)
	}
	return e.evaluateExpression(ctx)
}

func (e *expression) evaluateLogicalExpression(ctx antlr.IExpressionContext) (val.EqlValue, error) {
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
			return val.NewEqlValue(param1.String() >= param2.String()), nil
		}
		return val.NewEqlValue(param1.String() > param2.String()), nil
	}

	if ctx.Compair().LESS_THAN() != nil {
		if ctx.Compair().EQUAL() != nil {
			return val.NewEqlValue(param1.String() <= param2.String()), nil
		}
		return val.NewEqlValue(param1.String() < param2.String()), nil
	}

	return val.NewEqlValue(param1.String() == param2.String()), nil
}

func (e *expression) evaluateExpression(ctx antlr.IExpressionContext) (val.EqlValue, error) {
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

func (e *expression) evaluateActSpec(ctx antlr.IActionSpecContext) (val.EqlValue, error) {
	return GetActSpec(ctx, e.store).Evaluate(e.input)
}

func (e *expression) evaluateFactor(ctx antlr.IFactorContext) (val.EqlValue, error) {
	if ctx.Number() != nil {
		return val.NewEqlValue(utils.GetNumber(ctx.Number())), nil
	}

	if ctx.TRUE() != nil {
		return val.NewEqlValue(true), nil
	}

	if ctx.FALSE() != nil {
		return val.NewEqlValue(false), nil
	}

	if ctx.STRING() != nil {
		str := ctx.STRING().GetText()
		if len(str) <= 2 {
			return val.NewEqlValue(""), nil
		}
		return val.NewEqlValue(str[1 : len(str)-1]), nil
	}

	if ctx.ActionSpec() != nil {
		return e.evaluateActSpec(ctx.ActionSpec())
	}

	if ctx.Cell() != nil {
		row, col, err := utils.GetRowAndColum(ctx.Cell())
		if err != nil {
			return nil, err
		}
		return val.NewEqlValue(e.input.Get(row, col)), nil
	}

	if ctx.IDENT() != nil {
		return e.store.Get(ctx.IDENT().GetText()), nil
	}

	return newExpression(ctx.Expression(), e.input, e.store).Evaluate()
}

func (e *expression) evaluateTerm(ctx antlr.ITermContext) (val.EqlValue, error) {
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
