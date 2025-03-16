package executor

import (
	"fmt"
	"golox/src/lexer"
)

type Environment struct {
	store map[string]any
	// --- reference to enclosing environment
	enclosing *Environment
}

func NewEnvironment(enclosing *Environment) *Environment {
	return &Environment{
		store:     make(map[string]any),
		enclosing: enclosing,
	}
}

func (env *Environment) Set(key string, value any) {
	env.store[key] = value
}

// --- for better error logging, do so in the caller to get
func (env *Environment) Get(key lexer.Token) (any, error) {
	val, ok := env.store[key.Literal()]
	if ok {
		return val, nil
	}

	// --- if we have enclosing environments, check there
	if env.enclosing != nil {
		return env.enclosing.Get(key)
	}

	return nil, NewRuntimeError(key, fmt.Sprintf("undefined variable name '%s'", key.Literal()))
}

// --- tries to reasign key to value, returning value if successfull and an error if failed
func (env *Environment) Assign(key lexer.Token, value any) (any, error) {
	_, ok := env.store[key.Literal()]
	if ok {
		env.Set(key.Literal(), value)
		return value, nil
	}

	// --- if we have enclosing environments, check there
	if env.enclosing != nil {
		return env.enclosing.Assign(key, value)
	}

	return nil, NewRuntimeError(key, fmt.Sprintf("invalid assignment: variable '%s' does not exist", key.Literal()))
}
