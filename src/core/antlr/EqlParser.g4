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
act: math;
math: SUM | AVG;
param: def (COLON def)? | number | actSpec | expression ;
def: IDENTIFIER DIGIT;
number: MINUS? DIGIT;
factor:
    number
    | def
    | actSpec
    | LPAREN expression RPAREN
    ;