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
	decl ast.FunctionStatement
}

func NewGoloxFunction(decl ast.FunctionStatement) *GoloxFunction {
	return &GoloxFunction{
		decl: decl,
	}
}

// --- assumes all arity checks have already been done, but maybe worth moving this here
func (fun *GoloxFunction) call(executor *Executor, args []any) (any, error) {
	env := NewEnvironment(executor.global)

	// --- bind the args with the respective params
	assert(len(fun.decl.Parameters) == len(args), "incorrect number of arguments for function call")
	for i, arg := range args {
		env.Set(fun.decl.Parameters[i].Literal(), arg)
	}

	return executor.execBlock(fun.decl.Body, env)
}

func (fun *GoloxFunction) arity() int {
	return len(fun.decl.Parameters)
}
