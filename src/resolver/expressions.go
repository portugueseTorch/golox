package resolver

import (
	"golox/src/ast"
	"golox/src/lexer"
	"golox/src/parser"
)

func (resolver *Resolver) resolveExpr(expr ast.Expr) (any, error) {
	switch s := expr.(type) {
	case *ast.Variable:
		return resolver.resolveVariableExpression(s)
	case *ast.Assignment:
		return resolver.resolveAssignmentExpression(s)
	case *ast.Binary:
		return resolver.resolveBinaryExpression(s)
	case *ast.Call:
		return resolver.resolveCallExpression(s)
	case *ast.Grouping:
		return resolver.resolveGroupingExpression(s)
	case *ast.Logical:
		return resolver.resolveLogicalExpression(s)
	case *ast.Unary:
		return resolver.resolveExpr(s.Expression)
	case *ast.Literal:
		return nil, nil
	}

	return nil, nil
}

func (resolver *Resolver) resolveLogicalExpression(s *ast.Logical) (any, error) {
	resolver.resolveExpr(s.Left)
	return resolver.resolveExpr(s.Right)
}

func (resolver *Resolver) resolveGroupingExpression(s *ast.Grouping) (any, error) {
	return resolver.resolveExpr(s.Expression)
}

func (resolver *Resolver) resolveCallExpression(s *ast.Call) (any, error) {
	resolver.resolveExpr(s.Callee)

	for _, arg := range s.Args {
		resolver.resolveExpr(arg)
	}
	return nil, nil
}

func (resolver *Resolver) resolveBinaryExpression(s *ast.Binary) (any, error) {
	resolver.resolveExpr(s.Left)
	return resolver.resolveExpr(s.Right)
}

func (resolver *Resolver) resolveAssignmentExpression(s *ast.Assignment) (any, error) {
	_, err := resolver.resolveExpr(s.Value)
	if err != nil {
		return nil, err
	}

	return resolver.resolveLocal(s.Name, s.Value)
}

func (resolver *Resolver) resolveVariableExpression(s *ast.Variable) (any, error) {
	curScope, scopeOk := resolver.scopes.peek()
	val, valExistsInScope := (*curScope)[s.Name.Literal()]

	// --- if the variable in question is not defined in the current scope, error out
	if scopeOk && valExistsInScope && !val {
		return nil, parser.NewParsingError(s.Name, "invalid variable expression: variable is not defined in the current scope")
	}

	return resolver.resolveLocal(s.Name, s)
}

func (resolver *Resolver) resolveLocal(name lexer.Token, expr ast.Expr) (any, error) {
	// --- walk up the scopes until variable is found, and execute
	for i := len(resolver.scopes.items) - 1; i >= 0; i-- {
		_, exists := resolver.scopes.items[i][name.Literal()]
		if exists {
			resolver.executor.Set(expr, len(resolver.scopes.items)-1-i)
		}
	}
	return nil, nil
}
