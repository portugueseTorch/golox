package ast

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
