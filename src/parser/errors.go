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
