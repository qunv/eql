run:
	antlr4 -Dlanguage=Go src/core/antlr/EqlLexer.g4 src/core/antlr/EqlParser.g4 -package antlr
