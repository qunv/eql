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

DIGIT                   : [0-9]+;
IDENTIFIER              : [a-zA-Z]+;

WS                      : [ \t\r\n]+ -> skip;

EOS                     : ([\r\n]+ | ';' | '/*' .*? '*/' | EOF) -> mode(DEFAULT_MODE);
