// Generated from c://Users//pbara//OneDrive//Escritorio//T-Swift-Compiler//server//SwiftGrammar.g4 by ANTLR 4.13.1

    import "Server2/interfaces"
    import "Server2/environment"
    import "Server2/expressions"
    import "Server2/instructions"
    import "strings"

import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SwiftGrammarParser}.
 */
public interface SwiftGrammarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#s}.
	 * @param ctx the parse tree
	 */
	void enterS(SwiftGrammarParser.SContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#s}.
	 * @param ctx the parse tree
	 */
	void exitS(SwiftGrammarParser.SContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(SwiftGrammarParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(SwiftGrammarParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#instruction}.
	 * @param ctx the parse tree
	 */
	void enterInstruction(SwiftGrammarParser.InstructionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#instruction}.
	 * @param ctx the parse tree
	 */
	void exitInstruction(SwiftGrammarParser.InstructionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#printstmt}.
	 * @param ctx the parse tree
	 */
	void enterPrintstmt(SwiftGrammarParser.PrintstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#printstmt}.
	 * @param ctx the parse tree
	 */
	void exitPrintstmt(SwiftGrammarParser.PrintstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#ifstmt}.
	 * @param ctx the parse tree
	 */
	void enterIfstmt(SwiftGrammarParser.IfstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#ifstmt}.
	 * @param ctx the parse tree
	 */
	void exitIfstmt(SwiftGrammarParser.IfstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#declarationstmt}.
	 * @param ctx the parse tree
	 */
	void enterDeclarationstmt(SwiftGrammarParser.DeclarationstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#declarationstmt}.
	 * @param ctx the parse tree
	 */
	void exitDeclarationstmt(SwiftGrammarParser.DeclarationstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#assignstmt}.
	 * @param ctx the parse tree
	 */
	void enterAssignstmt(SwiftGrammarParser.AssignstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#assignstmt}.
	 * @param ctx the parse tree
	 */
	void exitAssignstmt(SwiftGrammarParser.AssignstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#fnstmt}.
	 * @param ctx the parse tree
	 */
	void enterFnstmt(SwiftGrammarParser.FnstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#fnstmt}.
	 * @param ctx the parse tree
	 */
	void exitFnstmt(SwiftGrammarParser.FnstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listParamsFunc}.
	 * @param ctx the parse tree
	 */
	void enterListParamsFunc(SwiftGrammarParser.ListParamsFuncContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listParamsFunc}.
	 * @param ctx the parse tree
	 */
	void exitListParamsFunc(SwiftGrammarParser.ListParamsFuncContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#parametro}.
	 * @param ctx the parse tree
	 */
	void enterParametro(SwiftGrammarParser.ParametroContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#parametro}.
	 * @param ctx the parse tree
	 */
	void exitParametro(SwiftGrammarParser.ParametroContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#callExp}.
	 * @param ctx the parse tree
	 */
	void enterCallExp(SwiftGrammarParser.CallExpContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#callExp}.
	 * @param ctx the parse tree
	 */
	void exitCallExp(SwiftGrammarParser.CallExpContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#callFunction}.
	 * @param ctx the parse tree
	 */
	void enterCallFunction(SwiftGrammarParser.CallFunctionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#callFunction}.
	 * @param ctx the parse tree
	 */
	void exitCallFunction(SwiftGrammarParser.CallFunctionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listParamsCall}.
	 * @param ctx the parse tree
	 */
	void enterListParamsCall(SwiftGrammarParser.ListParamsCallContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listParamsCall}.
	 * @param ctx the parse tree
	 */
	void exitListParamsCall(SwiftGrammarParser.ListParamsCallContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#types}.
	 * @param ctx the parse tree
	 */
	void enterTypes(SwiftGrammarParser.TypesContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#types}.
	 * @param ctx the parse tree
	 */
	void exitTypes(SwiftGrammarParser.TypesContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr(SwiftGrammarParser.ExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr(SwiftGrammarParser.ExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listParams}.
	 * @param ctx the parse tree
	 */
	void enterListParams(SwiftGrammarParser.ListParamsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listParams}.
	 * @param ctx the parse tree
	 */
	void exitListParams(SwiftGrammarParser.ListParamsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listAccessArray}.
	 * @param ctx the parse tree
	 */
	void enterListAccessArray(SwiftGrammarParser.ListAccessArrayContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listAccessArray}.
	 * @param ctx the parse tree
	 */
	void exitListAccessArray(SwiftGrammarParser.ListAccessArrayContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listArray}.
	 * @param ctx the parse tree
	 */
	void enterListArray(SwiftGrammarParser.ListArrayContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listArray}.
	 * @param ctx the parse tree
	 */
	void exitListArray(SwiftGrammarParser.ListArrayContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#exprComa}.
	 * @param ctx the parse tree
	 */
	void enterExprComa(SwiftGrammarParser.ExprComaContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#exprComa}.
	 * @param ctx the parse tree
	 */
	void exitExprComa(SwiftGrammarParser.ExprComaContext ctx);
}