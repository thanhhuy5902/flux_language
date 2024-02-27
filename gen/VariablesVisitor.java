// Generated from //wsl.localhost/Ubuntu-22.04/home/justintran/repos/flux_language/parser/Variables.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link Variables}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface VariablesVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link Variables#var_type}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitVar_type(Variables.Var_typeContext ctx);
	/**
	 * Visit a parse tree produced by {@link Variables#var_name}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitVar_name(Variables.Var_nameContext ctx);
	/**
	 * Visit a parse tree produced by {@link Variables#var_value}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitVar_value(Variables.Var_valueContext ctx);
	/**
	 * Visit a parse tree produced by {@link Variables#default_var_declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDefault_var_declaration(Variables.Default_var_declarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link Variables#text_var_declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitText_var_declaration(Variables.Text_var_declarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link Variables#number_var_declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNumber_var_declaration(Variables.Number_var_declarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link Variables#boolean_var_declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBoolean_var_declaration(Variables.Boolean_var_declarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link Variables#ipv4_var_declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIpv4_var_declaration(Variables.Ipv4_var_declarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link Variables#var_declaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitVar_declaration(Variables.Var_declarationContext ctx);
}