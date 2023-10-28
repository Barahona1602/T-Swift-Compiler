// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// SwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type SwiftGrammarListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterWhilestmt is called when entering the whilestmt production.
	EnterWhilestmt(c *WhilestmtContext)

	// EnterDeclarationstmt is called when entering the declarationstmt production.
	EnterDeclarationstmt(c *DeclarationstmtContext)

	// EnterAssignstmt is called when entering the assignstmt production.
	EnterAssignstmt(c *AssignstmtContext)

	// EnterForstmt is called when entering the forstmt production.
	EnterForstmt(c *ForstmtContext)

	// EnterGuardstmt is called when entering the guardstmt production.
	EnterGuardstmt(c *GuardstmtContext)

	// EnterBreakstmt is called when entering the breakstmt production.
	EnterBreakstmt(c *BreakstmtContext)

	// EnterContinuestmt is called when entering the continuestmt production.
	EnterContinuestmt(c *ContinuestmtContext)

	// EnterStructCreation is called when entering the structCreation production.
	EnterStructCreation(c *StructCreationContext)

	// EnterListStructDec is called when entering the listStructDec production.
	EnterListStructDec(c *ListStructDecContext)

	// EnterFnstmt is called when entering the fnstmt production.
	EnterFnstmt(c *FnstmtContext)

	// EnterListParamsFunc is called when entering the listParamsFunc production.
	EnterListParamsFunc(c *ListParamsFuncContext)

	// EnterParametro is called when entering the parametro production.
	EnterParametro(c *ParametroContext)

	// EnterCallExp is called when entering the callExp production.
	EnterCallExp(c *CallExpContext)

	// EnterCallFunction is called when entering the callFunction production.
	EnterCallFunction(c *CallFunctionContext)

	// EnterListParamsCall is called when entering the listParamsCall production.
	EnterListParamsCall(c *ListParamsCallContext)

	// EnterTypes is called when entering the types production.
	EnterTypes(c *TypesContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterListParams is called when entering the listParams production.
	EnterListParams(c *ListParamsContext)

	// EnterListAccessArray is called when entering the listAccessArray production.
	EnterListAccessArray(c *ListAccessArrayContext)

	// EnterListArray is called when entering the listArray production.
	EnterListArray(c *ListArrayContext)

	// EnterExprComa is called when entering the exprComa production.
	EnterExprComa(c *ExprComaContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitWhilestmt is called when exiting the whilestmt production.
	ExitWhilestmt(c *WhilestmtContext)

	// ExitDeclarationstmt is called when exiting the declarationstmt production.
	ExitDeclarationstmt(c *DeclarationstmtContext)

	// ExitAssignstmt is called when exiting the assignstmt production.
	ExitAssignstmt(c *AssignstmtContext)

	// ExitForstmt is called when exiting the forstmt production.
	ExitForstmt(c *ForstmtContext)

	// ExitGuardstmt is called when exiting the guardstmt production.
	ExitGuardstmt(c *GuardstmtContext)

	// ExitBreakstmt is called when exiting the breakstmt production.
	ExitBreakstmt(c *BreakstmtContext)

	// ExitContinuestmt is called when exiting the continuestmt production.
	ExitContinuestmt(c *ContinuestmtContext)

	// ExitStructCreation is called when exiting the structCreation production.
	ExitStructCreation(c *StructCreationContext)

	// ExitListStructDec is called when exiting the listStructDec production.
	ExitListStructDec(c *ListStructDecContext)

	// ExitFnstmt is called when exiting the fnstmt production.
	ExitFnstmt(c *FnstmtContext)

	// ExitListParamsFunc is called when exiting the listParamsFunc production.
	ExitListParamsFunc(c *ListParamsFuncContext)

	// ExitParametro is called when exiting the parametro production.
	ExitParametro(c *ParametroContext)

	// ExitCallExp is called when exiting the callExp production.
	ExitCallExp(c *CallExpContext)

	// ExitCallFunction is called when exiting the callFunction production.
	ExitCallFunction(c *CallFunctionContext)

	// ExitListParamsCall is called when exiting the listParamsCall production.
	ExitListParamsCall(c *ListParamsCallContext)

	// ExitTypes is called when exiting the types production.
	ExitTypes(c *TypesContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitListParams is called when exiting the listParams production.
	ExitListParams(c *ListParamsContext)

	// ExitListAccessArray is called when exiting the listAccessArray production.
	ExitListAccessArray(c *ListAccessArrayContext)

	// ExitListArray is called when exiting the listArray production.
	ExitListArray(c *ListArrayContext)

	// ExitExprComa is called when exiting the exprComa production.
	ExitExprComa(c *ExprComaContext)
}
