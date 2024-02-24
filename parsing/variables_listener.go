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

	// EnterDefault_var_declaration is called when entering the default_var_declaration production.
	EnterDefault_var_declaration(c *Default_var_declarationContext)

	// EnterText_var_declaration is called when entering the text_var_declaration production.
	EnterText_var_declaration(c *Text_var_declarationContext)

	// EnterNumber_var_declaration is called when entering the number_var_declaration production.
	EnterNumber_var_declaration(c *Number_var_declarationContext)

	// EnterBoolean_var_declaration is called when entering the boolean_var_declaration production.
	EnterBoolean_var_declaration(c *Boolean_var_declarationContext)

	// EnterIpv4_var_declaration is called when entering the ipv4_var_declaration production.
	EnterIpv4_var_declaration(c *Ipv4_var_declarationContext)

	// EnterVar_declaration is called when entering the var_declaration production.
	EnterVar_declaration(c *Var_declarationContext)

	// ExitVar_type is called when exiting the var_type production.
	ExitVar_type(c *Var_typeContext)

	// ExitVar_name is called when exiting the var_name production.
	ExitVar_name(c *Var_nameContext)

	// ExitVar_value is called when exiting the var_value production.
	ExitVar_value(c *Var_valueContext)

	// ExitDefault_var_declaration is called when exiting the default_var_declaration production.
	ExitDefault_var_declaration(c *Default_var_declarationContext)

	// ExitText_var_declaration is called when exiting the text_var_declaration production.
	ExitText_var_declaration(c *Text_var_declarationContext)

	// ExitNumber_var_declaration is called when exiting the number_var_declaration production.
	ExitNumber_var_declaration(c *Number_var_declarationContext)

	// ExitBoolean_var_declaration is called when exiting the boolean_var_declaration production.
	ExitBoolean_var_declaration(c *Boolean_var_declarationContext)

	// ExitIpv4_var_declaration is called when exiting the ipv4_var_declaration production.
	ExitIpv4_var_declaration(c *Ipv4_var_declarationContext)

	// ExitVar_declaration is called when exiting the var_declaration production.
	ExitVar_declaration(c *Var_declarationContext)
}
