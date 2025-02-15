package ast

import (
	"fmt"
)

func PrintAST(root Expr) {
	pretty := printAst(root)
	fmt.Println(pretty)
}

func printAst(node Expr) string {
	switch n := node.(type) {
	case *Binary:
		return displayBinary(n)
	default:
		return ""
	}
}

func displayBinary(bin *Binary) string {

}
