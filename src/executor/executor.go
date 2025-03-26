package executor

import (
	"fmt"
	"golox/src/ast"
	"strconv"
)

type Executor struct {
	statements []ast.Stmt
	env        *Environment
}

func NewExecutor(stmt []ast.Stmt, env *Environment) *Executor {
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

func (exec *Executor) reset(env *Environment) {
	exec.env = env
}

func (exec *Executor) execStatement(stmt ast.Stmt) (any, error) {
	switch s := stmt.(type) {
	case *ast.ExpressionStatement:
		return exec.execExpressionStatement(s)
	case *ast.ConditionalStatement:
		return exec.execConditionalStatement(s)
	case *ast.PrintStatement:
		return exec.execPrintStatement(s)
	case *ast.VariableStatement:
		return exec.execVariableStatement(s)
	case *ast.BlockStatement:
		return exec.execBlockStatement(s, NewEnvironment(exec.env))
	case *ast.WhileStatement:
		return exec.execWhileStatement(s)
	case *ast.ForStatement:
		return exec.execForStatement(s)
	}

	return nil, nil
}

func (exec *Executor) execForStatement(s *ast.ForStatement) (any, error) {
	switch init := s.Initializer.(type) {
	case *ast.VariableStatement:
		value, err := exec.execExpr(*init.Initializer)
		if err != nil {
			return nil, err
		}
		exec.env.Set(init.Name.Literal(), value)
	case *ast.ExpressionStatement:
		_, err := exec.execExpr(*&init.Expression)
		if err != nil {
			return nil, err
		}
	}

	cond, err := exec.execExpr(s.Condition)
	if err != nil {
		return nil, err
	}

	for isTruthy(cond) {
		_, err := exec.execStatement(s.Body)
		if err != nil {
			return nil, err
		}

		_, err = exec.execExpr(s.Increment)
		if err != nil {
			return nil, err
		}

		cond, err = exec.execExpr(s.Condition)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (exec *Executor) execWhileStatement(s *ast.WhileStatement) (any, error) {
	cond, err := exec.execExpr(s.Condition)
	if err != nil {
		return nil, err
	}

	for isTruthy(cond) {
		_, err := exec.execStatement(s.Body)
		if err != nil {
			return nil, err
		}

		cond, err = exec.execExpr(s.Condition)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (exec *Executor) execConditionalStatement(s *ast.ConditionalStatement) (any, error) {
	condition, err := exec.execExpr(s.Condition)
	if err != nil {
		return nil, err
	}

	if isTruthy(condition) {
		return exec.execStatement(s.IfBranch)
	} else {
		return exec.execStatement(s.ElseBranch)
	}
}

func (exec *Executor) execBlockStatement(s *ast.BlockStatement, env *Environment) (any, error) {
	previous := exec.env
	defer exec.reset(previous)

	// --- execute each statement individually. On error reset the state and return
	for _, statement := range s.Statements {
		exec.env = env

		_, err := exec.execStatement(statement)
		if err != nil {
			return nil, err
		}
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
	case *ast.Logical:
		return exec.execLogical(*e)
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

	return nil, nil
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
