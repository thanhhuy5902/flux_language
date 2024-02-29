// Code generated from Math.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Math
import "github.com/antlr4-go/antlr/v4"

// MathListener is a complete listener for a parse tree produced by Math.
type MathListener interface {
	antlr.ParseTreeListener

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
