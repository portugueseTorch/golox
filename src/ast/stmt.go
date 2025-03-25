package ast

import "golox/src/lexer"

type Stmt interface {
	stmtMarker()
}

// expression statements
type ExpressionStatement struct {
	Expression Expr
}

func NewExpressionStatement(expr Expr) *ExpressionStatement {
	return &ExpressionStatement{
		Expression: expr,
	}
}

func (t *ExpressionStatement) stmtMarker() {}

// conditional statements
type ConditionalStatement struct {
	Condition  Expr
	IfBranch   Stmt
	ElseBranch Stmt
}

func NewConditionalStatement(condition Expr, ifBranch Stmt, elseBranch Stmt) *ConditionalStatement {
	return &ConditionalStatement{
		Condition:  condition,
		IfBranch:   ifBranch,
		ElseBranch: elseBranch,
	}
}

func (t *ConditionalStatement) stmtMarker() {}

// print statement
type PrintStatement struct {
	Expression Expr
}

func NewPrintStatement(expr Expr) *PrintStatement {
	return &PrintStatement{
		Expression: expr,
	}
}

func (t *PrintStatement) stmtMarker() {}

// variable statement
type VariableStatement struct {
	Name lexer.Token
	// nil if no initializer exists
	Initializer *Expr
}

func NewVariableStatement(name lexer.Token, init *Expr) *VariableStatement {
	return &VariableStatement{
		Name:        name,
		Initializer: init,
	}
}

func (t *VariableStatement) stmtMarker() {}

// block - group of statements
type BlockStatement struct {
	Statements []Stmt
}

func NewBlockStatement(statements []Stmt) *BlockStatement {
	return &BlockStatement{
		Statements: statements,
	}
}

func (t *BlockStatement) stmtMarker() {}
