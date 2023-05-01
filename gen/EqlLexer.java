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
		ABS=16, AVG=17, ADD=18, DIVIDE=19, EQ=20, CONCAT=21, DIGIT=22, IDENTIFIER=23, 
		WS=24, EOS=25;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY", 
			"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "SUM", "ABS", 
			"AVG", "ADD", "DIVIDE", "EQ", "CONCAT", "DIGIT", "IDENTIFIER", "WS", 
			"EOS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'+'", "'-'", "'/'", "'*'", "'('", "')'", "'{'", "'}'", "'['", 
			"']'", "'='", "';'", "':'", "','", "'SUM'", "'ABS'", "'AVG'", "'ADD'", 
			"'DIVIDE'", "'EQ'", "'CONCAT'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY", 
			"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "SUM", "ABS", 
			"AVG", "ADD", "DIVIDE", "EQ", "CONCAT", "DIGIT", "IDENTIFIER", "WS", 
			"EOS"
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
		"\u0004\u0000\u0019\u0097\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002"+
		"\u0001\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002"+
		"\u0004\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002"+
		"\u0007\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002"+
		"\u000b\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e"+
		"\u0002\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011"+
		"\u0002\u0012\u0007\u0012\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014"+
		"\u0002\u0015\u0007\u0015\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017"+
		"\u0002\u0018\u0007\u0018\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001"+
		"\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007"+
		"\u0001\b\u0001\b\u0001\t\u0001\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b"+
		"\u0001\f\u0001\f\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0015\u0004\u0015r\b\u0015\u000b\u0015\f\u0015s\u0001\u0016\u0004\u0016"+
		"w\b\u0016\u000b\u0016\f\u0016x\u0001\u0017\u0004\u0017|\b\u0017\u000b"+
		"\u0017\f\u0017}\u0001\u0017\u0001\u0017\u0001\u0018\u0004\u0018\u0083"+
		"\b\u0018\u000b\u0018\f\u0018\u0084\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0005\u0018\u008c\b\u0018\n\u0018\f\u0018\u008f"+
		"\t\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0003\u0018\u0094\b\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u008d\u0000\u0019\u0001\u0001\u0003\u0002"+
		"\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006\r\u0007\u000f\b\u0011\t\u0013"+
		"\n\u0015\u000b\u0017\f\u0019\r\u001b\u000e\u001d\u000f\u001f\u0010!\u0011"+
		"#\u0012%\u0013\'\u0014)\u0015+\u0016-\u0017/\u00181\u0019\u0001\u0000"+
		"\u0004\u0001\u000009\u0002\u0000AZaz\u0003\u0000\t\n\r\r  \u0002\u0000"+
		"\n\n\r\r\u009e\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000"+
		"\u0000\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001\u0000"+
		"\u0000\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000\u0000"+
		"\u0000\u0000\r\u0001\u0000\u0000\u0000\u0000\u000f\u0001\u0000\u0000\u0000"+
		"\u0000\u0011\u0001\u0000\u0000\u0000\u0000\u0013\u0001\u0000\u0000\u0000"+
		"\u0000\u0015\u0001\u0000\u0000\u0000\u0000\u0017\u0001\u0000\u0000\u0000"+
		"\u0000\u0019\u0001\u0000\u0000\u0000\u0000\u001b\u0001\u0000\u0000\u0000"+
		"\u0000\u001d\u0001\u0000\u0000\u0000\u0000\u001f\u0001\u0000\u0000\u0000"+
		"\u0000!\u0001\u0000\u0000\u0000\u0000#\u0001\u0000\u0000\u0000\u0000%"+
		"\u0001\u0000\u0000\u0000\u0000\'\u0001\u0000\u0000\u0000\u0000)\u0001"+
		"\u0000\u0000\u0000\u0000+\u0001\u0000\u0000\u0000\u0000-\u0001\u0000\u0000"+
		"\u0000\u0000/\u0001\u0000\u0000\u0000\u00001\u0001\u0000\u0000\u0000\u0001"+
		"3\u0001\u0000\u0000\u0000\u00035\u0001\u0000\u0000\u0000\u00057\u0001"+
		"\u0000\u0000\u0000\u00079\u0001\u0000\u0000\u0000\t;\u0001\u0000\u0000"+
		"\u0000\u000b=\u0001\u0000\u0000\u0000\r?\u0001\u0000\u0000\u0000\u000f"+
		"A\u0001\u0000\u0000\u0000\u0011C\u0001\u0000\u0000\u0000\u0013E\u0001"+
		"\u0000\u0000\u0000\u0015G\u0001\u0000\u0000\u0000\u0017I\u0001\u0000\u0000"+
		"\u0000\u0019K\u0001\u0000\u0000\u0000\u001bM\u0001\u0000\u0000\u0000\u001d"+
		"O\u0001\u0000\u0000\u0000\u001fS\u0001\u0000\u0000\u0000!W\u0001\u0000"+
		"\u0000\u0000#[\u0001\u0000\u0000\u0000%_\u0001\u0000\u0000\u0000\'f\u0001"+
		"\u0000\u0000\u0000)i\u0001\u0000\u0000\u0000+q\u0001\u0000\u0000\u0000"+
		"-v\u0001\u0000\u0000\u0000/{\u0001\u0000\u0000\u00001\u0093\u0001\u0000"+
		"\u0000\u000034\u0005+\u0000\u00004\u0002\u0001\u0000\u0000\u000056\u0005"+
		"-\u0000\u00006\u0004\u0001\u0000\u0000\u000078\u0005/\u0000\u00008\u0006"+
		"\u0001\u0000\u0000\u00009:\u0005*\u0000\u0000:\b\u0001\u0000\u0000\u0000"+
		";<\u0005(\u0000\u0000<\n\u0001\u0000\u0000\u0000=>\u0005)\u0000\u0000"+
		">\f\u0001\u0000\u0000\u0000?@\u0005{\u0000\u0000@\u000e\u0001\u0000\u0000"+
		"\u0000AB\u0005}\u0000\u0000B\u0010\u0001\u0000\u0000\u0000CD\u0005[\u0000"+
		"\u0000D\u0012\u0001\u0000\u0000\u0000EF\u0005]\u0000\u0000F\u0014\u0001"+
		"\u0000\u0000\u0000GH\u0005=\u0000\u0000H\u0016\u0001\u0000\u0000\u0000"+
		"IJ\u0005;\u0000\u0000J\u0018\u0001\u0000\u0000\u0000KL\u0005:\u0000\u0000"+
		"L\u001a\u0001\u0000\u0000\u0000MN\u0005,\u0000\u0000N\u001c\u0001\u0000"+
		"\u0000\u0000OP\u0005S\u0000\u0000PQ\u0005U\u0000\u0000QR\u0005M\u0000"+
		"\u0000R\u001e\u0001\u0000\u0000\u0000ST\u0005A\u0000\u0000TU\u0005B\u0000"+
		"\u0000UV\u0005S\u0000\u0000V \u0001\u0000\u0000\u0000WX\u0005A\u0000\u0000"+
		"XY\u0005V\u0000\u0000YZ\u0005G\u0000\u0000Z\"\u0001\u0000\u0000\u0000"+
		"[\\\u0005A\u0000\u0000\\]\u0005D\u0000\u0000]^\u0005D\u0000\u0000^$\u0001"+
		"\u0000\u0000\u0000_`\u0005D\u0000\u0000`a\u0005I\u0000\u0000ab\u0005V"+
		"\u0000\u0000bc\u0005I\u0000\u0000cd\u0005D\u0000\u0000de\u0005E\u0000"+
		"\u0000e&\u0001\u0000\u0000\u0000fg\u0005E\u0000\u0000gh\u0005Q\u0000\u0000"+
		"h(\u0001\u0000\u0000\u0000ij\u0005C\u0000\u0000jk\u0005O\u0000\u0000k"+
		"l\u0005N\u0000\u0000lm\u0005C\u0000\u0000mn\u0005A\u0000\u0000no\u0005"+
		"T\u0000\u0000o*\u0001\u0000\u0000\u0000pr\u0007\u0000\u0000\u0000qp\u0001"+
		"\u0000\u0000\u0000rs\u0001\u0000\u0000\u0000sq\u0001\u0000\u0000\u0000"+
		"st\u0001\u0000\u0000\u0000t,\u0001\u0000\u0000\u0000uw\u0007\u0001\u0000"+
		"\u0000vu\u0001\u0000\u0000\u0000wx\u0001\u0000\u0000\u0000xv\u0001\u0000"+
		"\u0000\u0000xy\u0001\u0000\u0000\u0000y.\u0001\u0000\u0000\u0000z|\u0007"+
		"\u0002\u0000\u0000{z\u0001\u0000\u0000\u0000|}\u0001\u0000\u0000\u0000"+
		"}{\u0001\u0000\u0000\u0000}~\u0001\u0000\u0000\u0000~\u007f\u0001\u0000"+
		"\u0000\u0000\u007f\u0080\u0006\u0017\u0000\u0000\u00800\u0001\u0000\u0000"+
		"\u0000\u0081\u0083\u0007\u0003\u0000\u0000\u0082\u0081\u0001\u0000\u0000"+
		"\u0000\u0083\u0084\u0001\u0000\u0000\u0000\u0084\u0082\u0001\u0000\u0000"+
		"\u0000\u0084\u0085\u0001\u0000\u0000\u0000\u0085\u0094\u0001\u0000\u0000"+
		"\u0000\u0086\u0094\u0005;\u0000\u0000\u0087\u0088\u0005/\u0000\u0000\u0088"+
		"\u0089\u0005*\u0000\u0000\u0089\u008d\u0001\u0000\u0000\u0000\u008a\u008c"+
		"\t\u0000\u0000\u0000\u008b\u008a\u0001\u0000\u0000\u0000\u008c\u008f\u0001"+
		"\u0000\u0000\u0000\u008d\u008e\u0001\u0000\u0000\u0000\u008d\u008b\u0001"+
		"\u0000\u0000\u0000\u008e\u0090\u0001\u0000\u0000\u0000\u008f\u008d\u0001"+
		"\u0000\u0000\u0000\u0090\u0091\u0005*\u0000\u0000\u0091\u0094\u0005/\u0000"+
		"\u0000\u0092\u0094\u0005\u0000\u0000\u0001\u0093\u0082\u0001\u0000\u0000"+
		"\u0000\u0093\u0086\u0001\u0000\u0000\u0000\u0093\u0087\u0001\u0000\u0000"+
		"\u0000\u0093\u0092\u0001\u0000\u0000\u0000\u0094\u0095\u0001\u0000\u0000"+
		"\u0000\u0095\u0096\u0006\u0018\u0001\u0000\u00962\u0001\u0000\u0000\u0000"+
		"\u0007\u0000sx}\u0084\u008d\u0093\u0002\u0006\u0000\u0000\u0002\u0000"+
		"\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}