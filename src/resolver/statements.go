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
	case *ast.ConditionalStatement:
		return resolver.resolveConditionalExpression(s)
	case *ast.PrintStatement:
		return resolver.resolvePrintExpression(s)
	case *ast.ReturnStatement:
		return resolver.resolveReturnExpression(s)
	case *ast.WhileStatement:
		return resolver.resolveWhileStatement(s)
	}

	return nil, nil
}

func (resolver *Resolver) resolveWhileStatement(s *ast.WhileStatement) (any, error) {
	resolver.resolveExpr(s.Condition)
	return resolver.resolveStmt(s.Body)
}

func (resolver *Resolver) resolveReturnExpression(s *ast.ReturnStatement) (any, error) {
	if s.Expression != nil {
		return resolver.resolveExpr(s.Expression)
	}
	return nil, nil
}

func (resolver *Resolver) resolvePrintExpression(s *ast.PrintStatement) (any, error) {
	return resolver.resolveExpr(s.Expression)
}

func (resolver *Resolver) resolveConditionalExpression(s *ast.ConditionalStatement) (any, error) {
	resolver.resolveStmt(s.IfBranch)
	resolver.resolveExpr(s.Condition)
	if s.ElseBranch != nil {
		resolver.resolveStmt(s.ElseBranch)
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

	for _, stmt := range s.Body {
		_, err := resolver.resolveStmt(stmt)
		if err != nil {
			return nil, err
		}
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
	for _, stmt := range s.Statements {
		_, err := resolver.resolveStmt(stmt)
		if err != nil {
			return nil, err
		}
	}
	resolver.endScope()
	return nil, nil
}
