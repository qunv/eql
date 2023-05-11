// Generated from /home/joseph/Documents/self/eql/core/antlr/EqlLexer.g4 by ANTLR 4.12.0
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
		ABS=16, AVG=17, ADD=18, DIVIDE=19, MULTIPLY=20, EQ=21, CONCAT=22, GT=23, 
		GTE=24, DIGIT=25, IDENTIFIER=26, WS=27, EOS=28;
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
			"AVG", "ADD", "DIVIDE", "MULTIPLY", "EQ", "CONCAT", "GT", "GTE", "DIGIT", 
			"IDENTIFIER", "WS", "EOS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'+'", "'-'", "'/'", "'*'", "'('", "')'", "'{'", "'}'", "'['", 
			"']'", "'='", "';'", "':'", "','", "'SUM'", "'ABS'", "'AVG'", "'ADD'", 
			"'DIVIDE'", "'MULTIPLY'", "'EQ'", "'CONCAT'", "'GT'", "'GTE'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY", 
			"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "SUM", "ABS", 
			"AVG", "ADD", "DIVIDE", "MULTIPLY", "EQ", "CONCAT", "GT", "GTE", "DIGIT", 
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
		"\u0004\u0000\u001c\u00ad\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002"+
		"\u0001\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002"+
		"\u0004\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002"+
		"\u0007\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002"+
		"\u000b\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e"+
		"\u0002\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011"+
		"\u0002\u0012\u0007\u0012\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014"+
		"\u0002\u0015\u0007\u0015\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017"+
		"\u0002\u0018\u0007\u0018\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a"+
		"\u0002\u001b\u0007\u001b\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001"+
		"\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007"+
		"\u0001\b\u0001\b\u0001\t\u0001\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b"+
		"\u0001\f\u0001\f\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0001\u0015\u0001\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0018\u0004\u0018\u0088"+
		"\b\u0018\u000b\u0018\f\u0018\u0089\u0001\u0019\u0004\u0019\u008d\b\u0019"+
		"\u000b\u0019\f\u0019\u008e\u0001\u001a\u0004\u001a\u0092\b\u001a\u000b"+
		"\u001a\f\u001a\u0093\u0001\u001a\u0001\u001a\u0001\u001b\u0004\u001b\u0099"+
		"\b\u001b\u000b\u001b\f\u001b\u009a\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0005\u001b\u00a2\b\u001b\n\u001b\f\u001b\u00a5"+
		"\t\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0003\u001b\u00aa\b\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u00a3\u0000\u001c\u0001\u0001\u0003\u0002"+
		"\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006\r\u0007\u000f\b\u0011\t\u0013"+
		"\n\u0015\u000b\u0017\f\u0019\r\u001b\u000e\u001d\u000f\u001f\u0010!\u0011"+
		"#\u0012%\u0013\'\u0014)\u0015+\u0016-\u0017/\u00181\u00193\u001a5\u001b"+
		"7\u001c\u0001\u0000\u0004\u0001\u000009\u0002\u0000AZaz\u0003\u0000\t"+
		"\n\r\r  \u0002\u0000\n\n\r\r\u00b4\u0000\u0001\u0001\u0000\u0000\u0000"+
		"\u0000\u0003\u0001\u0000\u0000\u0000\u0000\u0005\u0001\u0000\u0000\u0000"+
		"\u0000\u0007\u0001\u0000\u0000\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000"+
		"\u000b\u0001\u0000\u0000\u0000\u0000\r\u0001\u0000\u0000\u0000\u0000\u000f"+
		"\u0001\u0000\u0000\u0000\u0000\u0011\u0001\u0000\u0000\u0000\u0000\u0013"+
		"\u0001\u0000\u0000\u0000\u0000\u0015\u0001\u0000\u0000\u0000\u0000\u0017"+
		"\u0001\u0000\u0000\u0000\u0000\u0019\u0001\u0000\u0000\u0000\u0000\u001b"+
		"\u0001\u0000\u0000\u0000\u0000\u001d\u0001\u0000\u0000\u0000\u0000\u001f"+
		"\u0001\u0000\u0000\u0000\u0000!\u0001\u0000\u0000\u0000\u0000#\u0001\u0000"+
		"\u0000\u0000\u0000%\u0001\u0000\u0000\u0000\u0000\'\u0001\u0000\u0000"+
		"\u0000\u0000)\u0001\u0000\u0000\u0000\u0000+\u0001\u0000\u0000\u0000\u0000"+
		"-\u0001\u0000\u0000\u0000\u0000/\u0001\u0000\u0000\u0000\u00001\u0001"+
		"\u0000\u0000\u0000\u00003\u0001\u0000\u0000\u0000\u00005\u0001\u0000\u0000"+
		"\u0000\u00007\u0001\u0000\u0000\u0000\u00019\u0001\u0000\u0000\u0000\u0003"+
		";\u0001\u0000\u0000\u0000\u0005=\u0001\u0000\u0000\u0000\u0007?\u0001"+
		"\u0000\u0000\u0000\tA\u0001\u0000\u0000\u0000\u000bC\u0001\u0000\u0000"+
		"\u0000\rE\u0001\u0000\u0000\u0000\u000fG\u0001\u0000\u0000\u0000\u0011"+
		"I\u0001\u0000\u0000\u0000\u0013K\u0001\u0000\u0000\u0000\u0015M\u0001"+
		"\u0000\u0000\u0000\u0017O\u0001\u0000\u0000\u0000\u0019Q\u0001\u0000\u0000"+
		"\u0000\u001bS\u0001\u0000\u0000\u0000\u001dU\u0001\u0000\u0000\u0000\u001f"+
		"Y\u0001\u0000\u0000\u0000!]\u0001\u0000\u0000\u0000#a\u0001\u0000\u0000"+
		"\u0000%e\u0001\u0000\u0000\u0000\'l\u0001\u0000\u0000\u0000)u\u0001\u0000"+
		"\u0000\u0000+x\u0001\u0000\u0000\u0000-\u007f\u0001\u0000\u0000\u0000"+
		"/\u0082\u0001\u0000\u0000\u00001\u0087\u0001\u0000\u0000\u00003\u008c"+
		"\u0001\u0000\u0000\u00005\u0091\u0001\u0000\u0000\u00007\u00a9\u0001\u0000"+
		"\u0000\u00009:\u0005+\u0000\u0000:\u0002\u0001\u0000\u0000\u0000;<\u0005"+
		"-\u0000\u0000<\u0004\u0001\u0000\u0000\u0000=>\u0005/\u0000\u0000>\u0006"+
		"\u0001\u0000\u0000\u0000?@\u0005*\u0000\u0000@\b\u0001\u0000\u0000\u0000"+
		"AB\u0005(\u0000\u0000B\n\u0001\u0000\u0000\u0000CD\u0005)\u0000\u0000"+
		"D\f\u0001\u0000\u0000\u0000EF\u0005{\u0000\u0000F\u000e\u0001\u0000\u0000"+
		"\u0000GH\u0005}\u0000\u0000H\u0010\u0001\u0000\u0000\u0000IJ\u0005[\u0000"+
		"\u0000J\u0012\u0001\u0000\u0000\u0000KL\u0005]\u0000\u0000L\u0014\u0001"+
		"\u0000\u0000\u0000MN\u0005=\u0000\u0000N\u0016\u0001\u0000\u0000\u0000"+
		"OP\u0005;\u0000\u0000P\u0018\u0001\u0000\u0000\u0000QR\u0005:\u0000\u0000"+
		"R\u001a\u0001\u0000\u0000\u0000ST\u0005,\u0000\u0000T\u001c\u0001\u0000"+
		"\u0000\u0000UV\u0005S\u0000\u0000VW\u0005U\u0000\u0000WX\u0005M\u0000"+
		"\u0000X\u001e\u0001\u0000\u0000\u0000YZ\u0005A\u0000\u0000Z[\u0005B\u0000"+
		"\u0000[\\\u0005S\u0000\u0000\\ \u0001\u0000\u0000\u0000]^\u0005A\u0000"+
		"\u0000^_\u0005V\u0000\u0000_`\u0005G\u0000\u0000`\"\u0001\u0000\u0000"+
		"\u0000ab\u0005A\u0000\u0000bc\u0005D\u0000\u0000cd\u0005D\u0000\u0000"+
		"d$\u0001\u0000\u0000\u0000ef\u0005D\u0000\u0000fg\u0005I\u0000\u0000g"+
		"h\u0005V\u0000\u0000hi\u0005I\u0000\u0000ij\u0005D\u0000\u0000jk\u0005"+
		"E\u0000\u0000k&\u0001\u0000\u0000\u0000lm\u0005M\u0000\u0000mn\u0005U"+
		"\u0000\u0000no\u0005L\u0000\u0000op\u0005T\u0000\u0000pq\u0005I\u0000"+
		"\u0000qr\u0005P\u0000\u0000rs\u0005L\u0000\u0000st\u0005Y\u0000\u0000"+
		"t(\u0001\u0000\u0000\u0000uv\u0005E\u0000\u0000vw\u0005Q\u0000\u0000w"+
		"*\u0001\u0000\u0000\u0000xy\u0005C\u0000\u0000yz\u0005O\u0000\u0000z{"+
		"\u0005N\u0000\u0000{|\u0005C\u0000\u0000|}\u0005A\u0000\u0000}~\u0005"+
		"T\u0000\u0000~,\u0001\u0000\u0000\u0000\u007f\u0080\u0005G\u0000\u0000"+
		"\u0080\u0081\u0005T\u0000\u0000\u0081.\u0001\u0000\u0000\u0000\u0082\u0083"+
		"\u0005G\u0000\u0000\u0083\u0084\u0005T\u0000\u0000\u0084\u0085\u0005E"+
		"\u0000\u0000\u00850\u0001\u0000\u0000\u0000\u0086\u0088\u0007\u0000\u0000"+
		"\u0000\u0087\u0086\u0001\u0000\u0000\u0000\u0088\u0089\u0001\u0000\u0000"+
		"\u0000\u0089\u0087\u0001\u0000\u0000\u0000\u0089\u008a\u0001\u0000\u0000"+
		"\u0000\u008a2\u0001\u0000\u0000\u0000\u008b\u008d\u0007\u0001\u0000\u0000"+
		"\u008c\u008b\u0001\u0000\u0000\u0000\u008d\u008e\u0001\u0000\u0000\u0000"+
		"\u008e\u008c\u0001\u0000\u0000\u0000\u008e\u008f\u0001\u0000\u0000\u0000"+
		"\u008f4\u0001\u0000\u0000\u0000\u0090\u0092\u0007\u0002\u0000\u0000\u0091"+
		"\u0090\u0001\u0000\u0000\u0000\u0092\u0093\u0001\u0000\u0000\u0000\u0093"+
		"\u0091\u0001\u0000\u0000\u0000\u0093\u0094\u0001\u0000\u0000\u0000\u0094"+
		"\u0095\u0001\u0000\u0000\u0000\u0095\u0096\u0006\u001a\u0000\u0000\u0096"+
		"6\u0001\u0000\u0000\u0000\u0097\u0099\u0007\u0003\u0000\u0000\u0098\u0097"+
		"\u0001\u0000\u0000\u0000\u0099\u009a\u0001\u0000\u0000\u0000\u009a\u0098"+
		"\u0001\u0000\u0000\u0000\u009a\u009b\u0001\u0000\u0000\u0000\u009b\u00aa"+
		"\u0001\u0000\u0000\u0000\u009c\u00aa\u0005;\u0000\u0000\u009d\u009e\u0005"+
		"/\u0000\u0000\u009e\u009f\u0005*\u0000\u0000\u009f\u00a3\u0001\u0000\u0000"+
		"\u0000\u00a0\u00a2\t\u0000\u0000\u0000\u00a1\u00a0\u0001\u0000\u0000\u0000"+
		"\u00a2\u00a5\u0001\u0000\u0000\u0000\u00a3\u00a4\u0001\u0000\u0000\u0000"+
		"\u00a3\u00a1\u0001\u0000\u0000\u0000\u00a4\u00a6\u0001\u0000\u0000\u0000"+
		"\u00a5\u00a3\u0001\u0000\u0000\u0000\u00a6\u00a7\u0005*\u0000\u0000\u00a7"+
		"\u00aa\u0005/\u0000\u0000\u00a8\u00aa\u0005\u0000\u0000\u0001\u00a9\u0098"+
		"\u0001\u0000\u0000\u0000\u00a9\u009c\u0001\u0000\u0000\u0000\u00a9\u009d"+
		"\u0001\u0000\u0000\u0000\u00a9\u00a8\u0001\u0000\u0000\u0000\u00aa\u00ab"+
		"\u0001\u0000\u0000\u0000\u00ab\u00ac\u0006\u001b\u0001\u0000\u00ac8\u0001"+
		"\u0000\u0000\u0000\u0007\u0000\u0089\u008e\u0093\u009a\u00a3\u00a9\u0002"+
		"\u0006\u0000\u0000\u0002\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}