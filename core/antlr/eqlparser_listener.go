// Code generated from core/antlr/EqlParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package antlr // EqlParser
import "github.com/antlr4-go/antlr/v4"

// EqlParserListener is a complete listener for a parse tree produced by EqlParser.
type EqlParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDeclarement is called when entering the declarement production.
	EnterDeclarement(c *DeclarementContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCompair is called when entering the compair production.
	EnterCompair(c *CompairContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterActionSpec is called when entering the actionSpec production.
	EnterActionSpec(c *ActionSpecContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterMath is called when entering the math production.
	EnterMath(c *MathContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterLogical is called when entering the logical production.
	EnterLogical(c *LogicalContext)

	// EnterInputRange is called when entering the inputRange production.
	EnterInputRange(c *InputRangeContext)

	// EnterCell is called when entering the cell production.
	EnterCell(c *CellContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDeclarement is called when exiting the declarement production.
	ExitDeclarement(c *DeclarementContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCompair is called when exiting the compair production.
	ExitCompair(c *CompairContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitActionSpec is called when exiting the actionSpec production.
	ExitActionSpec(c *ActionSpecContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitMath is called when exiting the math production.
	ExitMath(c *MathContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitLogical is called when exiting the logical production.
	ExitLogical(c *LogicalContext)

	// ExitInputRange is called when exiting the inputRange production.
	ExitInputRange(c *InputRangeContext)

	// ExitCell is called when exiting the cell production.
	ExitCell(c *CellContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)
}
