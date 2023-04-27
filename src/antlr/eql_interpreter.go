package antlr

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type EqlInterpreter struct {
	*BaseEqlParserListener
	input [][]float64
	vars  map[string]float64
}

func getColId(alias rune) int {
	return int(alias - 65)
}

func NewGoGoInterpreter() *EqlInterpreter {
	return &EqlInterpreter{
		vars: make(map[string]float64),
		input: [][]float64{
			//A, B, C, D
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
		},
	}
}

func (e *EqlInterpreter) ExitVarDecl(ctx *VarDeclContext) {
	identifier := ctx.IDENTIFIER().GetText()
	value := e.evaluateExpr(ctx.Expression(), e.vars)
	e.vars[identifier] = value
}

func (e *EqlInterpreter) ExitExpression(ctx *ExpressionContext) {
	terms := ctx.AllTerm()
	result := e.evaluateTerm(terms[0], e.vars)
	for i := 1; i < len(terms); i++ {
		op := ctx.GetChild(2*i - 1).(*antlr.TerminalNodeImpl).GetSymbol().GetTokenType()
		operand := e.evaluateTerm(terms[i].(*TermContext), e.vars)
		if op == EqlParserPLUS {
			result += operand
		} else {
			result -= operand
		}
	}
	fmt.Println(result)
}

func (e *EqlInterpreter) evaluateActSpec(ctx IActSpecContext) float64 {
	params := ctx.AllParam()
	result := 0.0
	for _, param := range params {
		if ctx.ACT().GetText() == "=SUM" {
			result += e.evaluateParam(param, func(values []float64) float64 {
				sum := 0.0
				for _, value := range values {
					sum += value
				}
				return sum
			})
		}
		if ctx.ACT().GetText() == "=AVG" {
			result += e.evaluateParam(param, func(values []float64) float64 {
				avg := 0.0
				for _, value := range values {
					avg += value / float64(len(values))
				}
				return avg
			}) / float64(len(params))
		}
	}
	return result
}

func (e *EqlInterpreter) evaluateParam(ctx IParamContext, f func([]float64) float64) float64 {
	if ctx.ActSpec() != nil {
		return e.evaluateActSpec(ctx.ActSpec())
	}

	if ctx.COLON() != nil {
		magic1 := ctx.Magic(0)
		row1, col1 := e.getRowAndColum(magic1)
		magic2 := ctx.Magic(1)
		row2, col2 := e.getRowAndColum(magic2)
		var values []float64
		for i := row1; i <= row2; i++ {
			values = append(values, e.input[i][col1], e.input[i][col2])
		}
		return f(values)
	}
	return e.evaluateMagic(ctx.Magic(0))
}

func (e *EqlInterpreter) evaluateMagic(ctx IMagicContext) float64 {
	row, col := e.getRowAndColum(ctx)
	return e.input[row][col]
}

func (e *EqlInterpreter) getRowAndColum(ctx IMagicContext) (int, int) {
	col := []rune(ctx.IDENTIFIER().GetText())
	row, _ := strconv.Atoi(ctx.NUMBER().GetText())
	return row, getColId(col[0])
}

func (e *EqlInterpreter) evaluateFactor(ctx IFactorContext, stack map[string]float64) float64 {
	if ctx.NUMBER() != nil {
		f, _ := strconv.ParseFloat(ctx.NUMBER().GetText(), 64)
		return f
	}
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		if v, ok := stack[name]; ok {
			return v
		} else {
			panic("Undefined variable: " + name)
		}
	}

	if ctx.ActSpec() != nil {
		return e.evaluateActSpec(ctx.ActSpec())
	}

	return e.evaluateExpr(ctx.Expression(), stack)
}

func (e *EqlInterpreter) evaluateTerm(ctx ITermContext, stack map[string]float64) float64 {
	factors := ctx.AllFactor()

	result := e.evaluateFactor(factors[0], stack)
	for i := 1; i < len(factors); i++ {
		op := ctx.GetChild(2*i - 1).(*antlr.TerminalNodeImpl).GetSymbol().GetTokenType()
		operand := e.evaluateFactor(factors[i].(*FactorContext), stack)
		if op == EqlParserMULT {
			result *= operand
		} else {
			result /= operand
		}
	}
	return result
}

func (e *EqlInterpreter) evaluateExpr(ctx IExpressionContext, stack map[string]float64) float64 {
	result := e.evaluateTerm(ctx.Term(0), stack)
	for i := 1; i < len(ctx.AllTerm()); i++ {
		op := ctx.GetChild(2*i - 1).(*antlr.TerminalNodeImpl).GetSymbol().GetTokenType()
		operand := e.evaluateTerm(ctx.Term(i), stack)
		if op == EqlParserPLUS {
			result += operand
		} else {
			result -= operand
		}
	}
	return result
}
