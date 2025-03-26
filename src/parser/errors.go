package parser

import (
	"fmt"
	"golox/src/lexer"
)

type ParsingError struct {
	Token lexer.Token
	Msg   string
}

func NewParsingError(token lexer.Token, msg string) ParsingError {
	return ParsingError{
		Token: token,
		Msg:   msg,
	}
}

func (err ParsingError) Error() string {
	return fmt.Sprintf("[ERROR]: parsing error at line %d: %s\n", err.Token.Line(), err.Msg)
}

/*
Ideally, once a parsing error is encountered we would like to keep parsing.
This allows us to report more errors to the user. The issue is that some gramatical
invariant has already been violated, so the parser is now in a panic state.
To allow (probably) enincumbered parsing, we try to skip to the next statement and continue
parsing from there
*/
func (parser *Parser) synchronize() {
	parser.next()

	for !parser.isAtEnd() {
		// if we hit a semicolon, this is likely the end of a statement
		if parser.prev().TokenType() == lexer.SEMICOLON {
			return
		}

		switch parser.prev().TokenType() {
		case lexer.CLASS, lexer.FUN, lexer.VAR, lexer.FOR, lexer.IF, lexer.WHILE, lexer.PRINT, lexer.RETURN:
			return
		}

		parser.next()
	}
}
