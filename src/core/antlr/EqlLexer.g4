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
AVG                     : 'AVG';
ABS                     : 'ABS';
ADD                     : 'ADD';

DIGIT                   : [0-9]+;
IDENTIFIER              : [a-zA-Z]+;

WS                      : [ \t\r\n]+ -> skip;

EOS                     : ([\r\n]+ | ';' | '/*' .*? '*/' | EOF) -> mode(DEFAULT_MODE);
