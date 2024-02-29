// Code generated from Main.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Main
import "github.com/antlr4-go/antlr/v4"

// BaseMainListener is a complete listener for a parse tree produced by Main.
type BaseMainListener struct{}

var _ MainListener = &BaseMainListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMainListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMainListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMainListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMainListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseMainListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseMainListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseMainListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseMainListener) ExitStatement(ctx *StatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMainListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMainListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseMainListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseMainListener) ExitBlock(ctx *BlockContext) {}

// EnterLoop_statement is called when production loop_statement is entered.
func (s *BaseMainListener) EnterLoop_statement(ctx *Loop_statementContext) {}

// ExitLoop_statement is called when production loop_statement is exited.
func (s *BaseMainListener) ExitLoop_statement(ctx *Loop_statementContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaseMainListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaseMainListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BaseMainListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BaseMainListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterData_type is called when production data_type is entered.
func (s *BaseMainListener) EnterData_type(ctx *Data_typeContext) {}

// ExitData_type is called when production data_type is exited.
func (s *BaseMainListener) ExitData_type(ctx *Data_typeContext) {}

// EnterFunc_declaration is called when production func_declaration is entered.
func (s *BaseMainListener) EnterFunc_declaration(ctx *Func_declarationContext) {}

// ExitFunc_declaration is called when production func_declaration is exited.
func (s *BaseMainListener) ExitFunc_declaration(ctx *Func_declarationContext) {}

// EnterVar_type is called when production var_type is entered.
func (s *BaseMainListener) EnterVar_type(ctx *Var_typeContext) {}

// ExitVar_type is called when production var_type is exited.
func (s *BaseMainListener) ExitVar_type(ctx *Var_typeContext) {}

// EnterVar_name is called when production var_name is entered.
func (s *BaseMainListener) EnterVar_name(ctx *Var_nameContext) {}

// ExitVar_name is called when production var_name is exited.
func (s *BaseMainListener) ExitVar_name(ctx *Var_nameContext) {}

// EnterVar_value is called when production var_value is entered.
func (s *BaseMainListener) EnterVar_value(ctx *Var_valueContext) {}

// ExitVar_value is called when production var_value is exited.
func (s *BaseMainListener) ExitVar_value(ctx *Var_valueContext) {}

// EnterOp_one_expression is called when production op_one_expression is entered.
func (s *BaseMainListener) EnterOp_one_expression(ctx *Op_one_expressionContext) {}

// ExitOp_one_expression is called when production op_one_expression is exited.
func (s *BaseMainListener) ExitOp_one_expression(ctx *Op_one_expressionContext) {}

// EnterOp_one_declaration is called when production op_one_declaration is entered.
func (s *BaseMainListener) EnterOp_one_declaration(ctx *Op_one_declarationContext) {}

// ExitOp_one_declaration is called when production op_one_declaration is exited.
func (s *BaseMainListener) ExitOp_one_declaration(ctx *Op_one_declarationContext) {}

// EnterOp_two_expression is called when production op_two_expression is entered.
func (s *BaseMainListener) EnterOp_two_expression(ctx *Op_two_expressionContext) {}

// ExitOp_two_expression is called when production op_two_expression is exited.
func (s *BaseMainListener) ExitOp_two_expression(ctx *Op_two_expressionContext) {}

// EnterString_var_declaration is called when production string_var_declaration is entered.
func (s *BaseMainListener) EnterString_var_declaration(ctx *String_var_declarationContext) {}

// ExitString_var_declaration is called when production string_var_declaration is exited.
func (s *BaseMainListener) ExitString_var_declaration(ctx *String_var_declarationContext) {}

// EnterNumber_var_declaration is called when production number_var_declaration is entered.
func (s *BaseMainListener) EnterNumber_var_declaration(ctx *Number_var_declarationContext) {}

// ExitNumber_var_declaration is called when production number_var_declaration is exited.
func (s *BaseMainListener) ExitNumber_var_declaration(ctx *Number_var_declarationContext) {}

// EnterBoolean_var_declaration is called when production boolean_var_declaration is entered.
func (s *BaseMainListener) EnterBoolean_var_declaration(ctx *Boolean_var_declarationContext) {}

// ExitBoolean_var_declaration is called when production boolean_var_declaration is exited.
func (s *BaseMainListener) ExitBoolean_var_declaration(ctx *Boolean_var_declarationContext) {}

// EnterSingle_var_declaration is called when production single_var_declaration is entered.
func (s *BaseMainListener) EnterSingle_var_declaration(ctx *Single_var_declarationContext) {}

// ExitSingle_var_declaration is called when production single_var_declaration is exited.
func (s *BaseMainListener) ExitSingle_var_declaration(ctx *Single_var_declarationContext) {}

// EnterArray_var_declaration is called when production array_var_declaration is entered.
func (s *BaseMainListener) EnterArray_var_declaration(ctx *Array_var_declarationContext) {}

// ExitArray_var_declaration is called when production array_var_declaration is exited.
func (s *BaseMainListener) ExitArray_var_declaration(ctx *Array_var_declarationContext) {}

// EnterVar_assignment is called when production var_assignment is entered.
func (s *BaseMainListener) EnterVar_assignment(ctx *Var_assignmentContext) {}

// ExitVar_assignment is called when production var_assignment is exited.
func (s *BaseMainListener) ExitVar_assignment(ctx *Var_assignmentContext) {}

// EnterOp_level1 is called when production op_level1 is entered.
func (s *BaseMainListener) EnterOp_level1(ctx *Op_level1Context) {}

// ExitOp_level1 is called when production op_level1 is exited.
func (s *BaseMainListener) ExitOp_level1(ctx *Op_level1Context) {}

// EnterOp_level2 is called when production op_level2 is entered.
func (s *BaseMainListener) EnterOp_level2(ctx *Op_level2Context) {}

// ExitOp_level2 is called when production op_level2 is exited.
func (s *BaseMainListener) ExitOp_level2(ctx *Op_level2Context) {}

// EnterOp_level3 is called when production op_level3 is entered.
func (s *BaseMainListener) EnterOp_level3(ctx *Op_level3Context) {}

// ExitOp_level3 is called when production op_level3 is exited.
func (s *BaseMainListener) ExitOp_level3(ctx *Op_level3Context) {}

// EnterOp_level4 is called when production op_level4 is entered.
func (s *BaseMainListener) EnterOp_level4(ctx *Op_level4Context) {}

// ExitOp_level4 is called when production op_level4 is exited.
func (s *BaseMainListener) ExitOp_level4(ctx *Op_level4Context) {}

// EnterOp_level5 is called when production op_level5 is entered.
func (s *BaseMainListener) EnterOp_level5(ctx *Op_level5Context) {}

// ExitOp_level5 is called when production op_level5 is exited.
func (s *BaseMainListener) ExitOp_level5(ctx *Op_level5Context) {}

// EnterNumberic_expression is called when production numberic_expression is entered.
func (s *BaseMainListener) EnterNumberic_expression(ctx *Numberic_expressionContext) {}

// ExitNumberic_expression is called when production numberic_expression is exited.
func (s *BaseMainListener) ExitNumberic_expression(ctx *Numberic_expressionContext) {}

// EnterLogical_expression is called when production logical_expression is entered.
func (s *BaseMainListener) EnterLogical_expression(ctx *Logical_expressionContext) {}

// ExitLogical_expression is called when production logical_expression is exited.
func (s *BaseMainListener) ExitLogical_expression(ctx *Logical_expressionContext) {}

// EnterComparative_expression is called when production comparative_expression is entered.
func (s *BaseMainListener) EnterComparative_expression(ctx *Comparative_expressionContext) {}

// ExitComparative_expression is called when production comparative_expression is exited.
func (s *BaseMainListener) ExitComparative_expression(ctx *Comparative_expressionContext) {}

// EnterGet_variable is called when production get_variable is entered.
func (s *BaseMainListener) EnterGet_variable(ctx *Get_variableContext) {}

// ExitGet_variable is called when production get_variable is exited.
func (s *BaseMainListener) ExitGet_variable(ctx *Get_variableContext) {}

// EnterMath_expression is called when production math_expression is entered.
func (s *BaseMainListener) EnterMath_expression(ctx *Math_expressionContext) {}

// ExitMath_expression is called when production math_expression is exited.
func (s *BaseMainListener) ExitMath_expression(ctx *Math_expressionContext) {}

// EnterGet_array_element is called when production get_array_element is entered.
func (s *BaseMainListener) EnterGet_array_element(ctx *Get_array_elementContext) {}

// ExitGet_array_element is called when production get_array_element is exited.
func (s *BaseMainListener) ExitGet_array_element(ctx *Get_array_elementContext) {}

// EnterGet_child is called when production get_child is entered.
func (s *BaseMainListener) EnterGet_child(ctx *Get_childContext) {}

// ExitGet_child is called when production get_child is exited.
func (s *BaseMainListener) ExitGet_child(ctx *Get_childContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseMainListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseMainListener) ExitFunction_call(ctx *Function_callContext) {}
