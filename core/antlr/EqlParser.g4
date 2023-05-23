parser grammar EqlParser;

options {
	tokenVocab = EqlLexer;
}

// Parser rules
program: statement+;

statement
    : expression
    | declarement
    | function
    ;

declarement
    : (cell | IDENT) COLON EQUAL expression
    ;
function
    : PRINT LPAREN expression RPAREN
    ;

expression
    : term ((PLUS | MINUS) term)*
    | term (compair term)?
    ;

compair: (GREATER_THAN | LESS_THAN)? EQUAL?;

term: factor ((MULT | DIV) factor)*;

actionSpec: type LPAREN (param (COMMA param)*) RPAREN;

param
    : inputRange
    | expression
    ;

type
    : math
    | operator
    | logical
    ;

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
    | LAMBDA
    ;


inputRange: cell COLON cell;

cell: ALPHABET INT;

number
    : INT
    | DECIMAL
    ;

factor
    : number
    | cell
    | actionSpec
    | TRUE
    | FALSE
    | IDENT
    | STRING
    | LPAREN expression RPAREN
    ;