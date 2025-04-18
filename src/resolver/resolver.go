package resolver

import (
	"go/ast"
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

func (resolver *Resolver) Resolve(stmts []ast.Stmt) {
	resolver.resolveStatements(stmts)
}

func (resolver *Resolver) resolveStatements(stmts []ast.Stmt) {
	resolver.beginScope()
	for _, stmt := range stmts {
		resolver.resolveStatement(stmt)
	}
	resolver.endScope()
}

func (resolver *Resolver) resolveStatement(stmt ast.Stmt) {
	switch s := stmt.(type) {
	case *ast.BlockStmt:
		resolver.resolveBlockStatement(s)
		break
	case *ast.VariableStatement:
		resolver.resolveVariableStatement(s)
		break
	}
}

func (resolver *Resolver) resolveBlockStatement(s *ast.BlockStmt) {
	resolver.beginScope()
	resolver.resolveStatements(s.List)
	resolver.endScope()
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
