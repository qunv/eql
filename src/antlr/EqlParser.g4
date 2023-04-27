parser grammar EqlParser;

options {
	tokenVocab = EqlLexer;
}

// Parser rules
program: statement+;

statement
    : declaration
    | expression
    ;

expression
    : term ((PLUS | MINUS) term)*
    ;

declaration
    : varDecl
    ;

term: factor ((MULT | DIV) factor)*;

varDecl: IDENTIFIER EQUAL expression SEMI;
actSpec: ACT LPAREN (param (COMMA param)*) RPAREN;
param: magic (COLON magic)? | actSpec;
magic: IDENTIFIER NUMBER;
factor:
    IDENTIFIER
    | NUMBER
    | actSpec
    | LPAREN expression RPAREN
    ;