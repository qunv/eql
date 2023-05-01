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

actionSpec: type LPAREN (param (SEMI param)*) RPAREN;
type: math | operator;
math
    : SUM
    | ABS
    ;
operator
    : AVG
    | ADD
    | EQ
    | CONCAT
    | DIVIDE
    ;

param
    : def
    | inputRange
    | number
    | actionSpec
    | expression
    ;
inputRange: def COLON def;
def: IDENTIFIER DIGIT;
number: MINUS? DIGIT;
factor:
    number
    | def
    | actionSpec
    | LPAREN expression RPAREN
    ;