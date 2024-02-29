<<<<<<< HEAD
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
=======
// Generated from /home/sooden/repos/flux_language/parser/Variables.g4 by ANTLR 4.13.1
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
	 * Visit a parse tree produced by {@link Variables#varType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitVarType(Variables.VarTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link Variables#varName}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitVarName(Variables.VarNameContext ctx);
	/**
	 * Visit a parse tree produced by {@link Variables#varValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitVarValue(Variables.VarValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link Variables#defaultVarDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDefaultVarDeclaration(Variables.DefaultVarDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link Variables#textVarDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTextVarDeclaration(Variables.TextVarDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link Variables#numberVarDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNumberVarDeclaration(Variables.NumberVarDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link Variables#booleanVarDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBooleanVarDeclaration(Variables.BooleanVarDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link Variables#ipv4VarDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIpv4VarDeclaration(Variables.Ipv4VarDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link Variables#varDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitVarDeclaration(Variables.VarDeclarationContext ctx);
>>>>>>> 4aa8e30 (Update 27/2/2024)
}