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
		return resolver.resolveExpressionExpression(s)
	}

	return nil, nil
}

func (resolver *Resolver) resolveExpressionExpression(s *ast.Assignment) (any, error) {
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
	for i := len(resolver.scopes.items) - 1; i >= 0; i++ {
		_, exists := resolver.scopes.items[i][name.Literal()]
		if exists {
			panic("Unimplemented")
		}
	}
	return nil, nil
}
