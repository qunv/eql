// Generated from /home/joseph/Documents/self/eql/src/core/antlr/EqlLexer.g4 by ANTLR 4.12.0
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
		LBRACKET=9, RBRACKET=10, EQUAL=11, SEMI=12, COLON=13, COMMA=14, SUM=15, 
		AVG=16, ABS=17, ADD=18, DIGIT=19, IDENTIFIER=20, WS=21, EOS=22;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY", 
			"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "SUM", "AVG", 
			"ABS", "ADD", "DIGIT", "IDENTIFIER", "WS", "EOS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'+'", "'-'", "'/'", "'*'", "'('", "')'", "'{'", "'}'", "'['", 
			"']'", "'='", "';'", "':'", "','", "'SUM'", "'AVG'", "'ABS'", "'ADD'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY", 
			"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "SUM", "AVG", 
			"ABS", "ADD", "DIGIT", "IDENTIFIER", "WS", "EOS"
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
		"\u0004\u0000\u0016\u0080\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002"+
		"\u0001\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002"+
		"\u0004\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002"+
		"\u0007\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002"+
		"\u000b\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e"+
		"\u0002\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011"+
		"\u0002\u0012\u0007\u0012\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014"+
		"\u0002\u0015\u0007\u0015\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001"+
		"\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007"+
		"\u0001\b\u0001\b\u0001\t\u0001\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b"+
		"\u0001\f\u0001\f\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0012\u0004\u0012[\b\u0012\u000b\u0012\f\u0012\\\u0001\u0013"+
		"\u0004\u0013`\b\u0013\u000b\u0013\f\u0013a\u0001\u0014\u0004\u0014e\b"+
		"\u0014\u000b\u0014\f\u0014f\u0001\u0014\u0001\u0014\u0001\u0015\u0004"+
		"\u0015l\b\u0015\u000b\u0015\f\u0015m\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0005\u0015u\b\u0015\n\u0015\f\u0015x\t\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0003\u0015}\b\u0015\u0001\u0015"+
		"\u0001\u0015\u0001v\u0000\u0016\u0001\u0001\u0003\u0002\u0005\u0003\u0007"+
		"\u0004\t\u0005\u000b\u0006\r\u0007\u000f\b\u0011\t\u0013\n\u0015\u000b"+
		"\u0017\f\u0019\r\u001b\u000e\u001d\u000f\u001f\u0010!\u0011#\u0012%\u0013"+
		"\'\u0014)\u0015+\u0016\u0001\u0000\u0004\u0001\u000009\u0002\u0000AZa"+
		"z\u0003\u0000\t\n\r\r  \u0002\u0000\n\n\r\r\u0087\u0000\u0001\u0001\u0000"+
		"\u0000\u0000\u0000\u0003\u0001\u0000\u0000\u0000\u0000\u0005\u0001\u0000"+
		"\u0000\u0000\u0000\u0007\u0001\u0000\u0000\u0000\u0000\t\u0001\u0000\u0000"+
		"\u0000\u0000\u000b\u0001\u0000\u0000\u0000\u0000\r\u0001\u0000\u0000\u0000"+
		"\u0000\u000f\u0001\u0000\u0000\u0000\u0000\u0011\u0001\u0000\u0000\u0000"+
		"\u0000\u0013\u0001\u0000\u0000\u0000\u0000\u0015\u0001\u0000\u0000\u0000"+
		"\u0000\u0017\u0001\u0000\u0000\u0000\u0000\u0019\u0001\u0000\u0000\u0000"+
		"\u0000\u001b\u0001\u0000\u0000\u0000\u0000\u001d\u0001\u0000\u0000\u0000"+
		"\u0000\u001f\u0001\u0000\u0000\u0000\u0000!\u0001\u0000\u0000\u0000\u0000"+
		"#\u0001\u0000\u0000\u0000\u0000%\u0001\u0000\u0000\u0000\u0000\'\u0001"+
		"\u0000\u0000\u0000\u0000)\u0001\u0000\u0000\u0000\u0000+\u0001\u0000\u0000"+
		"\u0000\u0001-\u0001\u0000\u0000\u0000\u0003/\u0001\u0000\u0000\u0000\u0005"+
		"1\u0001\u0000\u0000\u0000\u00073\u0001\u0000\u0000\u0000\t5\u0001\u0000"+
		"\u0000\u0000\u000b7\u0001\u0000\u0000\u0000\r9\u0001\u0000\u0000\u0000"+
		"\u000f;\u0001\u0000\u0000\u0000\u0011=\u0001\u0000\u0000\u0000\u0013?"+
		"\u0001\u0000\u0000\u0000\u0015A\u0001\u0000\u0000\u0000\u0017C\u0001\u0000"+
		"\u0000\u0000\u0019E\u0001\u0000\u0000\u0000\u001bG\u0001\u0000\u0000\u0000"+
		"\u001dI\u0001\u0000\u0000\u0000\u001fM\u0001\u0000\u0000\u0000!Q\u0001"+
		"\u0000\u0000\u0000#U\u0001\u0000\u0000\u0000%Z\u0001\u0000\u0000\u0000"+
		"\'_\u0001\u0000\u0000\u0000)d\u0001\u0000\u0000\u0000+|\u0001\u0000\u0000"+
		"\u0000-.\u0005+\u0000\u0000.\u0002\u0001\u0000\u0000\u0000/0\u0005-\u0000"+
		"\u00000\u0004\u0001\u0000\u0000\u000012\u0005/\u0000\u00002\u0006\u0001"+
		"\u0000\u0000\u000034\u0005*\u0000\u00004\b\u0001\u0000\u0000\u000056\u0005"+
		"(\u0000\u00006\n\u0001\u0000\u0000\u000078\u0005)\u0000\u00008\f\u0001"+
		"\u0000\u0000\u00009:\u0005{\u0000\u0000:\u000e\u0001\u0000\u0000\u0000"+
		";<\u0005}\u0000\u0000<\u0010\u0001\u0000\u0000\u0000=>\u0005[\u0000\u0000"+
		">\u0012\u0001\u0000\u0000\u0000?@\u0005]\u0000\u0000@\u0014\u0001\u0000"+
		"\u0000\u0000AB\u0005=\u0000\u0000B\u0016\u0001\u0000\u0000\u0000CD\u0005"+
		";\u0000\u0000D\u0018\u0001\u0000\u0000\u0000EF\u0005:\u0000\u0000F\u001a"+
		"\u0001\u0000\u0000\u0000GH\u0005,\u0000\u0000H\u001c\u0001\u0000\u0000"+
		"\u0000IJ\u0005S\u0000\u0000JK\u0005U\u0000\u0000KL\u0005M\u0000\u0000"+
		"L\u001e\u0001\u0000\u0000\u0000MN\u0005A\u0000\u0000NO\u0005V\u0000\u0000"+
		"OP\u0005G\u0000\u0000P \u0001\u0000\u0000\u0000QR\u0005A\u0000\u0000R"+
		"S\u0005B\u0000\u0000ST\u0005S\u0000\u0000T\"\u0001\u0000\u0000\u0000U"+
		"V\u0005A\u0000\u0000VW\u0005D\u0000\u0000WX\u0005D\u0000\u0000X$\u0001"+
		"\u0000\u0000\u0000Y[\u0007\u0000\u0000\u0000ZY\u0001\u0000\u0000\u0000"+
		"[\\\u0001\u0000\u0000\u0000\\Z\u0001\u0000\u0000\u0000\\]\u0001\u0000"+
		"\u0000\u0000]&\u0001\u0000\u0000\u0000^`\u0007\u0001\u0000\u0000_^\u0001"+
		"\u0000\u0000\u0000`a\u0001\u0000\u0000\u0000a_\u0001\u0000\u0000\u0000"+
		"ab\u0001\u0000\u0000\u0000b(\u0001\u0000\u0000\u0000ce\u0007\u0002\u0000"+
		"\u0000dc\u0001\u0000\u0000\u0000ef\u0001\u0000\u0000\u0000fd\u0001\u0000"+
		"\u0000\u0000fg\u0001\u0000\u0000\u0000gh\u0001\u0000\u0000\u0000hi\u0006"+
		"\u0014\u0000\u0000i*\u0001\u0000\u0000\u0000jl\u0007\u0003\u0000\u0000"+
		"kj\u0001\u0000\u0000\u0000lm\u0001\u0000\u0000\u0000mk\u0001\u0000\u0000"+
		"\u0000mn\u0001\u0000\u0000\u0000n}\u0001\u0000\u0000\u0000o}\u0005;\u0000"+
		"\u0000pq\u0005/\u0000\u0000qr\u0005*\u0000\u0000rv\u0001\u0000\u0000\u0000"+
		"su\t\u0000\u0000\u0000ts\u0001\u0000\u0000\u0000ux\u0001\u0000\u0000\u0000"+
		"vw\u0001\u0000\u0000\u0000vt\u0001\u0000\u0000\u0000wy\u0001\u0000\u0000"+
		"\u0000xv\u0001\u0000\u0000\u0000yz\u0005*\u0000\u0000z}\u0005/\u0000\u0000"+
		"{}\u0005\u0000\u0000\u0001|k\u0001\u0000\u0000\u0000|o\u0001\u0000\u0000"+
		"\u0000|p\u0001\u0000\u0000\u0000|{\u0001\u0000\u0000\u0000}~\u0001\u0000"+
		"\u0000\u0000~\u007f\u0006\u0015\u0001\u0000\u007f,\u0001\u0000\u0000\u0000"+
		"\u0007\u0000\\afmv|\u0002\u0006\u0000\u0000\u0002\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}