// Code generated from Variables.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Variables
import "github.com/antlr4-go/antlr/v4"

// BaseVariablesListener is a complete listener for a parse tree produced by Variables.
type BaseVariablesListener struct{}

var _ VariablesListener = &BaseVariablesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVariablesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVariablesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVariablesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVariablesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterVar_type is called when production var_type is entered.
func (s *BaseVariablesListener) EnterVar_type(ctx *Var_typeContext) {}

// ExitVar_type is called when production var_type is exited.
func (s *BaseVariablesListener) ExitVar_type(ctx *Var_typeContext) {}

// EnterVar_name is called when production var_name is entered.
func (s *BaseVariablesListener) EnterVar_name(ctx *Var_nameContext) {}

// ExitVar_name is called when production var_name is exited.
func (s *BaseVariablesListener) ExitVar_name(ctx *Var_nameContext) {}

// EnterVar_value is called when production var_value is entered.
func (s *BaseVariablesListener) EnterVar_value(ctx *Var_valueContext) {}

// ExitVar_value is called when production var_value is exited.
func (s *BaseVariablesListener) ExitVar_value(ctx *Var_valueContext) {}

// EnterOp_one_expression is called when production op_one_expression is entered.
func (s *BaseVariablesListener) EnterOp_one_expression(ctx *Op_one_expressionContext) {}

// ExitOp_one_expression is called when production op_one_expression is exited.
func (s *BaseVariablesListener) ExitOp_one_expression(ctx *Op_one_expressionContext) {}

// EnterOp_one_declaration is called when production op_one_declaration is entered.
func (s *BaseVariablesListener) EnterOp_one_declaration(ctx *Op_one_declarationContext) {}

// ExitOp_one_declaration is called when production op_one_declaration is exited.
func (s *BaseVariablesListener) ExitOp_one_declaration(ctx *Op_one_declarationContext) {}

// EnterOp_two_expression is called when production op_two_expression is entered.
func (s *BaseVariablesListener) EnterOp_two_expression(ctx *Op_two_expressionContext) {}

// ExitOp_two_expression is called when production op_two_expression is exited.
func (s *BaseVariablesListener) ExitOp_two_expression(ctx *Op_two_expressionContext) {}

// EnterString_var_declaration is called when production string_var_declaration is entered.
func (s *BaseVariablesListener) EnterString_var_declaration(ctx *String_var_declarationContext) {}

// ExitString_var_declaration is called when production string_var_declaration is exited.
func (s *BaseVariablesListener) ExitString_var_declaration(ctx *String_var_declarationContext) {}

// EnterNumber_var_declaration is called when production number_var_declaration is entered.
func (s *BaseVariablesListener) EnterNumber_var_declaration(ctx *Number_var_declarationContext) {}

// ExitNumber_var_declaration is called when production number_var_declaration is exited.
func (s *BaseVariablesListener) ExitNumber_var_declaration(ctx *Number_var_declarationContext) {}

// EnterBoolean_var_declaration is called when production boolean_var_declaration is entered.
func (s *BaseVariablesListener) EnterBoolean_var_declaration(ctx *Boolean_var_declarationContext) {}

// ExitBoolean_var_declaration is called when production boolean_var_declaration is exited.
func (s *BaseVariablesListener) ExitBoolean_var_declaration(ctx *Boolean_var_declarationContext) {}

// EnterSingle_var_declaration is called when production single_var_declaration is entered.
func (s *BaseVariablesListener) EnterSingle_var_declaration(ctx *Single_var_declarationContext) {}

// ExitSingle_var_declaration is called when production single_var_declaration is exited.
func (s *BaseVariablesListener) ExitSingle_var_declaration(ctx *Single_var_declarationContext) {}

// EnterArray_var_declaration is called when production array_var_declaration is entered.
func (s *BaseVariablesListener) EnterArray_var_declaration(ctx *Array_var_declarationContext) {}

// ExitArray_var_declaration is called when production array_var_declaration is exited.
func (s *BaseVariablesListener) ExitArray_var_declaration(ctx *Array_var_declarationContext) {}

// EnterVar_assignment is called when production var_assignment is entered.
func (s *BaseVariablesListener) EnterVar_assignment(ctx *Var_assignmentContext) {}

// ExitVar_assignment is called when production var_assignment is exited.
func (s *BaseVariablesListener) ExitVar_assignment(ctx *Var_assignmentContext) {}

// EnterOp_level1 is called when production op_level1 is entered.
func (s *BaseVariablesListener) EnterOp_level1(ctx *Op_level1Context) {}

// ExitOp_level1 is called when production op_level1 is exited.
func (s *BaseVariablesListener) ExitOp_level1(ctx *Op_level1Context) {}

// EnterOp_level2 is called when production op_level2 is entered.
func (s *BaseVariablesListener) EnterOp_level2(ctx *Op_level2Context) {}

// ExitOp_level2 is called when production op_level2 is exited.
func (s *BaseVariablesListener) ExitOp_level2(ctx *Op_level2Context) {}

// EnterOp_level3 is called when production op_level3 is entered.
func (s *BaseVariablesListener) EnterOp_level3(ctx *Op_level3Context) {}

// ExitOp_level3 is called when production op_level3 is exited.
func (s *BaseVariablesListener) ExitOp_level3(ctx *Op_level3Context) {}

// EnterOp_level4 is called when production op_level4 is entered.
func (s *BaseVariablesListener) EnterOp_level4(ctx *Op_level4Context) {}

// ExitOp_level4 is called when production op_level4 is exited.
func (s *BaseVariablesListener) ExitOp_level4(ctx *Op_level4Context) {}

// EnterOp_level5 is called when production op_level5 is entered.
func (s *BaseVariablesListener) EnterOp_level5(ctx *Op_level5Context) {}

// ExitOp_level5 is called when production op_level5 is exited.
func (s *BaseVariablesListener) ExitOp_level5(ctx *Op_level5Context) {}

// EnterNumberic_expression is called when production numberic_expression is entered.
func (s *BaseVariablesListener) EnterNumberic_expression(ctx *Numberic_expressionContext) {}

// ExitNumberic_expression is called when production numberic_expression is exited.
func (s *BaseVariablesListener) ExitNumberic_expression(ctx *Numberic_expressionContext) {}

// EnterLogical_expression is called when production logical_expression is entered.
func (s *BaseVariablesListener) EnterLogical_expression(ctx *Logical_expressionContext) {}

// ExitLogical_expression is called when production logical_expression is exited.
func (s *BaseVariablesListener) ExitLogical_expression(ctx *Logical_expressionContext) {}

// EnterComparative_expression is called when production comparative_expression is entered.
func (s *BaseVariablesListener) EnterComparative_expression(ctx *Comparative_expressionContext) {}

// ExitComparative_expression is called when production comparative_expression is exited.
func (s *BaseVariablesListener) ExitComparative_expression(ctx *Comparative_expressionContext) {}

// EnterGet_variable is called when production get_variable is entered.
func (s *BaseVariablesListener) EnterGet_variable(ctx *Get_variableContext) {}

// ExitGet_variable is called when production get_variable is exited.
func (s *BaseVariablesListener) ExitGet_variable(ctx *Get_variableContext) {}

// EnterMath_expression is called when production math_expression is entered.
func (s *BaseVariablesListener) EnterMath_expression(ctx *Math_expressionContext) {}

// ExitMath_expression is called when production math_expression is exited.
func (s *BaseVariablesListener) ExitMath_expression(ctx *Math_expressionContext) {}

// EnterGet_array_element is called when production get_array_element is entered.
func (s *BaseVariablesListener) EnterGet_array_element(ctx *Get_array_elementContext) {}

// ExitGet_array_element is called when production get_array_element is exited.
func (s *BaseVariablesListener) ExitGet_array_element(ctx *Get_array_elementContext) {}

// EnterGet_child is called when production get_child is entered.
func (s *BaseVariablesListener) EnterGet_child(ctx *Get_childContext) {}

// ExitGet_child is called when production get_child is exited.
func (s *BaseVariablesListener) ExitGet_child(ctx *Get_childContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseVariablesListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseVariablesListener) ExitFunction_call(ctx *Function_callContext) {}
