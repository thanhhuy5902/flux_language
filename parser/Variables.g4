parser grammar Variables;
import Math;

options {
    tokenVocab = Primitives;
}


var_type
    : TEXT_TYPE
    | NUMBER_TYPE
    | BOOLEAN_TYPE
    | IPV4_TYPE
    ;


var_name
    : VAR_IDENTIFIER
    ;

var_value
    : IPV4
    | TEXT
    | NUMBER
    | BOOLEAN
    ;
op_one_expression
  : OP_INCREMENT
  | OP_DECREMENT
  |OP_ADD_ASSIGN
  |OP_SUB_ASSIGN
  |OP_MUL_ASSIGN
  |OP_DIV_ASSIGN
  ;

op_one_declaration
  : var_name OP_INCREMENT
  | var_name OP_DECREMENT
  | NUMBER_TYPE var_name OP_INCREMENT
  | NUMBER_TYPE var_name OP_DECREMENT
  | NUMBER_TYPE var_name OP_EQUAL NUMBER
  | NUMBER_TYPE var_name OP_ADD_ASSIGN NUMBER
  | NUMBER_TYPE var_name OP_SUB_ASSIGN NUMBER
  | NUMBER_TYPE var_name OP_MUL_ASSIGN NUMBER
  | NUMBER_TYPE var_name OP_DIV_ASSIGN NUMBE
  ;

op_two_expression
  : OP_ADD
  | OP_SUB
  | OP_MUL
  | OP_DIV
  | OP_MOD
  | OP_POW
  ;

//op_two_declaration
//
//    ;

string_var_declaration
    : TEXT_TYPE var_name L_BLOCK NEWLINE* TEXT NEWLINE* R_BLOCK
    ;


number_var_declaration
    : NUMBER_TYPE var_name L_BLOCK NEWLINE* NUMBER NEWLINE* R_BLOCK
    | NUMBER_TYPE var_name L_BLOCK NEWLINE* math_expression NEWLINE* R_BLOCK
    ;

boolean_var_declaration
    : BOOLEAN_TYPE var_name L_BLOCK NEWLINE* BOOLEAN NEWLINE* R_BLOCK
    | BOOLEAN_TYPE var_name L_BLOCK NEWLINE* logical_expression NEWLINE* R_BLOCK
    | BOOLEAN_TYPE var_name L_BLOCK NEWLINE* comparative_expression NEWLINE* R_BLOCK
    ;

single_var_declaration
    : string_var_declaration
    | number_var_declaration
    | boolean_var_declaration
    |var_type var_name L_BLOCK NEWLINE* R_BLOCK
    ;

array_var_declaration
    : var_type var_name L_BLOCK NEWLINE* var_value (COMMA NEWLINE* var_value)* NEWLINE* R_BLOCK
    | var_type var_name L_BLOCK NEWLINE* math_expression (COMMA NEWLINE* math_expression)* NEWLINE R_BLOCK
    ;

var_assignment
    : var_name L_BLOCK NEWLINE* var_value NEWLINE* R_BLOCK
    | var_name L_BLOCK NEWLINE* math_expression NEWLINE* R_BLOCK
    ;