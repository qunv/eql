run:
	antlr4 -Dlanguage=Go src/antlr/EqlLexer.g4 src/antlr/EqlParser.g4 -package antlr
