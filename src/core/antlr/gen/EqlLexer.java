// Generated from /home/quannguyen20/Documents/self/eql/src/antlr/EqlLexer.g4 by ANTLR 4.12.0
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class EqlLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.12.0", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PLUS=1, MINUS=2, DIV=3, MULT=4, LPAREN=5, RPAREN=6, LCURLY=7, RCURLY=8, 
		LBRACKET=9, RBRACKET=10, EQUAL=11, SEMI=12, COLON=13, COMMA=14, ACT=15, 
		NUMBER=16, IDENTIFIER=17, WS=18, EOS=19;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY", 
			"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "ACT", "SUM", 
			"AVG", "NUMBER", "IDENTIFIER", "WS", "EOS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'+'", "'-'", "'/'", "'*'", "'('", "')'", "'{'", "'}'", "'['", 
			"']'", "'='", "';'", "':'", "','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY", 
			"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "ACT", "NUMBER", 
			"IDENTIFIER", "WS", "EOS"
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


	public EqlLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "EqlLexer.g4"; }

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
		"\u0004\u0000\u0013|\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002"+
		"\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002"+
		"\u0012\u0007\u0012\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0001"+
		"\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001"+
		"\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001"+
		"\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\t\u0001"+
		"\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\r\u0001"+
		"\r\u0001\u000e\u0001\u000e\u0003\u000eJ\b\u000e\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0011\u0004\u0011W\b\u0011\u000b\u0011"+
		"\f\u0011X\u0001\u0012\u0004\u0012\\\b\u0012\u000b\u0012\f\u0012]\u0001"+
		"\u0013\u0004\u0013a\b\u0013\u000b\u0013\f\u0013b\u0001\u0013\u0001\u0013"+
		"\u0001\u0014\u0004\u0014h\b\u0014\u000b\u0014\f\u0014i\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0005\u0014q\b\u0014\n\u0014"+
		"\f\u0014t\t\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0003\u0014y\b\u0014"+
		"\u0001\u0014\u0001\u0014\u0001r\u0000\u0015\u0001\u0001\u0003\u0002\u0005"+
		"\u0003\u0007\u0004\t\u0005\u000b\u0006\r\u0007\u000f\b\u0011\t\u0013\n"+
		"\u0015\u000b\u0017\f\u0019\r\u001b\u000e\u001d\u000f\u001f\u0000!\u0000"+
		"#\u0010%\u0011\'\u0012)\u0013\u0001\u0000\u0004\u0001\u000009\u0002\u0000"+
		"AZaz\u0003\u0000\t\n\r\r  \u0002\u0000\n\n\r\r\u0082\u0000\u0001\u0001"+
		"\u0000\u0000\u0000\u0000\u0003\u0001\u0000\u0000\u0000\u0000\u0005\u0001"+
		"\u0000\u0000\u0000\u0000\u0007\u0001\u0000\u0000\u0000\u0000\t\u0001\u0000"+
		"\u0000\u0000\u0000\u000b\u0001\u0000\u0000\u0000\u0000\r\u0001\u0000\u0000"+
		"\u0000\u0000\u000f\u0001\u0000\u0000\u0000\u0000\u0011\u0001\u0000\u0000"+
		"\u0000\u0000\u0013\u0001\u0000\u0000\u0000\u0000\u0015\u0001\u0000\u0000"+
		"\u0000\u0000\u0017\u0001\u0000\u0000\u0000\u0000\u0019\u0001\u0000\u0000"+
		"\u0000\u0000\u001b\u0001\u0000\u0000\u0000\u0000\u001d\u0001\u0000\u0000"+
		"\u0000\u0000#\u0001\u0000\u0000\u0000\u0000%\u0001\u0000\u0000\u0000\u0000"+
		"\'\u0001\u0000\u0000\u0000\u0000)\u0001\u0000\u0000\u0000\u0001+\u0001"+
		"\u0000\u0000\u0000\u0003-\u0001\u0000\u0000\u0000\u0005/\u0001\u0000\u0000"+
		"\u0000\u00071\u0001\u0000\u0000\u0000\t3\u0001\u0000\u0000\u0000\u000b"+
		"5\u0001\u0000\u0000\u0000\r7\u0001\u0000\u0000\u0000\u000f9\u0001\u0000"+
		"\u0000\u0000\u0011;\u0001\u0000\u0000\u0000\u0013=\u0001\u0000\u0000\u0000"+
		"\u0015?\u0001\u0000\u0000\u0000\u0017A\u0001\u0000\u0000\u0000\u0019C"+
		"\u0001\u0000\u0000\u0000\u001bE\u0001\u0000\u0000\u0000\u001dI\u0001\u0000"+
		"\u0000\u0000\u001fK\u0001\u0000\u0000\u0000!P\u0001\u0000\u0000\u0000"+
		"#V\u0001\u0000\u0000\u0000%[\u0001\u0000\u0000\u0000\'`\u0001\u0000\u0000"+
		"\u0000)x\u0001\u0000\u0000\u0000+,\u0005+\u0000\u0000,\u0002\u0001\u0000"+
		"\u0000\u0000-.\u0005-\u0000\u0000.\u0004\u0001\u0000\u0000\u0000/0\u0005"+
		"/\u0000\u00000\u0006\u0001\u0000\u0000\u000012\u0005*\u0000\u00002\b\u0001"+
		"\u0000\u0000\u000034\u0005(\u0000\u00004\n\u0001\u0000\u0000\u000056\u0005"+
		")\u0000\u00006\f\u0001\u0000\u0000\u000078\u0005{\u0000\u00008\u000e\u0001"+
		"\u0000\u0000\u00009:\u0005}\u0000\u0000:\u0010\u0001\u0000\u0000\u0000"+
		";<\u0005[\u0000\u0000<\u0012\u0001\u0000\u0000\u0000=>\u0005]\u0000\u0000"+
		">\u0014\u0001\u0000\u0000\u0000?@\u0005=\u0000\u0000@\u0016\u0001\u0000"+
		"\u0000\u0000AB\u0005;\u0000\u0000B\u0018\u0001\u0000\u0000\u0000CD\u0005"+
		":\u0000\u0000D\u001a\u0001\u0000\u0000\u0000EF\u0005,\u0000\u0000F\u001c"+
		"\u0001\u0000\u0000\u0000GJ\u0003\u001f\u000f\u0000HJ\u0003!\u0010\u0000"+
		"IG\u0001\u0000\u0000\u0000IH\u0001\u0000\u0000\u0000J\u001e\u0001\u0000"+
		"\u0000\u0000KL\u0005=\u0000\u0000LM\u0005S\u0000\u0000MN\u0005U\u0000"+
		"\u0000NO\u0005M\u0000\u0000O \u0001\u0000\u0000\u0000PQ\u0005=\u0000\u0000"+
		"QR\u0005A\u0000\u0000RS\u0005V\u0000\u0000ST\u0005G\u0000\u0000T\"\u0001"+
		"\u0000\u0000\u0000UW\u0007\u0000\u0000\u0000VU\u0001\u0000\u0000\u0000"+
		"WX\u0001\u0000\u0000\u0000XV\u0001\u0000\u0000\u0000XY\u0001\u0000\u0000"+
		"\u0000Y$\u0001\u0000\u0000\u0000Z\\\u0007\u0001\u0000\u0000[Z\u0001\u0000"+
		"\u0000\u0000\\]\u0001\u0000\u0000\u0000][\u0001\u0000\u0000\u0000]^\u0001"+
		"\u0000\u0000\u0000^&\u0001\u0000\u0000\u0000_a\u0007\u0002\u0000\u0000"+
		"`_\u0001\u0000\u0000\u0000ab\u0001\u0000\u0000\u0000b`\u0001\u0000\u0000"+
		"\u0000bc\u0001\u0000\u0000\u0000cd\u0001\u0000\u0000\u0000de\u0006\u0013"+
		"\u0000\u0000e(\u0001\u0000\u0000\u0000fh\u0007\u0003\u0000\u0000gf\u0001"+
		"\u0000\u0000\u0000hi\u0001\u0000\u0000\u0000ig\u0001\u0000\u0000\u0000"+
		"ij\u0001\u0000\u0000\u0000jy\u0001\u0000\u0000\u0000ky\u0005;\u0000\u0000"+
		"lm\u0005/\u0000\u0000mn\u0005*\u0000\u0000nr\u0001\u0000\u0000\u0000o"+
		"q\t\u0000\u0000\u0000po\u0001\u0000\u0000\u0000qt\u0001\u0000\u0000\u0000"+
		"rs\u0001\u0000\u0000\u0000rp\u0001\u0000\u0000\u0000su\u0001\u0000\u0000"+
		"\u0000tr\u0001\u0000\u0000\u0000uv\u0005*\u0000\u0000vy\u0005/\u0000\u0000"+
		"wy\u0005\u0000\u0000\u0001xg\u0001\u0000\u0000\u0000xk\u0001\u0000\u0000"+
		"\u0000xl\u0001\u0000\u0000\u0000xw\u0001\u0000\u0000\u0000yz\u0001\u0000"+
		"\u0000\u0000z{\u0006\u0014\u0001\u0000{*\u0001\u0000\u0000\u0000\b\u0000"+
		"IX]birx\u0002\u0006\u0000\u0000\u0002\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}