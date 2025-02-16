package ast

import (
	"fmt"
	"strings"
)

func PrintAST(root Expr) {
	pretty := printAst(root)
	fmt.Println(pretty)
}

func printAst(node Expr) string {
	switch n := node.(type) {
	case *Binary:
		return displayBinary(n)
	case *Grouping:
		return displayGrouping(n)
	case *Literal:
		return displayLiteral(n)
	case *Unary:
		return displayUnary(n)
	default:
		return ""
	}
}

func displayUnary(n *Unary) string {
	operator := n.Operator.Type()
	expression := printAst(n.Expression)

	return parenthesize(operator, expression)
}

func displayLiteral(n *Literal) string {
	if n.Value == nil {
		return "<nil>"
	}

	return fmt.Sprintf("%v", n.Value)
}

func displayGrouping(group *Grouping) string {
	expression := printAst(group.Expression)

	return parenthesize("GROUPING", expression)
}

func displayBinary(bin *Binary) string {
	operator := bin.Operator.Type()
	left := printAst(bin.Left)
	right := printAst(bin.Right)

	return parenthesize(operator, left, right)
}

func parenthesize(args ...string) string {
	return "(" + strings.Join(args, " ") + ")"
}
