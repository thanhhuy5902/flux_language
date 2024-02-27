lexer grammar Primitives;
//declare a variable type string: example text a {huy}
TEXT: '"' ~["]* '"'
    | '\''.*?'\'';
TEXT_TYPE: 'text';

BOOLEAN: 'true' | 'false';
BOOLEAN_TYPE: 'boolean';




NUMBER: '-'? (
            DIGIT+
            | DIGIT+ '.' DIGIT+
            | DIGIT+ '.' DIGIT+ SI_POSTFIX?
            | DIGIT+ SI_POSTFIX?
            | DIGIT + SI_POSTFIX + DIGIT?
            | DIGIT + SI_POSTFIX + DIGIT + SI_POSTFIX?
            | DIGIT + SI_POSTFIX + DIGIT + SI_POSTFIX + DIGIT+
            );

UNIT: (NUMBER + SI_POSTFIX) | (NUMBER + SI_POSTFIX + NUMBER);

SI_POSTFIX: 'K' | 'M' | 'G' | 'T' | 'P' | 'E' | 'Z' | 'Y' | 'k' | 'm' | 'g' | 't' | 'p' | 'e' | 'z' | 'y';

NUMBER_TYPE: 'num';



DIGIT
    : [0-9]
    ;
OCTET
    : '0'
    | ([1-9] DIGIT?) // 1 to 99
    | '1' DIGIT DIGIT // 100 to 199
    | '2' [0-4] DIGIT // 200 to 249
    | '25' [0-5] // 250 to 255
    ;
IPV4: OCTET '.' OCTET '.' OCTET '.' OCTET;
IPV4_TYPE: 'ipv4';

ARRAY: (TEXT |NUMBER|BOOLEAN|IPV4)*;
ARRAY_TYPE: 'array';

VAR_IDENTIFIER : [a-z][a-zA-Z0-9]*;
COMMON_IDENTIFIER : [a-zA-Z][a-zA-Z0-9]*; //class

// declare array


ARITHMETIC_OPERATOR: '+' | '-' | '*' | '/' | '%'| '^';
LOGICAL_OPERATOR: 'xor'|'and' | 'or' | 'not' | '=' | '!=' | '<' | '>' | '<=' | '>=';
L_BLOCK: '{';
R_BLOCK: '}';
L_PAR: '(';
R_PAR: ')';
L_BRACKET: '[';
R_BRACKET: ']';
DOT: '.';
COLON: ':';
SEMICOLON: ';';
COMMA: ',';
AT: '@';

// operators
OP_PLUS: '+';
OP_MINUS: '-';
OP_MULTIPLY: '*';
OP_DIVIDE: '/';
OP_MOD: '%';
OP_POWER: '^';
OP_EQUAL: '=';
OP_NOT_EQUAL: '!=';
OP_LESS: '<';
OP_GREATER: '>';
OP_LESS_EQUAL: '<=';
OP_GREATER_EQUAL: '>=';
OP_AND: 'and';
OP_OR: 'or';
OP_XOR: 'xor';
OP_NOT: 'not';

WS: [ \t\r\n]+ -> skip;