// Code generated from src/antlr/EqlParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package antlr // EqlParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// EqlParserListener is a complete listener for a parse tree produced by EqlParser.
type EqlParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterActSpec is called when entering the actSpec production.
	EnterActSpec(c *ActSpecContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterMagic is called when entering the magic production.
	EnterMagic(c *MagicContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitActSpec is called when exiting the actSpec production.
	ExitActSpec(c *ActSpecContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitMagic is called when exiting the magic production.
	ExitMagic(c *MagicContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)
}
