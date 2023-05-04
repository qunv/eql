// Generated from /home/quannguyen20/Documents/self/eql/src/core/antlr/EqlLexer.g4 by ANTLR 4.12.0
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
		ABS=16, AVG=17, ADD=18, DIVIDE=19, MULTIPLY=20, EQ=21, CONCAT=22, DIGIT=23, 
		IDENTIFIER=24, WS=25, EOS=26;
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
			"AVG", "ADD", "DIVIDE", "MULTIPLY", "EQ", "CONCAT", "DIGIT", "IDENTIFIER", 
			"WS", "EOS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'+'", "'-'", "'/'", "'*'", "'('", "')'", "'{'", "'}'", "'['", 
			"']'", "'='", "';'", "':'", "','", "'SUM'", "'ABS'", "'AVG'", "'ADD'", 
			"'DIVIDE'", "'MULTIPLY'", "'EQ'", "'CONCAT'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY", 
			"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "SUM", "ABS", 
			"AVG", "ADD", "DIVIDE", "MULTIPLY", "EQ", "CONCAT", "DIGIT", "IDENTIFIER", 
			"WS", "EOS"
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
		"\u0004\u0000\u001a\u00a2\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002"+
		"\u0001\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002"+
		"\u0004\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002"+
		"\u0007\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002"+
		"\u000b\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e"+
		"\u0002\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011"+
		"\u0002\u0012\u0007\u0012\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014"+
		"\u0002\u0015\u0007\u0015\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017"+
		"\u0002\u0018\u0007\u0018\u0002\u0019\u0007\u0019\u0001\u0000\u0001\u0000"+
		"\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003"+
		"\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006"+
		"\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\t\u0001\t\u0001\n\u0001"+
		"\n\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\r\u0001\r\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0016"+
		"\u0004\u0016}\b\u0016\u000b\u0016\f\u0016~\u0001\u0017\u0004\u0017\u0082"+
		"\b\u0017\u000b\u0017\f\u0017\u0083\u0001\u0018\u0004\u0018\u0087\b\u0018"+
		"\u000b\u0018\f\u0018\u0088\u0001\u0018\u0001\u0018\u0001\u0019\u0004\u0019"+
		"\u008e\b\u0019\u000b\u0019\f\u0019\u008f\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0005\u0019\u0097\b\u0019\n\u0019\f\u0019"+
		"\u009a\t\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0003\u0019\u009f\b"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0098\u0000\u001a\u0001\u0001\u0003"+
		"\u0002\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006\r\u0007\u000f\b\u0011"+
		"\t\u0013\n\u0015\u000b\u0017\f\u0019\r\u001b\u000e\u001d\u000f\u001f\u0010"+
		"!\u0011#\u0012%\u0013\'\u0014)\u0015+\u0016-\u0017/\u00181\u00193\u001a"+
		"\u0001\u0000\u0004\u0001\u000009\u0002\u0000AZaz\u0003\u0000\t\n\r\r "+
		" \u0002\u0000\n\n\r\r\u00a9\u0000\u0001\u0001\u0000\u0000\u0000\u0000"+
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
		"\u0000\u0000\u00003\u0001\u0000\u0000\u0000\u00015\u0001\u0000\u0000\u0000"+
		"\u00037\u0001\u0000\u0000\u0000\u00059\u0001\u0000\u0000\u0000\u0007;"+
		"\u0001\u0000\u0000\u0000\t=\u0001\u0000\u0000\u0000\u000b?\u0001\u0000"+
		"\u0000\u0000\rA\u0001\u0000\u0000\u0000\u000fC\u0001\u0000\u0000\u0000"+
		"\u0011E\u0001\u0000\u0000\u0000\u0013G\u0001\u0000\u0000\u0000\u0015I"+
		"\u0001\u0000\u0000\u0000\u0017K\u0001\u0000\u0000\u0000\u0019M\u0001\u0000"+
		"\u0000\u0000\u001bO\u0001\u0000\u0000\u0000\u001dQ\u0001\u0000\u0000\u0000"+
		"\u001fU\u0001\u0000\u0000\u0000!Y\u0001\u0000\u0000\u0000#]\u0001\u0000"+
		"\u0000\u0000%a\u0001\u0000\u0000\u0000\'h\u0001\u0000\u0000\u0000)q\u0001"+
		"\u0000\u0000\u0000+t\u0001\u0000\u0000\u0000-|\u0001\u0000\u0000\u0000"+
		"/\u0081\u0001\u0000\u0000\u00001\u0086\u0001\u0000\u0000\u00003\u009e"+
		"\u0001\u0000\u0000\u000056\u0005+\u0000\u00006\u0002\u0001\u0000\u0000"+
		"\u000078\u0005-\u0000\u00008\u0004\u0001\u0000\u0000\u00009:\u0005/\u0000"+
		"\u0000:\u0006\u0001\u0000\u0000\u0000;<\u0005*\u0000\u0000<\b\u0001\u0000"+
		"\u0000\u0000=>\u0005(\u0000\u0000>\n\u0001\u0000\u0000\u0000?@\u0005)"+
		"\u0000\u0000@\f\u0001\u0000\u0000\u0000AB\u0005{\u0000\u0000B\u000e\u0001"+
		"\u0000\u0000\u0000CD\u0005}\u0000\u0000D\u0010\u0001\u0000\u0000\u0000"+
		"EF\u0005[\u0000\u0000F\u0012\u0001\u0000\u0000\u0000GH\u0005]\u0000\u0000"+
		"H\u0014\u0001\u0000\u0000\u0000IJ\u0005=\u0000\u0000J\u0016\u0001\u0000"+
		"\u0000\u0000KL\u0005;\u0000\u0000L\u0018\u0001\u0000\u0000\u0000MN\u0005"+
		":\u0000\u0000N\u001a\u0001\u0000\u0000\u0000OP\u0005,\u0000\u0000P\u001c"+
		"\u0001\u0000\u0000\u0000QR\u0005S\u0000\u0000RS\u0005U\u0000\u0000ST\u0005"+
		"M\u0000\u0000T\u001e\u0001\u0000\u0000\u0000UV\u0005A\u0000\u0000VW\u0005"+
		"B\u0000\u0000WX\u0005S\u0000\u0000X \u0001\u0000\u0000\u0000YZ\u0005A"+
		"\u0000\u0000Z[\u0005V\u0000\u0000[\\\u0005G\u0000\u0000\\\"\u0001\u0000"+
		"\u0000\u0000]^\u0005A\u0000\u0000^_\u0005D\u0000\u0000_`\u0005D\u0000"+
		"\u0000`$\u0001\u0000\u0000\u0000ab\u0005D\u0000\u0000bc\u0005I\u0000\u0000"+
		"cd\u0005V\u0000\u0000de\u0005I\u0000\u0000ef\u0005D\u0000\u0000fg\u0005"+
		"E\u0000\u0000g&\u0001\u0000\u0000\u0000hi\u0005M\u0000\u0000ij\u0005U"+
		"\u0000\u0000jk\u0005L\u0000\u0000kl\u0005T\u0000\u0000lm\u0005I\u0000"+
		"\u0000mn\u0005P\u0000\u0000no\u0005L\u0000\u0000op\u0005Y\u0000\u0000"+
		"p(\u0001\u0000\u0000\u0000qr\u0005E\u0000\u0000rs\u0005Q\u0000\u0000s"+
		"*\u0001\u0000\u0000\u0000tu\u0005C\u0000\u0000uv\u0005O\u0000\u0000vw"+
		"\u0005N\u0000\u0000wx\u0005C\u0000\u0000xy\u0005A\u0000\u0000yz\u0005"+
		"T\u0000\u0000z,\u0001\u0000\u0000\u0000{}\u0007\u0000\u0000\u0000|{\u0001"+
		"\u0000\u0000\u0000}~\u0001\u0000\u0000\u0000~|\u0001\u0000\u0000\u0000"+
		"~\u007f\u0001\u0000\u0000\u0000\u007f.\u0001\u0000\u0000\u0000\u0080\u0082"+
		"\u0007\u0001\u0000\u0000\u0081\u0080\u0001\u0000\u0000\u0000\u0082\u0083"+
		"\u0001\u0000\u0000\u0000\u0083\u0081\u0001\u0000\u0000\u0000\u0083\u0084"+
		"\u0001\u0000\u0000\u0000\u00840\u0001\u0000\u0000\u0000\u0085\u0087\u0007"+
		"\u0002\u0000\u0000\u0086\u0085\u0001\u0000\u0000\u0000\u0087\u0088\u0001"+
		"\u0000\u0000\u0000\u0088\u0086\u0001\u0000\u0000\u0000\u0088\u0089\u0001"+
		"\u0000\u0000\u0000\u0089\u008a\u0001\u0000\u0000\u0000\u008a\u008b\u0006"+
		"\u0018\u0000\u0000\u008b2\u0001\u0000\u0000\u0000\u008c\u008e\u0007\u0003"+
		"\u0000\u0000\u008d\u008c\u0001\u0000\u0000\u0000\u008e\u008f\u0001\u0000"+
		"\u0000\u0000\u008f\u008d\u0001\u0000\u0000\u0000\u008f\u0090\u0001\u0000"+
		"\u0000\u0000\u0090\u009f\u0001\u0000\u0000\u0000\u0091\u009f\u0005;\u0000"+
		"\u0000\u0092\u0093\u0005/\u0000\u0000\u0093\u0094\u0005*\u0000\u0000\u0094"+
		"\u0098\u0001\u0000\u0000\u0000\u0095\u0097\t\u0000\u0000\u0000\u0096\u0095"+
		"\u0001\u0000\u0000\u0000\u0097\u009a\u0001\u0000\u0000\u0000\u0098\u0099"+
		"\u0001\u0000\u0000\u0000\u0098\u0096\u0001\u0000\u0000\u0000\u0099\u009b"+
		"\u0001\u0000\u0000\u0000\u009a\u0098\u0001\u0000\u0000\u0000\u009b\u009c"+
		"\u0005*\u0000\u0000\u009c\u009f\u0005/\u0000\u0000\u009d\u009f\u0005\u0000"+
		"\u0000\u0001\u009e\u008d\u0001\u0000\u0000\u0000\u009e\u0091\u0001\u0000"+
		"\u0000\u0000\u009e\u0092\u0001\u0000\u0000\u0000\u009e\u009d\u0001\u0000"+
		"\u0000\u0000\u009f\u00a0\u0001\u0000\u0000\u0000\u00a0\u00a1\u0006\u0019"+
		"\u0001\u0000\u00a14\u0001\u0000\u0000\u0000\u0007\u0000~\u0083\u0088\u008f"+
		"\u0098\u009e\u0002\u0006\u0000\u0000\u0002\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}