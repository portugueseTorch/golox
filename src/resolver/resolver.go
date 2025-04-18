package resolver

import (
	"golox/src/ast"
	"golox/src/executor"
)

type Resolver struct {
	executor *executor.Executor
	scopes   Stack[map[string]bool]
}

func NewResolver(exec *executor.Executor) Resolver {
	return Resolver{
		executor: exec,
		scopes:   NewStack[map[string]bool](),
	}
}

func (resolver *Resolver) Resolve(stmts []ast.Stmt) (any, error) {
	return resolver.resolveStatements(stmts)
}

func (resolver *Resolver) beginScope() {
	resolver.scopes.push(make(map[string]bool))
}

func (resolver *Resolver) endScope() {
	resolver.scopes.pop()
}

// --- declares the variable in the top inner-most scope
func (resolver *Resolver) declare(name string) {
	curScope, ok := resolver.scopes.peek()
	if !ok {
		return
	}

	(*curScope)[name] = false
}

// --- defines the variable in the top inner-most scope
func (resolver *Resolver) define(name string) {
	curScope, ok := resolver.scopes.peek()
	if !ok {
		return
	}

	(*curScope)[name] = true
}
