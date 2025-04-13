package executor

import (
	"golox/src/ast"
)

func assert(condition bool, msg string) {
	if !condition {
		panic(msg)
	}

	return
}

type Callable interface {
	call(executor *Executor, args []any) (any, error)
	arity() int
}

type GoloxFunction struct {
	decl    ast.FunctionStatement
	closure *Environment
}

func NewGoloxFunction(decl ast.FunctionStatement, env *Environment) *GoloxFunction {
	return &GoloxFunction{
		decl:    decl,
		closure: env,
	}
}

// --- assumes all arity checks have already been done, but maybe worth moving this here
func (fun *GoloxFunction) call(executor *Executor, args []any) (r any, e error) {
	defer func() {
		if raw := recover(); raw != nil {
			// --- if the panic holds a ReturnValue, handle it, otherwise re-panic
			if returnValue, isReturnValue := raw.(ReturnValue); isReturnValue {
				r = returnValue.val
				e = nil
			} else {
				panic(raw)
			}
		}
	}()

	env := NewEnvironment(fun.closure)

	// --- bind the args with the respective params
	assert(len(fun.decl.Parameters) == len(args), "incorrect number of arguments for function call")
	for i, arg := range args {
		env.Set(fun.decl.Parameters[i].Literal(), arg)
	}

	_, err := executor.execBlock(fun.decl.Body, env)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (fun *GoloxFunction) arity() int {
	return len(fun.decl.Parameters)
}
