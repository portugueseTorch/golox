package executor

import (
	"fmt"
	"golox/src/ast"
	"strconv"
)

// main executor function
func Execute(stmt []ast.Stmt) (any, error) {
	for _, s := range stmt {
		_, err := execStatement(s)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func execStatement(stmt ast.Stmt) (any, error) {
	switch s := stmt.(type) {
	case *ast.ExpressionStatement:
		return execExpressionStatement(s)
	case *ast.PrintStatement:
		return execPrintStatement(s)
	}

	panic("unimplemented statement type")
}

func execExpressionStatement(s *ast.ExpressionStatement) (any, error) {
	_, err := execExpr(s.Expression)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func execPrintStatement(s *ast.PrintStatement) (any, error) {
	expr, err := execExpr(s.Expression)
	if err != nil {
		return nil, err
	}

	// --- print with stringify
	fmt.Println(Stringify(expr))

	return nil, nil
}

func execExpr(expr ast.Expr) (any, error) {
	switch e := expr.(type) {
	case *ast.Binary:
		return execBinary(*e)
	case *ast.Unary:
		return execUnary(*e)
	case *ast.Grouping:
		return execGrouping(*e)
	case *ast.Literal:
		return execLiteral(*e)
	}

	panic("unreachable")
}

func Stringify(result any) string {
	switch t := result.(type) {
	case *string:
		return *t
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case string:
		return t
	}

	return fmt.Sprint(result)
}
