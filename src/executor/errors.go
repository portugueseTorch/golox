package executor

import (
	"fmt"
	"golox/src/lexer"
)

type RuntimeError struct {
	Token lexer.Token
	Msg   string
}

func NewRuntimeError(token lexer.Token, msg string) RuntimeError {
	return RuntimeError{
		Token: token,
		Msg:   msg,
	}
}

func (err RuntimeError) Error() string {
	return fmt.Sprintf("[ERROR]: runtime error at line %d: %s", err.Token.Line(), err.Msg)
}
