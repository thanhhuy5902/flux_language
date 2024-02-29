// Code generated from Math.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Math
import "github.com/antlr4-go/antlr/v4"

// BaseMathListener is a complete listener for a parse tree produced by Math.
type BaseMathListener struct{}

var _ MathListener = &BaseMathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterOp_level1 is called when production op_level1 is entered.
func (s *BaseMathListener) EnterOp_level1(ctx *Op_level1Context) {}

// ExitOp_level1 is called when production op_level1 is exited.
func (s *BaseMathListener) ExitOp_level1(ctx *Op_level1Context) {}

// EnterOp_level2 is called when production op_level2 is entered.
func (s *BaseMathListener) EnterOp_level2(ctx *Op_level2Context) {}

// ExitOp_level2 is called when production op_level2 is exited.
func (s *BaseMathListener) ExitOp_level2(ctx *Op_level2Context) {}

// EnterOp_level3 is called when production op_level3 is entered.
func (s *BaseMathListener) EnterOp_level3(ctx *Op_level3Context) {}

// ExitOp_level3 is called when production op_level3 is exited.
func (s *BaseMathListener) ExitOp_level3(ctx *Op_level3Context) {}

// EnterOp_level4 is called when production op_level4 is entered.
func (s *BaseMathListener) EnterOp_level4(ctx *Op_level4Context) {}

// ExitOp_level4 is called when production op_level4 is exited.
func (s *BaseMathListener) ExitOp_level4(ctx *Op_level4Context) {}

// EnterOp_level5 is called when production op_level5 is entered.
func (s *BaseMathListener) EnterOp_level5(ctx *Op_level5Context) {}

// ExitOp_level5 is called when production op_level5 is exited.
func (s *BaseMathListener) ExitOp_level5(ctx *Op_level5Context) {}

// EnterNumberic_expression is called when production numberic_expression is entered.
func (s *BaseMathListener) EnterNumberic_expression(ctx *Numberic_expressionContext) {}

// ExitNumberic_expression is called when production numberic_expression is exited.
func (s *BaseMathListener) ExitNumberic_expression(ctx *Numberic_expressionContext) {}

// EnterLogical_expression is called when production logical_expression is entered.
func (s *BaseMathListener) EnterLogical_expression(ctx *Logical_expressionContext) {}

// ExitLogical_expression is called when production logical_expression is exited.
func (s *BaseMathListener) ExitLogical_expression(ctx *Logical_expressionContext) {}

// EnterComparative_expression is called when production comparative_expression is entered.
func (s *BaseMathListener) EnterComparative_expression(ctx *Comparative_expressionContext) {}

// ExitComparative_expression is called when production comparative_expression is exited.
func (s *BaseMathListener) ExitComparative_expression(ctx *Comparative_expressionContext) {}

// EnterGet_variable is called when production get_variable is entered.
func (s *BaseMathListener) EnterGet_variable(ctx *Get_variableContext) {}

// ExitGet_variable is called when production get_variable is exited.
func (s *BaseMathListener) ExitGet_variable(ctx *Get_variableContext) {}

// EnterMath_expression is called when production math_expression is entered.
func (s *BaseMathListener) EnterMath_expression(ctx *Math_expressionContext) {}

// ExitMath_expression is called when production math_expression is exited.
func (s *BaseMathListener) ExitMath_expression(ctx *Math_expressionContext) {}

// EnterGet_array_element is called when production get_array_element is entered.
func (s *BaseMathListener) EnterGet_array_element(ctx *Get_array_elementContext) {}

// ExitGet_array_element is called when production get_array_element is exited.
func (s *BaseMathListener) ExitGet_array_element(ctx *Get_array_elementContext) {}

// EnterGet_child is called when production get_child is entered.
func (s *BaseMathListener) EnterGet_child(ctx *Get_childContext) {}

// ExitGet_child is called when production get_child is exited.
func (s *BaseMathListener) ExitGet_child(ctx *Get_childContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseMathListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseMathListener) ExitFunction_call(ctx *Function_callContext) {}
