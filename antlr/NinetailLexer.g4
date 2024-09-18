lexer grammar NinetailLexer;

INPUT: '@input';
OUTPUT: '@output';
META: '@meta';

TYPE: 'string' | 'number' | 'map';

VAR: [a-zA-Z_][a-zA-Z_0-9]*;

ASSIGN: '=';

STRING: '"' (~["\\] | '\\' ["\\bfnrt])* '"';

WS: [ \t\r\n]+ -> skip;
