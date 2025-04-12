package ast

import "golox/src/lexer"

type Stmt interface {
	stmtMarker()
}

// funcDecl - function declarations
type FunctionStatement struct {
	Name       lexer.Token
	Parameters []lexer.Token
	Body       []Stmt
}

func NewFunctionStatement(name lexer.Token, parms []lexer.Token, body []Stmt) *FunctionStatement {
	return &FunctionStatement{
		Name:       name,
		Parameters: parms,
		Body:       body,
	}
}

func (t *FunctionStatement) stmtMarker() {}

// for statements
type ForStatement struct {
	// runs once before execution, can be a statement or a variable declaration for convenience - optional
	Initializer Stmt
	// condition that gets checked in the begining of each iteration - optional
	Condition Expr
	// arbitrary expression that does some work at the end of each iteration (usually incrementing iterator) - optional
	Increment Expr
	Body      Stmt
}

func NewForStatement(init Stmt, cond Expr, incr Expr, body Stmt) *ForStatement {
	return &ForStatement{
		Initializer: init,
		Condition:   cond,
		Increment:   incr,
		Body:        body,
	}
}

func (t *ForStatement) stmtMarker() {}

// while statements
type WhileStatement struct {
	Condition Expr
	Body      Stmt
}

func NewWhileStatement(condition Expr, body Stmt) *WhileStatement {
	return &WhileStatement{
		Condition: condition,
		Body:      body,
	}
}

func (t *WhileStatement) stmtMarker() {}

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
