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
		R_BLOCK=18, AT=19, WS=20, ARRAY_TYPE=21, L_BRACKET=22, COMMA=23, R_BRACKET=24;
	public static final int
		RULE_var_type = 0, RULE_var_name = 1, RULE_var_value = 2, RULE_default_var_declaration = 3, 
		RULE_text_var_declaration = 4, RULE_number_var_declaration = 5, RULE_boolean_var_declaration = 6, 
		RULE_ipv4_var_declaration = 7, RULE_array_var_declaration = 8, RULE_var_declaration = 9;
	private static String[] makeRuleNames() {
		return new String[] {
			"var_type", "var_name", "var_value", "default_var_declaration", "text_var_declaration", 
			"number_var_declaration", "boolean_var_declaration", "ipv4_var_declaration", 
			"array_var_declaration", "var_declaration"
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
			"R_BLOCK", "AT", "WS", "ARRAY_TYPE", "L_BRACKET", "COMMA", "R_BRACKET"
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
		public TerminalNode ARRAY_TYPE() { return getToken(Variables.ARRAY_TYPE, 0); }
		public Var_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var_type; }
	}

	public final Var_typeContext var_type() throws RecognitionException {
		Var_typeContext _localctx = new Var_typeContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_var_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(20);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 2101780L) != 0)) ) {
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
	}

	public final Var_nameContext var_name() throws RecognitionException {
		Var_nameContext _localctx = new Var_nameContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_var_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(22);
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
	}

	public final Var_valueContext var_value() throws RecognitionException {
		Var_valueContext _localctx = new Var_valueContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_var_value);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(24);
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
	}

	public final Default_var_declarationContext default_var_declaration() throws RecognitionException {
		Default_var_declarationContext _localctx = new Default_var_declarationContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_default_var_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(26);
			var_type();
			setState(27);
			var_name();
			setState(28);
			match(L_BLOCK);
			setState(29);
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
	}

	public final Text_var_declarationContext text_var_declaration() throws RecognitionException {
		Text_var_declarationContext _localctx = new Text_var_declarationContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_text_var_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(31);
			match(TEXT_TYPE);
			setState(32);
			var_name();
			setState(33);
			match(L_BLOCK);
			setState(34);
			match(TEXT);
			setState(35);
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
	}

	public final Number_var_declarationContext number_var_declaration() throws RecognitionException {
		Number_var_declarationContext _localctx = new Number_var_declarationContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_number_var_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(37);
			match(NUMBER_TYPE);
			setState(38);
			var_name();
			setState(39);
			match(L_BLOCK);
			setState(40);
			match(NUMBER);
			setState(41);
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
	}

	public final Boolean_var_declarationContext boolean_var_declaration() throws RecognitionException {
		Boolean_var_declarationContext _localctx = new Boolean_var_declarationContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_boolean_var_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(43);
			match(BOOLEAN_TYPE);
			setState(44);
			var_name();
			setState(45);
			match(L_BLOCK);
			setState(46);
			match(BOOLEAN);
			setState(47);
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
	}

	public final Ipv4_var_declarationContext ipv4_var_declaration() throws RecognitionException {
		Ipv4_var_declarationContext _localctx = new Ipv4_var_declarationContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_ipv4_var_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(49);
			match(IPV4_TYPE);
			setState(50);
			var_name();
			setState(51);
			match(L_BLOCK);
			setState(52);
			match(IPV4);
			setState(53);
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
	public static class Array_var_declarationContext extends ParserRuleContext {
		public TerminalNode ARRAY_TYPE() { return getToken(Variables.ARRAY_TYPE, 0); }
		public Var_nameContext var_name() {
			return getRuleContext(Var_nameContext.class,0);
		}
		public TerminalNode L_BRACKET() { return getToken(Variables.L_BRACKET, 0); }
		public List<Var_valueContext> var_value() {
			return getRuleContexts(Var_valueContext.class);
		}
		public Var_valueContext var_value(int i) {
			return getRuleContext(Var_valueContext.class,i);
		}
		public TerminalNode R_BRACKET() { return getToken(Variables.R_BRACKET, 0); }
		public List<TerminalNode> COMMA() { return getTokens(Variables.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(Variables.COMMA, i);
		}
		public Array_var_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array_var_declaration; }
	}

	public final Array_var_declarationContext array_var_declaration() throws RecognitionException {
		Array_var_declarationContext _localctx = new Array_var_declarationContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_array_var_declaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(55);
			match(ARRAY_TYPE);
			setState(56);
			var_name();
			setState(57);
			match(L_BRACKET);
			setState(58);
			var_value();
			setState(63);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(59);
				match(COMMA);
				setState(60);
				var_value();
				}
				}
				setState(65);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(66);
			match(R_BRACKET);
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
		public Array_var_declarationContext array_var_declaration() {
			return getRuleContext(Array_var_declarationContext.class,0);
		}
		public Var_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var_declaration; }
	}

	public final Var_declarationContext var_declaration() throws RecognitionException {
		Var_declarationContext _localctx = new Var_declarationContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_var_declaration);
		try {
			setState(74);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(68);
				default_var_declaration();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(69);
				text_var_declaration();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(70);
				number_var_declaration();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(71);
				boolean_var_declaration();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(72);
				ipv4_var_declaration();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(73);
				array_var_declaration();
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
		"\u0004\u0001\u0018M\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001"+
		"\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0005\b>"+
		"\b\b\n\b\f\bA\t\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0003\tK\b\t\u0001\t\u0000\u0000\n\u0000\u0002\u0004\u0006"+
		"\b\n\f\u000e\u0010\u0012\u0000\u0002\u0005\u0000\u0002\u0002\u0004\u0004"+
		"\t\t\f\f\u0015\u0015\u0004\u0000\u0001\u0001\u0003\u0003\u0006\u0006\u000b"+
		"\u000bH\u0000\u0014\u0001\u0000\u0000\u0000\u0002\u0016\u0001\u0000\u0000"+
		"\u0000\u0004\u0018\u0001\u0000\u0000\u0000\u0006\u001a\u0001\u0000\u0000"+
		"\u0000\b\u001f\u0001\u0000\u0000\u0000\n%\u0001\u0000\u0000\u0000\f+\u0001"+
		"\u0000\u0000\u0000\u000e1\u0001\u0000\u0000\u0000\u00107\u0001\u0000\u0000"+
		"\u0000\u0012J\u0001\u0000\u0000\u0000\u0014\u0015\u0007\u0000\u0000\u0000"+
		"\u0015\u0001\u0001\u0000\u0000\u0000\u0016\u0017\u0005\r\u0000\u0000\u0017"+
		"\u0003\u0001\u0000\u0000\u0000\u0018\u0019\u0007\u0001\u0000\u0000\u0019"+
		"\u0005\u0001\u0000\u0000\u0000\u001a\u001b\u0003\u0000\u0000\u0000\u001b"+
		"\u001c\u0003\u0002\u0001\u0000\u001c\u001d\u0005\u0011\u0000\u0000\u001d"+
		"\u001e\u0005\u0012\u0000\u0000\u001e\u0007\u0001\u0000\u0000\u0000\u001f"+
		" \u0005\u0002\u0000\u0000 !\u0003\u0002\u0001\u0000!\"\u0005\u0011\u0000"+
		"\u0000\"#\u0005\u0001\u0000\u0000#$\u0005\u0012\u0000\u0000$\t\u0001\u0000"+
		"\u0000\u0000%&\u0005\t\u0000\u0000&\'\u0003\u0002\u0001\u0000\'(\u0005"+
		"\u0011\u0000\u0000()\u0005\u0006\u0000\u0000)*\u0005\u0012\u0000\u0000"+
		"*\u000b\u0001\u0000\u0000\u0000+,\u0005\u0004\u0000\u0000,-\u0003\u0002"+
		"\u0001\u0000-.\u0005\u0011\u0000\u0000./\u0005\u0003\u0000\u0000/0\u0005"+
		"\u0012\u0000\u00000\r\u0001\u0000\u0000\u000012\u0005\f\u0000\u000023"+
		"\u0003\u0002\u0001\u000034\u0005\u0011\u0000\u000045\u0005\u000b\u0000"+
		"\u000056\u0005\u0012\u0000\u00006\u000f\u0001\u0000\u0000\u000078\u0005"+
		"\u0015\u0000\u000089\u0003\u0002\u0001\u00009:\u0005\u0016\u0000\u0000"+
		":?\u0003\u0004\u0002\u0000;<\u0005\u0017\u0000\u0000<>\u0003\u0004\u0002"+
		"\u0000=;\u0001\u0000\u0000\u0000>A\u0001\u0000\u0000\u0000?=\u0001\u0000"+
		"\u0000\u0000?@\u0001\u0000\u0000\u0000@B\u0001\u0000\u0000\u0000A?\u0001"+
		"\u0000\u0000\u0000BC\u0005\u0018\u0000\u0000C\u0011\u0001\u0000\u0000"+
		"\u0000DK\u0003\u0006\u0003\u0000EK\u0003\b\u0004\u0000FK\u0003\n\u0005"+
		"\u0000GK\u0003\f\u0006\u0000HK\u0003\u000e\u0007\u0000IK\u0003\u0010\b"+
		"\u0000JD\u0001\u0000\u0000\u0000JE\u0001\u0000\u0000\u0000JF\u0001\u0000"+
		"\u0000\u0000JG\u0001\u0000\u0000\u0000JH\u0001\u0000\u0000\u0000JI\u0001"+
		"\u0000\u0000\u0000K\u0013\u0001\u0000\u0000\u0000\u0002?J";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}