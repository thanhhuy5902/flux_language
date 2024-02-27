lexer grammar Primitives;

TEXT: '"' ~["]* '"'
    | '\''.*?'\'';

TEXT_TYPE: 'text';

BOOLEAN: 'true' | 'false';
BOOLEAN_TYPE: 'boolean';


//I want to add SI postfixes in number values. for example : 2k1 = 2001, 2M2K = 20002000, 4K5.12 = 4500.12
//the number can be negative or positive, decimal or integer
DIGIT
    : [0-9]
    ;

NUMBER: '-'? (DIGIT+
            | DIGIT+ '.' DIGIT+
            | DIGIT+ '.' DIGIT+ SI_POSTFIX?
            | DIGIT+ SI_POSTFIX?
            | DIGIT + SI_POSTFIX + DIGIT?
            | DIGIT + SI_POSTFIX + DIGIT + SI_POSTFIX?
            | DIGIT + SI_POSTFIX + DIGIT + SI_POSTFIX + DIGIT?
            );

UNIT: (NUMBER + SI_POSTFIX) | (NUMBER + SI_POSTFIX + NUMBER);

SI_POSTFIX: 'K' | 'M' | 'G' | 'T' | 'P' | 'E' | 'Z' | 'Y' | 'k' | 'm' | 'g' | 't' | 'p' | 'e' | 'z' | 'y';

NUMBER_TYPE: 'num';




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