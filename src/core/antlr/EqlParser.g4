parser grammar EqlParser;

options {
	tokenVocab = EqlLexer;
}

// Parser rules
program: statement+;

statement
    : expression
    ;

expression
    : term ((PLUS | MINUS) term)*
    ;

term: factor ((MULT | DIV) factor)*;

actSpec: act LPAREN (param (SEMI param)*) RPAREN;
act: SUM | AVG;
param: magic (COLON magic)? | actSpec;
magic: IDENTIFIER NUMBER;
factor:
    NUMBER
    | magic
    | actSpec
    | LPAREN expression RPAREN
    ;