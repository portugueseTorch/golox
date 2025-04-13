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

type ReturnValue struct {
	val any
}

func NewReturnValue(val any) ReturnValue {
	return ReturnValue{val: val}
}

func (err ReturnValue) Error() string {
	return fmt.Sprintf("%s", err.val)
}
