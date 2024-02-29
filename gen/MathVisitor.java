// Generated from //wsl.localhost/Ubuntu-22.04/home/sooden/repos/flux_language/parser/Math.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link Math}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface MathVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link Math#op_level1}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOp_level1(Math.Op_level1Context ctx);
	/**
	 * Visit a parse tree produced by {@link Math#op_level2}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOp_level2(Math.Op_level2Context ctx);
	/**
	 * Visit a parse tree produced by {@link Math#op_level3}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOp_level3(Math.Op_level3Context ctx);
	/**
	 * Visit a parse tree produced by {@link Math#op_level4}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOp_level4(Math.Op_level4Context ctx);
	/**
	 * Visit a parse tree produced by {@link Math#op_level5}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOp_level5(Math.Op_level5Context ctx);
}