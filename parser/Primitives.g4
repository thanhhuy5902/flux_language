lexer grammar Primitives;

TEXT: '"' ~["]* '"'
    | '\''.*?'\'';

TEXT_TYPE: 'text';

BOOLEAN: 'true' | 'false';
BOOLEAN_TYPE: 'boolean';

NUMBER: '-'? [0-9]+('.'[0-9]*)?;
NUMBER_TYPE: 'num';

DIGIT: [0-9];
OCTET
    : '0'
    | ([1-9] DIGIT?) // 1 to 99
    | '1' DIGIT DIGIT // 100 to 199
    | '2' [0-4] DIGIT // 200 to 249
    | '25' [0-5] // 250 to 255
    ;
IPV4: OCTET '.' OCTET '.' OCTET '.' OCTET;
IPV4_TYPE: 'ipv4';

VAR_IDENTIFIER : [a-z][a-zA-Z0-9]*;
COMMON_IDENTIFIER : [a-zA-Z][a-zA-Z0-9]*; //class

ARITHMETIC_OPERATOR: '+' | '-' | '*' | '/' | '%'| '^';
LOGICAL_OPERATOR: 'xor'|'and' | 'or' | 'not' | '=' | '!=' | '<' | '>' | '<=' | '>=';
L_BLOCK: '{';
R_BLOCK: '}';
AT: '@';
WS: [ \t\r\n]+ -> skip;