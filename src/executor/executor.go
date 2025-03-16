package executor

import (
	"fmt"
	"golox/src/ast"
	"strconv"
)

type Executor struct {
	statements []ast.Stmt
	env        Environment
}

func NewExecutor(stmt []ast.Stmt, env Environment) *Executor {
	return &Executor{
		statements: stmt,
		env:        env,
	}
}

// main executor function
func (exec *Executor) Execute() (any, error) {
	for _, s := range exec.statements {
		_, err := exec.execStatement(s)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (exec *Executor) execStatement(stmt ast.Stmt) (any, error) {
	switch s := stmt.(type) {
	case *ast.ExpressionStatement:
		return exec.execExpressionStatement(s)
	case *ast.PrintStatement:
		return exec.execPrintStatement(s)
	case *ast.VariableStatement:
		return exec.execVariableStatement(s)
	}

	return nil, nil
}

func (exec *Executor) execVariableStatement(s *ast.VariableStatement) (any, error) {
	var init any = nil
	// --- if the variable has an initializer
	if s.Initializer != nil {
		var err error = nil
		init, err = exec.execExpr(*s.Initializer)
		if err != nil {
			return nil, err
		}
	}

	exec.env.Set(s.Name.Literal(), init)
	return nil, nil
}

func (exec *Executor) execExpressionStatement(s *ast.ExpressionStatement) (any, error) {
	_, err := exec.execExpr(s.Expression)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (exec *Executor) execPrintStatement(s *ast.PrintStatement) (any, error) {
	expr, err := exec.execExpr(s.Expression)
	if err != nil {
		return nil, err
	}

	// --- print with stringify
	fmt.Println(Stringify(expr))

	return nil, nil
}

func (exec *Executor) execExpr(expr ast.Expr) (any, error) {
	switch e := expr.(type) {
	case *ast.Binary:
		return exec.execBinary(*e)
	case *ast.Unary:
		return exec.execUnary(*e)
	case *ast.Grouping:
		return exec.execGrouping(*e)
	case *ast.Literal:
		return exec.execLiteral(*e)
	case *ast.Variable:
		return exec.execVariable(*e)
	case *ast.Assignment:
		return exec.execAssignment(*e)
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
