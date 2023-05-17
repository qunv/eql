lexer grammar EqlLexer;

// Key word

PLUS                    : '+';
MINUS                   : '-';
DIV                     : '/';
MULT                    : '*';
LPAREN                  : '(';
RPAREN                  : ')';
LCURLY                  : '{';
RCURLY                  : '}';
LBRACKET                : '[';
RBRACKET                : ']';
EQUAL                   : '=';
SEMI                    : ';';
COLON                   : ':';
COMMA                   : ',';
DOT                     : '.';
GREATER_THAN            : '>';
LESS_THAN               : '<';
TRUE                    : 'TRUE';
FALSE                   : 'FALSE';

// math
SUM                     : 'SUM';
ABS                     : 'ABS';

// operator
AVG                     : 'AVG';
ADD                     : 'ADD';
DIVIDE                  : 'DIVIDE';
MULTIPLY                : 'MULTIPLY';
EQ                      : 'EQ';
CONCAT                  : 'CONCAT';
GT                      : 'GT';
GTE                     : 'GTE';

INT                     : MINUS? DIGITS;
DECIMAL                 : MINUS? DIGITS DOT DIGITS;
DIGIT                   : [0-9];


STRING : '"' ( ESC_SEQ | ~('\\'|'"') )* '"';
fragment ESC_SEQ : '\\' [btnfr"'\\];
fragment DIGITS: DIGIT+;

// logical
IF                      : 'IF';

IDENTIFIER              : [a-zA-Z]+;

WS                      : [ \t\r\n]+ -> skip;

EOS                     : ([\r\n]+ | ';' | '/*' .*? '*/' | EOF) -> mode(DEFAULT_MODE);
