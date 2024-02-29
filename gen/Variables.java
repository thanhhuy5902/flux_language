<<<<<<< HEAD
// Generated from //wsl.localhost/Ubuntu-22.04/home/justintran/repos/flux_language/parser/Variables.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class Variables extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		TEXT=1, TEXT_TYPE=2, BOOLEAN=3, BOOLEAN_TYPE=4, DIGIT=5, NUMBER=6, UNIT=7, 
		SI_POSTFIX=8, NUMBER_TYPE=9, OCTET=10, IPV4=11, IPV4_TYPE=12, VAR_IDENTIFIER=13, 
		COMMON_IDENTIFIER=14, ARITHMETIC_OPERATOR=15, LOGICAL_OPERATOR=16, L_BLOCK=17, 
		R_BLOCK=18, AT=19, WS=20;
	public static final int
		RULE_var_type = 0, RULE_var_name = 1, RULE_var_value = 2, RULE_default_var_declaration = 3, 
		RULE_text_var_declaration = 4, RULE_number_var_declaration = 5, RULE_boolean_var_declaration = 6, 
		RULE_ipv4_var_declaration = 7, RULE_var_declaration = 8;
	private static String[] makeRuleNames() {
		return new String[] {
			"var_type", "var_name", "var_value", "default_var_declaration", "text_var_declaration", 
			"number_var_declaration", "boolean_var_declaration", "ipv4_var_declaration", 
			"var_declaration"
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
			"R_BLOCK", "AT", "WS"
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
	public String getGrammarFileName() { return "Variables.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Variables(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Var_typeContext extends ParserRuleContext {
		public TerminalNode TEXT_TYPE() { return getToken(Variables.TEXT_TYPE, 0); }
		public TerminalNode NUMBER_TYPE() { return getToken(Variables.NUMBER_TYPE, 0); }
		public TerminalNode BOOLEAN_TYPE() { return getToken(Variables.BOOLEAN_TYPE, 0); }
		public TerminalNode IPV4_TYPE() { return getToken(Variables.IPV4_TYPE, 0); }
		public Var_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var_type; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterVar_type(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitVar_type(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitVar_type(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Var_typeContext var_type() throws RecognitionException {
		Var_typeContext _localctx = new Var_typeContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_var_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(18);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 4628L) != 0)) ) {
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
	public static class Var_nameContext extends ParserRuleContext {
		public TerminalNode VAR_IDENTIFIER() { return getToken(Variables.VAR_IDENTIFIER, 0); }
		public Var_nameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var_name; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterVar_name(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitVar_name(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitVar_name(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Var_nameContext var_name() throws RecognitionException {
		Var_nameContext _localctx = new Var_nameContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_var_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(20);
			match(VAR_IDENTIFIER);
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
	public static class Var_valueContext extends ParserRuleContext {
		public TerminalNode IPV4() { return getToken(Variables.IPV4, 0); }
		public TerminalNode TEXT() { return getToken(Variables.TEXT, 0); }
		public TerminalNode NUMBER() { return getToken(Variables.NUMBER, 0); }
		public TerminalNode BOOLEAN() { return getToken(Variables.BOOLEAN, 0); }
		public Var_valueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var_value; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterVar_value(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitVar_value(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitVar_value(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Var_valueContext var_value() throws RecognitionException {
		Var_valueContext _localctx = new Var_valueContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_var_value);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(22);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 2122L) != 0)) ) {
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
	public static class Default_var_declarationContext extends ParserRuleContext {
		public Var_typeContext var_type() {
			return getRuleContext(Var_typeContext.class,0);
		}
		public Var_nameContext var_name() {
			return getRuleContext(Var_nameContext.class,0);
		}
		public TerminalNode L_BLOCK() { return getToken(Variables.L_BLOCK, 0); }
		public TerminalNode R_BLOCK() { return getToken(Variables.R_BLOCK, 0); }
		public Default_var_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_default_var_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterDefault_var_declaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitDefault_var_declaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitDefault_var_declaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Default_var_declarationContext default_var_declaration() throws RecognitionException {
		Default_var_declarationContext _localctx = new Default_var_declarationContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_default_var_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(24);
			var_type();
			setState(25);
			var_name();
			setState(26);
			match(L_BLOCK);
			setState(27);
			match(R_BLOCK);
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
	public static class Text_var_declarationContext extends ParserRuleContext {
		public TerminalNode TEXT_TYPE() { return getToken(Variables.TEXT_TYPE, 0); }
		public Var_nameContext var_name() {
			return getRuleContext(Var_nameContext.class,0);
		}
		public TerminalNode L_BLOCK() { return getToken(Variables.L_BLOCK, 0); }
		public TerminalNode TEXT() { return getToken(Variables.TEXT, 0); }
		public TerminalNode R_BLOCK() { return getToken(Variables.R_BLOCK, 0); }
		public Text_var_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_text_var_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterText_var_declaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitText_var_declaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitText_var_declaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Text_var_declarationContext text_var_declaration() throws RecognitionException {
		Text_var_declarationContext _localctx = new Text_var_declarationContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_text_var_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(29);
			match(TEXT_TYPE);
			setState(30);
			var_name();
			setState(31);
			match(L_BLOCK);
			setState(32);
			match(TEXT);
			setState(33);
			match(R_BLOCK);
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
	public static class Number_var_declarationContext extends ParserRuleContext {
		public TerminalNode NUMBER_TYPE() { return getToken(Variables.NUMBER_TYPE, 0); }
		public Var_nameContext var_name() {
			return getRuleContext(Var_nameContext.class,0);
		}
		public TerminalNode L_BLOCK() { return getToken(Variables.L_BLOCK, 0); }
		public TerminalNode NUMBER() { return getToken(Variables.NUMBER, 0); }
		public TerminalNode R_BLOCK() { return getToken(Variables.R_BLOCK, 0); }
		public Number_var_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_number_var_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterNumber_var_declaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitNumber_var_declaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitNumber_var_declaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Number_var_declarationContext number_var_declaration() throws RecognitionException {
		Number_var_declarationContext _localctx = new Number_var_declarationContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_number_var_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(35);
			match(NUMBER_TYPE);
			setState(36);
			var_name();
			setState(37);
			match(L_BLOCK);
			setState(38);
			match(NUMBER);
			setState(39);
			match(R_BLOCK);
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
	public static class Boolean_var_declarationContext extends ParserRuleContext {
		public TerminalNode BOOLEAN_TYPE() { return getToken(Variables.BOOLEAN_TYPE, 0); }
		public Var_nameContext var_name() {
			return getRuleContext(Var_nameContext.class,0);
		}
		public TerminalNode L_BLOCK() { return getToken(Variables.L_BLOCK, 0); }
		public TerminalNode BOOLEAN() { return getToken(Variables.BOOLEAN, 0); }
		public TerminalNode R_BLOCK() { return getToken(Variables.R_BLOCK, 0); }
		public Boolean_var_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_boolean_var_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterBoolean_var_declaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitBoolean_var_declaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitBoolean_var_declaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Boolean_var_declarationContext boolean_var_declaration() throws RecognitionException {
		Boolean_var_declarationContext _localctx = new Boolean_var_declarationContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_boolean_var_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(41);
			match(BOOLEAN_TYPE);
			setState(42);
			var_name();
			setState(43);
			match(L_BLOCK);
			setState(44);
			match(BOOLEAN);
			setState(45);
			match(R_BLOCK);
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
	public static class Ipv4_var_declarationContext extends ParserRuleContext {
		public TerminalNode IPV4_TYPE() { return getToken(Variables.IPV4_TYPE, 0); }
		public Var_nameContext var_name() {
			return getRuleContext(Var_nameContext.class,0);
		}
		public TerminalNode L_BLOCK() { return getToken(Variables.L_BLOCK, 0); }
		public TerminalNode IPV4() { return getToken(Variables.IPV4, 0); }
		public TerminalNode R_BLOCK() { return getToken(Variables.R_BLOCK, 0); }
		public Ipv4_var_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ipv4_var_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterIpv4_var_declaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitIpv4_var_declaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitIpv4_var_declaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Ipv4_var_declarationContext ipv4_var_declaration() throws RecognitionException {
		Ipv4_var_declarationContext _localctx = new Ipv4_var_declarationContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_ipv4_var_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(47);
			match(IPV4_TYPE);
			setState(48);
			var_name();
			setState(49);
			match(L_BLOCK);
			setState(50);
			match(IPV4);
			setState(51);
			match(R_BLOCK);
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
	public static class Var_declarationContext extends ParserRuleContext {
		public Default_var_declarationContext default_var_declaration() {
			return getRuleContext(Default_var_declarationContext.class,0);
		}
		public Text_var_declarationContext text_var_declaration() {
			return getRuleContext(Text_var_declarationContext.class,0);
		}
		public Number_var_declarationContext number_var_declaration() {
			return getRuleContext(Number_var_declarationContext.class,0);
		}
		public Boolean_var_declarationContext boolean_var_declaration() {
			return getRuleContext(Boolean_var_declarationContext.class,0);
		}
		public Ipv4_var_declarationContext ipv4_var_declaration() {
			return getRuleContext(Ipv4_var_declarationContext.class,0);
		}
		public Var_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var_declaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterVar_declaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitVar_declaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitVar_declaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Var_declarationContext var_declaration() throws RecognitionException {
		Var_declarationContext _localctx = new Var_declarationContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_var_declaration);
		try {
			setState(58);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(53);
				default_var_declaration();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(54);
				text_var_declaration();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(55);
				number_var_declaration();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(56);
				boolean_var_declaration();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(57);
				ipv4_var_declaration();
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

	public static final String _serializedATN =
		"\u0004\u0001\u0014=\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003\b;\b\b\u0001\b\u0000\u0000"+
		"\t\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0000\u0002\u0004\u0000\u0002"+
		"\u0002\u0004\u0004\t\t\f\f\u0004\u0000\u0001\u0001\u0003\u0003\u0006\u0006"+
		"\u000b\u000b7\u0000\u0012\u0001\u0000\u0000\u0000\u0002\u0014\u0001\u0000"+
		"\u0000\u0000\u0004\u0016\u0001\u0000\u0000\u0000\u0006\u0018\u0001\u0000"+
		"\u0000\u0000\b\u001d\u0001\u0000\u0000\u0000\n#\u0001\u0000\u0000\u0000"+
		"\f)\u0001\u0000\u0000\u0000\u000e/\u0001\u0000\u0000\u0000\u0010:\u0001"+
		"\u0000\u0000\u0000\u0012\u0013\u0007\u0000\u0000\u0000\u0013\u0001\u0001"+
		"\u0000\u0000\u0000\u0014\u0015\u0005\r\u0000\u0000\u0015\u0003\u0001\u0000"+
		"\u0000\u0000\u0016\u0017\u0007\u0001\u0000\u0000\u0017\u0005\u0001\u0000"+
		"\u0000\u0000\u0018\u0019\u0003\u0000\u0000\u0000\u0019\u001a\u0003\u0002"+
		"\u0001\u0000\u001a\u001b\u0005\u0011\u0000\u0000\u001b\u001c\u0005\u0012"+
		"\u0000\u0000\u001c\u0007\u0001\u0000\u0000\u0000\u001d\u001e\u0005\u0002"+
		"\u0000\u0000\u001e\u001f\u0003\u0002\u0001\u0000\u001f \u0005\u0011\u0000"+
		"\u0000 !\u0005\u0001\u0000\u0000!\"\u0005\u0012\u0000\u0000\"\t\u0001"+
		"\u0000\u0000\u0000#$\u0005\t\u0000\u0000$%\u0003\u0002\u0001\u0000%&\u0005"+
		"\u0011\u0000\u0000&\'\u0005\u0006\u0000\u0000\'(\u0005\u0012\u0000\u0000"+
		"(\u000b\u0001\u0000\u0000\u0000)*\u0005\u0004\u0000\u0000*+\u0003\u0002"+
		"\u0001\u0000+,\u0005\u0011\u0000\u0000,-\u0005\u0003\u0000\u0000-.\u0005"+
		"\u0012\u0000\u0000.\r\u0001\u0000\u0000\u0000/0\u0005\f\u0000\u000001"+
		"\u0003\u0002\u0001\u000012\u0005\u0011\u0000\u000023\u0005\u000b\u0000"+
		"\u000034\u0005\u0012\u0000\u00004\u000f\u0001\u0000\u0000\u00005;\u0003"+
		"\u0006\u0003\u00006;\u0003\b\u0004\u00007;\u0003\n\u0005\u00008;\u0003"+
		"\f\u0006\u00009;\u0003\u000e\u0007\u0000:5\u0001\u0000\u0000\u0000:6\u0001"+
		"\u0000\u0000\u0000:7\u0001\u0000\u0000\u0000:8\u0001\u0000\u0000\u0000"+
		":9\u0001\u0000\u0000\u0000;\u0011\u0001\u0000\u0000\u0000\u0001:";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
=======
// Generated from /home/sooden/repos/flux_language/parser/Variables.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class Variables extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		TEXT=1, TEXT_TYPE=2, BOOLEAN=3, BOOLEAN_TYPE=4, DIGIT=5, NUMBER=6, UNIT=7, 
		SI_POSTFIX=8, NUMBER_TYPE=9, OCTET=10, IPV4=11, IPV4_TYPE=12, VAR_IDENTIFIER=13, 
		COMMON_IDENTIFIER=14, ARITHMETIC_OPERATOR=15, LOGICAL_OPERATOR=16, L_BLOCK=17, 
		R_BLOCK=18, AT=19, WS=20;
	public static final int
		RULE_varType = 0, RULE_varName = 1, RULE_varValue = 2, RULE_defaultVarDeclaration = 3, 
		RULE_textVarDeclaration = 4, RULE_numberVarDeclaration = 5, RULE_booleanVarDeclaration = 6, 
		RULE_ipv4VarDeclaration = 7, RULE_varDeclaration = 8;
	private static String[] makeRuleNames() {
		return new String[] {
			"varType", "varName", "varValue", "defaultVarDeclaration", "textVarDeclaration", 
			"numberVarDeclaration", "booleanVarDeclaration", "ipv4VarDeclaration", 
			"varDeclaration"
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
			"R_BLOCK", "AT", "WS"
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
	public String getGrammarFileName() { return "Variables.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Variables(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VarTypeContext extends ParserRuleContext {
		public TerminalNode TEXT_TYPE() { return getToken(Variables.TEXT_TYPE, 0); }
		public TerminalNode NUMBER_TYPE() { return getToken(Variables.NUMBER_TYPE, 0); }
		public TerminalNode BOOLEAN_TYPE() { return getToken(Variables.BOOLEAN_TYPE, 0); }
		public TerminalNode IPV4_TYPE() { return getToken(Variables.IPV4_TYPE, 0); }
		public VarTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varType; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterVarType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitVarType(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitVarType(this);
			else return visitor.visitChildren(this);
		}
	}

	public final VarTypeContext varType() throws RecognitionException {
		VarTypeContext _localctx = new VarTypeContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_varType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(18);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 4628L) != 0)) ) {
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
	public static class VarNameContext extends ParserRuleContext {
		public TerminalNode VAR_IDENTIFIER() { return getToken(Variables.VAR_IDENTIFIER, 0); }
		public VarNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varName; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterVarName(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitVarName(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitVarName(this);
			else return visitor.visitChildren(this);
		}
	}

	public final VarNameContext varName() throws RecognitionException {
		VarNameContext _localctx = new VarNameContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_varName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(20);
			match(VAR_IDENTIFIER);
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
	public static class VarValueContext extends ParserRuleContext {
		public TerminalNode IPV4() { return getToken(Variables.IPV4, 0); }
		public TerminalNode TEXT() { return getToken(Variables.TEXT, 0); }
		public TerminalNode NUMBER() { return getToken(Variables.NUMBER, 0); }
		public TerminalNode BOOLEAN() { return getToken(Variables.BOOLEAN, 0); }
		public VarValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varValue; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterVarValue(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitVarValue(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitVarValue(this);
			else return visitor.visitChildren(this);
		}
	}

	public final VarValueContext varValue() throws RecognitionException {
		VarValueContext _localctx = new VarValueContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_varValue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(22);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 2122L) != 0)) ) {
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
	public static class DefaultVarDeclarationContext extends ParserRuleContext {
		public VarTypeContext varType() {
			return getRuleContext(VarTypeContext.class,0);
		}
		public VarNameContext varName() {
			return getRuleContext(VarNameContext.class,0);
		}
		public TerminalNode L_BLOCK() { return getToken(Variables.L_BLOCK, 0); }
		public TerminalNode R_BLOCK() { return getToken(Variables.R_BLOCK, 0); }
		public DefaultVarDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_defaultVarDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterDefaultVarDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitDefaultVarDeclaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitDefaultVarDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final DefaultVarDeclarationContext defaultVarDeclaration() throws RecognitionException {
		DefaultVarDeclarationContext _localctx = new DefaultVarDeclarationContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_defaultVarDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(24);
			varType();
			setState(25);
			varName();
			setState(26);
			match(L_BLOCK);
			setState(27);
			match(R_BLOCK);
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
	public static class TextVarDeclarationContext extends ParserRuleContext {
		public TerminalNode TEXT_TYPE() { return getToken(Variables.TEXT_TYPE, 0); }
		public VarNameContext varName() {
			return getRuleContext(VarNameContext.class,0);
		}
		public TerminalNode L_BLOCK() { return getToken(Variables.L_BLOCK, 0); }
		public TerminalNode TEXT() { return getToken(Variables.TEXT, 0); }
		public TerminalNode R_BLOCK() { return getToken(Variables.R_BLOCK, 0); }
		public TextVarDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_textVarDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterTextVarDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitTextVarDeclaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitTextVarDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TextVarDeclarationContext textVarDeclaration() throws RecognitionException {
		TextVarDeclarationContext _localctx = new TextVarDeclarationContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_textVarDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(29);
			match(TEXT_TYPE);
			setState(30);
			varName();
			setState(31);
			match(L_BLOCK);
			setState(32);
			match(TEXT);
			setState(33);
			match(R_BLOCK);
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
	public static class NumberVarDeclarationContext extends ParserRuleContext {
		public TerminalNode NUMBER_TYPE() { return getToken(Variables.NUMBER_TYPE, 0); }
		public VarNameContext varName() {
			return getRuleContext(VarNameContext.class,0);
		}
		public TerminalNode L_BLOCK() { return getToken(Variables.L_BLOCK, 0); }
		public TerminalNode NUMBER() { return getToken(Variables.NUMBER, 0); }
		public TerminalNode R_BLOCK() { return getToken(Variables.R_BLOCK, 0); }
		public NumberVarDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_numberVarDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterNumberVarDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitNumberVarDeclaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitNumberVarDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final NumberVarDeclarationContext numberVarDeclaration() throws RecognitionException {
		NumberVarDeclarationContext _localctx = new NumberVarDeclarationContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_numberVarDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(35);
			match(NUMBER_TYPE);
			setState(36);
			varName();
			setState(37);
			match(L_BLOCK);
			setState(38);
			match(NUMBER);
			setState(39);
			match(R_BLOCK);
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
	public static class BooleanVarDeclarationContext extends ParserRuleContext {
		public TerminalNode BOOLEAN_TYPE() { return getToken(Variables.BOOLEAN_TYPE, 0); }
		public VarNameContext varName() {
			return getRuleContext(VarNameContext.class,0);
		}
		public TerminalNode L_BLOCK() { return getToken(Variables.L_BLOCK, 0); }
		public TerminalNode BOOLEAN() { return getToken(Variables.BOOLEAN, 0); }
		public TerminalNode R_BLOCK() { return getToken(Variables.R_BLOCK, 0); }
		public BooleanVarDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_booleanVarDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterBooleanVarDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitBooleanVarDeclaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitBooleanVarDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final BooleanVarDeclarationContext booleanVarDeclaration() throws RecognitionException {
		BooleanVarDeclarationContext _localctx = new BooleanVarDeclarationContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_booleanVarDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(41);
			match(BOOLEAN_TYPE);
			setState(42);
			varName();
			setState(43);
			match(L_BLOCK);
			setState(44);
			match(BOOLEAN);
			setState(45);
			match(R_BLOCK);
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
	public static class Ipv4VarDeclarationContext extends ParserRuleContext {
		public TerminalNode IPV4_TYPE() { return getToken(Variables.IPV4_TYPE, 0); }
		public VarNameContext varName() {
			return getRuleContext(VarNameContext.class,0);
		}
		public TerminalNode L_BLOCK() { return getToken(Variables.L_BLOCK, 0); }
		public TerminalNode IPV4() { return getToken(Variables.IPV4, 0); }
		public TerminalNode R_BLOCK() { return getToken(Variables.R_BLOCK, 0); }
		public Ipv4VarDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ipv4VarDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterIpv4VarDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitIpv4VarDeclaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitIpv4VarDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Ipv4VarDeclarationContext ipv4VarDeclaration() throws RecognitionException {
		Ipv4VarDeclarationContext _localctx = new Ipv4VarDeclarationContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_ipv4VarDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(47);
			match(IPV4_TYPE);
			setState(48);
			varName();
			setState(49);
			match(L_BLOCK);
			setState(50);
			match(IPV4);
			setState(51);
			match(R_BLOCK);
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
	public static class VarDeclarationContext extends ParserRuleContext {
		public DefaultVarDeclarationContext defaultVarDeclaration() {
			return getRuleContext(DefaultVarDeclarationContext.class,0);
		}
		public TextVarDeclarationContext textVarDeclaration() {
			return getRuleContext(TextVarDeclarationContext.class,0);
		}
		public NumberVarDeclarationContext numberVarDeclaration() {
			return getRuleContext(NumberVarDeclarationContext.class,0);
		}
		public BooleanVarDeclarationContext booleanVarDeclaration() {
			return getRuleContext(BooleanVarDeclarationContext.class,0);
		}
		public Ipv4VarDeclarationContext ipv4VarDeclaration() {
			return getRuleContext(Ipv4VarDeclarationContext.class,0);
		}
		public VarDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).enterVarDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VariablesListener ) ((VariablesListener)listener).exitVarDeclaration(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof VariablesVisitor ) return ((VariablesVisitor<? extends T>)visitor).visitVarDeclaration(this);
			else return visitor.visitChildren(this);
		}
	}

	public final VarDeclarationContext varDeclaration() throws RecognitionException {
		VarDeclarationContext _localctx = new VarDeclarationContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_varDeclaration);
		try {
			setState(58);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(53);
				defaultVarDeclaration();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(54);
				textVarDeclaration();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(55);
				numberVarDeclaration();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(56);
				booleanVarDeclaration();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(57);
				ipv4VarDeclaration();
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

	public static final String _serializedATN =
		"\u0004\u0001\u0014=\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003\b;\b\b\u0001\b\u0000\u0000"+
		"\t\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0000\u0002\u0004\u0000\u0002"+
		"\u0002\u0004\u0004\t\t\f\f\u0004\u0000\u0001\u0001\u0003\u0003\u0006\u0006"+
		"\u000b\u000b7\u0000\u0012\u0001\u0000\u0000\u0000\u0002\u0014\u0001\u0000"+
		"\u0000\u0000\u0004\u0016\u0001\u0000\u0000\u0000\u0006\u0018\u0001\u0000"+
		"\u0000\u0000\b\u001d\u0001\u0000\u0000\u0000\n#\u0001\u0000\u0000\u0000"+
		"\f)\u0001\u0000\u0000\u0000\u000e/\u0001\u0000\u0000\u0000\u0010:\u0001"+
		"\u0000\u0000\u0000\u0012\u0013\u0007\u0000\u0000\u0000\u0013\u0001\u0001"+
		"\u0000\u0000\u0000\u0014\u0015\u0005\r\u0000\u0000\u0015\u0003\u0001\u0000"+
		"\u0000\u0000\u0016\u0017\u0007\u0001\u0000\u0000\u0017\u0005\u0001\u0000"+
		"\u0000\u0000\u0018\u0019\u0003\u0000\u0000\u0000\u0019\u001a\u0003\u0002"+
		"\u0001\u0000\u001a\u001b\u0005\u0011\u0000\u0000\u001b\u001c\u0005\u0012"+
		"\u0000\u0000\u001c\u0007\u0001\u0000\u0000\u0000\u001d\u001e\u0005\u0002"+
		"\u0000\u0000\u001e\u001f\u0003\u0002\u0001\u0000\u001f \u0005\u0011\u0000"+
		"\u0000 !\u0005\u0001\u0000\u0000!\"\u0005\u0012\u0000\u0000\"\t\u0001"+
		"\u0000\u0000\u0000#$\u0005\t\u0000\u0000$%\u0003\u0002\u0001\u0000%&\u0005"+
		"\u0011\u0000\u0000&\'\u0005\u0006\u0000\u0000\'(\u0005\u0012\u0000\u0000"+
		"(\u000b\u0001\u0000\u0000\u0000)*\u0005\u0004\u0000\u0000*+\u0003\u0002"+
		"\u0001\u0000+,\u0005\u0011\u0000\u0000,-\u0005\u0003\u0000\u0000-.\u0005"+
		"\u0012\u0000\u0000.\r\u0001\u0000\u0000\u0000/0\u0005\f\u0000\u000001"+
		"\u0003\u0002\u0001\u000012\u0005\u0011\u0000\u000023\u0005\u000b\u0000"+
		"\u000034\u0005\u0012\u0000\u00004\u000f\u0001\u0000\u0000\u00005;\u0003"+
		"\u0006\u0003\u00006;\u0003\b\u0004\u00007;\u0003\n\u0005\u00008;\u0003"+
		"\f\u0006\u00009;\u0003\u000e\u0007\u0000:5\u0001\u0000\u0000\u0000:6\u0001"+
		"\u0000\u0000\u0000:7\u0001\u0000\u0000\u0000:8\u0001\u0000\u0000\u0000"+
		":9\u0001\u0000\u0000\u0000;\u0011\u0001\u0000\u0000\u0000\u0001:";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
>>>>>>> 4aa8e30 (Update 27/2/2024)
}