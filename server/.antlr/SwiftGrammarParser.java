// Generated from c://Users//pbara//OneDrive//Escritorio//T-Swift-Compiler//server//SwiftGrammar.g4 by ANTLR 4.13.1

    import "Server2/interfaces"
    import "Server2/environment"
    import "Server2/expressions"
    import "Server2/instructions"
    import "strings"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class SwiftGrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT=1, FLOAT=2, BOOL=3, STR=4, CHAR=5, VAR=6, LET=7, VOID=8, TRU=9, FAL=10, 
		PRINT=11, IF=12, ELSE=13, WHILE=14, FOR=15, IN=16, SWITCH=17, CASE=18, 
		DEFAULT=19, BREAK=20, RETURN=21, CONTINUE=22, GUARD=23, FUNC=24, NIL=25, 
		STRUCT=26, MUTATING=27, SELF=28, INOUT=29, APPEND=30, REMOVELAST=31, REMOVE=32, 
		AT=33, ISEMPTY=34, COUNT=35, NUMBER=36, STRING=37, ID=38, DIF=39, IG_IG=40, 
		NOT=41, OR=42, AND=43, IG=44, MAY_IG=45, MEN_IG=46, SUM_IG=47, SUB_IG=48, 
		MAYOR=49, MENOR=50, MUL=51, DIV=52, ADD=53, SUB=54, MOD=55, PARIZQ=56, 
		PARDER=57, LLAVEIZQ=58, LLAVEDER=59, D_PTS=60, CORIZQ=61, CORDER=62, COMA=63, 
		PUNTO=64, COMILLA=65, FLECHA=66, GUIONBAJO=67, AMP=68, WHITESPACE=69, 
		COMMENT=70, LINE_COMMENT=71;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_printstmt = 3, 
		RULE_ifstmt = 4, RULE_whilestmt = 5, RULE_declarationstmt = 6, RULE_assignstmt = 7, 
		RULE_forstmt = 8, RULE_guardstmt = 9, RULE_breakstmt = 10, RULE_continuestmt = 11, 
		RULE_structCreation = 12, RULE_listStructDec = 13, RULE_fnstmt = 14, RULE_listParamsFunc = 15, 
		RULE_parametro = 16, RULE_callExp = 17, RULE_callFunction = 18, RULE_listParamsCall = 19, 
		RULE_types = 20, RULE_expr = 21, RULE_listParams = 22, RULE_listAccessArray = 23, 
		RULE_listArray = 24, RULE_exprComa = 25;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "printstmt", "ifstmt", "whilestmt", "declarationstmt", 
			"assignstmt", "forstmt", "guardstmt", "breakstmt", "continuestmt", "structCreation", 
			"listStructDec", "fnstmt", "listParamsFunc", "parametro", "callExp", 
			"callFunction", "listParamsCall", "types", "expr", "listParams", "listAccessArray", 
			"listArray", "exprComa"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "'var'", 
			"'let'", "'void'", "'true'", "'false'", "'print'", "'if'", "'else'", 
			"'while'", "'for'", "'in'", "'switch'", "'case'", "'default'", "'break'", 
			"'return'", "'continue'", "'guard'", "'func'", "'nil'", "'struct'", "'mutating'", 
			"'self'", "'inout'", "'append'", "'removeLast'", "'remove'", "'at'", 
			"'isEmpty'", "'count'", null, null, null, "'!='", "'=='", "'!'", "'||'", 
			"'&&'", "'='", "'>='", "'<='", "'+='", "'-='", "'>'", "'<'", "'*'", "'/'", 
			"'+'", "'-'", "'%'", "'('", "')'", "'{'", "'}'", "':'", "'['", "']'", 
			"','", "'.'", "'\"'", "'->'", "'_'", "'&'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "STR", "CHAR", "VAR", "LET", "VOID", "TRU", 
			"FAL", "PRINT", "IF", "ELSE", "WHILE", "FOR", "IN", "SWITCH", "CASE", 
			"DEFAULT", "BREAK", "RETURN", "CONTINUE", "GUARD", "FUNC", "NIL", "STRUCT", 
			"MUTATING", "SELF", "INOUT", "APPEND", "REMOVELAST", "REMOVE", "AT", 
			"ISEMPTY", "COUNT", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", 
			"OR", "AND", "IG", "MAY_IG", "MEN_IG", "SUM_IG", "SUB_IG", "MAYOR", "MENOR", 
			"MUL", "DIV", "ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", 
			"D_PTS", "CORIZQ", "CORDER", "COMA", "PUNTO", "COMILLA", "FLECHA", "GUIONBAJO", 
			"AMP", "WHITESPACE", "COMMENT", "LINE_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "SwiftGrammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SwiftGrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SContext extends ParserRuleContext {
		public []interface{} code;
		public BlockContext block;
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode EOF() { return getToken(SwiftGrammarParser.EOF, 0); }
		public SContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_s; }
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(52);
			((SContext)_localctx).block = block();
			setState(53);
			match(EOF);
			   
			        _localctx.code = ((SContext)_localctx).block.blk
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public []interface{} blk;
		public InstructionContext instruction;
		public List<InstructionContext> ins = new ArrayList<InstructionContext>();
		public List<InstructionContext> instruction() {
			return getRuleContexts(InstructionContext.class);
		}
		public InstructionContext instruction(int i) {
			return getRuleContext(InstructionContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_block);

		    _localctx.blk = []interface{}{}
		    var listInt []IInstructionContext
		  
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(57); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(56);
				((BlockContext)_localctx).instruction = instruction();
				((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
				}
				}
				setState(59); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 274975480000L) != 0) );

			        listInt = localctx.(*BlockContext).GetIns()
			        for _, e := range listInt {
			            _localctx.blk = append(_localctx.blk, e.GetInst())
			        }
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InstructionContext extends ParserRuleContext {
		public interfaces.Instruction inst;
		public PrintstmtContext printstmt;
		public IfstmtContext ifstmt;
		public DeclarationstmtContext declarationstmt;
		public WhilestmtContext whilestmt;
		public AssignstmtContext assignstmt;
		public ForstmtContext forstmt;
		public GuardstmtContext guardstmt;
		public BreakstmtContext breakstmt;
		public ContinuestmtContext continuestmt;
		public StructCreationContext structCreation;
		public FnstmtContext fnstmt;
		public CallFunctionContext callFunction;
		public PrintstmtContext printstmt() {
			return getRuleContext(PrintstmtContext.class,0);
		}
		public IfstmtContext ifstmt() {
			return getRuleContext(IfstmtContext.class,0);
		}
		public DeclarationstmtContext declarationstmt() {
			return getRuleContext(DeclarationstmtContext.class,0);
		}
		public WhilestmtContext whilestmt() {
			return getRuleContext(WhilestmtContext.class,0);
		}
		public AssignstmtContext assignstmt() {
			return getRuleContext(AssignstmtContext.class,0);
		}
		public ForstmtContext forstmt() {
			return getRuleContext(ForstmtContext.class,0);
		}
		public GuardstmtContext guardstmt() {
			return getRuleContext(GuardstmtContext.class,0);
		}
		public BreakstmtContext breakstmt() {
			return getRuleContext(BreakstmtContext.class,0);
		}
		public ContinuestmtContext continuestmt() {
			return getRuleContext(ContinuestmtContext.class,0);
		}
		public StructCreationContext structCreation() {
			return getRuleContext(StructCreationContext.class,0);
		}
		public FnstmtContext fnstmt() {
			return getRuleContext(FnstmtContext.class,0);
		}
		public CallFunctionContext callFunction() {
			return getRuleContext(CallFunctionContext.class,0);
		}
		public InstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruction; }
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruction);
		try {
			setState(99);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(63);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(66);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ifinst 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(69);
				((InstructionContext)_localctx).declarationstmt = declarationstmt();
				 _localctx.inst = ((InstructionContext)_localctx).declarationstmt.dec 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(72);
				((InstructionContext)_localctx).whilestmt = whilestmt();
				 _localctx.inst = ((InstructionContext)_localctx).whilestmt.whl 
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(75);
				((InstructionContext)_localctx).assignstmt = assignstmt();
				 _localctx.inst = ((InstructionContext)_localctx).assignstmt.asg 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(78);
				((InstructionContext)_localctx).forstmt = forstmt();
				 _localctx.inst = ((InstructionContext)_localctx).forstmt.fr 
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(81);
				((InstructionContext)_localctx).guardstmt = guardstmt();
				 _localctx.inst = ((InstructionContext)_localctx).guardstmt.grd 
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(84);
				((InstructionContext)_localctx).breakstmt = breakstmt();
				 _localctx.inst = ((InstructionContext)_localctx).breakstmt.brk 
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(87);
				((InstructionContext)_localctx).continuestmt = continuestmt();
				 _localctx.inst = ((InstructionContext)_localctx).continuestmt.cnt 
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(90);
				((InstructionContext)_localctx).structCreation = structCreation();
				 _localctx.inst = ((InstructionContext)_localctx).structCreation.dec 
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(93);
				((InstructionContext)_localctx).fnstmt = fnstmt();
				 _localctx.inst = ((InstructionContext)_localctx).fnstmt.fn 
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(96);
				((InstructionContext)_localctx).callFunction = callFunction();
				 _localctx.inst = ((InstructionContext)_localctx).callFunction.cf 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrintstmtContext extends ParserRuleContext {
		public interfaces.Instruction prnt;
		public Token PRINT;
		public ExprContext expr;
		public ExprComaContext exprComa;
		public TerminalNode PRINT() { return getToken(SwiftGrammarParser.PRINT, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public ExprComaContext exprComa() {
			return getRuleContext(ExprComaContext.class,0);
		}
		public PrintstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printstmt; }
	}

	public final PrintstmtContext printstmt() throws RecognitionException {
		PrintstmtContext _localctx = new PrintstmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_printstmt);
		try {
			setState(113);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(101);
				((PrintstmtContext)_localctx).PRINT = match(PRINT);
				setState(102);
				match(PARIZQ);
				setState(103);
				((PrintstmtContext)_localctx).expr = expr(0);
				setState(104);
				match(PARDER);
				 _localctx.prnt = instructions.NewPrint((((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getLine():0),(((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getCharPositionInLine():0),((PrintstmtContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(107);
				((PrintstmtContext)_localctx).PRINT = match(PRINT);
				setState(108);
				match(PARIZQ);
				setState(109);
				((PrintstmtContext)_localctx).exprComa = exprComa(0);
				setState(110);
				match(PARDER);
				 _localctx.prnt = instructions.NewPrint((((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getLine():0),(((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getCharPositionInLine():0),((PrintstmtContext)_localctx).exprComa.t)
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IfstmtContext extends ParserRuleContext {
		public interfaces.Instruction ifinst;
		public Token IF;
		public ExprContext expr;
		public BlockContext block;
		public BlockContext e1;
		public BlockContext e2;
		public IfstmtContext ifstmt;
		public TerminalNode IF() { return getToken(SwiftGrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<TerminalNode> LLAVEIZQ() { return getTokens(SwiftGrammarParser.LLAVEIZQ); }
		public TerminalNode LLAVEIZQ(int i) {
			return getToken(SwiftGrammarParser.LLAVEIZQ, i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public List<TerminalNode> LLAVEDER() { return getTokens(SwiftGrammarParser.LLAVEDER); }
		public TerminalNode LLAVEDER(int i) {
			return getToken(SwiftGrammarParser.LLAVEDER, i);
		}
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public IfstmtContext ifstmt() {
			return getRuleContext(IfstmtContext.class,0);
		}
		public IfstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifstmt; }
	}

	public final IfstmtContext ifstmt() throws RecognitionException {
		IfstmtContext _localctx = new IfstmtContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_ifstmt);
		try {
			setState(142);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(115);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(116);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(117);
				match(LLAVEIZQ);
				setState(118);
				((IfstmtContext)_localctx).block = block();
				setState(119);
				match(LLAVEDER);
				 _localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).block.blk, nil) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(122);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(123);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(124);
				match(LLAVEIZQ);
				setState(125);
				((IfstmtContext)_localctx).e1 = block();
				setState(126);
				match(LLAVEDER);
				setState(127);
				match(ELSE);
				setState(128);
				match(LLAVEIZQ);
				setState(129);
				((IfstmtContext)_localctx).e2 = block();
				setState(130);
				match(LLAVEDER);
				 _localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).e1.blk, ((IfstmtContext)_localctx).e2.blk) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(133);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(134);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(135);
				match(LLAVEIZQ);
				setState(136);
				((IfstmtContext)_localctx).block = block();
				setState(137);
				match(LLAVEDER);
				setState(138);
				match(ELSE);
				setState(139);
				((IfstmtContext)_localctx).ifstmt = ifstmt();
				 _localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).block.blk, []interface{}{((IfstmtContext)_localctx).ifstmt.ifinst}) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class WhilestmtContext extends ParserRuleContext {
		public interfaces.Instruction whl;
		public Token WHILE;
		public ExprContext expr;
		public BlockContext block;
		public TerminalNode WHILE() { return getToken(SwiftGrammarParser.WHILE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public WhilestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whilestmt; }
	}

	public final WhilestmtContext whilestmt() throws RecognitionException {
		WhilestmtContext _localctx = new WhilestmtContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(144);
			((WhilestmtContext)_localctx).WHILE = match(WHILE);
			setState(145);
			((WhilestmtContext)_localctx).expr = expr(0);
			setState(146);
			match(LLAVEIZQ);
			setState(147);
			((WhilestmtContext)_localctx).block = block();
			setState(148);
			match(LLAVEDER);
			 _localctx.whl = instructions.NewWhile((((WhilestmtContext)_localctx).WHILE!=null?((WhilestmtContext)_localctx).WHILE.getLine():0), (((WhilestmtContext)_localctx).WHILE!=null?((WhilestmtContext)_localctx).WHILE.getCharPositionInLine():0), ((WhilestmtContext)_localctx).expr.e, ((WhilestmtContext)_localctx).block.blk) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DeclarationstmtContext extends ParserRuleContext {
		public interfaces.Instruction dec;
		public Token VAR;
		public Token ID;
		public TypesContext types;
		public ExprContext expr;
		public Token LET;
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode D_PTS() { return getToken(SwiftGrammarParser.D_PTS, 0); }
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LET() { return getToken(SwiftGrammarParser.LET, 0); }
		public DeclarationstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarationstmt; }
	}

	public final DeclarationstmtContext declarationstmt() throws RecognitionException {
		DeclarationstmtContext _localctx = new DeclarationstmtContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_declarationstmt);
		try {
			setState(191);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(151);
				((DeclarationstmtContext)_localctx).VAR = match(VAR);
				setState(152);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(153);
				match(D_PTS);
				setState(154);
				((DeclarationstmtContext)_localctx).types = types();
				setState(155);
				match(IG);
				setState(156);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				 _localctx.dec = instructions.NewDeclaration((((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getLine():0), (((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), ((DeclarationstmtContext)_localctx).types.ty, ((DeclarationstmtContext)_localctx).expr.e, true) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(159);
				((DeclarationstmtContext)_localctx).VAR = match(VAR);
				setState(160);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(161);
				match(IG);
				setState(162);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				 _localctx.dec = instructions.NewDeclaration((((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getLine():0), (((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), environment.UNKNOWN, ((DeclarationstmtContext)_localctx).expr.e, true) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(165);
				((DeclarationstmtContext)_localctx).VAR = match(VAR);
				setState(166);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(167);
				match(D_PTS);
				setState(168);
				((DeclarationstmtContext)_localctx).types = types();
				 _localctx.dec = instructions.NewDeclaration((((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getLine():0), (((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), ((DeclarationstmtContext)_localctx).types.ty, nil, true) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(171);
				((DeclarationstmtContext)_localctx).LET = match(LET);
				setState(172);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(173);
				match(D_PTS);
				setState(174);
				((DeclarationstmtContext)_localctx).types = types();
				setState(175);
				match(IG);
				setState(176);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				 _localctx.dec = instructions.NewDeclaration((((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getLine():0), (((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), ((DeclarationstmtContext)_localctx).types.ty, ((DeclarationstmtContext)_localctx).expr.e, false) 
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(179);
				((DeclarationstmtContext)_localctx).LET = match(LET);
				setState(180);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(181);
				match(D_PTS);
				setState(182);
				((DeclarationstmtContext)_localctx).types = types();
				 _localctx.dec = instructions.NewDeclaration((((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getLine():0), (((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), ((DeclarationstmtContext)_localctx).types.ty, nil, false) 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(185);
				((DeclarationstmtContext)_localctx).LET = match(LET);
				setState(186);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(187);
				match(IG);
				setState(188);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				 _localctx.dec = instructions.NewDeclaration((((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getLine():0), (((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), environment.UNKNOWN, ((DeclarationstmtContext)_localctx).expr.e, false) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AssignstmtContext extends ParserRuleContext {
		public interfaces.Instruction asg;
		public Token ID;
		public Token op;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public TerminalNode SUB_IG() { return getToken(SwiftGrammarParser.SUB_IG, 0); }
		public TerminalNode SUM_IG() { return getToken(SwiftGrammarParser.SUM_IG, 0); }
		public AssignstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignstmt; }
	}

	public final AssignstmtContext assignstmt() throws RecognitionException {
		AssignstmtContext _localctx = new AssignstmtContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_assignstmt);
		int _la;
		try {
			setState(203);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(193);
				((AssignstmtContext)_localctx).ID = match(ID);
				setState(194);
				((AssignstmtContext)_localctx).op = match(IG);
				setState(195);
				((AssignstmtContext)_localctx).expr = expr(0);
				 _localctx.asg = instructions.NewAssignment((((AssignstmtContext)_localctx).ID!=null?((AssignstmtContext)_localctx).ID.getLine():0), (((AssignstmtContext)_localctx).ID!=null?((AssignstmtContext)_localctx).ID.getCharPositionInLine():0), (((AssignstmtContext)_localctx).ID!=null?((AssignstmtContext)_localctx).ID.getText():null), ((AssignstmtContext)_localctx).expr.e) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(198);
				((AssignstmtContext)_localctx).ID = match(ID);
				setState(199);
				((AssignstmtContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==SUM_IG || _la==SUB_IG) ) {
					((AssignstmtContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(200);
				((AssignstmtContext)_localctx).expr = expr(0);
				 ((AssignstmtContext)_localctx).asg =  instructions.NewImplicitAssignment((((AssignstmtContext)_localctx).ID!=null?((AssignstmtContext)_localctx).ID.getLine():0), (((AssignstmtContext)_localctx).ID!=null?((AssignstmtContext)_localctx).ID.getCharPositionInLine():0), (((AssignstmtContext)_localctx).ID!=null?((AssignstmtContext)_localctx).ID.getText():null), (((AssignstmtContext)_localctx).op!=null?((AssignstmtContext)_localctx).op.getText():null), ((AssignstmtContext)_localctx).expr.e); 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ForstmtContext extends ParserRuleContext {
		public interfaces.Instruction fr;
		public Token FOR;
		public Token ID;
		public ExprContext expr;
		public BlockContext block;
		public TerminalNode FOR() { return getToken(SwiftGrammarParser.FOR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode IN() { return getToken(SwiftGrammarParser.IN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public ForstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forstmt; }
	}

	public final ForstmtContext forstmt() throws RecognitionException {
		ForstmtContext _localctx = new ForstmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_forstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(205);
			((ForstmtContext)_localctx).FOR = match(FOR);
			setState(206);
			((ForstmtContext)_localctx).ID = match(ID);
			setState(207);
			match(IN);
			setState(208);
			((ForstmtContext)_localctx).expr = expr(0);
			setState(209);
			match(LLAVEIZQ);
			setState(210);
			((ForstmtContext)_localctx).block = block();
			setState(211);
			match(LLAVEDER);
			 _localctx.fr = instructions.NewForIn((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), ((ForstmtContext)_localctx).expr.e, ((ForstmtContext)_localctx).block.blk) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class GuardstmtContext extends ParserRuleContext {
		public interfaces.Instruction grd;
		public Token GUARD;
		public ExprContext expr;
		public BlockContext block;
		public TerminalNode GUARD() { return getToken(SwiftGrammarParser.GUARD, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public GuardstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guardstmt; }
	}

	public final GuardstmtContext guardstmt() throws RecognitionException {
		GuardstmtContext _localctx = new GuardstmtContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_guardstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(214);
			((GuardstmtContext)_localctx).GUARD = match(GUARD);
			setState(215);
			((GuardstmtContext)_localctx).expr = expr(0);
			setState(216);
			match(ELSE);
			setState(217);
			match(LLAVEIZQ);
			setState(218);
			((GuardstmtContext)_localctx).block = block();
			setState(219);
			match(LLAVEDER);
			 _localctx.grd = instructions.NewGuard((((GuardstmtContext)_localctx).GUARD!=null?((GuardstmtContext)_localctx).GUARD.getLine():0), (((GuardstmtContext)_localctx).GUARD!=null?((GuardstmtContext)_localctx).GUARD.getCharPositionInLine():0), ((GuardstmtContext)_localctx).expr.e, ((GuardstmtContext)_localctx).block.blk) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BreakstmtContext extends ParserRuleContext {
		public interfaces.Instruction brk;
		public Token BREAK;
		public TerminalNode BREAK() { return getToken(SwiftGrammarParser.BREAK, 0); }
		public BreakstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_breakstmt; }
	}

	public final BreakstmtContext breakstmt() throws RecognitionException {
		BreakstmtContext _localctx = new BreakstmtContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_breakstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(222);
			((BreakstmtContext)_localctx).BREAK = match(BREAK);
			 _localctx.brk = instructions.NewBreak((((BreakstmtContext)_localctx).BREAK!=null?((BreakstmtContext)_localctx).BREAK.getLine():0), (((BreakstmtContext)_localctx).BREAK!=null?((BreakstmtContext)_localctx).BREAK.getCharPositionInLine():0), nil) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ContinuestmtContext extends ParserRuleContext {
		public interfaces.Instruction cnt;
		public Token CONTINUE;
		public TerminalNode CONTINUE() { return getToken(SwiftGrammarParser.CONTINUE, 0); }
		public ContinuestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continuestmt; }
	}

	public final ContinuestmtContext continuestmt() throws RecognitionException {
		ContinuestmtContext _localctx = new ContinuestmtContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_continuestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(225);
			((ContinuestmtContext)_localctx).CONTINUE = match(CONTINUE);
			 _localctx.cnt = instructions.NewContinue((((ContinuestmtContext)_localctx).CONTINUE!=null?((ContinuestmtContext)_localctx).CONTINUE.getLine():0), (((ContinuestmtContext)_localctx).CONTINUE!=null?((ContinuestmtContext)_localctx).CONTINUE.getCharPositionInLine():0)) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructCreationContext extends ParserRuleContext {
		public interfaces.Instruction dec;
		public Token STRUCT;
		public Token ID;
		public ListStructDecContext listStructDec;
		public TerminalNode STRUCT() { return getToken(SwiftGrammarParser.STRUCT, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public ListStructDecContext listStructDec() {
			return getRuleContext(ListStructDecContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public StructCreationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structCreation; }
	}

	public final StructCreationContext structCreation() throws RecognitionException {
		StructCreationContext _localctx = new StructCreationContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_structCreation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(228);
			((StructCreationContext)_localctx).STRUCT = match(STRUCT);
			setState(229);
			((StructCreationContext)_localctx).ID = match(ID);
			setState(230);
			match(LLAVEIZQ);
			setState(231);
			((StructCreationContext)_localctx).listStructDec = listStructDec(0);
			setState(232);
			match(LLAVEDER);
			 _localctx.dec = instructions.NewStruct((((StructCreationContext)_localctx).STRUCT!=null?((StructCreationContext)_localctx).STRUCT.getLine():0), (((StructCreationContext)_localctx).STRUCT!=null?((StructCreationContext)_localctx).STRUCT.getCharPositionInLine():0), (((StructCreationContext)_localctx).ID!=null?((StructCreationContext)_localctx).ID.getText():null), ((StructCreationContext)_localctx).listStructDec.l) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListStructDecContext extends ParserRuleContext {
		public []interface{} l;
		public ListStructDecContext list;
		public Token ID;
		public TypesContext types;
		public Token id1;
		public Token id2;
		public List<TerminalNode> ID() { return getTokens(SwiftGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(SwiftGrammarParser.ID, i);
		}
		public TerminalNode D_PTS() { return getToken(SwiftGrammarParser.D_PTS, 0); }
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode LET() { return getToken(SwiftGrammarParser.LET, 0); }
		public ListStructDecContext listStructDec() {
			return getRuleContext(ListStructDecContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ListStructDecContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listStructDec; }
	}

	public final ListStructDecContext listStructDec() throws RecognitionException {
		return listStructDec(0);
	}

	private ListStructDecContext listStructDec(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListStructDecContext _localctx = new ListStructDecContext(_ctx, _parentState);
		ListStructDecContext _prevctx = _localctx;
		int _startState = 26;
		enterRecursionRule(_localctx, 26, RULE_listStructDec, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(248);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				{
				setState(236);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==LET) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(237);
				((ListStructDecContext)_localctx).ID = match(ID);
				setState(238);
				match(D_PTS);
				setState(239);
				((ListStructDecContext)_localctx).types = types();

				                        var arr []interface{}
				                        newParams := environment.NewStructType((((ListStructDecContext)_localctx).ID!=null?((ListStructDecContext)_localctx).ID.getText():null), ((ListStructDecContext)_localctx).types.ty, "")
				                        arr = append(arr, newParams)
				                        _localctx.l = arr
				                    
				}
				break;
			case 2:
				{
				setState(242);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==LET) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(243);
				((ListStructDecContext)_localctx).id1 = match(ID);
				setState(244);
				match(D_PTS);
				setState(245);
				((ListStructDecContext)_localctx).id2 = match(ID);

				                                                var arr []interface{}
				                                                newParams := environment.NewStructType((((ListStructDecContext)_localctx).id1!=null?((ListStructDecContext)_localctx).id1.getText():null),environment.UNKNOWN , (((ListStructDecContext)_localctx).id2!=null?((ListStructDecContext)_localctx).id2.getText():null))
				                                                arr = append(((ListStructDecContext)_localctx).list.l, newParams)
				                                                _localctx.l = arr
				                                            
				}
				break;
			case 3:
				{
				 _localctx.l = []interface{}{} 
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(285);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(283);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
					case 1:
						{
						_localctx = new ListStructDecContext(_parentctx, _parentState);
						_localctx.list = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_listStructDec);
						setState(250);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(252);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==COMA) {
							{
							setState(251);
							match(COMA);
							}
						}

						setState(254);
						_la = _input.LA(1);
						if ( !(_la==VAR || _la==LET) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(255);
						((ListStructDecContext)_localctx).ID = match(ID);
						setState(256);
						match(D_PTS);
						setState(257);
						((ListStructDecContext)_localctx).types = types();

						                                                          var arr []interface{}
						                                                          newParams := environment.NewStructType((((ListStructDecContext)_localctx).ID!=null?((ListStructDecContext)_localctx).ID.getText():null), ((ListStructDecContext)_localctx).types.ty, "")
						                                                          arr = append(((ListStructDecContext)_localctx).list.l, newParams)
						                                                          _localctx.l = arr
						                                                      
						}
						break;
					case 2:
						{
						_localctx = new ListStructDecContext(_parentctx, _parentState);
						_localctx.list = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_listStructDec);
						setState(260);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(262);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==COMA) {
							{
							setState(261);
							match(COMA);
							}
						}

						setState(264);
						_la = _input.LA(1);
						if ( !(_la==VAR || _la==LET) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(265);
						((ListStructDecContext)_localctx).id1 = match(ID);
						setState(266);
						match(D_PTS);
						setState(267);
						((ListStructDecContext)_localctx).id2 = match(ID);

						                                                          var arr []interface{}
						                                                          newParams := environment.NewStructType((((ListStructDecContext)_localctx).id1!=null?((ListStructDecContext)_localctx).id1.getText():null), environment.UNKNOWN, (((ListStructDecContext)_localctx).id2!=null?((ListStructDecContext)_localctx).id2.getText():null))
						                                                          arr = append(((ListStructDecContext)_localctx).list.l, newParams)
						                                                          _localctx.l = arr
						                                                      
						}
						break;
					case 3:
						{
						_localctx = new ListStructDecContext(_parentctx, _parentState);
						_localctx.list = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_listStructDec);
						setState(269);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(271);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==COMA) {
							{
							setState(270);
							match(COMA);
							}
						}

						setState(273);
						_la = _input.LA(1);
						if ( !(_la==VAR || _la==LET) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(274);
						((ListStructDecContext)_localctx).id1 = match(ID);
						setState(277);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==D_PTS) {
							{
							setState(275);
							match(D_PTS);
							setState(276);
							((ListStructDecContext)_localctx).types = types();
							}
						}

						setState(279);
						match(IG);
						setState(280);
						expr(0);

						                                                          var arr []interface{}
						                                                          newParams := environment.NewStructType((((ListStructDecContext)_localctx).id1!=null?((ListStructDecContext)_localctx).id1.getText():null), ((ListStructDecContext)_localctx).types.ty, "")
						                                                          arr = append(((ListStructDecContext)_localctx).list.l, newParams)
						                                                          _localctx.l = arr
						                                                      
						}
						break;
					}
					} 
				}
				setState(287);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FnstmtContext extends ParserRuleContext {
		public interfaces.Instruction fn;
		public Token FUNC;
		public Token ID;
		public ListParamsFuncContext listParamsFunc;
		public TypesContext types;
		public BlockContext block;
		public TerminalNode FUNC() { return getToken(SwiftGrammarParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ListParamsFuncContext listParamsFunc() {
			return getRuleContext(ListParamsFuncContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public TerminalNode FLECHA() { return getToken(SwiftGrammarParser.FLECHA, 0); }
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public FnstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fnstmt; }
	}

	public final FnstmtContext fnstmt() throws RecognitionException {
		FnstmtContext _localctx = new FnstmtContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_fnstmt);
		try {
			setState(310);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(288);
				((FnstmtContext)_localctx).FUNC = match(FUNC);
				setState(289);
				((FnstmtContext)_localctx).ID = match(ID);
				setState(290);
				match(PARIZQ);
				setState(291);
				((FnstmtContext)_localctx).listParamsFunc = listParamsFunc(0);
				setState(292);
				match(PARDER);
				setState(293);
				match(FLECHA);
				setState(294);
				((FnstmtContext)_localctx).types = types();
				setState(295);
				match(LLAVEIZQ);
				setState(296);
				((FnstmtContext)_localctx).block = block();
				setState(297);
				match(LLAVEDER);
				 _localctx.fn = instructions.NewFunction((((FnstmtContext)_localctx).FUNC!=null?((FnstmtContext)_localctx).FUNC.getLine():0), (((FnstmtContext)_localctx).FUNC!=null?((FnstmtContext)_localctx).FUNC.getCharPositionInLine():0), (((FnstmtContext)_localctx).ID!=null?((FnstmtContext)_localctx).ID.getText():null), ((FnstmtContext)_localctx).listParamsFunc.l, ((FnstmtContext)_localctx).types.ty, ((FnstmtContext)_localctx).block.blk) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(300);
				((FnstmtContext)_localctx).FUNC = match(FUNC);
				setState(301);
				((FnstmtContext)_localctx).ID = match(ID);
				setState(302);
				match(PARIZQ);
				setState(303);
				((FnstmtContext)_localctx).listParamsFunc = listParamsFunc(0);
				setState(304);
				match(PARDER);
				setState(305);
				match(LLAVEIZQ);
				setState(306);
				((FnstmtContext)_localctx).block = block();
				setState(307);
				match(LLAVEDER);
				 _localctx.fn = instructions.NewFunction((((FnstmtContext)_localctx).FUNC!=null?((FnstmtContext)_localctx).FUNC.getLine():0), (((FnstmtContext)_localctx).FUNC!=null?((FnstmtContext)_localctx).FUNC.getCharPositionInLine():0), (((FnstmtContext)_localctx).ID!=null?((FnstmtContext)_localctx).ID.getText():null), ((FnstmtContext)_localctx).listParamsFunc.l, environment.NIL, ((FnstmtContext)_localctx).block.blk) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListParamsFuncContext extends ParserRuleContext {
		public []interface{} l;
		public ListParamsFuncContext list;
		public ParametroContext parametro;
		public ParametroContext parametro() {
			return getRuleContext(ParametroContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ListParamsFuncContext listParamsFunc() {
			return getRuleContext(ListParamsFuncContext.class,0);
		}
		public ListParamsFuncContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listParamsFunc; }
	}

	public final ListParamsFuncContext listParamsFunc() throws RecognitionException {
		return listParamsFunc(0);
	}

	private ListParamsFuncContext listParamsFunc(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListParamsFuncContext _localctx = new ListParamsFuncContext(_ctx, _parentState);
		ListParamsFuncContext _prevctx = _localctx;
		int _startState = 30;
		enterRecursionRule(_localctx, 30, RULE_listParamsFunc, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(317);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				{
				setState(313);
				((ListParamsFuncContext)_localctx).parametro = parametro();

				            _localctx.l = []interface{}{}
				            _localctx.l = append(_localctx.l, ((ListParamsFuncContext)_localctx).parametro.p)
				        
				}
				break;
			case 2:
				{

				        _localctx.l = []interface{}{}
				    
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(326);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListParamsFuncContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listParamsFunc);
					setState(319);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(320);
					match(COMA);
					setState(321);
					((ListParamsFuncContext)_localctx).parametro = parametro();

					                                          var arr []interface{}
					                                          arr = append(((ListParamsFuncContext)_localctx).list.l, ((ListParamsFuncContext)_localctx).parametro.p)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(328);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParametroContext extends ParserRuleContext {
		public interfaces.Instruction p;
		public Token ID;
		public TypesContext types;
		public Token exte;
		public List<TerminalNode> ID() { return getTokens(SwiftGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(SwiftGrammarParser.ID, i);
		}
		public TerminalNode D_PTS() { return getToken(SwiftGrammarParser.D_PTS, 0); }
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public TerminalNode GUIONBAJO() { return getToken(SwiftGrammarParser.GUIONBAJO, 0); }
		public ParametroContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametro; }
	}

	public final ParametroContext parametro() throws RecognitionException {
		ParametroContext _localctx = new ParametroContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_parametro);
		int _la;
		try {
			setState(340);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(329);
				((ParametroContext)_localctx).ID = match(ID);
				setState(330);
				match(D_PTS);
				setState(331);
				((ParametroContext)_localctx).types = types();
				 _localctx.p = instructions.NewParamsDeclaration((((ParametroContext)_localctx).ID!=null?((ParametroContext)_localctx).ID.getLine():0),(((ParametroContext)_localctx).ID!=null?((ParametroContext)_localctx).ID.getCharPositionInLine():0),(((ParametroContext)_localctx).ID!=null?((ParametroContext)_localctx).ID.getText():null), ((ParametroContext)_localctx).types.ty)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(334);
				((ParametroContext)_localctx).exte = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==ID || _la==GUIONBAJO) ) {
					((ParametroContext)_localctx).exte = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(335);
				((ParametroContext)_localctx).ID = match(ID);
				setState(336);
				match(D_PTS);
				setState(337);
				((ParametroContext)_localctx).types = types();
				 _localctx.p = instructions.NewParamsDeclaration((((ParametroContext)_localctx).ID!=null?((ParametroContext)_localctx).ID.getLine():0),(((ParametroContext)_localctx).ID!=null?((ParametroContext)_localctx).ID.getCharPositionInLine():0), (((ParametroContext)_localctx).ID!=null?((ParametroContext)_localctx).ID.getText():null), ((ParametroContext)_localctx).types.ty)
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CallExpContext extends ParserRuleContext {
		public interfaces.Expression cfe;
		public Token ID;
		public ListParamsCallContext listParamsCall;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ListParamsCallContext listParamsCall() {
			return getRuleContext(ListParamsCallContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public CallExpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_callExp; }
	}

	public final CallExpContext callExp() throws RecognitionException {
		CallExpContext _localctx = new CallExpContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_callExp);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(342);
			((CallExpContext)_localctx).ID = match(ID);
			setState(343);
			match(PARIZQ);
			setState(344);
			((CallExpContext)_localctx).listParamsCall = listParamsCall(0);
			setState(345);
			match(PARDER);
			 _localctx.cfe = expressions.NewCallExp((((CallExpContext)_localctx).ID!=null?((CallExpContext)_localctx).ID.getLine():0), (((CallExpContext)_localctx).ID!=null?((CallExpContext)_localctx).ID.getCharPositionInLine():0), (((CallExpContext)_localctx).ID!=null?((CallExpContext)_localctx).ID.getText():null), ((CallExpContext)_localctx).listParamsCall.l) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CallFunctionContext extends ParserRuleContext {
		public interfaces.Instruction cf;
		public Token ID;
		public ListParamsCallContext listParamsCall;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ListParamsCallContext listParamsCall() {
			return getRuleContext(ListParamsCallContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public CallFunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_callFunction; }
	}

	public final CallFunctionContext callFunction() throws RecognitionException {
		CallFunctionContext _localctx = new CallFunctionContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_callFunction);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(348);
			((CallFunctionContext)_localctx).ID = match(ID);
			setState(349);
			match(PARIZQ);
			setState(350);
			((CallFunctionContext)_localctx).listParamsCall = listParamsCall(0);
			setState(351);
			match(PARDER);
			 _localctx.cf = instructions.NewCallFunc((((CallFunctionContext)_localctx).ID!=null?((CallFunctionContext)_localctx).ID.getLine():0), (((CallFunctionContext)_localctx).ID!=null?((CallFunctionContext)_localctx).ID.getCharPositionInLine():0), (((CallFunctionContext)_localctx).ID!=null?((CallFunctionContext)_localctx).ID.getText():null), ((CallFunctionContext)_localctx).listParamsCall.l) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListParamsCallContext extends ParserRuleContext {
		public []interface{} l;
		public ListParamsCallContext list;
		public ExprContext expr;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ListParamsCallContext listParamsCall() {
			return getRuleContext(ListParamsCallContext.class,0);
		}
		public ListParamsCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listParamsCall; }
	}

	public final ListParamsCallContext listParamsCall() throws RecognitionException {
		return listParamsCall(0);
	}

	private ListParamsCallContext listParamsCall(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListParamsCallContext _localctx = new ListParamsCallContext(_ctx, _parentState);
		ListParamsCallContext _prevctx = _localctx;
		int _startState = 38;
		enterRecursionRule(_localctx, 38, RULE_listParamsCall, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(359);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				{
				setState(355);
				((ListParamsCallContext)_localctx).expr = expr(0);

				            _localctx.l = []interface{}{}
				            _localctx.l = append(_localctx.l, ((ListParamsCallContext)_localctx).expr.e)
				        
				}
				break;
			case 2:
				{

				        _localctx.l = []interface{}{}
				    
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(368);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListParamsCallContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listParamsCall);
					setState(361);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(362);
					match(COMA);
					setState(363);
					((ListParamsCallContext)_localctx).expr = expr(0);

					                                              var arr []interface{}
					                                              arr = append(((ListParamsCallContext)_localctx).list.l, ((ListParamsCallContext)_localctx).expr.e)
					                                              _localctx.l = arr
					                                          
					}
					} 
				}
				setState(370);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypesContext extends ParserRuleContext {
		public environment.TipoExpresion ty;
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode STR() { return getToken(SwiftGrammarParser.STR, 0); }
		public TerminalNode BOOL() { return getToken(SwiftGrammarParser.BOOL, 0); }
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public List<TerminalNode> COMILLA() { return getTokens(SwiftGrammarParser.COMILLA); }
		public TerminalNode COMILLA(int i) {
			return getToken(SwiftGrammarParser.COMILLA, i);
		}
		public TerminalNode NIL() { return getToken(SwiftGrammarParser.NIL, 0); }
		public TerminalNode STRUCT() { return getToken(SwiftGrammarParser.STRUCT, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TypesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_types; }
	}

	public final TypesContext types() throws RecognitionException {
		TypesContext _localctx = new TypesContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_types);
		try {
			setState(394);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(371);
				match(INT);
				 _localctx.ty = environment.INTEGER 
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(373);
				match(FLOAT);
				 _localctx.ty = environment.FLOAT 
				}
				break;
			case STR:
				enterOuterAlt(_localctx, 3);
				{
				setState(375);
				match(STR);
				 _localctx.ty = environment.STRING 
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(377);
				match(BOOL);
				 _localctx.ty = environment.BOOLEAN 
				}
				break;
			case CORIZQ:
				enterOuterAlt(_localctx, 5);
				{
				setState(379);
				match(CORIZQ);
				setState(380);
				types();
				setState(381);
				match(CORDER);
				 _localctx.ty = environment.ARRAY 
				}
				break;
			case COMILLA:
				enterOuterAlt(_localctx, 6);
				{
				setState(384);
				match(COMILLA);
				setState(385);
				match(STR);
				setState(386);
				match(COMILLA);
				 _localctx.ty = environment.STR 
				}
				break;
			case NIL:
				enterOuterAlt(_localctx, 7);
				{
				setState(388);
				match(NIL);
				 _localctx.ty = environment.NIL 
				}
				break;
			case STRUCT:
				enterOuterAlt(_localctx, 8);
				{
				setState(390);
				match(STRUCT);
				 _localctx.ty = environment.STRUCT 
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 9);
				{
				setState(392);
				match(ID);
				 _localctx.ty = environment.UNKNOWN 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprContext extends ParserRuleContext {
		public interfaces.Expression e;
		public ExprContext expuno;
		public ExprContext left;
		public Token SUB;
		public ExprContext opDe;
		public ExprContext expr;
		public Token NOT;
		public ExprContext right;
		public CallExpContext callExp;
		public Token CORIZQ;
		public ListArrayContext list;
		public ListParamsContext listParams;
		public Token NUMBER;
		public Token STRING;
		public Token TRU;
		public Token FAL;
		public Token ID;
		public Token NIL;
		public ExprContext expdos;
		public Token op;
		public TerminalNode SUB() { return getToken(SwiftGrammarParser.SUB, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode NOT() { return getToken(SwiftGrammarParser.NOT, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public CallExpContext callExp() {
			return getRuleContext(CallExpContext.class,0);
		}
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public ListArrayContext listArray() {
			return getRuleContext(ListArrayContext.class,0);
		}
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public TerminalNode NUMBER() { return getToken(SwiftGrammarParser.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(SwiftGrammarParser.STRING, 0); }
		public TerminalNode TRU() { return getToken(SwiftGrammarParser.TRU, 0); }
		public TerminalNode FAL() { return getToken(SwiftGrammarParser.FAL, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public List<TerminalNode> PUNTO() { return getTokens(SwiftGrammarParser.PUNTO); }
		public TerminalNode PUNTO(int i) {
			return getToken(SwiftGrammarParser.PUNTO, i);
		}
		public TerminalNode COUNT() { return getToken(SwiftGrammarParser.COUNT, 0); }
		public TerminalNode NIL() { return getToken(SwiftGrammarParser.NIL, 0); }
		public TerminalNode SUB_IG() { return getToken(SwiftGrammarParser.SUB_IG, 0); }
		public TerminalNode SUM_IG() { return getToken(SwiftGrammarParser.SUM_IG, 0); }
		public TerminalNode MUL() { return getToken(SwiftGrammarParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(SwiftGrammarParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(SwiftGrammarParser.MOD, 0); }
		public TerminalNode ADD() { return getToken(SwiftGrammarParser.ADD, 0); }
		public TerminalNode MAY_IG() { return getToken(SwiftGrammarParser.MAY_IG, 0); }
		public TerminalNode MAYOR() { return getToken(SwiftGrammarParser.MAYOR, 0); }
		public TerminalNode MEN_IG() { return getToken(SwiftGrammarParser.MEN_IG, 0); }
		public TerminalNode MENOR() { return getToken(SwiftGrammarParser.MENOR, 0); }
		public TerminalNode IG_IG() { return getToken(SwiftGrammarParser.IG_IG, 0); }
		public TerminalNode DIF() { return getToken(SwiftGrammarParser.DIF, 0); }
		public TerminalNode AND() { return getToken(SwiftGrammarParser.AND, 0); }
		public TerminalNode OR() { return getToken(SwiftGrammarParser.OR, 0); }
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 42;
		enterRecursionRule(_localctx, 42, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(438);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				{
				setState(397);
				((ExprContext)_localctx).SUB = match(SUB);
				setState(398);
				((ExprContext)_localctx).opDe = ((ExprContext)_localctx).expr = expr(22);
				_localctx.e = expressions.NewOperation((((ExprContext)_localctx).SUB!=null?((ExprContext)_localctx).SUB.getLine():0),(((ExprContext)_localctx).SUB!=null?((ExprContext)_localctx).SUB.getCharPositionInLine():0),((ExprContext)_localctx).opDe.e,"NEGACION",nil)
				}
				break;
			case 2:
				{
				setState(401);
				((ExprContext)_localctx).NOT = match(NOT);
				setState(402);
				((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(12);
				_localctx.e = expressions.NewOperation((((ExprContext)_localctx).NOT!=null?((ExprContext)_localctx).NOT.getLine():0), (((ExprContext)_localctx).NOT!=null?((ExprContext)_localctx).NOT.getCharPositionInLine():0), ((ExprContext)_localctx).right.e, (((ExprContext)_localctx).NOT!=null?((ExprContext)_localctx).NOT.getText():null) ,nil)
				}
				break;
			case 3:
				{
				setState(405);
				match(PARIZQ);
				setState(406);
				((ExprContext)_localctx).expr = expr(0);
				setState(407);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case 4:
				{
				setState(410);
				((ExprContext)_localctx).callExp = callExp();
				 _localctx.e = ((ExprContext)_localctx).callExp.cfe 
				}
				break;
			case 5:
				{
				setState(413);
				((ExprContext)_localctx).CORIZQ = match(CORIZQ);
				setState(414);
				match(CORDER);
				 _localctx.e = expressions.NewArray((((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getLine():0), (((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getCharPositionInLine():0), nil) 
				}
				break;
			case 6:
				{
				setState(416);
				((ExprContext)_localctx).list = listArray(0);
				 _localctx.e = ((ExprContext)_localctx).list.p
				}
				break;
			case 7:
				{
				setState(419);
				((ExprContext)_localctx).CORIZQ = match(CORIZQ);
				setState(420);
				((ExprContext)_localctx).listParams = listParams(0);
				setState(421);
				match(CORDER);
				 _localctx.e = expressions.NewArray((((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getLine():0), (((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getCharPositionInLine():0), ((ExprContext)_localctx).listParams.l) 
				}
				break;
			case 8:
				{
				setState(424);
				((ExprContext)_localctx).NUMBER = match(NUMBER);

				        if (strings.Contains((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null),".")){
				            num,err := strconv.ParseFloat((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null), 64);
				            if err!= nil{
				                fmt.Println(err)
				            }
				            _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getLine():0),(((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getCharPositionInLine():0),num,environment.FLOAT)
				        }else{
				            num,err := strconv.Atoi((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null))
				            if err!= nil{
				                fmt.Println(err)
				            }
				            _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getLine():0),(((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getCharPositionInLine():0),num,environment.INTEGER)
				        }
				    
				}
				break;
			case 9:
				{
				setState(426);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case 10:
				{
				setState(428);
				((ExprContext)_localctx).TRU = match(TRU);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getLine():0), (((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case 11:
				{
				setState(430);
				((ExprContext)_localctx).FAL = match(FAL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getLine():0), (((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case 12:
				{
				setState(432);
				((ExprContext)_localctx).ID = match(ID);
				setState(433);
				match(PUNTO);
				setState(434);
				match(COUNT);
				 _localctx.e = expressions.NewCount((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null)) 
				}
				break;
			case 13:
				{
				setState(436);
				((ExprContext)_localctx).NIL = match(NIL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getLine():0), (((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getCharPositionInLine():0), "nil", environment.NIL) 
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(489);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(487);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.expuno = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(440);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(441);
						match(PUNTO);
						setState(442);
						match(PUNTO);
						setState(443);
						match(PUNTO);
						setState(444);
						((ExprContext)_localctx).expdos = ((ExprContext)_localctx).expr = expr(22);
						 _localctx.e = expressions.NewRange((((ExprContext)_localctx).expuno!=null?(((ExprContext)_localctx).expuno.start):null).GetLine(),(((ExprContext)_localctx).expuno!=null?(((ExprContext)_localctx).expuno.start):null).GetColumn(), ((ExprContext)_localctx).expuno.e, ((ExprContext)_localctx).expdos.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(447);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(448);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==SUM_IG || _la==SUB_IG) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(449);
						((ExprContext)_localctx).expr = expr(21);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getLine():0), (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getCharPositionInLine():0), nil, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).expr.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(452);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(453);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 42784196460019712L) != 0)) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(454);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(20);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(457);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(458);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==ADD || _la==SUB) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(459);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(19);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(462);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(463);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MAY_IG || _la==MAYOR) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(464);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(18);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(467);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(468);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MEN_IG || _la==MENOR) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(469);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(17);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(472);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(473);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==DIF || _la==IG_IG) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(474);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(16);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 8:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(477);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(478);
						((ExprContext)_localctx).op = match(AND);
						setState(479);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(15);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 9:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(482);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(483);
						((ExprContext)_localctx).op = match(OR);
						setState(484);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(14);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(491);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListParamsContext extends ParserRuleContext {
		public []interface{} l;
		public ListParamsContext list;
		public ExprContext expr;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public ListParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listParams; }
	}

	public final ListParamsContext listParams() throws RecognitionException {
		return listParams(0);
	}

	private ListParamsContext listParams(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListParamsContext _localctx = new ListParamsContext(_ctx, _parentState);
		ListParamsContext _prevctx = _localctx;
		int _startState = 44;
		enterRecursionRule(_localctx, 44, RULE_listParams, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(493);
			((ListParamsContext)_localctx).expr = expr(0);

			            _localctx.l = []interface{}{}
			            _localctx.l = append(_localctx.l, ((ListParamsContext)_localctx).expr.e)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(503);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,23,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListParamsContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listParams);
					setState(496);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(497);
					match(COMA);
					setState(498);
					((ListParamsContext)_localctx).expr = expr(0);

					                                          var arr []interface{}
					                                          arr = append(((ListParamsContext)_localctx).list.l, ((ListParamsContext)_localctx).expr.e)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(505);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,23,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListAccessArrayContext extends ParserRuleContext {
		public []interface{} l;
		public ListAccessArrayContext list;
		public ExprContext expr;
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public ListAccessArrayContext listAccessArray() {
			return getRuleContext(ListAccessArrayContext.class,0);
		}
		public ListAccessArrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listAccessArray; }
	}

	public final ListAccessArrayContext listAccessArray() throws RecognitionException {
		return listAccessArray(0);
	}

	private ListAccessArrayContext listAccessArray(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListAccessArrayContext _localctx = new ListAccessArrayContext(_ctx, _parentState);
		ListAccessArrayContext _prevctx = _localctx;
		int _startState = 46;
		enterRecursionRule(_localctx, 46, RULE_listAccessArray, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(507);
			match(CORIZQ);
			setState(508);
			((ListAccessArrayContext)_localctx).expr = expr(0);
			setState(509);
			match(CORDER);

			            _localctx.l = []interface{}{}
			            _localctx.l = append(_localctx.l, ((ListAccessArrayContext)_localctx).expr.e)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(520);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListAccessArrayContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listAccessArray);
					setState(512);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(513);
					match(CORIZQ);
					setState(514);
					((ListAccessArrayContext)_localctx).expr = expr(0);
					setState(515);
					match(CORDER);

					                                          var arr []interface{}
					                                          arr = append(((ListAccessArrayContext)_localctx).list.l, ((ListAccessArrayContext)_localctx).expr.e)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(522);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListArrayContext extends ParserRuleContext {
		public interfaces.Expression p;
		public ListArrayContext list;
		public Token ID;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public ListArrayContext listArray() {
			return getRuleContext(ListArrayContext.class,0);
		}
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ListArrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listArray; }
	}

	public final ListArrayContext listArray() throws RecognitionException {
		return listArray(0);
	}

	private ListArrayContext listArray(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListArrayContext _localctx = new ListArrayContext(_ctx, _parentState);
		ListArrayContext _prevctx = _localctx;
		int _startState = 48;
		enterRecursionRule(_localctx, 48, RULE_listArray, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(524);
			((ListArrayContext)_localctx).ID = match(ID);
			 _localctx.p = expressions.NewCallVar((((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getCharPositionInLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getText():null))
			}
			_ctx.stop = _input.LT(-1);
			setState(543);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(541);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
					case 1:
						{
						_localctx = new ListArrayContext(_parentctx, _parentState);
						_localctx.list = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_listArray);
						setState(527);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(528);
						match(CORIZQ);
						setState(529);
						((ListArrayContext)_localctx).expr = expr(0);
						setState(530);
						match(CORDER);
						 _localctx.p = expressions.NewArrayAccess((((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetLine(), (((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetColumn(), ((ListArrayContext)_localctx).list.p, ((ListArrayContext)_localctx).expr.e) 
						}
						break;
					case 2:
						{
						_localctx = new ListArrayContext(_parentctx, _parentState);
						_localctx.list = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_listArray);
						setState(533);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(534);
						types();
						setState(535);
						match(IG);
						setState(536);
						match(CORIZQ);
						setState(537);
						((ListArrayContext)_localctx).expr = expr(0);
						setState(538);
						match(CORDER);
						 _localctx.p = expressions.NewArrayAccess((((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetLine(), (((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetColumn(), ((ListArrayContext)_localctx).list.p, ((ListArrayContext)_localctx).expr.e) 
						}
						break;
					}
					} 
				}
				setState(545);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprComaContext extends ParserRuleContext {
		public interfaces.Expression t;
		public ExprComaContext left;
		public ExprContext expr;
		public Token op;
		public ExprContext right;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ExprComaContext exprComa() {
			return getRuleContext(ExprComaContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ExprComaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exprComa; }
	}

	public final ExprComaContext exprComa() throws RecognitionException {
		return exprComa(0);
	}

	private ExprComaContext exprComa(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprComaContext _localctx = new ExprComaContext(_ctx, _parentState);
		ExprComaContext _prevctx = _localctx;
		int _startState = 50;
		enterRecursionRule(_localctx, 50, RULE_exprComa, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(547);
			((ExprComaContext)_localctx).expr = expr(0);
			 _localctx.t = ((ExprComaContext)_localctx).expr.e 
			}
			_ctx.stop = _input.LT(-1);
			setState(557);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ExprComaContext(_parentctx, _parentState);
					_localctx.left = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_exprComa);
					setState(550);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(551);
					((ExprComaContext)_localctx).op = match(COMA);
					setState(552);
					((ExprComaContext)_localctx).right = ((ExprComaContext)_localctx).expr = expr(0);
					 _localctx.t = expressions.NewOperation((((ExprComaContext)_localctx).left!=null?(((ExprComaContext)_localctx).left.start):null).GetLine(), (((ExprComaContext)_localctx).left!=null?(((ExprComaContext)_localctx).left.start):null).GetColumn(), ((ExprComaContext)_localctx).left.t, (((ExprComaContext)_localctx).op!=null?((ExprComaContext)_localctx).op.getText():null), ((ExprComaContext)_localctx).right.e) 
					}
					} 
				}
				setState(559);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 13:
			return listStructDec_sempred((ListStructDecContext)_localctx, predIndex);
		case 15:
			return listParamsFunc_sempred((ListParamsFuncContext)_localctx, predIndex);
		case 19:
			return listParamsCall_sempred((ListParamsCallContext)_localctx, predIndex);
		case 21:
			return expr_sempred((ExprContext)_localctx, predIndex);
		case 22:
			return listParams_sempred((ListParamsContext)_localctx, predIndex);
		case 23:
			return listAccessArray_sempred((ListAccessArrayContext)_localctx, predIndex);
		case 24:
			return listArray_sempred((ListArrayContext)_localctx, predIndex);
		case 25:
			return exprComa_sempred((ExprComaContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean listStructDec_sempred(ListStructDecContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 6);
		case 1:
			return precpred(_ctx, 5);
		case 2:
			return precpred(_ctx, 4);
		}
		return true;
	}
	private boolean listParamsFunc_sempred(ListParamsFuncContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean listParamsCall_sempred(ListParamsCallContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 5:
			return precpred(_ctx, 21);
		case 6:
			return precpred(_ctx, 20);
		case 7:
			return precpred(_ctx, 19);
		case 8:
			return precpred(_ctx, 18);
		case 9:
			return precpred(_ctx, 17);
		case 10:
			return precpred(_ctx, 16);
		case 11:
			return precpred(_ctx, 15);
		case 12:
			return precpred(_ctx, 14);
		case 13:
			return precpred(_ctx, 13);
		}
		return true;
	}
	private boolean listParams_sempred(ListParamsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 14:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listAccessArray_sempred(ListAccessArrayContext _localctx, int predIndex) {
		switch (predIndex) {
		case 15:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listArray_sempred(ListArrayContext _localctx, int predIndex) {
		switch (predIndex) {
		case 16:
			return precpred(_ctx, 3);
		case 17:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean exprComa_sempred(ExprComaContext _localctx, int predIndex) {
		switch (predIndex) {
		case 18:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001G\u0231\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0001\u0004\u0001:\b\u0001\u000b\u0001\f\u0001;\u0001\u0001\u0001"+
		"\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0003\u0002d\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0003\u0003r\b\u0003\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0003\u0004\u008f\b\u0004\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0003\u0006\u00c0\b\u0006\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0003\u0007\u00cc\b\u0007\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001"+
		"\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\f"+
		"\u0001\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0003\r\u00f9"+
		"\b\r\u0001\r\u0001\r\u0003\r\u00fd\b\r\u0001\r\u0001\r\u0001\r\u0001\r"+
		"\u0001\r\u0001\r\u0001\r\u0001\r\u0003\r\u0107\b\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0003\r\u0110\b\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0003\r\u0116\b\r\u0001\r\u0001\r\u0001\r\u0001\r\u0005\r\u011c"+
		"\b\r\n\r\f\r\u011f\t\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0003\u000e\u0137\b\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0003\u000f\u013e\b\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0005\u000f\u0145\b\u000f\n\u000f\f\u000f\u0148"+
		"\t\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0003"+
		"\u0010\u0155\b\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0003\u0013\u0168\b\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0005\u0013\u016f\b\u0013\n\u0013\f\u0013\u0172\t\u0013"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0003\u0014"+
		"\u018b\b\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0003\u0015\u01b7\b\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0005\u0015\u01e8\b\u0015\n\u0015\f\u0015\u01eb"+
		"\t\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0005\u0016\u01f6\b\u0016\n"+
		"\u0016\f\u0016\u01f9\t\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0005\u0017\u0207\b\u0017\n\u0017\f\u0017"+
		"\u020a\t\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0005\u0018\u021e\b\u0018\n\u0018\f\u0018\u0221\t\u0018\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0005\u0019\u022c\b\u0019\n\u0019\f\u0019"+
		"\u022f\t\u0019\u0001\u0019\u0000\b\u001a\u001e&*,.02\u001a\u0000\u0002"+
		"\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e"+
		" \"$&(*,.02\u0000\b\u0001\u0000/0\u0001\u0000\u0006\u0007\u0002\u0000"+
		"&&CC\u0002\u00003477\u0001\u000056\u0002\u0000--11\u0002\u0000..22\u0001"+
		"\u0000\'(\u025c\u00004\u0001\u0000\u0000\u0000\u00029\u0001\u0000\u0000"+
		"\u0000\u0004c\u0001\u0000\u0000\u0000\u0006q\u0001\u0000\u0000\u0000\b"+
		"\u008e\u0001\u0000\u0000\u0000\n\u0090\u0001\u0000\u0000\u0000\f\u00bf"+
		"\u0001\u0000\u0000\u0000\u000e\u00cb\u0001\u0000\u0000\u0000\u0010\u00cd"+
		"\u0001\u0000\u0000\u0000\u0012\u00d6\u0001\u0000\u0000\u0000\u0014\u00de"+
		"\u0001\u0000\u0000\u0000\u0016\u00e1\u0001\u0000\u0000\u0000\u0018\u00e4"+
		"\u0001\u0000\u0000\u0000\u001a\u00f8\u0001\u0000\u0000\u0000\u001c\u0136"+
		"\u0001\u0000\u0000\u0000\u001e\u013d\u0001\u0000\u0000\u0000 \u0154\u0001"+
		"\u0000\u0000\u0000\"\u0156\u0001\u0000\u0000\u0000$\u015c\u0001\u0000"+
		"\u0000\u0000&\u0167\u0001\u0000\u0000\u0000(\u018a\u0001\u0000\u0000\u0000"+
		"*\u01b6\u0001\u0000\u0000\u0000,\u01ec\u0001\u0000\u0000\u0000.\u01fa"+
		"\u0001\u0000\u0000\u00000\u020b\u0001\u0000\u0000\u00002\u0222\u0001\u0000"+
		"\u0000\u000045\u0003\u0002\u0001\u000056\u0005\u0000\u0000\u000167\u0006"+
		"\u0000\uffff\uffff\u00007\u0001\u0001\u0000\u0000\u00008:\u0003\u0004"+
		"\u0002\u000098\u0001\u0000\u0000\u0000:;\u0001\u0000\u0000\u0000;9\u0001"+
		"\u0000\u0000\u0000;<\u0001\u0000\u0000\u0000<=\u0001\u0000\u0000\u0000"+
		"=>\u0006\u0001\uffff\uffff\u0000>\u0003\u0001\u0000\u0000\u0000?@\u0003"+
		"\u0006\u0003\u0000@A\u0006\u0002\uffff\uffff\u0000Ad\u0001\u0000\u0000"+
		"\u0000BC\u0003\b\u0004\u0000CD\u0006\u0002\uffff\uffff\u0000Dd\u0001\u0000"+
		"\u0000\u0000EF\u0003\f\u0006\u0000FG\u0006\u0002\uffff\uffff\u0000Gd\u0001"+
		"\u0000\u0000\u0000HI\u0003\n\u0005\u0000IJ\u0006\u0002\uffff\uffff\u0000"+
		"Jd\u0001\u0000\u0000\u0000KL\u0003\u000e\u0007\u0000LM\u0006\u0002\uffff"+
		"\uffff\u0000Md\u0001\u0000\u0000\u0000NO\u0003\u0010\b\u0000OP\u0006\u0002"+
		"\uffff\uffff\u0000Pd\u0001\u0000\u0000\u0000QR\u0003\u0012\t\u0000RS\u0006"+
		"\u0002\uffff\uffff\u0000Sd\u0001\u0000\u0000\u0000TU\u0003\u0014\n\u0000"+
		"UV\u0006\u0002\uffff\uffff\u0000Vd\u0001\u0000\u0000\u0000WX\u0003\u0016"+
		"\u000b\u0000XY\u0006\u0002\uffff\uffff\u0000Yd\u0001\u0000\u0000\u0000"+
		"Z[\u0003\u0018\f\u0000[\\\u0006\u0002\uffff\uffff\u0000\\d\u0001\u0000"+
		"\u0000\u0000]^\u0003\u001c\u000e\u0000^_\u0006\u0002\uffff\uffff\u0000"+
		"_d\u0001\u0000\u0000\u0000`a\u0003$\u0012\u0000ab\u0006\u0002\uffff\uffff"+
		"\u0000bd\u0001\u0000\u0000\u0000c?\u0001\u0000\u0000\u0000cB\u0001\u0000"+
		"\u0000\u0000cE\u0001\u0000\u0000\u0000cH\u0001\u0000\u0000\u0000cK\u0001"+
		"\u0000\u0000\u0000cN\u0001\u0000\u0000\u0000cQ\u0001\u0000\u0000\u0000"+
		"cT\u0001\u0000\u0000\u0000cW\u0001\u0000\u0000\u0000cZ\u0001\u0000\u0000"+
		"\u0000c]\u0001\u0000\u0000\u0000c`\u0001\u0000\u0000\u0000d\u0005\u0001"+
		"\u0000\u0000\u0000ef\u0005\u000b\u0000\u0000fg\u00058\u0000\u0000gh\u0003"+
		"*\u0015\u0000hi\u00059\u0000\u0000ij\u0006\u0003\uffff\uffff\u0000jr\u0001"+
		"\u0000\u0000\u0000kl\u0005\u000b\u0000\u0000lm\u00058\u0000\u0000mn\u0003"+
		"2\u0019\u0000no\u00059\u0000\u0000op\u0006\u0003\uffff\uffff\u0000pr\u0001"+
		"\u0000\u0000\u0000qe\u0001\u0000\u0000\u0000qk\u0001\u0000\u0000\u0000"+
		"r\u0007\u0001\u0000\u0000\u0000st\u0005\f\u0000\u0000tu\u0003*\u0015\u0000"+
		"uv\u0005:\u0000\u0000vw\u0003\u0002\u0001\u0000wx\u0005;\u0000\u0000x"+
		"y\u0006\u0004\uffff\uffff\u0000y\u008f\u0001\u0000\u0000\u0000z{\u0005"+
		"\f\u0000\u0000{|\u0003*\u0015\u0000|}\u0005:\u0000\u0000}~\u0003\u0002"+
		"\u0001\u0000~\u007f\u0005;\u0000\u0000\u007f\u0080\u0005\r\u0000\u0000"+
		"\u0080\u0081\u0005:\u0000\u0000\u0081\u0082\u0003\u0002\u0001\u0000\u0082"+
		"\u0083\u0005;\u0000\u0000\u0083\u0084\u0006\u0004\uffff\uffff\u0000\u0084"+
		"\u008f\u0001\u0000\u0000\u0000\u0085\u0086\u0005\f\u0000\u0000\u0086\u0087"+
		"\u0003*\u0015\u0000\u0087\u0088\u0005:\u0000\u0000\u0088\u0089\u0003\u0002"+
		"\u0001\u0000\u0089\u008a\u0005;\u0000\u0000\u008a\u008b\u0005\r\u0000"+
		"\u0000\u008b\u008c\u0003\b\u0004\u0000\u008c\u008d\u0006\u0004\uffff\uffff"+
		"\u0000\u008d\u008f\u0001\u0000\u0000\u0000\u008es\u0001\u0000\u0000\u0000"+
		"\u008ez\u0001\u0000\u0000\u0000\u008e\u0085\u0001\u0000\u0000\u0000\u008f"+
		"\t\u0001\u0000\u0000\u0000\u0090\u0091\u0005\u000e\u0000\u0000\u0091\u0092"+
		"\u0003*\u0015\u0000\u0092\u0093\u0005:\u0000\u0000\u0093\u0094\u0003\u0002"+
		"\u0001\u0000\u0094\u0095\u0005;\u0000\u0000\u0095\u0096\u0006\u0005\uffff"+
		"\uffff\u0000\u0096\u000b\u0001\u0000\u0000\u0000\u0097\u0098\u0005\u0006"+
		"\u0000\u0000\u0098\u0099\u0005&\u0000\u0000\u0099\u009a\u0005<\u0000\u0000"+
		"\u009a\u009b\u0003(\u0014\u0000\u009b\u009c\u0005,\u0000\u0000\u009c\u009d"+
		"\u0003*\u0015\u0000\u009d\u009e\u0006\u0006\uffff\uffff\u0000\u009e\u00c0"+
		"\u0001\u0000\u0000\u0000\u009f\u00a0\u0005\u0006\u0000\u0000\u00a0\u00a1"+
		"\u0005&\u0000\u0000\u00a1\u00a2\u0005,\u0000\u0000\u00a2\u00a3\u0003*"+
		"\u0015\u0000\u00a3\u00a4\u0006\u0006\uffff\uffff\u0000\u00a4\u00c0\u0001"+
		"\u0000\u0000\u0000\u00a5\u00a6\u0005\u0006\u0000\u0000\u00a6\u00a7\u0005"+
		"&\u0000\u0000\u00a7\u00a8\u0005<\u0000\u0000\u00a8\u00a9\u0003(\u0014"+
		"\u0000\u00a9\u00aa\u0006\u0006\uffff\uffff\u0000\u00aa\u00c0\u0001\u0000"+
		"\u0000\u0000\u00ab\u00ac\u0005\u0007\u0000\u0000\u00ac\u00ad\u0005&\u0000"+
		"\u0000\u00ad\u00ae\u0005<\u0000\u0000\u00ae\u00af\u0003(\u0014\u0000\u00af"+
		"\u00b0\u0005,\u0000\u0000\u00b0\u00b1\u0003*\u0015\u0000\u00b1\u00b2\u0006"+
		"\u0006\uffff\uffff\u0000\u00b2\u00c0\u0001\u0000\u0000\u0000\u00b3\u00b4"+
		"\u0005\u0007\u0000\u0000\u00b4\u00b5\u0005&\u0000\u0000\u00b5\u00b6\u0005"+
		"<\u0000\u0000\u00b6\u00b7\u0003(\u0014\u0000\u00b7\u00b8\u0006\u0006\uffff"+
		"\uffff\u0000\u00b8\u00c0\u0001\u0000\u0000\u0000\u00b9\u00ba\u0005\u0007"+
		"\u0000\u0000\u00ba\u00bb\u0005&\u0000\u0000\u00bb\u00bc\u0005,\u0000\u0000"+
		"\u00bc\u00bd\u0003*\u0015\u0000\u00bd\u00be\u0006\u0006\uffff\uffff\u0000"+
		"\u00be\u00c0\u0001\u0000\u0000\u0000\u00bf\u0097\u0001\u0000\u0000\u0000"+
		"\u00bf\u009f\u0001\u0000\u0000\u0000\u00bf\u00a5\u0001\u0000\u0000\u0000"+
		"\u00bf\u00ab\u0001\u0000\u0000\u0000\u00bf\u00b3\u0001\u0000\u0000\u0000"+
		"\u00bf\u00b9\u0001\u0000\u0000\u0000\u00c0\r\u0001\u0000\u0000\u0000\u00c1"+
		"\u00c2\u0005&\u0000\u0000\u00c2\u00c3\u0005,\u0000\u0000\u00c3\u00c4\u0003"+
		"*\u0015\u0000\u00c4\u00c5\u0006\u0007\uffff\uffff\u0000\u00c5\u00cc\u0001"+
		"\u0000\u0000\u0000\u00c6\u00c7\u0005&\u0000\u0000\u00c7\u00c8\u0007\u0000"+
		"\u0000\u0000\u00c8\u00c9\u0003*\u0015\u0000\u00c9\u00ca\u0006\u0007\uffff"+
		"\uffff\u0000\u00ca\u00cc\u0001\u0000\u0000\u0000\u00cb\u00c1\u0001\u0000"+
		"\u0000\u0000\u00cb\u00c6\u0001\u0000\u0000\u0000\u00cc\u000f\u0001\u0000"+
		"\u0000\u0000\u00cd\u00ce\u0005\u000f\u0000\u0000\u00ce\u00cf\u0005&\u0000"+
		"\u0000\u00cf\u00d0\u0005\u0010\u0000\u0000\u00d0\u00d1\u0003*\u0015\u0000"+
		"\u00d1\u00d2\u0005:\u0000\u0000\u00d2\u00d3\u0003\u0002\u0001\u0000\u00d3"+
		"\u00d4\u0005;\u0000\u0000\u00d4\u00d5\u0006\b\uffff\uffff\u0000\u00d5"+
		"\u0011\u0001\u0000\u0000\u0000\u00d6\u00d7\u0005\u0017\u0000\u0000\u00d7"+
		"\u00d8\u0003*\u0015\u0000\u00d8\u00d9\u0005\r\u0000\u0000\u00d9\u00da"+
		"\u0005:\u0000\u0000\u00da\u00db\u0003\u0002\u0001\u0000\u00db\u00dc\u0005"+
		";\u0000\u0000\u00dc\u00dd\u0006\t\uffff\uffff\u0000\u00dd\u0013\u0001"+
		"\u0000\u0000\u0000\u00de\u00df\u0005\u0014\u0000\u0000\u00df\u00e0\u0006"+
		"\n\uffff\uffff\u0000\u00e0\u0015\u0001\u0000\u0000\u0000\u00e1\u00e2\u0005"+
		"\u0016\u0000\u0000\u00e2\u00e3\u0006\u000b\uffff\uffff\u0000\u00e3\u0017"+
		"\u0001\u0000\u0000\u0000\u00e4\u00e5\u0005\u001a\u0000\u0000\u00e5\u00e6"+
		"\u0005&\u0000\u0000\u00e6\u00e7\u0005:\u0000\u0000\u00e7\u00e8\u0003\u001a"+
		"\r\u0000\u00e8\u00e9\u0005;\u0000\u0000\u00e9\u00ea\u0006\f\uffff\uffff"+
		"\u0000\u00ea\u0019\u0001\u0000\u0000\u0000\u00eb\u00ec\u0006\r\uffff\uffff"+
		"\u0000\u00ec\u00ed\u0007\u0001\u0000\u0000\u00ed\u00ee\u0005&\u0000\u0000"+
		"\u00ee\u00ef\u0005<\u0000\u0000\u00ef\u00f0\u0003(\u0014\u0000\u00f0\u00f1"+
		"\u0006\r\uffff\uffff\u0000\u00f1\u00f9\u0001\u0000\u0000\u0000\u00f2\u00f3"+
		"\u0007\u0001\u0000\u0000\u00f3\u00f4\u0005&\u0000\u0000\u00f4\u00f5\u0005"+
		"<\u0000\u0000\u00f5\u00f6\u0005&\u0000\u0000\u00f6\u00f9\u0006\r\uffff"+
		"\uffff\u0000\u00f7\u00f9\u0006\r\uffff\uffff\u0000\u00f8\u00eb\u0001\u0000"+
		"\u0000\u0000\u00f8\u00f2\u0001\u0000\u0000\u0000\u00f8\u00f7\u0001\u0000"+
		"\u0000\u0000\u00f9\u011d\u0001\u0000\u0000\u0000\u00fa\u00fc\n\u0006\u0000"+
		"\u0000\u00fb\u00fd\u0005?\u0000\u0000\u00fc\u00fb\u0001\u0000\u0000\u0000"+
		"\u00fc\u00fd\u0001\u0000\u0000\u0000\u00fd\u00fe\u0001\u0000\u0000\u0000"+
		"\u00fe\u00ff\u0007\u0001\u0000\u0000\u00ff\u0100\u0005&\u0000\u0000\u0100"+
		"\u0101\u0005<\u0000\u0000\u0101\u0102\u0003(\u0014\u0000\u0102\u0103\u0006"+
		"\r\uffff\uffff\u0000\u0103\u011c\u0001\u0000\u0000\u0000\u0104\u0106\n"+
		"\u0005\u0000\u0000\u0105\u0107\u0005?\u0000\u0000\u0106\u0105\u0001\u0000"+
		"\u0000\u0000\u0106\u0107\u0001\u0000\u0000\u0000\u0107\u0108\u0001\u0000"+
		"\u0000\u0000\u0108\u0109\u0007\u0001\u0000\u0000\u0109\u010a\u0005&\u0000"+
		"\u0000\u010a\u010b\u0005<\u0000\u0000\u010b\u010c\u0005&\u0000\u0000\u010c"+
		"\u011c\u0006\r\uffff\uffff\u0000\u010d\u010f\n\u0004\u0000\u0000\u010e"+
		"\u0110\u0005?\u0000\u0000\u010f\u010e\u0001\u0000\u0000\u0000\u010f\u0110"+
		"\u0001\u0000\u0000\u0000\u0110\u0111\u0001\u0000\u0000\u0000\u0111\u0112"+
		"\u0007\u0001\u0000\u0000\u0112\u0115\u0005&\u0000\u0000\u0113\u0114\u0005"+
		"<\u0000\u0000\u0114\u0116\u0003(\u0014\u0000\u0115\u0113\u0001\u0000\u0000"+
		"\u0000\u0115\u0116\u0001\u0000\u0000\u0000\u0116\u0117\u0001\u0000\u0000"+
		"\u0000\u0117\u0118\u0005,\u0000\u0000\u0118\u0119\u0003*\u0015\u0000\u0119"+
		"\u011a\u0006\r\uffff\uffff\u0000\u011a\u011c\u0001\u0000\u0000\u0000\u011b"+
		"\u00fa\u0001\u0000\u0000\u0000\u011b\u0104\u0001\u0000\u0000\u0000\u011b"+
		"\u010d\u0001\u0000\u0000\u0000\u011c\u011f\u0001\u0000\u0000\u0000\u011d"+
		"\u011b\u0001\u0000\u0000\u0000\u011d\u011e\u0001\u0000\u0000\u0000\u011e"+
		"\u001b\u0001\u0000\u0000\u0000\u011f\u011d\u0001\u0000\u0000\u0000\u0120"+
		"\u0121\u0005\u0018\u0000\u0000\u0121\u0122\u0005&\u0000\u0000\u0122\u0123"+
		"\u00058\u0000\u0000\u0123\u0124\u0003\u001e\u000f\u0000\u0124\u0125\u0005"+
		"9\u0000\u0000\u0125\u0126\u0005B\u0000\u0000\u0126\u0127\u0003(\u0014"+
		"\u0000\u0127\u0128\u0005:\u0000\u0000\u0128\u0129\u0003\u0002\u0001\u0000"+
		"\u0129\u012a\u0005;\u0000\u0000\u012a\u012b\u0006\u000e\uffff\uffff\u0000"+
		"\u012b\u0137\u0001\u0000\u0000\u0000\u012c\u012d\u0005\u0018\u0000\u0000"+
		"\u012d\u012e\u0005&\u0000\u0000\u012e\u012f\u00058\u0000\u0000\u012f\u0130"+
		"\u0003\u001e\u000f\u0000\u0130\u0131\u00059\u0000\u0000\u0131\u0132\u0005"+
		":\u0000\u0000\u0132\u0133\u0003\u0002\u0001\u0000\u0133\u0134\u0005;\u0000"+
		"\u0000\u0134\u0135\u0006\u000e\uffff\uffff\u0000\u0135\u0137\u0001\u0000"+
		"\u0000\u0000\u0136\u0120\u0001\u0000\u0000\u0000\u0136\u012c\u0001\u0000"+
		"\u0000\u0000\u0137\u001d\u0001\u0000\u0000\u0000\u0138\u0139\u0006\u000f"+
		"\uffff\uffff\u0000\u0139\u013a\u0003 \u0010\u0000\u013a\u013b\u0006\u000f"+
		"\uffff\uffff\u0000\u013b\u013e\u0001\u0000\u0000\u0000\u013c\u013e\u0006"+
		"\u000f\uffff\uffff\u0000\u013d\u0138\u0001\u0000\u0000\u0000\u013d\u013c"+
		"\u0001\u0000\u0000\u0000\u013e\u0146\u0001\u0000\u0000\u0000\u013f\u0140"+
		"\n\u0003\u0000\u0000\u0140\u0141\u0005?\u0000\u0000\u0141\u0142\u0003"+
		" \u0010\u0000\u0142\u0143\u0006\u000f\uffff\uffff\u0000\u0143\u0145\u0001"+
		"\u0000\u0000\u0000\u0144\u013f\u0001\u0000\u0000\u0000\u0145\u0148\u0001"+
		"\u0000\u0000\u0000\u0146\u0144\u0001\u0000\u0000\u0000\u0146\u0147\u0001"+
		"\u0000\u0000\u0000\u0147\u001f\u0001\u0000\u0000\u0000\u0148\u0146\u0001"+
		"\u0000\u0000\u0000\u0149\u014a\u0005&\u0000\u0000\u014a\u014b\u0005<\u0000"+
		"\u0000\u014b\u014c\u0003(\u0014\u0000\u014c\u014d\u0006\u0010\uffff\uffff"+
		"\u0000\u014d\u0155\u0001\u0000\u0000\u0000\u014e\u014f\u0007\u0002\u0000"+
		"\u0000\u014f\u0150\u0005&\u0000\u0000\u0150\u0151\u0005<\u0000\u0000\u0151"+
		"\u0152\u0003(\u0014\u0000\u0152\u0153\u0006\u0010\uffff\uffff\u0000\u0153"+
		"\u0155\u0001\u0000\u0000\u0000\u0154\u0149\u0001\u0000\u0000\u0000\u0154"+
		"\u014e\u0001\u0000\u0000\u0000\u0155!\u0001\u0000\u0000\u0000\u0156\u0157"+
		"\u0005&\u0000\u0000\u0157\u0158\u00058\u0000\u0000\u0158\u0159\u0003&"+
		"\u0013\u0000\u0159\u015a\u00059\u0000\u0000\u015a\u015b\u0006\u0011\uffff"+
		"\uffff\u0000\u015b#\u0001\u0000\u0000\u0000\u015c\u015d\u0005&\u0000\u0000"+
		"\u015d\u015e\u00058\u0000\u0000\u015e\u015f\u0003&\u0013\u0000\u015f\u0160"+
		"\u00059\u0000\u0000\u0160\u0161\u0006\u0012\uffff\uffff\u0000\u0161%\u0001"+
		"\u0000\u0000\u0000\u0162\u0163\u0006\u0013\uffff\uffff\u0000\u0163\u0164"+
		"\u0003*\u0015\u0000\u0164\u0165\u0006\u0013\uffff\uffff\u0000\u0165\u0168"+
		"\u0001\u0000\u0000\u0000\u0166\u0168\u0006\u0013\uffff\uffff\u0000\u0167"+
		"\u0162\u0001\u0000\u0000\u0000\u0167\u0166\u0001\u0000\u0000\u0000\u0168"+
		"\u0170\u0001\u0000\u0000\u0000\u0169\u016a\n\u0003\u0000\u0000\u016a\u016b"+
		"\u0005?\u0000\u0000\u016b\u016c\u0003*\u0015\u0000\u016c\u016d\u0006\u0013"+
		"\uffff\uffff\u0000\u016d\u016f\u0001\u0000\u0000\u0000\u016e\u0169\u0001"+
		"\u0000\u0000\u0000\u016f\u0172\u0001\u0000\u0000\u0000\u0170\u016e\u0001"+
		"\u0000\u0000\u0000\u0170\u0171\u0001\u0000\u0000\u0000\u0171\'\u0001\u0000"+
		"\u0000\u0000\u0172\u0170\u0001\u0000\u0000\u0000\u0173\u0174\u0005\u0001"+
		"\u0000\u0000\u0174\u018b\u0006\u0014\uffff\uffff\u0000\u0175\u0176\u0005"+
		"\u0002\u0000\u0000\u0176\u018b\u0006\u0014\uffff\uffff\u0000\u0177\u0178"+
		"\u0005\u0004\u0000\u0000\u0178\u018b\u0006\u0014\uffff\uffff\u0000\u0179"+
		"\u017a\u0005\u0003\u0000\u0000\u017a\u018b\u0006\u0014\uffff\uffff\u0000"+
		"\u017b\u017c\u0005=\u0000\u0000\u017c\u017d\u0003(\u0014\u0000\u017d\u017e"+
		"\u0005>\u0000\u0000\u017e\u017f\u0006\u0014\uffff\uffff\u0000\u017f\u018b"+
		"\u0001\u0000\u0000\u0000\u0180\u0181\u0005A\u0000\u0000\u0181\u0182\u0005"+
		"\u0004\u0000\u0000\u0182\u0183\u0005A\u0000\u0000\u0183\u018b\u0006\u0014"+
		"\uffff\uffff\u0000\u0184\u0185\u0005\u0019\u0000\u0000\u0185\u018b\u0006"+
		"\u0014\uffff\uffff\u0000\u0186\u0187\u0005\u001a\u0000\u0000\u0187\u018b"+
		"\u0006\u0014\uffff\uffff\u0000\u0188\u0189\u0005&\u0000\u0000\u0189\u018b"+
		"\u0006\u0014\uffff\uffff\u0000\u018a\u0173\u0001\u0000\u0000\u0000\u018a"+
		"\u0175\u0001\u0000\u0000\u0000\u018a\u0177\u0001\u0000\u0000\u0000\u018a"+
		"\u0179\u0001\u0000\u0000\u0000\u018a\u017b\u0001\u0000\u0000\u0000\u018a"+
		"\u0180\u0001\u0000\u0000\u0000\u018a\u0184\u0001\u0000\u0000\u0000\u018a"+
		"\u0186\u0001\u0000\u0000\u0000\u018a\u0188\u0001\u0000\u0000\u0000\u018b"+
		")\u0001\u0000\u0000\u0000\u018c\u018d\u0006\u0015\uffff\uffff\u0000\u018d"+
		"\u018e\u00056\u0000\u0000\u018e\u018f\u0003*\u0015\u0016\u018f\u0190\u0006"+
		"\u0015\uffff\uffff\u0000\u0190\u01b7\u0001\u0000\u0000\u0000\u0191\u0192"+
		"\u0005)\u0000\u0000\u0192\u0193\u0003*\u0015\f\u0193\u0194\u0006\u0015"+
		"\uffff\uffff\u0000\u0194\u01b7\u0001\u0000\u0000\u0000\u0195\u0196\u0005"+
		"8\u0000\u0000\u0196\u0197\u0003*\u0015\u0000\u0197\u0198\u00059\u0000"+
		"\u0000\u0198\u0199\u0006\u0015\uffff\uffff\u0000\u0199\u01b7\u0001\u0000"+
		"\u0000\u0000\u019a\u019b\u0003\"\u0011\u0000\u019b\u019c\u0006\u0015\uffff"+
		"\uffff\u0000\u019c\u01b7\u0001\u0000\u0000\u0000\u019d\u019e\u0005=\u0000"+
		"\u0000\u019e\u019f\u0005>\u0000\u0000\u019f\u01b7\u0006\u0015\uffff\uffff"+
		"\u0000\u01a0\u01a1\u00030\u0018\u0000\u01a1\u01a2\u0006\u0015\uffff\uffff"+
		"\u0000\u01a2\u01b7\u0001\u0000\u0000\u0000\u01a3\u01a4\u0005=\u0000\u0000"+
		"\u01a4\u01a5\u0003,\u0016\u0000\u01a5\u01a6\u0005>\u0000\u0000\u01a6\u01a7"+
		"\u0006\u0015\uffff\uffff\u0000\u01a7\u01b7\u0001\u0000\u0000\u0000\u01a8"+
		"\u01a9\u0005$\u0000\u0000\u01a9\u01b7\u0006\u0015\uffff\uffff\u0000\u01aa"+
		"\u01ab\u0005%\u0000\u0000\u01ab\u01b7\u0006\u0015\uffff\uffff\u0000\u01ac"+
		"\u01ad\u0005\t\u0000\u0000\u01ad\u01b7\u0006\u0015\uffff\uffff\u0000\u01ae"+
		"\u01af\u0005\n\u0000\u0000\u01af\u01b7\u0006\u0015\uffff\uffff\u0000\u01b0"+
		"\u01b1\u0005&\u0000\u0000\u01b1\u01b2\u0005@\u0000\u0000\u01b2\u01b3\u0005"+
		"#\u0000\u0000\u01b3\u01b7\u0006\u0015\uffff\uffff\u0000\u01b4\u01b5\u0005"+
		"\u0019\u0000\u0000\u01b5\u01b7\u0006\u0015\uffff\uffff\u0000\u01b6\u018c"+
		"\u0001\u0000\u0000\u0000\u01b6\u0191\u0001\u0000\u0000\u0000\u01b6\u0195"+
		"\u0001\u0000\u0000\u0000\u01b6\u019a\u0001\u0000\u0000\u0000\u01b6\u019d"+
		"\u0001\u0000\u0000\u0000\u01b6\u01a0\u0001\u0000\u0000\u0000\u01b6\u01a3"+
		"\u0001\u0000\u0000\u0000\u01b6\u01a8\u0001\u0000\u0000\u0000\u01b6\u01aa"+
		"\u0001\u0000\u0000\u0000\u01b6\u01ac\u0001\u0000\u0000\u0000\u01b6\u01ae"+
		"\u0001\u0000\u0000\u0000\u01b6\u01b0\u0001\u0000\u0000\u0000\u01b6\u01b4"+
		"\u0001\u0000\u0000\u0000\u01b7\u01e9\u0001\u0000\u0000\u0000\u01b8\u01b9"+
		"\n\u0015\u0000\u0000\u01b9\u01ba\u0005@\u0000\u0000\u01ba\u01bb\u0005"+
		"@\u0000\u0000\u01bb\u01bc\u0005@\u0000\u0000\u01bc\u01bd\u0003*\u0015"+
		"\u0016\u01bd\u01be\u0006\u0015\uffff\uffff\u0000\u01be\u01e8\u0001\u0000"+
		"\u0000\u0000\u01bf\u01c0\n\u0014\u0000\u0000\u01c0\u01c1\u0007\u0000\u0000"+
		"\u0000\u01c1\u01c2\u0003*\u0015\u0015\u01c2\u01c3\u0006\u0015\uffff\uffff"+
		"\u0000\u01c3\u01e8\u0001\u0000\u0000\u0000\u01c4\u01c5\n\u0013\u0000\u0000"+
		"\u01c5\u01c6\u0007\u0003\u0000\u0000\u01c6\u01c7\u0003*\u0015\u0014\u01c7"+
		"\u01c8\u0006\u0015\uffff\uffff\u0000\u01c8\u01e8\u0001\u0000\u0000\u0000"+
		"\u01c9\u01ca\n\u0012\u0000\u0000\u01ca\u01cb\u0007\u0004\u0000\u0000\u01cb"+
		"\u01cc\u0003*\u0015\u0013\u01cc\u01cd\u0006\u0015\uffff\uffff\u0000\u01cd"+
		"\u01e8\u0001\u0000\u0000\u0000\u01ce\u01cf\n\u0011\u0000\u0000\u01cf\u01d0"+
		"\u0007\u0005\u0000\u0000\u01d0\u01d1\u0003*\u0015\u0012\u01d1\u01d2\u0006"+
		"\u0015\uffff\uffff\u0000\u01d2\u01e8\u0001\u0000\u0000\u0000\u01d3\u01d4"+
		"\n\u0010\u0000\u0000\u01d4\u01d5\u0007\u0006\u0000\u0000\u01d5\u01d6\u0003"+
		"*\u0015\u0011\u01d6\u01d7\u0006\u0015\uffff\uffff\u0000\u01d7\u01e8\u0001"+
		"\u0000\u0000\u0000\u01d8\u01d9\n\u000f\u0000\u0000\u01d9\u01da\u0007\u0007"+
		"\u0000\u0000\u01da\u01db\u0003*\u0015\u0010\u01db\u01dc\u0006\u0015\uffff"+
		"\uffff\u0000\u01dc\u01e8\u0001\u0000\u0000\u0000\u01dd\u01de\n\u000e\u0000"+
		"\u0000\u01de\u01df\u0005+\u0000\u0000\u01df\u01e0\u0003*\u0015\u000f\u01e0"+
		"\u01e1\u0006\u0015\uffff\uffff\u0000\u01e1\u01e8\u0001\u0000\u0000\u0000"+
		"\u01e2\u01e3\n\r\u0000\u0000\u01e3\u01e4\u0005*\u0000\u0000\u01e4\u01e5"+
		"\u0003*\u0015\u000e\u01e5\u01e6\u0006\u0015\uffff\uffff\u0000\u01e6\u01e8"+
		"\u0001\u0000\u0000\u0000\u01e7\u01b8\u0001\u0000\u0000\u0000\u01e7\u01bf"+
		"\u0001\u0000\u0000\u0000\u01e7\u01c4\u0001\u0000\u0000\u0000\u01e7\u01c9"+
		"\u0001\u0000\u0000\u0000\u01e7\u01ce\u0001\u0000\u0000\u0000\u01e7\u01d3"+
		"\u0001\u0000\u0000\u0000\u01e7\u01d8\u0001\u0000\u0000\u0000\u01e7\u01dd"+
		"\u0001\u0000\u0000\u0000\u01e7\u01e2\u0001\u0000\u0000\u0000\u01e8\u01eb"+
		"\u0001\u0000\u0000\u0000\u01e9\u01e7\u0001\u0000\u0000\u0000\u01e9\u01ea"+
		"\u0001\u0000\u0000\u0000\u01ea+\u0001\u0000\u0000\u0000\u01eb\u01e9\u0001"+
		"\u0000\u0000\u0000\u01ec\u01ed\u0006\u0016\uffff\uffff\u0000\u01ed\u01ee"+
		"\u0003*\u0015\u0000\u01ee\u01ef\u0006\u0016\uffff\uffff\u0000\u01ef\u01f7"+
		"\u0001\u0000\u0000\u0000\u01f0\u01f1\n\u0002\u0000\u0000\u01f1\u01f2\u0005"+
		"?\u0000\u0000\u01f2\u01f3\u0003*\u0015\u0000\u01f3\u01f4\u0006\u0016\uffff"+
		"\uffff\u0000\u01f4\u01f6\u0001\u0000\u0000\u0000\u01f5\u01f0\u0001\u0000"+
		"\u0000\u0000\u01f6\u01f9\u0001\u0000\u0000\u0000\u01f7\u01f5\u0001\u0000"+
		"\u0000\u0000\u01f7\u01f8\u0001\u0000\u0000\u0000\u01f8-\u0001\u0000\u0000"+
		"\u0000\u01f9\u01f7\u0001\u0000\u0000\u0000\u01fa\u01fb\u0006\u0017\uffff"+
		"\uffff\u0000\u01fb\u01fc\u0005=\u0000\u0000\u01fc\u01fd\u0003*\u0015\u0000"+
		"\u01fd\u01fe\u0005>\u0000\u0000\u01fe\u01ff\u0006\u0017\uffff\uffff\u0000"+
		"\u01ff\u0208\u0001\u0000\u0000\u0000\u0200\u0201\n\u0002\u0000\u0000\u0201"+
		"\u0202\u0005=\u0000\u0000\u0202\u0203\u0003*\u0015\u0000\u0203\u0204\u0005"+
		">\u0000\u0000\u0204\u0205\u0006\u0017\uffff\uffff\u0000\u0205\u0207\u0001"+
		"\u0000\u0000\u0000\u0206\u0200\u0001\u0000\u0000\u0000\u0207\u020a\u0001"+
		"\u0000\u0000\u0000\u0208\u0206\u0001\u0000\u0000\u0000\u0208\u0209\u0001"+
		"\u0000\u0000\u0000\u0209/\u0001\u0000\u0000\u0000\u020a\u0208\u0001\u0000"+
		"\u0000\u0000\u020b\u020c\u0006\u0018\uffff\uffff\u0000\u020c\u020d\u0005"+
		"&\u0000\u0000\u020d\u020e\u0006\u0018\uffff\uffff\u0000\u020e\u021f\u0001"+
		"\u0000\u0000\u0000\u020f\u0210\n\u0003\u0000\u0000\u0210\u0211\u0005="+
		"\u0000\u0000\u0211\u0212\u0003*\u0015\u0000\u0212\u0213\u0005>\u0000\u0000"+
		"\u0213\u0214\u0006\u0018\uffff\uffff\u0000\u0214\u021e\u0001\u0000\u0000"+
		"\u0000\u0215\u0216\n\u0002\u0000\u0000\u0216\u0217\u0003(\u0014\u0000"+
		"\u0217\u0218\u0005,\u0000\u0000\u0218\u0219\u0005=\u0000\u0000\u0219\u021a"+
		"\u0003*\u0015\u0000\u021a\u021b\u0005>\u0000\u0000\u021b\u021c\u0006\u0018"+
		"\uffff\uffff\u0000\u021c\u021e\u0001\u0000\u0000\u0000\u021d\u020f\u0001"+
		"\u0000\u0000\u0000\u021d\u0215\u0001\u0000\u0000\u0000\u021e\u0221\u0001"+
		"\u0000\u0000\u0000\u021f\u021d\u0001\u0000\u0000\u0000\u021f\u0220\u0001"+
		"\u0000\u0000\u0000\u02201\u0001\u0000\u0000\u0000\u0221\u021f\u0001\u0000"+
		"\u0000\u0000\u0222\u0223\u0006\u0019\uffff\uffff\u0000\u0223\u0224\u0003"+
		"*\u0015\u0000\u0224\u0225\u0006\u0019\uffff\uffff\u0000\u0225\u022d\u0001"+
		"\u0000\u0000\u0000\u0226\u0227\n\u0002\u0000\u0000\u0227\u0228\u0005?"+
		"\u0000\u0000\u0228\u0229\u0003*\u0015\u0000\u0229\u022a\u0006\u0019\uffff"+
		"\uffff\u0000\u022a\u022c\u0001\u0000\u0000\u0000\u022b\u0226\u0001\u0000"+
		"\u0000\u0000\u022c\u022f\u0001\u0000\u0000\u0000\u022d\u022b\u0001\u0000"+
		"\u0000\u0000\u022d\u022e\u0001\u0000\u0000\u0000\u022e3\u0001\u0000\u0000"+
		"\u0000\u022f\u022d\u0001\u0000\u0000\u0000\u001c;cq\u008e\u00bf\u00cb"+
		"\u00f8\u00fc\u0106\u010f\u0115\u011b\u011d\u0136\u013d\u0146\u0154\u0167"+
		"\u0170\u018a\u01b6\u01e7\u01e9\u01f7\u0208\u021d\u021f\u022d";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}