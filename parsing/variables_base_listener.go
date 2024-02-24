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

// EnterDefault_var_declaration is called when production default_var_declaration is entered.
func (s *BaseVariablesListener) EnterDefault_var_declaration(ctx *Default_var_declarationContext) {}

// ExitDefault_var_declaration is called when production default_var_declaration is exited.
func (s *BaseVariablesListener) ExitDefault_var_declaration(ctx *Default_var_declarationContext) {}

// EnterText_var_declaration is called when production text_var_declaration is entered.
func (s *BaseVariablesListener) EnterText_var_declaration(ctx *Text_var_declarationContext) {}

// ExitText_var_declaration is called when production text_var_declaration is exited.
func (s *BaseVariablesListener) ExitText_var_declaration(ctx *Text_var_declarationContext) {}

// EnterNumber_var_declaration is called when production number_var_declaration is entered.
func (s *BaseVariablesListener) EnterNumber_var_declaration(ctx *Number_var_declarationContext) {}

// ExitNumber_var_declaration is called when production number_var_declaration is exited.
func (s *BaseVariablesListener) ExitNumber_var_declaration(ctx *Number_var_declarationContext) {}

// EnterBoolean_var_declaration is called when production boolean_var_declaration is entered.
func (s *BaseVariablesListener) EnterBoolean_var_declaration(ctx *Boolean_var_declarationContext) {}

// ExitBoolean_var_declaration is called when production boolean_var_declaration is exited.
func (s *BaseVariablesListener) ExitBoolean_var_declaration(ctx *Boolean_var_declarationContext) {}

// EnterIpv4_var_declaration is called when production ipv4_var_declaration is entered.
func (s *BaseVariablesListener) EnterIpv4_var_declaration(ctx *Ipv4_var_declarationContext) {}

// ExitIpv4_var_declaration is called when production ipv4_var_declaration is exited.
func (s *BaseVariablesListener) ExitIpv4_var_declaration(ctx *Ipv4_var_declarationContext) {}

// EnterVar_declaration is called when production var_declaration is entered.
func (s *BaseVariablesListener) EnterVar_declaration(ctx *Var_declarationContext) {}

// ExitVar_declaration is called when production var_declaration is exited.
func (s *BaseVariablesListener) ExitVar_declaration(ctx *Var_declarationContext) {}
