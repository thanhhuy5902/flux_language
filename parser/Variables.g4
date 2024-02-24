parser grammar Variables;
options {
  tokenVocab = Primitives;
}

var_type
    :TEXT_TYPE
    |NUMBER_TYPE
    |BOOLEAN_TYPE
    |IPV4_TYPE
    ;

var_name
    :VAR_IDENTIFIER
    ;

var_value
    :IPV4
    |TEXT
    |NUMBER
    |BOOLEAN
    ;

default_var_declaration
    :var_type var_name L_BLOCK R_BLOCK

    ;

text_var_declaration

    :TEXT_TYPE var_name L_BLOCK TEXT R_BLOCK
    ;

number_var_declaration
    :NUMBER_TYPE var_name L_BLOCK NUMBER R_BLOCK
    ;

boolean_var_declaration
    :BOOLEAN_TYPE var_name L_BLOCK BOOLEAN R_BLOCK
    ;

ipv4_var_declaration
    :IPV4_TYPE var_name L_BLOCK IPV4 R_BLOCK
    ;

var_declaration
    :default_var_declaration
    |text_var_declaration
    |number_var_declaration
    |boolean_var_declaration
    |ipv4_var_declaration
    ;


