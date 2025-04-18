package resolver

import (
	"golox/src/ast"
)

func (resolver *Resolver) resolveStatements(stmts []ast.Stmt) (any, error) {
	resolver.beginScope()
	for _, stmt := range stmts {
		_, err := resolver.resolveStmt(stmt)
		if err != nil {
			return nil, err
		}
	}
	resolver.endScope()
	return nil, nil
}

func (resolver *Resolver) resolveStmt(stmt ast.Stmt) (any, error) {
	switch s := stmt.(type) {
	case *ast.BlockStatement:
		return resolver.resolveBlockStatement(s)
	case *ast.VariableStatement:
		return resolver.resolveVariableStatement(s)
	case *ast.FunctionStatement:
		return resolver.resolveFunctionStatement(s)
	case *ast.ExpressionStatement:
		return resolver.resolveExpr(s.Expression)
	}

	return nil, nil
}

func (resolver *Resolver) resolveFunctionStatement(s *ast.FunctionStatement) (any, error) {
	resolver.declare(s.Name.Literal())
	resolver.define(s.Name.Literal())

	return resolver.resolveFunction(s)
}

func (resolver *Resolver) resolveFunction(s *ast.FunctionStatement) (any, error) {
	resolver.beginScope()

	for _, tok := range s.Parameters {
		resolver.declare(tok.Literal())
		resolver.define(tok.Literal())
	}

	_, err := resolver.resolveStatements(s.Body)
	if err != nil {
		return nil, err
	}

	resolver.endScope()
	return nil, nil
}

func (resolver *Resolver) resolveVariableStatement(s *ast.VariableStatement) (any, error) {
	resolver.declare(s.Name.Literal())
	// --- if there is an initializer, resolve it
	if s.Initializer != nil {
		_, err := resolver.resolveExpr(*s.Initializer)
		if err != nil {
			return nil, err
		}
	}
	resolver.define(s.Name.Literal())
	return nil, nil
}

func (resolver *Resolver) resolveBlockStatement(s *ast.BlockStatement) (any, error) {
	resolver.beginScope()
	_, err := resolver.resolveStatements(s.Statements)
	if err != nil {
		return nil, err
	}
	resolver.endScope()
	return nil, nil
}
