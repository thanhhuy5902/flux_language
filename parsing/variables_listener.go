// Code generated from Variables.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Variables
import "github.com/antlr4-go/antlr/v4"

// VariablesListener is a complete listener for a parse tree produced by Variables.
type VariablesListener interface {
	antlr.ParseTreeListener

	// EnterVar_type is called when entering the var_type production.
	EnterVar_type(c *Var_typeContext)

	// EnterVar_name is called when entering the var_name production.
	EnterVar_name(c *Var_nameContext)

	// EnterVar_value is called when entering the var_value production.
	EnterVar_value(c *Var_valueContext)

	// EnterOp_one_expression is called when entering the op_one_expression production.
	EnterOp_one_expression(c *Op_one_expressionContext)

	// EnterOp_one_declaration is called when entering the op_one_declaration production.
	EnterOp_one_declaration(c *Op_one_declarationContext)

	// EnterOp_two_expression is called when entering the op_two_expression production.
	EnterOp_two_expression(c *Op_two_expressionContext)

	// EnterString_var_declaration is called when entering the string_var_declaration production.
	EnterString_var_declaration(c *String_var_declarationContext)

	// EnterNumber_var_declaration is called when entering the number_var_declaration production.
	EnterNumber_var_declaration(c *Number_var_declarationContext)

	// EnterBoolean_var_declaration is called when entering the boolean_var_declaration production.
	EnterBoolean_var_declaration(c *Boolean_var_declarationContext)

	// EnterSingle_var_declaration is called when entering the single_var_declaration production.
	EnterSingle_var_declaration(c *Single_var_declarationContext)

	// EnterArray_var_declaration is called when entering the array_var_declaration production.
	EnterArray_var_declaration(c *Array_var_declarationContext)

	// EnterVar_assignment is called when entering the var_assignment production.
	EnterVar_assignment(c *Var_assignmentContext)

	// EnterOp_level1 is called when entering the op_level1 production.
	EnterOp_level1(c *Op_level1Context)

	// EnterOp_level2 is called when entering the op_level2 production.
	EnterOp_level2(c *Op_level2Context)

	// EnterOp_level3 is called when entering the op_level3 production.
	EnterOp_level3(c *Op_level3Context)

	// EnterOp_level4 is called when entering the op_level4 production.
	EnterOp_level4(c *Op_level4Context)

	// EnterOp_level5 is called when entering the op_level5 production.
	EnterOp_level5(c *Op_level5Context)

	// EnterNumberic_expression is called when entering the numberic_expression production.
	EnterNumberic_expression(c *Numberic_expressionContext)

	// EnterLogical_expression is called when entering the logical_expression production.
	EnterLogical_expression(c *Logical_expressionContext)

	// EnterComparative_expression is called when entering the comparative_expression production.
	EnterComparative_expression(c *Comparative_expressionContext)

	// EnterGet_variable is called when entering the get_variable production.
	EnterGet_variable(c *Get_variableContext)

	// EnterMath_expression is called when entering the math_expression production.
	EnterMath_expression(c *Math_expressionContext)

	// EnterGet_array_element is called when entering the get_array_element production.
	EnterGet_array_element(c *Get_array_elementContext)

	// EnterGet_child is called when entering the get_child production.
	EnterGet_child(c *Get_childContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// ExitVar_type is called when exiting the var_type production.
	ExitVar_type(c *Var_typeContext)

	// ExitVar_name is called when exiting the var_name production.
	ExitVar_name(c *Var_nameContext)

	// ExitVar_value is called when exiting the var_value production.
	ExitVar_value(c *Var_valueContext)

	// ExitOp_one_expression is called when exiting the op_one_expression production.
	ExitOp_one_expression(c *Op_one_expressionContext)

	// ExitOp_one_declaration is called when exiting the op_one_declaration production.
	ExitOp_one_declaration(c *Op_one_declarationContext)

	// ExitOp_two_expression is called when exiting the op_two_expression production.
	ExitOp_two_expression(c *Op_two_expressionContext)

	// ExitString_var_declaration is called when exiting the string_var_declaration production.
	ExitString_var_declaration(c *String_var_declarationContext)

	// ExitNumber_var_declaration is called when exiting the number_var_declaration production.
	ExitNumber_var_declaration(c *Number_var_declarationContext)

	// ExitBoolean_var_declaration is called when exiting the boolean_var_declaration production.
	ExitBoolean_var_declaration(c *Boolean_var_declarationContext)

	// ExitSingle_var_declaration is called when exiting the single_var_declaration production.
	ExitSingle_var_declaration(c *Single_var_declarationContext)

	// ExitArray_var_declaration is called when exiting the array_var_declaration production.
	ExitArray_var_declaration(c *Array_var_declarationContext)

	// ExitVar_assignment is called when exiting the var_assignment production.
	ExitVar_assignment(c *Var_assignmentContext)

	// ExitOp_level1 is called when exiting the op_level1 production.
	ExitOp_level1(c *Op_level1Context)

	// ExitOp_level2 is called when exiting the op_level2 production.
	ExitOp_level2(c *Op_level2Context)

	// ExitOp_level3 is called when exiting the op_level3 production.
	ExitOp_level3(c *Op_level3Context)

	// ExitOp_level4 is called when exiting the op_level4 production.
	ExitOp_level4(c *Op_level4Context)

	// ExitOp_level5 is called when exiting the op_level5 production.
	ExitOp_level5(c *Op_level5Context)

	// ExitNumberic_expression is called when exiting the numberic_expression production.
	ExitNumberic_expression(c *Numberic_expressionContext)

	// ExitLogical_expression is called when exiting the logical_expression production.
	ExitLogical_expression(c *Logical_expressionContext)

	// ExitComparative_expression is called when exiting the comparative_expression production.
	ExitComparative_expression(c *Comparative_expressionContext)

	// ExitGet_variable is called when exiting the get_variable production.
	ExitGet_variable(c *Get_variableContext)

	// ExitMath_expression is called when exiting the math_expression production.
	ExitMath_expression(c *Math_expressionContext)

	// ExitGet_array_element is called when exiting the get_array_element production.
	ExitGet_array_element(c *Get_array_elementContext)

	// ExitGet_child is called when exiting the get_child production.
	ExitGet_child(c *Get_childContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)
}