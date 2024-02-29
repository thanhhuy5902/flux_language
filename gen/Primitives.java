// Generated from //wsl.localhost/Ubuntu-22.04/home/sooden/repos/flux_language/parser/Primitives.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class Primitives extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		TEXT=1, TEXT_TYPE=2, BOOLEAN=3, BOOLEAN_TYPE=4, NUMBER=5, NUMBER_TYPE=6, 
		NULL=7, DIGIT=8, OCTET=9, IPV4=10, IPV4_TYPE=11, ARRAY=12, ARRAY_TYPE=13, 
		LOOP=14, IF=15, ELSE=16, FUNC=17, RETURN=18, RETURN_TYPE=19, L_BLOCK=20, 
		R_BLOCK=21, L_PAREN=22, R_PAREN=23, L_SQUARE=24, R_SQUARE=25, DOT=26, 
		COLON=27, SEMICOLON=28, COMMA=29, AT=30, OP_PLUS=31, OP_MINUS=32, OP_MULTIPLY=33, 
		OP_DIVIDE=34, OP_MOD=35, OP_POWER=36, OP_EQUAL=37, OP_NOT_EQUAL=38, OP_LESS=39, 
		OP_GREATER=40, OP_LESS_EQUAL=41, OP_GREATER_EQUAL=42, OP_AND=43, OP_OR=44, 
		OP_XOR=45, OP_NOT=46, OP_INCREMENT=47, OP_DECREMENT=48, OP_ADD_ASSIGN=49, 
		OP_SUB_ASSIGN=50, OP_MUL_ASSIGN=51, OP_DIV_ASSIGN=52, OP_MOD_ASSIGN=53, 
		VAR_IDENTIFIER=54, COMMON_IDENTIFIER=55, NEWLINE=56, WS=57;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"TEXT", "TEXT_TYPE", "BOOLEAN", "BOOLEAN_TYPE", "NUMBER", "NUMBER_TYPE", 
			"NULL", "DIGIT", "OCTET", "IPV4", "IPV4_TYPE", "ARRAY", "ARRAY_TYPE", 
			"LOOP", "IF", "ELSE", "FUNC", "RETURN", "RETURN_TYPE", "L_BLOCK", "R_BLOCK", 
			"L_PAREN", "R_PAREN", "L_SQUARE", "R_SQUARE", "DOT", "COLON", "SEMICOLON", 
			"COMMA", "AT", "OP_PLUS", "OP_MINUS", "OP_MULTIPLY", "OP_DIVIDE", "OP_MOD", 
			"OP_POWER", "OP_EQUAL", "OP_NOT_EQUAL", "OP_LESS", "OP_GREATER", "OP_LESS_EQUAL", 
			"OP_GREATER_EQUAL", "OP_AND", "OP_OR", "OP_XOR", "OP_NOT", "OP_INCREMENT", 
			"OP_DECREMENT", "OP_ADD_ASSIGN", "OP_SUB_ASSIGN", "OP_MUL_ASSIGN", "OP_DIV_ASSIGN", 
			"OP_MOD_ASSIGN", "VAR_IDENTIFIER", "COMMON_IDENTIFIER", "NEWLINE", "WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'text'", null, "'boolean'", null, "'num'", "'na'", null, 
			null, null, "'ipv4'", null, "'array'", "'loop'", "'if'", "'else'", "'fun'", 
			"'return'", "'->'", "'{'", "'}'", "'('", "')'", "'['", "']'", "'.'", 
			"':'", "';'", "','", "'@'", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", 
			"'='", "'!='", "'<'", "'>'", "'<='", "'>='", "'and'", "'or'", "'xor'", 
			"'not'", "'++'", "'--'", "'+='", "'-='", "'*='", "'/='", "'%='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "TEXT", "TEXT_TYPE", "BOOLEAN", "BOOLEAN_TYPE", "NUMBER", "NUMBER_TYPE", 
			"NULL", "DIGIT", "OCTET", "IPV4", "IPV4_TYPE", "ARRAY", "ARRAY_TYPE", 
			"LOOP", "IF", "ELSE", "FUNC", "RETURN", "RETURN_TYPE", "L_BLOCK", "R_BLOCK", 
			"L_PAREN", "R_PAREN", "L_SQUARE", "R_SQUARE", "DOT", "COLON", "SEMICOLON", 
			"COMMA", "AT", "OP_PLUS", "OP_MINUS", "OP_MULTIPLY", "OP_DIVIDE", "OP_MOD", 
			"OP_POWER", "OP_EQUAL", "OP_NOT_EQUAL", "OP_LESS", "OP_GREATER", "OP_LESS_EQUAL", 
			"OP_GREATER_EQUAL", "OP_AND", "OP_OR", "OP_XOR", "OP_NOT", "OP_INCREMENT", 
			"OP_DECREMENT", "OP_ADD_ASSIGN", "OP_SUB_ASSIGN", "OP_MUL_ASSIGN", "OP_DIV_ASSIGN", 
			"OP_MOD_ASSIGN", "VAR_IDENTIFIER", "COMMON_IDENTIFIER", "NEWLINE", "WS"
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


	public Primitives(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Primitives.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u00009\u016f\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002"+
		"\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002"+
		"\u0012\u0007\u0012\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002"+
		"\u0015\u0007\u0015\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002"+
		"\u0018\u0007\u0018\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002"+
		"\u001b\u0007\u001b\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002"+
		"\u001e\u0007\u001e\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007"+
		"!\u0002\"\u0007\"\u0002#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007"+
		"&\u0002\'\u0007\'\u0002(\u0007(\u0002)\u0007)\u0002*\u0007*\u0002+\u0007"+
		"+\u0002,\u0007,\u0002-\u0007-\u0002.\u0007.\u0002/\u0007/\u00020\u0007"+
		"0\u00021\u00071\u00022\u00072\u00023\u00073\u00024\u00074\u00025\u0007"+
		"5\u00026\u00076\u00027\u00077\u00028\u00078\u0001\u0000\u0001\u0000\u0005"+
		"\u0000v\b\u0000\n\u0000\f\u0000y\t\u0000\u0001\u0000\u0001\u0000\u0001"+
		"\u0000\u0005\u0000~\b\u0000\n\u0000\f\u0000\u0081\t\u0000\u0001\u0000"+
		"\u0003\u0000\u0084\b\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002\u0094\b\u0002"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0004\u0003\u0004\u009f\b\u0004\u0001\u0004"+
		"\u0004\u0004\u00a2\b\u0004\u000b\u0004\f\u0004\u00a3\u0001\u0004\u0001"+
		"\u0004\u0005\u0004\u00a8\b\u0004\n\u0004\f\u0004\u00ab\t\u0004\u0003\u0004"+
		"\u00ad\b\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001"+
		"\b\u0003\b\u00bb\b\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003\b\u00c8\b\b\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0005\u000b"+
		"\u00db\b\u000b\n\u000b\f\u000b\u00de\t\u000b\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0014"+
		"\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0016\u0001\u0016\u0001\u0017"+
		"\u0001\u0017\u0001\u0018\u0001\u0018\u0001\u0019\u0001\u0019\u0001\u001a"+
		"\u0001\u001a\u0001\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0001\u001d"+
		"\u0001\u001d\u0001\u001e\u0001\u001e\u0001\u001f\u0001\u001f\u0001 \u0001"+
		" \u0001!\u0001!\u0001\"\u0001\"\u0001#\u0001#\u0001$\u0001$\u0001%\u0001"+
		"%\u0001%\u0001&\u0001&\u0001\'\u0001\'\u0001(\u0001(\u0001(\u0001)\u0001"+
		")\u0001)\u0001*\u0001*\u0001*\u0001*\u0001+\u0001+\u0001+\u0001,\u0001"+
		",\u0001,\u0001,\u0001-\u0001-\u0001-\u0001-\u0001.\u0001.\u0001.\u0001"+
		"/\u0001/\u0001/\u00010\u00010\u00010\u00011\u00011\u00011\u00012\u0001"+
		"2\u00012\u00013\u00013\u00013\u00014\u00014\u00014\u00015\u00015\u0005"+
		"5\u0158\b5\n5\f5\u015b\t5\u00016\u00016\u00056\u015f\b6\n6\f6\u0162\t"+
		"6\u00017\u00037\u0165\b7\u00017\u00017\u00018\u00048\u016a\b8\u000b8\f"+
		"8\u016b\u00018\u00018\u0001\u007f\u00009\u0001\u0001\u0003\u0002\u0005"+
		"\u0003\u0007\u0004\t\u0005\u000b\u0006\r\u0007\u000f\b\u0011\t\u0013\n"+
		"\u0015\u000b\u0017\f\u0019\r\u001b\u000e\u001d\u000f\u001f\u0010!\u0011"+
		"#\u0012%\u0013\'\u0014)\u0015+\u0016-\u0017/\u00181\u00193\u001a5\u001b"+
		"7\u001c9\u001d;\u001e=\u001f? A!C\"E#G$I%K&M\'O(Q)S*U+W,Y-[.]/_0a1c2e"+
		"3g4i5k6m7o8q9\u0001\u0000\t\u0001\u0000\"\"\u0001\u000009\u0001\u0000"+
		"19\u0001\u000004\u0001\u000005\u0001\u0000az\u0003\u000009AZaz\u0001\u0000"+
		"AZ\u0003\u0000\t\n\r\r  \u0183\u0000\u0001\u0001\u0000\u0000\u0000\u0000"+
		"\u0003\u0001\u0000\u0000\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0000"+
		"\u0007\u0001\u0000\u0000\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b"+
		"\u0001\u0000\u0000\u0000\u0000\r\u0001\u0000\u0000\u0000\u0000\u000f\u0001"+
		"\u0000\u0000\u0000\u0000\u0011\u0001\u0000\u0000\u0000\u0000\u0013\u0001"+
		"\u0000\u0000\u0000\u0000\u0015\u0001\u0000\u0000\u0000\u0000\u0017\u0001"+
		"\u0000\u0000\u0000\u0000\u0019\u0001\u0000\u0000\u0000\u0000\u001b\u0001"+
		"\u0000\u0000\u0000\u0000\u001d\u0001\u0000\u0000\u0000\u0000\u001f\u0001"+
		"\u0000\u0000\u0000\u0000!\u0001\u0000\u0000\u0000\u0000#\u0001\u0000\u0000"+
		"\u0000\u0000%\u0001\u0000\u0000\u0000\u0000\'\u0001\u0000\u0000\u0000"+
		"\u0000)\u0001\u0000\u0000\u0000\u0000+\u0001\u0000\u0000\u0000\u0000-"+
		"\u0001\u0000\u0000\u0000\u0000/\u0001\u0000\u0000\u0000\u00001\u0001\u0000"+
		"\u0000\u0000\u00003\u0001\u0000\u0000\u0000\u00005\u0001\u0000\u0000\u0000"+
		"\u00007\u0001\u0000\u0000\u0000\u00009\u0001\u0000\u0000\u0000\u0000;"+
		"\u0001\u0000\u0000\u0000\u0000=\u0001\u0000\u0000\u0000\u0000?\u0001\u0000"+
		"\u0000\u0000\u0000A\u0001\u0000\u0000\u0000\u0000C\u0001\u0000\u0000\u0000"+
		"\u0000E\u0001\u0000\u0000\u0000\u0000G\u0001\u0000\u0000\u0000\u0000I"+
		"\u0001\u0000\u0000\u0000\u0000K\u0001\u0000\u0000\u0000\u0000M\u0001\u0000"+
		"\u0000\u0000\u0000O\u0001\u0000\u0000\u0000\u0000Q\u0001\u0000\u0000\u0000"+
		"\u0000S\u0001\u0000\u0000\u0000\u0000U\u0001\u0000\u0000\u0000\u0000W"+
		"\u0001\u0000\u0000\u0000\u0000Y\u0001\u0000\u0000\u0000\u0000[\u0001\u0000"+
		"\u0000\u0000\u0000]\u0001\u0000\u0000\u0000\u0000_\u0001\u0000\u0000\u0000"+
		"\u0000a\u0001\u0000\u0000\u0000\u0000c\u0001\u0000\u0000\u0000\u0000e"+
		"\u0001\u0000\u0000\u0000\u0000g\u0001\u0000\u0000\u0000\u0000i\u0001\u0000"+
		"\u0000\u0000\u0000k\u0001\u0000\u0000\u0000\u0000m\u0001\u0000\u0000\u0000"+
		"\u0000o\u0001\u0000\u0000\u0000\u0000q\u0001\u0000\u0000\u0000\u0001\u0083"+
		"\u0001\u0000\u0000\u0000\u0003\u0085\u0001\u0000\u0000\u0000\u0005\u0093"+
		"\u0001\u0000\u0000\u0000\u0007\u0095\u0001\u0000\u0000\u0000\t\u009e\u0001"+
		"\u0000\u0000\u0000\u000b\u00ae\u0001\u0000\u0000\u0000\r\u00b2\u0001\u0000"+
		"\u0000\u0000\u000f\u00b5\u0001\u0000\u0000\u0000\u0011\u00c7\u0001\u0000"+
		"\u0000\u0000\u0013\u00c9\u0001\u0000\u0000\u0000\u0015\u00d1\u0001\u0000"+
		"\u0000\u0000\u0017\u00dc\u0001\u0000\u0000\u0000\u0019\u00df\u0001\u0000"+
		"\u0000\u0000\u001b\u00e5\u0001\u0000\u0000\u0000\u001d\u00ea\u0001\u0000"+
		"\u0000\u0000\u001f\u00ed\u0001\u0000\u0000\u0000!\u00f2\u0001\u0000\u0000"+
		"\u0000#\u00f6\u0001\u0000\u0000\u0000%\u00fd\u0001\u0000\u0000\u0000\'"+
		"\u0100\u0001\u0000\u0000\u0000)\u0102\u0001\u0000\u0000\u0000+\u0104\u0001"+
		"\u0000\u0000\u0000-\u0106\u0001\u0000\u0000\u0000/\u0108\u0001\u0000\u0000"+
		"\u00001\u010a\u0001\u0000\u0000\u00003\u010c\u0001\u0000\u0000\u00005"+
		"\u010e\u0001\u0000\u0000\u00007\u0110\u0001\u0000\u0000\u00009\u0112\u0001"+
		"\u0000\u0000\u0000;\u0114\u0001\u0000\u0000\u0000=\u0116\u0001\u0000\u0000"+
		"\u0000?\u0118\u0001\u0000\u0000\u0000A\u011a\u0001\u0000\u0000\u0000C"+
		"\u011c\u0001\u0000\u0000\u0000E\u011e\u0001\u0000\u0000\u0000G\u0120\u0001"+
		"\u0000\u0000\u0000I\u0122\u0001\u0000\u0000\u0000K\u0124\u0001\u0000\u0000"+
		"\u0000M\u0127\u0001\u0000\u0000\u0000O\u0129\u0001\u0000\u0000\u0000Q"+
		"\u012b\u0001\u0000\u0000\u0000S\u012e\u0001\u0000\u0000\u0000U\u0131\u0001"+
		"\u0000\u0000\u0000W\u0135\u0001\u0000\u0000\u0000Y\u0138\u0001\u0000\u0000"+
		"\u0000[\u013c\u0001\u0000\u0000\u0000]\u0140\u0001\u0000\u0000\u0000_"+
		"\u0143\u0001\u0000\u0000\u0000a\u0146\u0001\u0000\u0000\u0000c\u0149\u0001"+
		"\u0000\u0000\u0000e\u014c\u0001\u0000\u0000\u0000g\u014f\u0001\u0000\u0000"+
		"\u0000i\u0152\u0001\u0000\u0000\u0000k\u0155\u0001\u0000\u0000\u0000m"+
		"\u015c\u0001\u0000\u0000\u0000o\u0164\u0001\u0000\u0000\u0000q\u0169\u0001"+
		"\u0000\u0000\u0000sw\u0005\"\u0000\u0000tv\b\u0000\u0000\u0000ut\u0001"+
		"\u0000\u0000\u0000vy\u0001\u0000\u0000\u0000wu\u0001\u0000\u0000\u0000"+
		"wx\u0001\u0000\u0000\u0000xz\u0001\u0000\u0000\u0000yw\u0001\u0000\u0000"+
		"\u0000z\u0084\u0005\"\u0000\u0000{\u007f\u0005\'\u0000\u0000|~\t\u0000"+
		"\u0000\u0000}|\u0001\u0000\u0000\u0000~\u0081\u0001\u0000\u0000\u0000"+
		"\u007f\u0080\u0001\u0000\u0000\u0000\u007f}\u0001\u0000\u0000\u0000\u0080"+
		"\u0082\u0001\u0000\u0000\u0000\u0081\u007f\u0001\u0000\u0000\u0000\u0082"+
		"\u0084\u0005\'\u0000\u0000\u0083s\u0001\u0000\u0000\u0000\u0083{\u0001"+
		"\u0000\u0000\u0000\u0084\u0002\u0001\u0000\u0000\u0000\u0085\u0086\u0005"+
		"t\u0000\u0000\u0086\u0087\u0005e\u0000\u0000\u0087\u0088\u0005x\u0000"+
		"\u0000\u0088\u0089\u0005t\u0000\u0000\u0089\u0004\u0001\u0000\u0000\u0000"+
		"\u008a\u008b\u0005t\u0000\u0000\u008b\u008c\u0005r\u0000\u0000\u008c\u008d"+
		"\u0005u\u0000\u0000\u008d\u0094\u0005e\u0000\u0000\u008e\u008f\u0005f"+
		"\u0000\u0000\u008f\u0090\u0005a\u0000\u0000\u0090\u0091\u0005l\u0000\u0000"+
		"\u0091\u0092\u0005s\u0000\u0000\u0092\u0094\u0005e\u0000\u0000\u0093\u008a"+
		"\u0001\u0000\u0000\u0000\u0093\u008e\u0001\u0000\u0000\u0000\u0094\u0006"+
		"\u0001\u0000\u0000\u0000\u0095\u0096\u0005b\u0000\u0000\u0096\u0097\u0005"+
		"o\u0000\u0000\u0097\u0098\u0005o\u0000\u0000\u0098\u0099\u0005l\u0000"+
		"\u0000\u0099\u009a\u0005e\u0000\u0000\u009a\u009b\u0005a\u0000\u0000\u009b"+
		"\u009c\u0005n\u0000\u0000\u009c\b\u0001\u0000\u0000\u0000\u009d\u009f"+
		"\u0005-\u0000\u0000\u009e\u009d\u0001\u0000\u0000\u0000\u009e\u009f\u0001"+
		"\u0000\u0000\u0000\u009f\u00a1\u0001\u0000\u0000\u0000\u00a0\u00a2\u0007"+
		"\u0001\u0000\u0000\u00a1\u00a0\u0001\u0000\u0000\u0000\u00a2\u00a3\u0001"+
		"\u0000\u0000\u0000\u00a3\u00a1\u0001\u0000\u0000\u0000\u00a3\u00a4\u0001"+
		"\u0000\u0000\u0000\u00a4\u00ac\u0001\u0000\u0000\u0000\u00a5\u00a9\u0005"+
		".\u0000\u0000\u00a6\u00a8\u0007\u0001\u0000\u0000\u00a7\u00a6\u0001\u0000"+
		"\u0000\u0000\u00a8\u00ab\u0001\u0000\u0000\u0000\u00a9\u00a7\u0001\u0000"+
		"\u0000\u0000\u00a9\u00aa\u0001\u0000\u0000\u0000\u00aa\u00ad\u0001\u0000"+
		"\u0000\u0000\u00ab\u00a9\u0001\u0000\u0000\u0000\u00ac\u00a5\u0001\u0000"+
		"\u0000\u0000\u00ac\u00ad\u0001\u0000\u0000\u0000\u00ad\n\u0001\u0000\u0000"+
		"\u0000\u00ae\u00af\u0005n\u0000\u0000\u00af\u00b0\u0005u\u0000\u0000\u00b0"+
		"\u00b1\u0005m\u0000\u0000\u00b1\f\u0001\u0000\u0000\u0000\u00b2\u00b3"+
		"\u0005n\u0000\u0000\u00b3\u00b4\u0005a\u0000\u0000\u00b4\u000e\u0001\u0000"+
		"\u0000\u0000\u00b5\u00b6\u0007\u0001\u0000\u0000\u00b6\u0010\u0001\u0000"+
		"\u0000\u0000\u00b7\u00c8\u00050\u0000\u0000\u00b8\u00ba\u0007\u0002\u0000"+
		"\u0000\u00b9\u00bb\u0003\u000f\u0007\u0000\u00ba\u00b9\u0001\u0000\u0000"+
		"\u0000\u00ba\u00bb\u0001\u0000\u0000\u0000\u00bb\u00c8\u0001\u0000\u0000"+
		"\u0000\u00bc\u00bd\u00051\u0000\u0000\u00bd\u00be\u0003\u000f\u0007\u0000"+
		"\u00be\u00bf\u0003\u000f\u0007\u0000\u00bf\u00c8\u0001\u0000\u0000\u0000"+
		"\u00c0\u00c1\u00052\u0000\u0000\u00c1\u00c2\u0007\u0003\u0000\u0000\u00c2"+
		"\u00c8\u0003\u000f\u0007\u0000\u00c3\u00c4\u00052\u0000\u0000\u00c4\u00c5"+
		"\u00055\u0000\u0000\u00c5\u00c6\u0001\u0000\u0000\u0000\u00c6\u00c8\u0007"+
		"\u0004\u0000\u0000\u00c7\u00b7\u0001\u0000\u0000\u0000\u00c7\u00b8\u0001"+
		"\u0000\u0000\u0000\u00c7\u00bc\u0001\u0000\u0000\u0000\u00c7\u00c0\u0001"+
		"\u0000\u0000\u0000\u00c7\u00c3\u0001\u0000\u0000\u0000\u00c8\u0012\u0001"+
		"\u0000\u0000\u0000\u00c9\u00ca\u0003\u0011\b\u0000\u00ca\u00cb\u0005."+
		"\u0000\u0000\u00cb\u00cc\u0003\u0011\b\u0000\u00cc\u00cd\u0005.\u0000"+
		"\u0000\u00cd\u00ce\u0003\u0011\b\u0000\u00ce\u00cf\u0005.\u0000\u0000"+
		"\u00cf\u00d0\u0003\u0011\b\u0000\u00d0\u0014\u0001\u0000\u0000\u0000\u00d1"+
		"\u00d2\u0005i\u0000\u0000\u00d2\u00d3\u0005p\u0000\u0000\u00d3\u00d4\u0005"+
		"v\u0000\u0000\u00d4\u00d5\u00054\u0000\u0000\u00d5\u0016\u0001\u0000\u0000"+
		"\u0000\u00d6\u00db\u0003\u0001\u0000\u0000\u00d7\u00db\u0003\t\u0004\u0000"+
		"\u00d8\u00db\u0003\u0005\u0002\u0000\u00d9\u00db\u0003\u0013\t\u0000\u00da"+
		"\u00d6\u0001\u0000\u0000\u0000\u00da\u00d7\u0001\u0000\u0000\u0000\u00da"+
		"\u00d8\u0001\u0000\u0000\u0000\u00da\u00d9\u0001\u0000\u0000\u0000\u00db"+
		"\u00de\u0001\u0000\u0000\u0000\u00dc\u00da\u0001\u0000\u0000\u0000\u00dc"+
		"\u00dd\u0001\u0000\u0000\u0000\u00dd\u0018\u0001\u0000\u0000\u0000\u00de"+
		"\u00dc\u0001\u0000\u0000\u0000\u00df\u00e0\u0005a\u0000\u0000\u00e0\u00e1"+
		"\u0005r\u0000\u0000\u00e1\u00e2\u0005r\u0000\u0000\u00e2\u00e3\u0005a"+
		"\u0000\u0000\u00e3\u00e4\u0005y\u0000\u0000\u00e4\u001a\u0001\u0000\u0000"+
		"\u0000\u00e5\u00e6\u0005l\u0000\u0000\u00e6\u00e7\u0005o\u0000\u0000\u00e7"+
		"\u00e8\u0005o\u0000\u0000\u00e8\u00e9\u0005p\u0000\u0000\u00e9\u001c\u0001"+
		"\u0000\u0000\u0000\u00ea\u00eb\u0005i\u0000\u0000\u00eb\u00ec\u0005f\u0000"+
		"\u0000\u00ec\u001e\u0001\u0000\u0000\u0000\u00ed\u00ee\u0005e\u0000\u0000"+
		"\u00ee\u00ef\u0005l\u0000\u0000\u00ef\u00f0\u0005s\u0000\u0000\u00f0\u00f1"+
		"\u0005e\u0000\u0000\u00f1 \u0001\u0000\u0000\u0000\u00f2\u00f3\u0005f"+
		"\u0000\u0000\u00f3\u00f4\u0005u\u0000\u0000\u00f4\u00f5\u0005n\u0000\u0000"+
		"\u00f5\"\u0001\u0000\u0000\u0000\u00f6\u00f7\u0005r\u0000\u0000\u00f7"+
		"\u00f8\u0005e\u0000\u0000\u00f8\u00f9\u0005t\u0000\u0000\u00f9\u00fa\u0005"+
		"u\u0000\u0000\u00fa\u00fb\u0005r\u0000\u0000\u00fb\u00fc\u0005n\u0000"+
		"\u0000\u00fc$\u0001\u0000\u0000\u0000\u00fd\u00fe\u0005-\u0000\u0000\u00fe"+
		"\u00ff\u0005>\u0000\u0000\u00ff&\u0001\u0000\u0000\u0000\u0100\u0101\u0005"+
		"{\u0000\u0000\u0101(\u0001\u0000\u0000\u0000\u0102\u0103\u0005}\u0000"+
		"\u0000\u0103*\u0001\u0000\u0000\u0000\u0104\u0105\u0005(\u0000\u0000\u0105"+
		",\u0001\u0000\u0000\u0000\u0106\u0107\u0005)\u0000\u0000\u0107.\u0001"+
		"\u0000\u0000\u0000\u0108\u0109\u0005[\u0000\u0000\u01090\u0001\u0000\u0000"+
		"\u0000\u010a\u010b\u0005]\u0000\u0000\u010b2\u0001\u0000\u0000\u0000\u010c"+
		"\u010d\u0005.\u0000\u0000\u010d4\u0001\u0000\u0000\u0000\u010e\u010f\u0005"+
		":\u0000\u0000\u010f6\u0001\u0000\u0000\u0000\u0110\u0111\u0005;\u0000"+
		"\u0000\u01118\u0001\u0000\u0000\u0000\u0112\u0113\u0005,\u0000\u0000\u0113"+
		":\u0001\u0000\u0000\u0000\u0114\u0115\u0005@\u0000\u0000\u0115<\u0001"+
		"\u0000\u0000\u0000\u0116\u0117\u0005+\u0000\u0000\u0117>\u0001\u0000\u0000"+
		"\u0000\u0118\u0119\u0005-\u0000\u0000\u0119@\u0001\u0000\u0000\u0000\u011a"+
		"\u011b\u0005*\u0000\u0000\u011bB\u0001\u0000\u0000\u0000\u011c\u011d\u0005"+
		"/\u0000\u0000\u011dD\u0001\u0000\u0000\u0000\u011e\u011f\u0005%\u0000"+
		"\u0000\u011fF\u0001\u0000\u0000\u0000\u0120\u0121\u0005^\u0000\u0000\u0121"+
		"H\u0001\u0000\u0000\u0000\u0122\u0123\u0005=\u0000\u0000\u0123J\u0001"+
		"\u0000\u0000\u0000\u0124\u0125\u0005!\u0000\u0000\u0125\u0126\u0005=\u0000"+
		"\u0000\u0126L\u0001\u0000\u0000\u0000\u0127\u0128\u0005<\u0000\u0000\u0128"+
		"N\u0001\u0000\u0000\u0000\u0129\u012a\u0005>\u0000\u0000\u012aP\u0001"+
		"\u0000\u0000\u0000\u012b\u012c\u0005<\u0000\u0000\u012c\u012d\u0005=\u0000"+
		"\u0000\u012dR\u0001\u0000\u0000\u0000\u012e\u012f\u0005>\u0000\u0000\u012f"+
		"\u0130\u0005=\u0000\u0000\u0130T\u0001\u0000\u0000\u0000\u0131\u0132\u0005"+
		"a\u0000\u0000\u0132\u0133\u0005n\u0000\u0000\u0133\u0134\u0005d\u0000"+
		"\u0000\u0134V\u0001\u0000\u0000\u0000\u0135\u0136\u0005o\u0000\u0000\u0136"+
		"\u0137\u0005r\u0000\u0000\u0137X\u0001\u0000\u0000\u0000\u0138\u0139\u0005"+
		"x\u0000\u0000\u0139\u013a\u0005o\u0000\u0000\u013a\u013b\u0005r\u0000"+
		"\u0000\u013bZ\u0001\u0000\u0000\u0000\u013c\u013d\u0005n\u0000\u0000\u013d"+
		"\u013e\u0005o\u0000\u0000\u013e\u013f\u0005t\u0000\u0000\u013f\\\u0001"+
		"\u0000\u0000\u0000\u0140\u0141\u0005+\u0000\u0000\u0141\u0142\u0005+\u0000"+
		"\u0000\u0142^\u0001\u0000\u0000\u0000\u0143\u0144\u0005-\u0000\u0000\u0144"+
		"\u0145\u0005-\u0000\u0000\u0145`\u0001\u0000\u0000\u0000\u0146\u0147\u0005"+
		"+\u0000\u0000\u0147\u0148\u0005=\u0000\u0000\u0148b\u0001\u0000\u0000"+
		"\u0000\u0149\u014a\u0005-\u0000\u0000\u014a\u014b\u0005=\u0000\u0000\u014b"+
		"d\u0001\u0000\u0000\u0000\u014c\u014d\u0005*\u0000\u0000\u014d\u014e\u0005"+
		"=\u0000\u0000\u014ef\u0001\u0000\u0000\u0000\u014f\u0150\u0005/\u0000"+
		"\u0000\u0150\u0151\u0005=\u0000\u0000\u0151h\u0001\u0000\u0000\u0000\u0152"+
		"\u0153\u0005%\u0000\u0000\u0153\u0154\u0005=\u0000\u0000\u0154j\u0001"+
		"\u0000\u0000\u0000\u0155\u0159\u0007\u0005\u0000\u0000\u0156\u0158\u0007"+
		"\u0006\u0000\u0000\u0157\u0156\u0001\u0000\u0000\u0000\u0158\u015b\u0001"+
		"\u0000\u0000\u0000\u0159\u0157\u0001\u0000\u0000\u0000\u0159\u015a\u0001"+
		"\u0000\u0000\u0000\u015al\u0001\u0000\u0000\u0000\u015b\u0159\u0001\u0000"+
		"\u0000\u0000\u015c\u0160\u0007\u0007\u0000\u0000\u015d\u015f\u0007\u0006"+
		"\u0000\u0000\u015e\u015d\u0001\u0000\u0000\u0000\u015f\u0162\u0001\u0000"+
		"\u0000\u0000\u0160\u015e\u0001\u0000\u0000\u0000\u0160\u0161\u0001\u0000"+
		"\u0000\u0000\u0161n\u0001\u0000\u0000\u0000\u0162\u0160\u0001\u0000\u0000"+
		"\u0000\u0163\u0165\u0005\r\u0000\u0000\u0164\u0163\u0001\u0000\u0000\u0000"+
		"\u0164\u0165\u0001\u0000\u0000\u0000\u0165\u0166\u0001\u0000\u0000\u0000"+
		"\u0166\u0167\u0005\n\u0000\u0000\u0167p\u0001\u0000\u0000\u0000\u0168"+
		"\u016a\u0007\b\u0000\u0000\u0169\u0168\u0001\u0000\u0000\u0000\u016a\u016b"+
		"\u0001\u0000\u0000\u0000\u016b\u0169\u0001\u0000\u0000\u0000\u016b\u016c"+
		"\u0001\u0000\u0000\u0000\u016c\u016d\u0001\u0000\u0000\u0000\u016d\u016e"+
		"\u00068\u0000\u0000\u016er\u0001\u0000\u0000\u0000\u0011\u0000w\u007f"+
		"\u0083\u0093\u009e\u00a3\u00a9\u00ac\u00ba\u00c7\u00da\u00dc\u0159\u0160"+
		"\u0164\u016b\u0001\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}