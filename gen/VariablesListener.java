// Generated from //wsl.localhost/Ubuntu-22.04/home/justintran/repos/flux_language/parser/Variables.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link Variables}.
 */
public interface VariablesListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link Variables#var_type}.
	 * @param ctx the parse tree
	 */
	void enterVar_type(Variables.Var_typeContext ctx);
	/**
	 * Exit a parse tree produced by {@link Variables#var_type}.
	 * @param ctx the parse tree
	 */
	void exitVar_type(Variables.Var_typeContext ctx);
	/**
	 * Enter a parse tree produced by {@link Variables#var_name}.
	 * @param ctx the parse tree
	 */
	void enterVar_name(Variables.Var_nameContext ctx);
	/**
	 * Exit a parse tree produced by {@link Variables#var_name}.
	 * @param ctx the parse tree
	 */
	void exitVar_name(Variables.Var_nameContext ctx);
	/**
	 * Enter a parse tree produced by {@link Variables#var_value}.
	 * @param ctx the parse tree
	 */
	void enterVar_value(Variables.Var_valueContext ctx);
	/**
	 * Exit a parse tree produced by {@link Variables#var_value}.
	 * @param ctx the parse tree
	 */
	void exitVar_value(Variables.Var_valueContext ctx);
	/**
	 * Enter a parse tree produced by {@link Variables#default_var_declaration}.
	 * @param ctx the parse tree
	 */
	void enterDefault_var_declaration(Variables.Default_var_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link Variables#default_var_declaration}.
	 * @param ctx the parse tree
	 */
	void exitDefault_var_declaration(Variables.Default_var_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link Variables#text_var_declaration}.
	 * @param ctx the parse tree
	 */
	void enterText_var_declaration(Variables.Text_var_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link Variables#text_var_declaration}.
	 * @param ctx the parse tree
	 */
	void exitText_var_declaration(Variables.Text_var_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link Variables#number_var_declaration}.
	 * @param ctx the parse tree
	 */
	void enterNumber_var_declaration(Variables.Number_var_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link Variables#number_var_declaration}.
	 * @param ctx the parse tree
	 */
	void exitNumber_var_declaration(Variables.Number_var_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link Variables#boolean_var_declaration}.
	 * @param ctx the parse tree
	 */
	void enterBoolean_var_declaration(Variables.Boolean_var_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link Variables#boolean_var_declaration}.
	 * @param ctx the parse tree
	 */
	void exitBoolean_var_declaration(Variables.Boolean_var_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link Variables#ipv4_var_declaration}.
	 * @param ctx the parse tree
	 */
	void enterIpv4_var_declaration(Variables.Ipv4_var_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link Variables#ipv4_var_declaration}.
	 * @param ctx the parse tree
	 */
	void exitIpv4_var_declaration(Variables.Ipv4_var_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link Variables#var_declaration}.
	 * @param ctx the parse tree
	 */
	void enterVar_declaration(Variables.Var_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link Variables#var_declaration}.
	 * @param ctx the parse tree
	 */
	void exitVar_declaration(Variables.Var_declarationContext ctx);
}