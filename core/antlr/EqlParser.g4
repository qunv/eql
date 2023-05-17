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
    | term (compair term)?
    ;

compair: (GREATER_THAN | LESS_THAN)? EQUAL?;

term: factor ((MULT | DIV) factor)*;

actionSpec: type LPAREN (param (COMMA param)*) RPAREN;

type: math | operator | logical;

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
    | MULTIPLY
    | GT
    | GTE
    ;

logical
    : IF
    ;

param
    : def
    | inputRange
    | number
    | STRING
    | actionSpec
    | expression
    ;

inputRange: def COLON def;

def: IDENTIFIER INT;

number
    : INT
    | DECIMAL
    ;

factor
    : number
    | def
    | actionSpec
    | LPAREN expression RPAREN
    | TRUE
    | FALSE
    ;