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
		RULE_listArray = 24, RULE_count = 25, RULE_exprComa = 26;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "printstmt", "ifstmt", "whilestmt", "declarationstmt", 
			"assignstmt", "forstmt", "guardstmt", "breakstmt", "continuestmt", "structCreation", 
			"listStructDec", "fnstmt", "listParamsFunc", "parametro", "callExp", 
			"callFunction", "listParamsCall", "types", "expr", "listParams", "listAccessArray", 
			"listArray", "count", "exprComa"
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
			setState(54);
			((SContext)_localctx).block = block();
			setState(55);
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
			setState(59); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(58);
				((BlockContext)_localctx).instruction = instruction();
				((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
				}
				}
				setState(61); 
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
			setState(101);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(65);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(68);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ifinst 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(71);
				((InstructionContext)_localctx).declarationstmt = declarationstmt();
				 _localctx.inst = ((InstructionContext)_localctx).declarationstmt.dec 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(74);
				((InstructionContext)_localctx).whilestmt = whilestmt();
				 _localctx.inst = ((InstructionContext)_localctx).whilestmt.whl 
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(77);
				((InstructionContext)_localctx).assignstmt = assignstmt();
				 _localctx.inst = ((InstructionContext)_localctx).assignstmt.asg 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(80);
				((InstructionContext)_localctx).forstmt = forstmt();
				 _localctx.inst = ((InstructionContext)_localctx).forstmt.fr 
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(83);
				((InstructionContext)_localctx).guardstmt = guardstmt();
				 _localctx.inst = ((InstructionContext)_localctx).guardstmt.grd 
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(86);
				((InstructionContext)_localctx).breakstmt = breakstmt();
				 _localctx.inst = ((InstructionContext)_localctx).breakstmt.brk 
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(89);
				((InstructionContext)_localctx).continuestmt = continuestmt();
				 _localctx.inst = ((InstructionContext)_localctx).continuestmt.cnt 
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(92);
				((InstructionContext)_localctx).structCreation = structCreation();
				 _localctx.inst = ((InstructionContext)_localctx).structCreation.dec 
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(95);
				((InstructionContext)_localctx).fnstmt = fnstmt();
				 _localctx.inst = ((InstructionContext)_localctx).fnstmt.fn 
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(98);
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
			setState(115);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(103);
				((PrintstmtContext)_localctx).PRINT = match(PRINT);
				setState(104);
				match(PARIZQ);
				setState(105);
				((PrintstmtContext)_localctx).expr = expr(0);
				setState(106);
				match(PARDER);
				 _localctx.prnt = instructions.NewPrint((((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getLine():0),(((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getCharPositionInLine():0),((PrintstmtContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(109);
				((PrintstmtContext)_localctx).PRINT = match(PRINT);
				setState(110);
				match(PARIZQ);
				setState(111);
				((PrintstmtContext)_localctx).exprComa = exprComa(0);
				setState(112);
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
			setState(144);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(117);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(118);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(119);
				match(LLAVEIZQ);
				setState(120);
				((IfstmtContext)_localctx).block = block();
				setState(121);
				match(LLAVEDER);
				 _localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).block.blk, nil) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(124);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(125);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(126);
				match(LLAVEIZQ);
				setState(127);
				((IfstmtContext)_localctx).e1 = block();
				setState(128);
				match(LLAVEDER);
				setState(129);
				match(ELSE);
				setState(130);
				match(LLAVEIZQ);
				setState(131);
				((IfstmtContext)_localctx).e2 = block();
				setState(132);
				match(LLAVEDER);
				 _localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).e1.blk, ((IfstmtContext)_localctx).e2.blk) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(135);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(136);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(137);
				match(LLAVEIZQ);
				setState(138);
				((IfstmtContext)_localctx).block = block();
				setState(139);
				match(LLAVEDER);
				setState(140);
				match(ELSE);
				setState(141);
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
			setState(146);
			((WhilestmtContext)_localctx).WHILE = match(WHILE);
			setState(147);
			((WhilestmtContext)_localctx).expr = expr(0);
			setState(148);
			match(LLAVEIZQ);
			setState(149);
			((WhilestmtContext)_localctx).block = block();
			setState(150);
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
			setState(193);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(153);
				((DeclarationstmtContext)_localctx).VAR = match(VAR);
				setState(154);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(155);
				match(D_PTS);
				setState(156);
				((DeclarationstmtContext)_localctx).types = types();
				setState(157);
				match(IG);
				setState(158);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				 _localctx.dec = instructions.NewDeclaration((((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getLine():0), (((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), ((DeclarationstmtContext)_localctx).types.ty, ((DeclarationstmtContext)_localctx).expr.e, true) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(161);
				((DeclarationstmtContext)_localctx).VAR = match(VAR);
				setState(162);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(163);
				match(IG);
				setState(164);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				 _localctx.dec = instructions.NewDeclaration((((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getLine():0), (((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), environment.UNKNOWN, ((DeclarationstmtContext)_localctx).expr.e, true) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(167);
				((DeclarationstmtContext)_localctx).VAR = match(VAR);
				setState(168);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(169);
				match(D_PTS);
				setState(170);
				((DeclarationstmtContext)_localctx).types = types();
				 _localctx.dec = instructions.NewDeclaration((((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getLine():0), (((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), ((DeclarationstmtContext)_localctx).types.ty, nil, true) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(173);
				((DeclarationstmtContext)_localctx).LET = match(LET);
				setState(174);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(175);
				match(D_PTS);
				setState(176);
				((DeclarationstmtContext)_localctx).types = types();
				setState(177);
				match(IG);
				setState(178);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				 _localctx.dec = instructions.NewDeclaration((((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getLine():0), (((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), ((DeclarationstmtContext)_localctx).types.ty, ((DeclarationstmtContext)_localctx).expr.e, false) 
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(181);
				((DeclarationstmtContext)_localctx).LET = match(LET);
				setState(182);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(183);
				match(D_PTS);
				setState(184);
				((DeclarationstmtContext)_localctx).types = types();
				 _localctx.dec = instructions.NewDeclaration((((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getLine():0), (((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), ((DeclarationstmtContext)_localctx).types.ty, nil, false) 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(187);
				((DeclarationstmtContext)_localctx).LET = match(LET);
				setState(188);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(189);
				match(IG);
				setState(190);
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
			setState(205);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(195);
				((AssignstmtContext)_localctx).ID = match(ID);
				setState(196);
				((AssignstmtContext)_localctx).op = match(IG);
				setState(197);
				((AssignstmtContext)_localctx).expr = expr(0);
				 _localctx.asg = instructions.NewAssignment((((AssignstmtContext)_localctx).ID!=null?((AssignstmtContext)_localctx).ID.getLine():0), (((AssignstmtContext)_localctx).ID!=null?((AssignstmtContext)_localctx).ID.getCharPositionInLine():0), (((AssignstmtContext)_localctx).ID!=null?((AssignstmtContext)_localctx).ID.getText():null), ((AssignstmtContext)_localctx).expr.e) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(200);
				((AssignstmtContext)_localctx).ID = match(ID);
				setState(201);
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
				setState(202);
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
			setState(207);
			((ForstmtContext)_localctx).FOR = match(FOR);
			setState(208);
			((ForstmtContext)_localctx).ID = match(ID);
			setState(209);
			match(IN);
			setState(210);
			((ForstmtContext)_localctx).expr = expr(0);
			setState(211);
			match(LLAVEIZQ);
			setState(212);
			((ForstmtContext)_localctx).block = block();
			setState(213);
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
			setState(216);
			((GuardstmtContext)_localctx).GUARD = match(GUARD);
			setState(217);
			((GuardstmtContext)_localctx).expr = expr(0);
			setState(218);
			match(ELSE);
			setState(219);
			match(LLAVEIZQ);
			setState(220);
			((GuardstmtContext)_localctx).block = block();
			setState(221);
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
			setState(224);
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
			setState(227);
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
			setState(230);
			((StructCreationContext)_localctx).STRUCT = match(STRUCT);
			setState(231);
			((StructCreationContext)_localctx).ID = match(ID);
			setState(232);
			match(LLAVEIZQ);
			setState(233);
			((StructCreationContext)_localctx).listStructDec = listStructDec(0);
			setState(234);
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
			setState(250);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				{
				setState(238);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==LET) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(239);
				((ListStructDecContext)_localctx).ID = match(ID);
				setState(240);
				match(D_PTS);
				setState(241);
				((ListStructDecContext)_localctx).types = types();

				                        var arr []interface{}
				                        newParams := environment.NewStructType((((ListStructDecContext)_localctx).ID!=null?((ListStructDecContext)_localctx).ID.getText():null), ((ListStructDecContext)_localctx).types.ty, "")
				                        arr = append(arr, newParams)
				                        _localctx.l = arr
				                    
				}
				break;
			case 2:
				{
				setState(244);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==LET) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(245);
				((ListStructDecContext)_localctx).id1 = match(ID);
				setState(246);
				match(D_PTS);
				setState(247);
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
			setState(287);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(285);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
					case 1:
						{
						_localctx = new ListStructDecContext(_parentctx, _parentState);
						_localctx.list = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_listStructDec);
						setState(252);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(254);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==COMA) {
							{
							setState(253);
							match(COMA);
							}
						}

						setState(256);
						_la = _input.LA(1);
						if ( !(_la==VAR || _la==LET) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(257);
						((ListStructDecContext)_localctx).ID = match(ID);
						setState(258);
						match(D_PTS);
						setState(259);
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
						setState(262);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(264);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==COMA) {
							{
							setState(263);
							match(COMA);
							}
						}

						setState(266);
						_la = _input.LA(1);
						if ( !(_la==VAR || _la==LET) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(267);
						((ListStructDecContext)_localctx).id1 = match(ID);
						setState(268);
						match(D_PTS);
						setState(269);
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
						setState(271);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(273);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==COMA) {
							{
							setState(272);
							match(COMA);
							}
						}

						setState(275);
						_la = _input.LA(1);
						if ( !(_la==VAR || _la==LET) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(276);
						((ListStructDecContext)_localctx).id1 = match(ID);
						setState(279);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==D_PTS) {
							{
							setState(277);
							match(D_PTS);
							setState(278);
							((ListStructDecContext)_localctx).types = types();
							}
						}

						setState(281);
						match(IG);
						setState(282);
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
				setState(289);
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
			setState(312);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(290);
				((FnstmtContext)_localctx).FUNC = match(FUNC);
				setState(291);
				((FnstmtContext)_localctx).ID = match(ID);
				setState(292);
				match(PARIZQ);
				setState(293);
				((FnstmtContext)_localctx).listParamsFunc = listParamsFunc(0);
				setState(294);
				match(PARDER);
				setState(295);
				match(FLECHA);
				setState(296);
				((FnstmtContext)_localctx).types = types();
				setState(297);
				match(LLAVEIZQ);
				setState(298);
				((FnstmtContext)_localctx).block = block();
				setState(299);
				match(LLAVEDER);
				 _localctx.fn = instructions.NewFunction((((FnstmtContext)_localctx).FUNC!=null?((FnstmtContext)_localctx).FUNC.getLine():0), (((FnstmtContext)_localctx).FUNC!=null?((FnstmtContext)_localctx).FUNC.getCharPositionInLine():0), (((FnstmtContext)_localctx).ID!=null?((FnstmtContext)_localctx).ID.getText():null), ((FnstmtContext)_localctx).listParamsFunc.l, ((FnstmtContext)_localctx).types.ty, ((FnstmtContext)_localctx).block.blk) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(302);
				((FnstmtContext)_localctx).FUNC = match(FUNC);
				setState(303);
				((FnstmtContext)_localctx).ID = match(ID);
				setState(304);
				match(PARIZQ);
				setState(305);
				((FnstmtContext)_localctx).listParamsFunc = listParamsFunc(0);
				setState(306);
				match(PARDER);
				setState(307);
				match(LLAVEIZQ);
				setState(308);
				((FnstmtContext)_localctx).block = block();
				setState(309);
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
			setState(319);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				{
				setState(315);
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
			setState(328);
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
					setState(321);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(322);
					match(COMA);
					setState(323);
					((ListParamsFuncContext)_localctx).parametro = parametro();

					                                          var arr []interface{}
					                                          arr = append(((ListParamsFuncContext)_localctx).list.l, ((ListParamsFuncContext)_localctx).parametro.p)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(330);
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
			setState(342);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(331);
				((ParametroContext)_localctx).ID = match(ID);
				setState(332);
				match(D_PTS);
				setState(333);
				((ParametroContext)_localctx).types = types();
				 _localctx.p = instructions.NewParamsDeclaration((((ParametroContext)_localctx).ID!=null?((ParametroContext)_localctx).ID.getLine():0),(((ParametroContext)_localctx).ID!=null?((ParametroContext)_localctx).ID.getCharPositionInLine():0),(((ParametroContext)_localctx).ID!=null?((ParametroContext)_localctx).ID.getText():null), ((ParametroContext)_localctx).types.ty)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(336);
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
				setState(337);
				((ParametroContext)_localctx).ID = match(ID);
				setState(338);
				match(D_PTS);
				setState(339);
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
			setState(344);
			((CallExpContext)_localctx).ID = match(ID);
			setState(345);
			match(PARIZQ);
			setState(346);
			((CallExpContext)_localctx).listParamsCall = listParamsCall(0);
			setState(347);
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
			setState(350);
			((CallFunctionContext)_localctx).ID = match(ID);
			setState(351);
			match(PARIZQ);
			setState(352);
			((CallFunctionContext)_localctx).listParamsCall = listParamsCall(0);
			setState(353);
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
			setState(361);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				{
				setState(357);
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
			setState(370);
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
					setState(363);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(364);
					match(COMA);
					setState(365);
					((ListParamsCallContext)_localctx).expr = expr(0);

					                                              var arr []interface{}
					                                              arr = append(((ListParamsCallContext)_localctx).list.l, ((ListParamsCallContext)_localctx).expr.e)
					                                              _localctx.l = arr
					                                          
					}
					} 
				}
				setState(372);
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
			setState(396);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(373);
				match(INT);
				 _localctx.ty = environment.INTEGER 
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(375);
				match(FLOAT);
				 _localctx.ty = environment.FLOAT 
				}
				break;
			case STR:
				enterOuterAlt(_localctx, 3);
				{
				setState(377);
				match(STR);
				 _localctx.ty = environment.STRING 
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(379);
				match(BOOL);
				 _localctx.ty = environment.BOOLEAN 
				}
				break;
			case CORIZQ:
				enterOuterAlt(_localctx, 5);
				{
				setState(381);
				match(CORIZQ);
				setState(382);
				types();
				setState(383);
				match(CORDER);
				 _localctx.ty = environment.ARRAY 
				}
				break;
			case COMILLA:
				enterOuterAlt(_localctx, 6);
				{
				setState(386);
				match(COMILLA);
				setState(387);
				match(STR);
				setState(388);
				match(COMILLA);
				 _localctx.ty = environment.STR 
				}
				break;
			case NIL:
				enterOuterAlt(_localctx, 7);
				{
				setState(390);
				match(NIL);
				 _localctx.ty = environment.NIL 
				}
				break;
			case STRUCT:
				enterOuterAlt(_localctx, 8);
				{
				setState(392);
				match(STRUCT);
				 _localctx.ty = environment.STRUCT 
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 9);
				{
				setState(394);
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
		public ExprContext ID2;
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
		public CountContext cnt;
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
		public CountContext count() {
			return getRuleContext(CountContext.class,0);
		}
		public TerminalNode NIL() { return getToken(SwiftGrammarParser.NIL, 0); }
		public List<TerminalNode> PUNTO() { return getTokens(SwiftGrammarParser.PUNTO); }
		public TerminalNode PUNTO(int i) {
			return getToken(SwiftGrammarParser.PUNTO, i);
		}
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
		public TerminalNode ISEMPTY() { return getToken(SwiftGrammarParser.ISEMPTY, 0); }
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
			setState(439);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				{
				setState(399);
				((ExprContext)_localctx).SUB = match(SUB);
				setState(400);
				((ExprContext)_localctx).opDe = ((ExprContext)_localctx).expr = expr(23);
				_localctx.e = expressions.NewOperation((((ExprContext)_localctx).SUB!=null?((ExprContext)_localctx).SUB.getLine():0),(((ExprContext)_localctx).SUB!=null?((ExprContext)_localctx).SUB.getCharPositionInLine():0),((ExprContext)_localctx).opDe.e,"NEGACION",nil)
				}
				break;
			case 2:
				{
				setState(403);
				((ExprContext)_localctx).NOT = match(NOT);
				setState(404);
				((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(13);
				_localctx.e = expressions.NewOperation((((ExprContext)_localctx).NOT!=null?((ExprContext)_localctx).NOT.getLine():0), (((ExprContext)_localctx).NOT!=null?((ExprContext)_localctx).NOT.getCharPositionInLine():0), ((ExprContext)_localctx).right.e, (((ExprContext)_localctx).NOT!=null?((ExprContext)_localctx).NOT.getText():null) ,nil)
				}
				break;
			case 3:
				{
				setState(407);
				match(PARIZQ);
				setState(408);
				((ExprContext)_localctx).expr = expr(0);
				setState(409);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case 4:
				{
				setState(412);
				((ExprContext)_localctx).callExp = callExp();
				 _localctx.e = ((ExprContext)_localctx).callExp.cfe 
				}
				break;
			case 5:
				{
				setState(415);
				((ExprContext)_localctx).CORIZQ = match(CORIZQ);
				setState(416);
				match(CORDER);
				 _localctx.e = expressions.NewArray((((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getLine():0), (((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getCharPositionInLine():0), nil) 
				}
				break;
			case 6:
				{
				setState(418);
				((ExprContext)_localctx).list = listArray(0);
				 _localctx.e = ((ExprContext)_localctx).list.p
				}
				break;
			case 7:
				{
				setState(421);
				((ExprContext)_localctx).CORIZQ = match(CORIZQ);
				setState(422);
				((ExprContext)_localctx).listParams = listParams(0);
				setState(423);
				match(CORDER);
				 _localctx.e = expressions.NewArray((((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getLine():0), (((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getCharPositionInLine():0), ((ExprContext)_localctx).listParams.l) 
				}
				break;
			case 8:
				{
				setState(426);
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
				setState(428);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case 10:
				{
				setState(430);
				((ExprContext)_localctx).TRU = match(TRU);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getLine():0), (((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case 11:
				{
				setState(432);
				((ExprContext)_localctx).FAL = match(FAL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getLine():0), (((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case 12:
				{
				setState(434);
				((ExprContext)_localctx).cnt = count(0);
				 _localctx.e = ((ExprContext)_localctx).cnt.cn 
				}
				break;
			case 13:
				{
				setState(437);
				((ExprContext)_localctx).NIL = match(NIL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getLine():0), (((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getCharPositionInLine():0), "nil", environment.NIL) 
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(494);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(492);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.expuno = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(441);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(442);
						match(PUNTO);
						setState(443);
						match(PUNTO);
						setState(444);
						match(PUNTO);
						setState(445);
						((ExprContext)_localctx).expdos = ((ExprContext)_localctx).expr = expr(23);
						 _localctx.e = expressions.NewRange((((ExprContext)_localctx).expuno!=null?(((ExprContext)_localctx).expuno.start):null).GetLine(),(((ExprContext)_localctx).expuno!=null?(((ExprContext)_localctx).expuno.start):null).GetColumn(), ((ExprContext)_localctx).expuno.e, ((ExprContext)_localctx).expdos.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(448);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(449);
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
						setState(450);
						((ExprContext)_localctx).expr = expr(22);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getLine():0), (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getCharPositionInLine():0), nil, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).expr.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(453);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(454);
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
						setState(455);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(21);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(458);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(459);
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
						setState(460);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(20);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(463);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(464);
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
						setState(465);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(19);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(468);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(469);
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
						setState(470);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(18);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(473);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(474);
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
						setState(475);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(17);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 8:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(478);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(479);
						((ExprContext)_localctx).op = match(AND);
						setState(480);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(16);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 9:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(483);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(484);
						((ExprContext)_localctx).op = match(OR);
						setState(485);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(15);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 10:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.ID2 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(488);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(489);
						match(PUNTO);
						setState(490);
						match(ISEMPTY);
						 _localctx.e = expressions.NewIsEmpty((((ExprContext)_localctx).ID2!=null?(((ExprContext)_localctx).ID2.start):null).GetLine(), (((ExprContext)_localctx).ID2!=null?(((ExprContext)_localctx).ID2.start):null).GetColumn(), ((ExprContext)_localctx).ID2.e) 
						}
						break;
					}
					} 
				}
				setState(496);
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
			setState(498);
			((ListParamsContext)_localctx).expr = expr(0);

			            _localctx.l = []interface{}{}
			            _localctx.l = append(_localctx.l, ((ListParamsContext)_localctx).expr.e)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(508);
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
					setState(501);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(502);
					match(COMA);
					setState(503);
					((ListParamsContext)_localctx).expr = expr(0);

					                                          var arr []interface{}
					                                          arr = append(((ListParamsContext)_localctx).list.l, ((ListParamsContext)_localctx).expr.e)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(510);
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
			setState(512);
			match(CORIZQ);
			setState(513);
			((ListAccessArrayContext)_localctx).expr = expr(0);
			setState(514);
			match(CORDER);

			            _localctx.l = []interface{}{}
			            _localctx.l = append(_localctx.l, ((ListAccessArrayContext)_localctx).expr.e)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(525);
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
					setState(517);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(518);
					match(CORIZQ);
					setState(519);
					((ListAccessArrayContext)_localctx).expr = expr(0);
					setState(520);
					match(CORDER);

					                                          var arr []interface{}
					                                          arr = append(((ListAccessArrayContext)_localctx).list.l, ((ListAccessArrayContext)_localctx).expr.e)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(527);
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
			setState(529);
			((ListArrayContext)_localctx).ID = match(ID);
			 _localctx.p = expressions.NewCallVar((((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getCharPositionInLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getText():null))
			}
			_ctx.stop = _input.LT(-1);
			setState(548);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(546);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
					case 1:
						{
						_localctx = new ListArrayContext(_parentctx, _parentState);
						_localctx.list = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_listArray);
						setState(532);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(533);
						match(CORIZQ);
						setState(534);
						((ListArrayContext)_localctx).expr = expr(0);
						setState(535);
						match(CORDER);
						 _localctx.p = expressions.NewArrayAccess((((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetLine(), (((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetColumn(), ((ListArrayContext)_localctx).list.p, ((ListArrayContext)_localctx).expr.e) 
						}
						break;
					case 2:
						{
						_localctx = new ListArrayContext(_parentctx, _parentState);
						_localctx.list = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_listArray);
						setState(538);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(539);
						types();
						setState(540);
						match(IG);
						setState(541);
						match(CORIZQ);
						setState(542);
						((ListArrayContext)_localctx).expr = expr(0);
						setState(543);
						match(CORDER);
						 _localctx.p = expressions.NewArrayAccess((((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetLine(), (((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetColumn(), ((ListArrayContext)_localctx).list.p, ((ListArrayContext)_localctx).expr.e) 
						}
						break;
					}
					} 
				}
				setState(550);
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
	public static class CountContext extends ParserRuleContext {
		public interfaces.Expression cn;
		public CountContext cnt;
		public Token ID;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode COUNT() { return getToken(SwiftGrammarParser.COUNT, 0); }
		public CountContext count() {
			return getRuleContext(CountContext.class,0);
		}
		public CountContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_count; }
	}

	public final CountContext count() throws RecognitionException {
		return count(0);
	}

	private CountContext count(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		CountContext _localctx = new CountContext(_ctx, _parentState);
		CountContext _prevctx = _localctx;
		int _startState = 50;
		enterRecursionRule(_localctx, 50, RULE_count, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(552);
			((CountContext)_localctx).ID = match(ID);
			 _localctx.cn = expressions.NewCallVar((((CountContext)_localctx).ID!=null?((CountContext)_localctx).ID.getLine():0), (((CountContext)_localctx).ID!=null?((CountContext)_localctx).ID.getCharPositionInLine():0), (((CountContext)_localctx).ID!=null?((CountContext)_localctx).ID.getText():null))
			}
			_ctx.stop = _input.LT(-1);
			setState(561);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new CountContext(_parentctx, _parentState);
					_localctx.cnt = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_count);
					setState(555);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(556);
					match(PUNTO);
					setState(557);
					match(COUNT);
					 _localctx.cn = expressions.NewCount((((CountContext)_localctx).cnt!=null?(((CountContext)_localctx).cnt.start):null).GetLine(), (((CountContext)_localctx).cnt!=null?(((CountContext)_localctx).cnt.start):null).GetColumn(),((CountContext)_localctx).cnt.cn) 
					}
					} 
				}
				setState(563);
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
		int _startState = 52;
		enterRecursionRule(_localctx, 52, RULE_exprComa, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(565);
			((ExprComaContext)_localctx).expr = expr(0);
			 _localctx.t = ((ExprComaContext)_localctx).expr.e 
			}
			_ctx.stop = _input.LT(-1);
			setState(575);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ExprComaContext(_parentctx, _parentState);
					_localctx.left = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_exprComa);
					setState(568);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(569);
					((ExprComaContext)_localctx).op = match(COMA);
					setState(570);
					((ExprComaContext)_localctx).right = ((ExprComaContext)_localctx).expr = expr(0);
					 _localctx.t = expressions.NewOperation((((ExprComaContext)_localctx).left!=null?(((ExprComaContext)_localctx).left.start):null).GetLine(), (((ExprComaContext)_localctx).left!=null?(((ExprComaContext)_localctx).left.start):null).GetColumn(), ((ExprComaContext)_localctx).left.t, (((ExprComaContext)_localctx).op!=null?((ExprComaContext)_localctx).op.getText():null), ((ExprComaContext)_localctx).right.e) 
					}
					} 
				}
				setState(577);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
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
			return count_sempred((CountContext)_localctx, predIndex);
		case 26:
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
			return precpred(_ctx, 22);
		case 6:
			return precpred(_ctx, 21);
		case 7:
			return precpred(_ctx, 20);
		case 8:
			return precpred(_ctx, 19);
		case 9:
			return precpred(_ctx, 18);
		case 10:
			return precpred(_ctx, 17);
		case 11:
			return precpred(_ctx, 16);
		case 12:
			return precpred(_ctx, 15);
		case 13:
			return precpred(_ctx, 14);
		case 14:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listParams_sempred(ListParamsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 15:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listAccessArray_sempred(ListAccessArrayContext _localctx, int predIndex) {
		switch (predIndex) {
		case 16:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listArray_sempred(ListArrayContext _localctx, int predIndex) {
		switch (predIndex) {
		case 17:
			return precpred(_ctx, 3);
		case 18:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean count_sempred(CountContext _localctx, int predIndex) {
		switch (predIndex) {
		case 19:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean exprComa_sempred(ExprComaContext _localctx, int predIndex) {
		switch (predIndex) {
		case 20:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001G\u0243\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0001\u0004\u0001<\b\u0001\u000b\u0001"+
		"\f\u0001=\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002f\b\u0002\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u0003"+
		"t\b\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004\u0091\b\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006"+
		"\u00c2\b\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007"+
		"\u00ce\b\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b"+
		"\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0003\r\u00fb\b\r\u0001\r\u0001\r\u0003\r\u00ff\b\r"+
		"\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0003"+
		"\r\u0109\b\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0003"+
		"\r\u0112\b\r\u0001\r\u0001\r\u0001\r\u0001\r\u0003\r\u0118\b\r\u0001\r"+
		"\u0001\r\u0001\r\u0001\r\u0005\r\u011e\b\r\n\r\f\r\u0121\t\r\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0003\u000e\u0139\b\u000e\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u0140\b\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0005\u000f"+
		"\u0147\b\u000f\n\u000f\f\u000f\u014a\t\u000f\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0003\u0010\u0157\b\u0010\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0003\u0013\u016a\b\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0005\u0013\u0171"+
		"\b\u0013\n\u0013\f\u0013\u0174\t\u0013\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0003\u0014\u018d\b\u0014\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0003\u0015\u01b8\b\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0005\u0015\u01ed\b\u0015\n\u0015\f\u0015\u01f0"+
		"\t\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0005\u0016\u01fb\b\u0016\n"+
		"\u0016\f\u0016\u01fe\t\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0005\u0017\u020c\b\u0017\n\u0017\f\u0017"+
		"\u020f\t\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0005\u0018\u0223\b\u0018\n\u0018\f\u0018\u0226\t\u0018\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0005\u0019\u0230\b\u0019\n\u0019\f\u0019\u0233\t\u0019"+
		"\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a"+
		"\u0001\u001a\u0001\u001a\u0001\u001a\u0005\u001a\u023e\b\u001a\n\u001a"+
		"\f\u001a\u0241\t\u001a\u0001\u001a\u0000\t\u001a\u001e&*,.024\u001b\u0000"+
		"\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c"+
		"\u001e \"$&(*,.024\u0000\b\u0001\u0000/0\u0001\u0000\u0006\u0007\u0002"+
		"\u0000&&CC\u0002\u00003477\u0001\u000056\u0002\u0000--11\u0002\u0000."+
		".22\u0001\u0000\'(\u026f\u00006\u0001\u0000\u0000\u0000\u0002;\u0001\u0000"+
		"\u0000\u0000\u0004e\u0001\u0000\u0000\u0000\u0006s\u0001\u0000\u0000\u0000"+
		"\b\u0090\u0001\u0000\u0000\u0000\n\u0092\u0001\u0000\u0000\u0000\f\u00c1"+
		"\u0001\u0000\u0000\u0000\u000e\u00cd\u0001\u0000\u0000\u0000\u0010\u00cf"+
		"\u0001\u0000\u0000\u0000\u0012\u00d8\u0001\u0000\u0000\u0000\u0014\u00e0"+
		"\u0001\u0000\u0000\u0000\u0016\u00e3\u0001\u0000\u0000\u0000\u0018\u00e6"+
		"\u0001\u0000\u0000\u0000\u001a\u00fa\u0001\u0000\u0000\u0000\u001c\u0138"+
		"\u0001\u0000\u0000\u0000\u001e\u013f\u0001\u0000\u0000\u0000 \u0156\u0001"+
		"\u0000\u0000\u0000\"\u0158\u0001\u0000\u0000\u0000$\u015e\u0001\u0000"+
		"\u0000\u0000&\u0169\u0001\u0000\u0000\u0000(\u018c\u0001\u0000\u0000\u0000"+
		"*\u01b7\u0001\u0000\u0000\u0000,\u01f1\u0001\u0000\u0000\u0000.\u01ff"+
		"\u0001\u0000\u0000\u00000\u0210\u0001\u0000\u0000\u00002\u0227\u0001\u0000"+
		"\u0000\u00004\u0234\u0001\u0000\u0000\u000067\u0003\u0002\u0001\u0000"+
		"78\u0005\u0000\u0000\u000189\u0006\u0000\uffff\uffff\u00009\u0001\u0001"+
		"\u0000\u0000\u0000:<\u0003\u0004\u0002\u0000;:\u0001\u0000\u0000\u0000"+
		"<=\u0001\u0000\u0000\u0000=;\u0001\u0000\u0000\u0000=>\u0001\u0000\u0000"+
		"\u0000>?\u0001\u0000\u0000\u0000?@\u0006\u0001\uffff\uffff\u0000@\u0003"+
		"\u0001\u0000\u0000\u0000AB\u0003\u0006\u0003\u0000BC\u0006\u0002\uffff"+
		"\uffff\u0000Cf\u0001\u0000\u0000\u0000DE\u0003\b\u0004\u0000EF\u0006\u0002"+
		"\uffff\uffff\u0000Ff\u0001\u0000\u0000\u0000GH\u0003\f\u0006\u0000HI\u0006"+
		"\u0002\uffff\uffff\u0000If\u0001\u0000\u0000\u0000JK\u0003\n\u0005\u0000"+
		"KL\u0006\u0002\uffff\uffff\u0000Lf\u0001\u0000\u0000\u0000MN\u0003\u000e"+
		"\u0007\u0000NO\u0006\u0002\uffff\uffff\u0000Of\u0001\u0000\u0000\u0000"+
		"PQ\u0003\u0010\b\u0000QR\u0006\u0002\uffff\uffff\u0000Rf\u0001\u0000\u0000"+
		"\u0000ST\u0003\u0012\t\u0000TU\u0006\u0002\uffff\uffff\u0000Uf\u0001\u0000"+
		"\u0000\u0000VW\u0003\u0014\n\u0000WX\u0006\u0002\uffff\uffff\u0000Xf\u0001"+
		"\u0000\u0000\u0000YZ\u0003\u0016\u000b\u0000Z[\u0006\u0002\uffff\uffff"+
		"\u0000[f\u0001\u0000\u0000\u0000\\]\u0003\u0018\f\u0000]^\u0006\u0002"+
		"\uffff\uffff\u0000^f\u0001\u0000\u0000\u0000_`\u0003\u001c\u000e\u0000"+
		"`a\u0006\u0002\uffff\uffff\u0000af\u0001\u0000\u0000\u0000bc\u0003$\u0012"+
		"\u0000cd\u0006\u0002\uffff\uffff\u0000df\u0001\u0000\u0000\u0000eA\u0001"+
		"\u0000\u0000\u0000eD\u0001\u0000\u0000\u0000eG\u0001\u0000\u0000\u0000"+
		"eJ\u0001\u0000\u0000\u0000eM\u0001\u0000\u0000\u0000eP\u0001\u0000\u0000"+
		"\u0000eS\u0001\u0000\u0000\u0000eV\u0001\u0000\u0000\u0000eY\u0001\u0000"+
		"\u0000\u0000e\\\u0001\u0000\u0000\u0000e_\u0001\u0000\u0000\u0000eb\u0001"+
		"\u0000\u0000\u0000f\u0005\u0001\u0000\u0000\u0000gh\u0005\u000b\u0000"+
		"\u0000hi\u00058\u0000\u0000ij\u0003*\u0015\u0000jk\u00059\u0000\u0000"+
		"kl\u0006\u0003\uffff\uffff\u0000lt\u0001\u0000\u0000\u0000mn\u0005\u000b"+
		"\u0000\u0000no\u00058\u0000\u0000op\u00034\u001a\u0000pq\u00059\u0000"+
		"\u0000qr\u0006\u0003\uffff\uffff\u0000rt\u0001\u0000\u0000\u0000sg\u0001"+
		"\u0000\u0000\u0000sm\u0001\u0000\u0000\u0000t\u0007\u0001\u0000\u0000"+
		"\u0000uv\u0005\f\u0000\u0000vw\u0003*\u0015\u0000wx\u0005:\u0000\u0000"+
		"xy\u0003\u0002\u0001\u0000yz\u0005;\u0000\u0000z{\u0006\u0004\uffff\uffff"+
		"\u0000{\u0091\u0001\u0000\u0000\u0000|}\u0005\f\u0000\u0000}~\u0003*\u0015"+
		"\u0000~\u007f\u0005:\u0000\u0000\u007f\u0080\u0003\u0002\u0001\u0000\u0080"+
		"\u0081\u0005;\u0000\u0000\u0081\u0082\u0005\r\u0000\u0000\u0082\u0083"+
		"\u0005:\u0000\u0000\u0083\u0084\u0003\u0002\u0001\u0000\u0084\u0085\u0005"+
		";\u0000\u0000\u0085\u0086\u0006\u0004\uffff\uffff\u0000\u0086\u0091\u0001"+
		"\u0000\u0000\u0000\u0087\u0088\u0005\f\u0000\u0000\u0088\u0089\u0003*"+
		"\u0015\u0000\u0089\u008a\u0005:\u0000\u0000\u008a\u008b\u0003\u0002\u0001"+
		"\u0000\u008b\u008c\u0005;\u0000\u0000\u008c\u008d\u0005\r\u0000\u0000"+
		"\u008d\u008e\u0003\b\u0004\u0000\u008e\u008f\u0006\u0004\uffff\uffff\u0000"+
		"\u008f\u0091\u0001\u0000\u0000\u0000\u0090u\u0001\u0000\u0000\u0000\u0090"+
		"|\u0001\u0000\u0000\u0000\u0090\u0087\u0001\u0000\u0000\u0000\u0091\t"+
		"\u0001\u0000\u0000\u0000\u0092\u0093\u0005\u000e\u0000\u0000\u0093\u0094"+
		"\u0003*\u0015\u0000\u0094\u0095\u0005:\u0000\u0000\u0095\u0096\u0003\u0002"+
		"\u0001\u0000\u0096\u0097\u0005;\u0000\u0000\u0097\u0098\u0006\u0005\uffff"+
		"\uffff\u0000\u0098\u000b\u0001\u0000\u0000\u0000\u0099\u009a\u0005\u0006"+
		"\u0000\u0000\u009a\u009b\u0005&\u0000\u0000\u009b\u009c\u0005<\u0000\u0000"+
		"\u009c\u009d\u0003(\u0014\u0000\u009d\u009e\u0005,\u0000\u0000\u009e\u009f"+
		"\u0003*\u0015\u0000\u009f\u00a0\u0006\u0006\uffff\uffff\u0000\u00a0\u00c2"+
		"\u0001\u0000\u0000\u0000\u00a1\u00a2\u0005\u0006\u0000\u0000\u00a2\u00a3"+
		"\u0005&\u0000\u0000\u00a3\u00a4\u0005,\u0000\u0000\u00a4\u00a5\u0003*"+
		"\u0015\u0000\u00a5\u00a6\u0006\u0006\uffff\uffff\u0000\u00a6\u00c2\u0001"+
		"\u0000\u0000\u0000\u00a7\u00a8\u0005\u0006\u0000\u0000\u00a8\u00a9\u0005"+
		"&\u0000\u0000\u00a9\u00aa\u0005<\u0000\u0000\u00aa\u00ab\u0003(\u0014"+
		"\u0000\u00ab\u00ac\u0006\u0006\uffff\uffff\u0000\u00ac\u00c2\u0001\u0000"+
		"\u0000\u0000\u00ad\u00ae\u0005\u0007\u0000\u0000\u00ae\u00af\u0005&\u0000"+
		"\u0000\u00af\u00b0\u0005<\u0000\u0000\u00b0\u00b1\u0003(\u0014\u0000\u00b1"+
		"\u00b2\u0005,\u0000\u0000\u00b2\u00b3\u0003*\u0015\u0000\u00b3\u00b4\u0006"+
		"\u0006\uffff\uffff\u0000\u00b4\u00c2\u0001\u0000\u0000\u0000\u00b5\u00b6"+
		"\u0005\u0007\u0000\u0000\u00b6\u00b7\u0005&\u0000\u0000\u00b7\u00b8\u0005"+
		"<\u0000\u0000\u00b8\u00b9\u0003(\u0014\u0000\u00b9\u00ba\u0006\u0006\uffff"+
		"\uffff\u0000\u00ba\u00c2\u0001\u0000\u0000\u0000\u00bb\u00bc\u0005\u0007"+
		"\u0000\u0000\u00bc\u00bd\u0005&\u0000\u0000\u00bd\u00be\u0005,\u0000\u0000"+
		"\u00be\u00bf\u0003*\u0015\u0000\u00bf\u00c0\u0006\u0006\uffff\uffff\u0000"+
		"\u00c0\u00c2\u0001\u0000\u0000\u0000\u00c1\u0099\u0001\u0000\u0000\u0000"+
		"\u00c1\u00a1\u0001\u0000\u0000\u0000\u00c1\u00a7\u0001\u0000\u0000\u0000"+
		"\u00c1\u00ad\u0001\u0000\u0000\u0000\u00c1\u00b5\u0001\u0000\u0000\u0000"+
		"\u00c1\u00bb\u0001\u0000\u0000\u0000\u00c2\r\u0001\u0000\u0000\u0000\u00c3"+
		"\u00c4\u0005&\u0000\u0000\u00c4\u00c5\u0005,\u0000\u0000\u00c5\u00c6\u0003"+
		"*\u0015\u0000\u00c6\u00c7\u0006\u0007\uffff\uffff\u0000\u00c7\u00ce\u0001"+
		"\u0000\u0000\u0000\u00c8\u00c9\u0005&\u0000\u0000\u00c9\u00ca\u0007\u0000"+
		"\u0000\u0000\u00ca\u00cb\u0003*\u0015\u0000\u00cb\u00cc\u0006\u0007\uffff"+
		"\uffff\u0000\u00cc\u00ce\u0001\u0000\u0000\u0000\u00cd\u00c3\u0001\u0000"+
		"\u0000\u0000\u00cd\u00c8\u0001\u0000\u0000\u0000\u00ce\u000f\u0001\u0000"+
		"\u0000\u0000\u00cf\u00d0\u0005\u000f\u0000\u0000\u00d0\u00d1\u0005&\u0000"+
		"\u0000\u00d1\u00d2\u0005\u0010\u0000\u0000\u00d2\u00d3\u0003*\u0015\u0000"+
		"\u00d3\u00d4\u0005:\u0000\u0000\u00d4\u00d5\u0003\u0002\u0001\u0000\u00d5"+
		"\u00d6\u0005;\u0000\u0000\u00d6\u00d7\u0006\b\uffff\uffff\u0000\u00d7"+
		"\u0011\u0001\u0000\u0000\u0000\u00d8\u00d9\u0005\u0017\u0000\u0000\u00d9"+
		"\u00da\u0003*\u0015\u0000\u00da\u00db\u0005\r\u0000\u0000\u00db\u00dc"+
		"\u0005:\u0000\u0000\u00dc\u00dd\u0003\u0002\u0001\u0000\u00dd\u00de\u0005"+
		";\u0000\u0000\u00de\u00df\u0006\t\uffff\uffff\u0000\u00df\u0013\u0001"+
		"\u0000\u0000\u0000\u00e0\u00e1\u0005\u0014\u0000\u0000\u00e1\u00e2\u0006"+
		"\n\uffff\uffff\u0000\u00e2\u0015\u0001\u0000\u0000\u0000\u00e3\u00e4\u0005"+
		"\u0016\u0000\u0000\u00e4\u00e5\u0006\u000b\uffff\uffff\u0000\u00e5\u0017"+
		"\u0001\u0000\u0000\u0000\u00e6\u00e7\u0005\u001a\u0000\u0000\u00e7\u00e8"+
		"\u0005&\u0000\u0000\u00e8\u00e9\u0005:\u0000\u0000\u00e9\u00ea\u0003\u001a"+
		"\r\u0000\u00ea\u00eb\u0005;\u0000\u0000\u00eb\u00ec\u0006\f\uffff\uffff"+
		"\u0000\u00ec\u0019\u0001\u0000\u0000\u0000\u00ed\u00ee\u0006\r\uffff\uffff"+
		"\u0000\u00ee\u00ef\u0007\u0001\u0000\u0000\u00ef\u00f0\u0005&\u0000\u0000"+
		"\u00f0\u00f1\u0005<\u0000\u0000\u00f1\u00f2\u0003(\u0014\u0000\u00f2\u00f3"+
		"\u0006\r\uffff\uffff\u0000\u00f3\u00fb\u0001\u0000\u0000\u0000\u00f4\u00f5"+
		"\u0007\u0001\u0000\u0000\u00f5\u00f6\u0005&\u0000\u0000\u00f6\u00f7\u0005"+
		"<\u0000\u0000\u00f7\u00f8\u0005&\u0000\u0000\u00f8\u00fb\u0006\r\uffff"+
		"\uffff\u0000\u00f9\u00fb\u0006\r\uffff\uffff\u0000\u00fa\u00ed\u0001\u0000"+
		"\u0000\u0000\u00fa\u00f4\u0001\u0000\u0000\u0000\u00fa\u00f9\u0001\u0000"+
		"\u0000\u0000\u00fb\u011f\u0001\u0000\u0000\u0000\u00fc\u00fe\n\u0006\u0000"+
		"\u0000\u00fd\u00ff\u0005?\u0000\u0000\u00fe\u00fd\u0001\u0000\u0000\u0000"+
		"\u00fe\u00ff\u0001\u0000\u0000\u0000\u00ff\u0100\u0001\u0000\u0000\u0000"+
		"\u0100\u0101\u0007\u0001\u0000\u0000\u0101\u0102\u0005&\u0000\u0000\u0102"+
		"\u0103\u0005<\u0000\u0000\u0103\u0104\u0003(\u0014\u0000\u0104\u0105\u0006"+
		"\r\uffff\uffff\u0000\u0105\u011e\u0001\u0000\u0000\u0000\u0106\u0108\n"+
		"\u0005\u0000\u0000\u0107\u0109\u0005?\u0000\u0000\u0108\u0107\u0001\u0000"+
		"\u0000\u0000\u0108\u0109\u0001\u0000\u0000\u0000\u0109\u010a\u0001\u0000"+
		"\u0000\u0000\u010a\u010b\u0007\u0001\u0000\u0000\u010b\u010c\u0005&\u0000"+
		"\u0000\u010c\u010d\u0005<\u0000\u0000\u010d\u010e\u0005&\u0000\u0000\u010e"+
		"\u011e\u0006\r\uffff\uffff\u0000\u010f\u0111\n\u0004\u0000\u0000\u0110"+
		"\u0112\u0005?\u0000\u0000\u0111\u0110\u0001\u0000\u0000\u0000\u0111\u0112"+
		"\u0001\u0000\u0000\u0000\u0112\u0113\u0001\u0000\u0000\u0000\u0113\u0114"+
		"\u0007\u0001\u0000\u0000\u0114\u0117\u0005&\u0000\u0000\u0115\u0116\u0005"+
		"<\u0000\u0000\u0116\u0118\u0003(\u0014\u0000\u0117\u0115\u0001\u0000\u0000"+
		"\u0000\u0117\u0118\u0001\u0000\u0000\u0000\u0118\u0119\u0001\u0000\u0000"+
		"\u0000\u0119\u011a\u0005,\u0000\u0000\u011a\u011b\u0003*\u0015\u0000\u011b"+
		"\u011c\u0006\r\uffff\uffff\u0000\u011c\u011e\u0001\u0000\u0000\u0000\u011d"+
		"\u00fc\u0001\u0000\u0000\u0000\u011d\u0106\u0001\u0000\u0000\u0000\u011d"+
		"\u010f\u0001\u0000\u0000\u0000\u011e\u0121\u0001\u0000\u0000\u0000\u011f"+
		"\u011d\u0001\u0000\u0000\u0000\u011f\u0120\u0001\u0000\u0000\u0000\u0120"+
		"\u001b\u0001\u0000\u0000\u0000\u0121\u011f\u0001\u0000\u0000\u0000\u0122"+
		"\u0123\u0005\u0018\u0000\u0000\u0123\u0124\u0005&\u0000\u0000\u0124\u0125"+
		"\u00058\u0000\u0000\u0125\u0126\u0003\u001e\u000f\u0000\u0126\u0127\u0005"+
		"9\u0000\u0000\u0127\u0128\u0005B\u0000\u0000\u0128\u0129\u0003(\u0014"+
		"\u0000\u0129\u012a\u0005:\u0000\u0000\u012a\u012b\u0003\u0002\u0001\u0000"+
		"\u012b\u012c\u0005;\u0000\u0000\u012c\u012d\u0006\u000e\uffff\uffff\u0000"+
		"\u012d\u0139\u0001\u0000\u0000\u0000\u012e\u012f\u0005\u0018\u0000\u0000"+
		"\u012f\u0130\u0005&\u0000\u0000\u0130\u0131\u00058\u0000\u0000\u0131\u0132"+
		"\u0003\u001e\u000f\u0000\u0132\u0133\u00059\u0000\u0000\u0133\u0134\u0005"+
		":\u0000\u0000\u0134\u0135\u0003\u0002\u0001\u0000\u0135\u0136\u0005;\u0000"+
		"\u0000\u0136\u0137\u0006\u000e\uffff\uffff\u0000\u0137\u0139\u0001\u0000"+
		"\u0000\u0000\u0138\u0122\u0001\u0000\u0000\u0000\u0138\u012e\u0001\u0000"+
		"\u0000\u0000\u0139\u001d\u0001\u0000\u0000\u0000\u013a\u013b\u0006\u000f"+
		"\uffff\uffff\u0000\u013b\u013c\u0003 \u0010\u0000\u013c\u013d\u0006\u000f"+
		"\uffff\uffff\u0000\u013d\u0140\u0001\u0000\u0000\u0000\u013e\u0140\u0006"+
		"\u000f\uffff\uffff\u0000\u013f\u013a\u0001\u0000\u0000\u0000\u013f\u013e"+
		"\u0001\u0000\u0000\u0000\u0140\u0148\u0001\u0000\u0000\u0000\u0141\u0142"+
		"\n\u0003\u0000\u0000\u0142\u0143\u0005?\u0000\u0000\u0143\u0144\u0003"+
		" \u0010\u0000\u0144\u0145\u0006\u000f\uffff\uffff\u0000\u0145\u0147\u0001"+
		"\u0000\u0000\u0000\u0146\u0141\u0001\u0000\u0000\u0000\u0147\u014a\u0001"+
		"\u0000\u0000\u0000\u0148\u0146\u0001\u0000\u0000\u0000\u0148\u0149\u0001"+
		"\u0000\u0000\u0000\u0149\u001f\u0001\u0000\u0000\u0000\u014a\u0148\u0001"+
		"\u0000\u0000\u0000\u014b\u014c\u0005&\u0000\u0000\u014c\u014d\u0005<\u0000"+
		"\u0000\u014d\u014e\u0003(\u0014\u0000\u014e\u014f\u0006\u0010\uffff\uffff"+
		"\u0000\u014f\u0157\u0001\u0000\u0000\u0000\u0150\u0151\u0007\u0002\u0000"+
		"\u0000\u0151\u0152\u0005&\u0000\u0000\u0152\u0153\u0005<\u0000\u0000\u0153"+
		"\u0154\u0003(\u0014\u0000\u0154\u0155\u0006\u0010\uffff\uffff\u0000\u0155"+
		"\u0157\u0001\u0000\u0000\u0000\u0156\u014b\u0001\u0000\u0000\u0000\u0156"+
		"\u0150\u0001\u0000\u0000\u0000\u0157!\u0001\u0000\u0000\u0000\u0158\u0159"+
		"\u0005&\u0000\u0000\u0159\u015a\u00058\u0000\u0000\u015a\u015b\u0003&"+
		"\u0013\u0000\u015b\u015c\u00059\u0000\u0000\u015c\u015d\u0006\u0011\uffff"+
		"\uffff\u0000\u015d#\u0001\u0000\u0000\u0000\u015e\u015f\u0005&\u0000\u0000"+
		"\u015f\u0160\u00058\u0000\u0000\u0160\u0161\u0003&\u0013\u0000\u0161\u0162"+
		"\u00059\u0000\u0000\u0162\u0163\u0006\u0012\uffff\uffff\u0000\u0163%\u0001"+
		"\u0000\u0000\u0000\u0164\u0165\u0006\u0013\uffff\uffff\u0000\u0165\u0166"+
		"\u0003*\u0015\u0000\u0166\u0167\u0006\u0013\uffff\uffff\u0000\u0167\u016a"+
		"\u0001\u0000\u0000\u0000\u0168\u016a\u0006\u0013\uffff\uffff\u0000\u0169"+
		"\u0164\u0001\u0000\u0000\u0000\u0169\u0168\u0001\u0000\u0000\u0000\u016a"+
		"\u0172\u0001\u0000\u0000\u0000\u016b\u016c\n\u0003\u0000\u0000\u016c\u016d"+
		"\u0005?\u0000\u0000\u016d\u016e\u0003*\u0015\u0000\u016e\u016f\u0006\u0013"+
		"\uffff\uffff\u0000\u016f\u0171\u0001\u0000\u0000\u0000\u0170\u016b\u0001"+
		"\u0000\u0000\u0000\u0171\u0174\u0001\u0000\u0000\u0000\u0172\u0170\u0001"+
		"\u0000\u0000\u0000\u0172\u0173\u0001\u0000\u0000\u0000\u0173\'\u0001\u0000"+
		"\u0000\u0000\u0174\u0172\u0001\u0000\u0000\u0000\u0175\u0176\u0005\u0001"+
		"\u0000\u0000\u0176\u018d\u0006\u0014\uffff\uffff\u0000\u0177\u0178\u0005"+
		"\u0002\u0000\u0000\u0178\u018d\u0006\u0014\uffff\uffff\u0000\u0179\u017a"+
		"\u0005\u0004\u0000\u0000\u017a\u018d\u0006\u0014\uffff\uffff\u0000\u017b"+
		"\u017c\u0005\u0003\u0000\u0000\u017c\u018d\u0006\u0014\uffff\uffff\u0000"+
		"\u017d\u017e\u0005=\u0000\u0000\u017e\u017f\u0003(\u0014\u0000\u017f\u0180"+
		"\u0005>\u0000\u0000\u0180\u0181\u0006\u0014\uffff\uffff\u0000\u0181\u018d"+
		"\u0001\u0000\u0000\u0000\u0182\u0183\u0005A\u0000\u0000\u0183\u0184\u0005"+
		"\u0004\u0000\u0000\u0184\u0185\u0005A\u0000\u0000\u0185\u018d\u0006\u0014"+
		"\uffff\uffff\u0000\u0186\u0187\u0005\u0019\u0000\u0000\u0187\u018d\u0006"+
		"\u0014\uffff\uffff\u0000\u0188\u0189\u0005\u001a\u0000\u0000\u0189\u018d"+
		"\u0006\u0014\uffff\uffff\u0000\u018a\u018b\u0005&\u0000\u0000\u018b\u018d"+
		"\u0006\u0014\uffff\uffff\u0000\u018c\u0175\u0001\u0000\u0000\u0000\u018c"+
		"\u0177\u0001\u0000\u0000\u0000\u018c\u0179\u0001\u0000\u0000\u0000\u018c"+
		"\u017b\u0001\u0000\u0000\u0000\u018c\u017d\u0001\u0000\u0000\u0000\u018c"+
		"\u0182\u0001\u0000\u0000\u0000\u018c\u0186\u0001\u0000\u0000\u0000\u018c"+
		"\u0188\u0001\u0000\u0000\u0000\u018c\u018a\u0001\u0000\u0000\u0000\u018d"+
		")\u0001\u0000\u0000\u0000\u018e\u018f\u0006\u0015\uffff\uffff\u0000\u018f"+
		"\u0190\u00056\u0000\u0000\u0190\u0191\u0003*\u0015\u0017\u0191\u0192\u0006"+
		"\u0015\uffff\uffff\u0000\u0192\u01b8\u0001\u0000\u0000\u0000\u0193\u0194"+
		"\u0005)\u0000\u0000\u0194\u0195\u0003*\u0015\r\u0195\u0196\u0006\u0015"+
		"\uffff\uffff\u0000\u0196\u01b8\u0001\u0000\u0000\u0000\u0197\u0198\u0005"+
		"8\u0000\u0000\u0198\u0199\u0003*\u0015\u0000\u0199\u019a\u00059\u0000"+
		"\u0000\u019a\u019b\u0006\u0015\uffff\uffff\u0000\u019b\u01b8\u0001\u0000"+
		"\u0000\u0000\u019c\u019d\u0003\"\u0011\u0000\u019d\u019e\u0006\u0015\uffff"+
		"\uffff\u0000\u019e\u01b8\u0001\u0000\u0000\u0000\u019f\u01a0\u0005=\u0000"+
		"\u0000\u01a0\u01a1\u0005>\u0000\u0000\u01a1\u01b8\u0006\u0015\uffff\uffff"+
		"\u0000\u01a2\u01a3\u00030\u0018\u0000\u01a3\u01a4\u0006\u0015\uffff\uffff"+
		"\u0000\u01a4\u01b8\u0001\u0000\u0000\u0000\u01a5\u01a6\u0005=\u0000\u0000"+
		"\u01a6\u01a7\u0003,\u0016\u0000\u01a7\u01a8\u0005>\u0000\u0000\u01a8\u01a9"+
		"\u0006\u0015\uffff\uffff\u0000\u01a9\u01b8\u0001\u0000\u0000\u0000\u01aa"+
		"\u01ab\u0005$\u0000\u0000\u01ab\u01b8\u0006\u0015\uffff\uffff\u0000\u01ac"+
		"\u01ad\u0005%\u0000\u0000\u01ad\u01b8\u0006\u0015\uffff\uffff\u0000\u01ae"+
		"\u01af\u0005\t\u0000\u0000\u01af\u01b8\u0006\u0015\uffff\uffff\u0000\u01b0"+
		"\u01b1\u0005\n\u0000\u0000\u01b1\u01b8\u0006\u0015\uffff\uffff\u0000\u01b2"+
		"\u01b3\u00032\u0019\u0000\u01b3\u01b4\u0006\u0015\uffff\uffff\u0000\u01b4"+
		"\u01b8\u0001\u0000\u0000\u0000\u01b5\u01b6\u0005\u0019\u0000\u0000\u01b6"+
		"\u01b8\u0006\u0015\uffff\uffff\u0000\u01b7\u018e\u0001\u0000\u0000\u0000"+
		"\u01b7\u0193\u0001\u0000\u0000\u0000\u01b7\u0197\u0001\u0000\u0000\u0000"+
		"\u01b7\u019c\u0001\u0000\u0000\u0000\u01b7\u019f\u0001\u0000\u0000\u0000"+
		"\u01b7\u01a2\u0001\u0000\u0000\u0000\u01b7\u01a5\u0001\u0000\u0000\u0000"+
		"\u01b7\u01aa\u0001\u0000\u0000\u0000\u01b7\u01ac\u0001\u0000\u0000\u0000"+
		"\u01b7\u01ae\u0001\u0000\u0000\u0000\u01b7\u01b0\u0001\u0000\u0000\u0000"+
		"\u01b7\u01b2\u0001\u0000\u0000\u0000\u01b7\u01b5\u0001\u0000\u0000\u0000"+
		"\u01b8\u01ee\u0001\u0000\u0000\u0000\u01b9\u01ba\n\u0016\u0000\u0000\u01ba"+
		"\u01bb\u0005@\u0000\u0000\u01bb\u01bc\u0005@\u0000\u0000\u01bc\u01bd\u0005"+
		"@\u0000\u0000\u01bd\u01be\u0003*\u0015\u0017\u01be\u01bf\u0006\u0015\uffff"+
		"\uffff\u0000\u01bf\u01ed\u0001\u0000\u0000\u0000\u01c0\u01c1\n\u0015\u0000"+
		"\u0000\u01c1\u01c2\u0007\u0000\u0000\u0000\u01c2\u01c3\u0003*\u0015\u0016"+
		"\u01c3\u01c4\u0006\u0015\uffff\uffff\u0000\u01c4\u01ed\u0001\u0000\u0000"+
		"\u0000\u01c5\u01c6\n\u0014\u0000\u0000\u01c6\u01c7\u0007\u0003\u0000\u0000"+
		"\u01c7\u01c8\u0003*\u0015\u0015\u01c8\u01c9\u0006\u0015\uffff\uffff\u0000"+
		"\u01c9\u01ed\u0001\u0000\u0000\u0000\u01ca\u01cb\n\u0013\u0000\u0000\u01cb"+
		"\u01cc\u0007\u0004\u0000\u0000\u01cc\u01cd\u0003*\u0015\u0014\u01cd\u01ce"+
		"\u0006\u0015\uffff\uffff\u0000\u01ce\u01ed\u0001\u0000\u0000\u0000\u01cf"+
		"\u01d0\n\u0012\u0000\u0000\u01d0\u01d1\u0007\u0005\u0000\u0000\u01d1\u01d2"+
		"\u0003*\u0015\u0013\u01d2\u01d3\u0006\u0015\uffff\uffff\u0000\u01d3\u01ed"+
		"\u0001\u0000\u0000\u0000\u01d4\u01d5\n\u0011\u0000\u0000\u01d5\u01d6\u0007"+
		"\u0006\u0000\u0000\u01d6\u01d7\u0003*\u0015\u0012\u01d7\u01d8\u0006\u0015"+
		"\uffff\uffff\u0000\u01d8\u01ed\u0001\u0000\u0000\u0000\u01d9\u01da\n\u0010"+
		"\u0000\u0000\u01da\u01db\u0007\u0007\u0000\u0000\u01db\u01dc\u0003*\u0015"+
		"\u0011\u01dc\u01dd\u0006\u0015\uffff\uffff\u0000\u01dd\u01ed\u0001\u0000"+
		"\u0000\u0000\u01de\u01df\n\u000f\u0000\u0000\u01df\u01e0\u0005+\u0000"+
		"\u0000\u01e0\u01e1\u0003*\u0015\u0010\u01e1\u01e2\u0006\u0015\uffff\uffff"+
		"\u0000\u01e2\u01ed\u0001\u0000\u0000\u0000\u01e3\u01e4\n\u000e\u0000\u0000"+
		"\u01e4\u01e5\u0005*\u0000\u0000\u01e5\u01e6\u0003*\u0015\u000f\u01e6\u01e7"+
		"\u0006\u0015\uffff\uffff\u0000\u01e7\u01ed\u0001\u0000\u0000\u0000\u01e8"+
		"\u01e9\n\u0002\u0000\u0000\u01e9\u01ea\u0005@\u0000\u0000\u01ea\u01eb"+
		"\u0005\"\u0000\u0000\u01eb\u01ed\u0006\u0015\uffff\uffff\u0000\u01ec\u01b9"+
		"\u0001\u0000\u0000\u0000\u01ec\u01c0\u0001\u0000\u0000\u0000\u01ec\u01c5"+
		"\u0001\u0000\u0000\u0000\u01ec\u01ca\u0001\u0000\u0000\u0000\u01ec\u01cf"+
		"\u0001\u0000\u0000\u0000\u01ec\u01d4\u0001\u0000\u0000\u0000\u01ec\u01d9"+
		"\u0001\u0000\u0000\u0000\u01ec\u01de\u0001\u0000\u0000\u0000\u01ec\u01e3"+
		"\u0001\u0000\u0000\u0000\u01ec\u01e8\u0001\u0000\u0000\u0000\u01ed\u01f0"+
		"\u0001\u0000\u0000\u0000\u01ee\u01ec\u0001\u0000\u0000\u0000\u01ee\u01ef"+
		"\u0001\u0000\u0000\u0000\u01ef+\u0001\u0000\u0000\u0000\u01f0\u01ee\u0001"+
		"\u0000\u0000\u0000\u01f1\u01f2\u0006\u0016\uffff\uffff\u0000\u01f2\u01f3"+
		"\u0003*\u0015\u0000\u01f3\u01f4\u0006\u0016\uffff\uffff\u0000\u01f4\u01fc"+
		"\u0001\u0000\u0000\u0000\u01f5\u01f6\n\u0002\u0000\u0000\u01f6\u01f7\u0005"+
		"?\u0000\u0000\u01f7\u01f8\u0003*\u0015\u0000\u01f8\u01f9\u0006\u0016\uffff"+
		"\uffff\u0000\u01f9\u01fb\u0001\u0000\u0000\u0000\u01fa\u01f5\u0001\u0000"+
		"\u0000\u0000\u01fb\u01fe\u0001\u0000\u0000\u0000\u01fc\u01fa\u0001\u0000"+
		"\u0000\u0000\u01fc\u01fd\u0001\u0000\u0000\u0000\u01fd-\u0001\u0000\u0000"+
		"\u0000\u01fe\u01fc\u0001\u0000\u0000\u0000\u01ff\u0200\u0006\u0017\uffff"+
		"\uffff\u0000\u0200\u0201\u0005=\u0000\u0000\u0201\u0202\u0003*\u0015\u0000"+
		"\u0202\u0203\u0005>\u0000\u0000\u0203\u0204\u0006\u0017\uffff\uffff\u0000"+
		"\u0204\u020d\u0001\u0000\u0000\u0000\u0205\u0206\n\u0002\u0000\u0000\u0206"+
		"\u0207\u0005=\u0000\u0000\u0207\u0208\u0003*\u0015\u0000\u0208\u0209\u0005"+
		">\u0000\u0000\u0209\u020a\u0006\u0017\uffff\uffff\u0000\u020a\u020c\u0001"+
		"\u0000\u0000\u0000\u020b\u0205\u0001\u0000\u0000\u0000\u020c\u020f\u0001"+
		"\u0000\u0000\u0000\u020d\u020b\u0001\u0000\u0000\u0000\u020d\u020e\u0001"+
		"\u0000\u0000\u0000\u020e/\u0001\u0000\u0000\u0000\u020f\u020d\u0001\u0000"+
		"\u0000\u0000\u0210\u0211\u0006\u0018\uffff\uffff\u0000\u0211\u0212\u0005"+
		"&\u0000\u0000\u0212\u0213\u0006\u0018\uffff\uffff\u0000\u0213\u0224\u0001"+
		"\u0000\u0000\u0000\u0214\u0215\n\u0003\u0000\u0000\u0215\u0216\u0005="+
		"\u0000\u0000\u0216\u0217\u0003*\u0015\u0000\u0217\u0218\u0005>\u0000\u0000"+
		"\u0218\u0219\u0006\u0018\uffff\uffff\u0000\u0219\u0223\u0001\u0000\u0000"+
		"\u0000\u021a\u021b\n\u0002\u0000\u0000\u021b\u021c\u0003(\u0014\u0000"+
		"\u021c\u021d\u0005,\u0000\u0000\u021d\u021e\u0005=\u0000\u0000\u021e\u021f"+
		"\u0003*\u0015\u0000\u021f\u0220\u0005>\u0000\u0000\u0220\u0221\u0006\u0018"+
		"\uffff\uffff\u0000\u0221\u0223\u0001\u0000\u0000\u0000\u0222\u0214\u0001"+
		"\u0000\u0000\u0000\u0222\u021a\u0001\u0000\u0000\u0000\u0223\u0226\u0001"+
		"\u0000\u0000\u0000\u0224\u0222\u0001\u0000\u0000\u0000\u0224\u0225\u0001"+
		"\u0000\u0000\u0000\u02251\u0001\u0000\u0000\u0000\u0226\u0224\u0001\u0000"+
		"\u0000\u0000\u0227\u0228\u0006\u0019\uffff\uffff\u0000\u0228\u0229\u0005"+
		"&\u0000\u0000\u0229\u022a\u0006\u0019\uffff\uffff\u0000\u022a\u0231\u0001"+
		"\u0000\u0000\u0000\u022b\u022c\n\u0002\u0000\u0000\u022c\u022d\u0005@"+
		"\u0000\u0000\u022d\u022e\u0005#\u0000\u0000\u022e\u0230\u0006\u0019\uffff"+
		"\uffff\u0000\u022f\u022b\u0001\u0000\u0000\u0000\u0230\u0233\u0001\u0000"+
		"\u0000\u0000\u0231\u022f\u0001\u0000\u0000\u0000\u0231\u0232\u0001\u0000"+
		"\u0000\u0000\u02323\u0001\u0000\u0000\u0000\u0233\u0231\u0001\u0000\u0000"+
		"\u0000\u0234\u0235\u0006\u001a\uffff\uffff\u0000\u0235\u0236\u0003*\u0015"+
		"\u0000\u0236\u0237\u0006\u001a\uffff\uffff\u0000\u0237\u023f\u0001\u0000"+
		"\u0000\u0000\u0238\u0239\n\u0002\u0000\u0000\u0239\u023a\u0005?\u0000"+
		"\u0000\u023a\u023b\u0003*\u0015\u0000\u023b\u023c\u0006\u001a\uffff\uffff"+
		"\u0000\u023c\u023e\u0001\u0000\u0000\u0000\u023d\u0238\u0001\u0000\u0000"+
		"\u0000\u023e\u0241\u0001\u0000\u0000\u0000\u023f\u023d\u0001\u0000\u0000"+
		"\u0000\u023f\u0240\u0001\u0000\u0000\u0000\u02405\u0001\u0000\u0000\u0000"+
		"\u0241\u023f\u0001\u0000\u0000\u0000\u001d=es\u0090\u00c1\u00cd\u00fa"+
		"\u00fe\u0108\u0111\u0117\u011d\u011f\u0138\u013f\u0148\u0156\u0169\u0172"+
		"\u018c\u01b7\u01ec\u01ee\u01fc\u020d\u0222\u0224\u0231\u023f";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}