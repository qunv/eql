run:
	antlr4 -Dlanguage=Go core/antlr/EqlLexer.g4 core/antlr/EqlParser.g4 -package antlr
