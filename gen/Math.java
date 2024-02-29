// Generated from //wsl.localhost/Ubuntu-22.04/home/sooden/repos/flux_language/parser/Math.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class Math extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		TEXT=1, TEXT_TYPE=2, BOOLEAN=3, BOOLEAN_TYPE=4, DIGIT=5, NUMBER=6, UNIT=7, 
		SI_POSTFIX=8, NUMBER_TYPE=9, OCTET=10, IPV4=11, IPV4_TYPE=12, VAR_IDENTIFIER=13, 
		COMMON_IDENTIFIER=14, ARITHMETIC_OPERATOR=15, LOGICAL_OPERATOR=16, L_BLOCK=17, 
		R_BLOCK=18, AT=19, WS=20, OP_MULTIPLY=21, OP_DIVIDE=22, OP_MODULO=23, 
		OP_POWER=24, OP_PLUS=25, OP_MINUS=26, OP_AND=27, OP_OR=28, OP_XOR=29, 
		OP_EQUAL=30, OP_NOT_EQUAL=31, OP_LESS=32, OP_LESS_OR_EQUAL=33, OP_GREATER=34, 
		OP_GREATER_OR_EQUAL=35, OP_NOT=36;
	public static final int
		RULE_op_level1 = 0, RULE_op_level2 = 1, RULE_op_level3 = 2, RULE_op_level4 = 3, 
		RULE_op_level5 = 4;
	private static String[] makeRuleNames() {
		return new String[] {
			"op_level1", "op_level2", "op_level3", "op_level4", "op_level5"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'text'", null, "'boolean'", null, null, null, null, "'num'", 
			null, null, "'ipv4'", null, null, null, null, "'{'", "'}'", "'@'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "TEXT", "TEXT_TYPE", "BOOLEAN", "BOOLEAN_TYPE", "DIGIT", "NUMBER", 
			"UNIT", "SI_POSTFIX", "NUMBER_TYPE", "OCTET", "IPV4", "IPV4_TYPE", "VAR_IDENTIFIER", 
			"COMMON_IDENTIFIER", "ARITHMETIC_OPERATOR", "LOGICAL_OPERATOR", "L_BLOCK", 
			"R_BLOCK", "AT", "WS", "OP_MULTIPLY", "OP_DIVIDE", "OP_MODULO", "OP_POWER", 
			"OP_PLUS", "OP_MINUS", "OP_AND", "OP_OR", "OP_XOR", "OP_EQUAL", "OP_NOT_EQUAL", 
			"OP_LESS", "OP_LESS_OR_EQUAL", "OP_GREATER", "OP_GREATER_OR_EQUAL", "OP_NOT"
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
	public String getGrammarFileName() { return "Math.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Math(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Op_level1Context extends ParserRuleContext {
		public TerminalNode OP_MULTIPLY() { return getToken(Math.OP_MULTIPLY, 0); }
		public TerminalNode OP_DIVIDE() { return getToken(Math.OP_DIVIDE, 0); }
		public TerminalNode OP_MODULO() { return getToken(Math.OP_MODULO, 0); }
		public TerminalNode OP_POWER() { return getToken(Math.OP_POWER, 0); }
		public Op_level1Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_level1; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterOp_level1(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitOp_level1(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitOp_level1(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Op_level1Context op_level1() throws RecognitionException {
		Op_level1Context _localctx = new Op_level1Context(_ctx, getState());
		enterRule(_localctx, 0, RULE_op_level1);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(10);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 31457280L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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
	public static class Op_level2Context extends ParserRuleContext {
		public TerminalNode OP_PLUS() { return getToken(Math.OP_PLUS, 0); }
		public TerminalNode OP_MINUS() { return getToken(Math.OP_MINUS, 0); }
		public Op_level2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_level2; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterOp_level2(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitOp_level2(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitOp_level2(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Op_level2Context op_level2() throws RecognitionException {
		Op_level2Context _localctx = new Op_level2Context(_ctx, getState());
		enterRule(_localctx, 2, RULE_op_level2);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(12);
			_la = _input.LA(1);
			if ( !(_la==OP_PLUS || _la==OP_MINUS) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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
	public static class Op_level3Context extends ParserRuleContext {
		public TerminalNode OP_AND() { return getToken(Math.OP_AND, 0); }
		public TerminalNode OP_OR() { return getToken(Math.OP_OR, 0); }
		public TerminalNode OP_XOR() { return getToken(Math.OP_XOR, 0); }
		public Op_level3Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_level3; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterOp_level3(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitOp_level3(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitOp_level3(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Op_level3Context op_level3() throws RecognitionException {
		Op_level3Context _localctx = new Op_level3Context(_ctx, getState());
		enterRule(_localctx, 4, RULE_op_level3);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(14);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 939524096L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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
	public static class Op_level4Context extends ParserRuleContext {
		public TerminalNode OP_EQUAL() { return getToken(Math.OP_EQUAL, 0); }
		public TerminalNode OP_NOT_EQUAL() { return getToken(Math.OP_NOT_EQUAL, 0); }
		public TerminalNode OP_LESS() { return getToken(Math.OP_LESS, 0); }
		public TerminalNode OP_LESS_OR_EQUAL() { return getToken(Math.OP_LESS_OR_EQUAL, 0); }
		public TerminalNode OP_GREATER() { return getToken(Math.OP_GREATER, 0); }
		public TerminalNode OP_GREATER_OR_EQUAL() { return getToken(Math.OP_GREATER_OR_EQUAL, 0); }
		public Op_level4Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_level4; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterOp_level4(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitOp_level4(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitOp_level4(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Op_level4Context op_level4() throws RecognitionException {
		Op_level4Context _localctx = new Op_level4Context(_ctx, getState());
		enterRule(_localctx, 6, RULE_op_level4);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(16);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 67645734912L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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
	public static class Op_level5Context extends ParserRuleContext {
		public TerminalNode OP_NOT() { return getToken(Math.OP_NOT, 0); }
		public Op_level5Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_level5; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterOp_level5(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitOp_level5(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitOp_level5(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Op_level5Context op_level5() throws RecognitionException {
		Op_level5Context _localctx = new Op_level5Context(_ctx, getState());
		enterRule(_localctx, 8, RULE_op_level5);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(18);
			match(OP_NOT);
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

	public static final String _serializedATN =
		"\u0004\u0001$\u0015\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0001"+
		"\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001"+
		"\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0000\u0000\u0005"+
		"\u0000\u0002\u0004\u0006\b\u0000\u0004\u0001\u0000\u0015\u0018\u0001\u0000"+
		"\u0019\u001a\u0001\u0000\u001b\u001d\u0001\u0000\u001e#\u000f\u0000\n"+
		"\u0001\u0000\u0000\u0000\u0002\f\u0001\u0000\u0000\u0000\u0004\u000e\u0001"+
		"\u0000\u0000\u0000\u0006\u0010\u0001\u0000\u0000\u0000\b\u0012\u0001\u0000"+
		"\u0000\u0000\n\u000b\u0007\u0000\u0000\u0000\u000b\u0001\u0001\u0000\u0000"+
		"\u0000\f\r\u0007\u0001\u0000\u0000\r\u0003\u0001\u0000\u0000\u0000\u000e"+
		"\u000f\u0007\u0002\u0000\u0000\u000f\u0005\u0001\u0000\u0000\u0000\u0010"+
		"\u0011\u0007\u0003\u0000\u0000\u0011\u0007\u0001\u0000\u0000\u0000\u0012"+
		"\u0013\u0005$\u0000\u0000\u0013\t\u0001\u0000\u0000\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}