package main

import (
	"bufio"
	"fmt"
	"golox/src/ast"
	"golox/src/lexer"
	"os"
)

func run(input string) {
	lexer := lexer.NewLexer(input)

	lexer.ScanTokens()
	if lexer.HasError() {
		return
	}

	for _, tok := range lexer.GetTokens() {
		fmt.Println(tok)
	}
}

func HandleFileInput(filePath string) {
	// --- load file into memory
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("[ERROR]:", err)
		return
	}

	run(string(file))
}

func HandleReplInput() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">> ")

		// --- handle empty input
		if !scanner.Scan() {
			fmt.Println("\nExiting...")
			return
		}

		input := scanner.Text()
		run(input)
	}
}

func main() {
	// args := os.Args
	//
	// if len(args) > 2 {
	// 	panic("[ERROR]: Usage: golox [file_path]")
	// } else if len(args) == 2 {
	// 	HandleFileInput(args[1])
	// } else {
	// 	HandleReplInput()
	// }

	expression := &ast.Binary{
		Left: &ast.Unary{
			Operator: lexer.NewToken(lexer.MINUS),
			Expression: &ast.Literal{
				Value: 123,
			},
		},
		Operator: lexer.NewToken(lexer.STAR),
		Right: &ast.Grouping{
			Expression: &ast.Binary{
				Left: &ast.Literal{
					Value: 45.32,
				},
				Operator: lexer.NewToken(lexer.SLASH),
				Right: &ast.Literal{
					Value: 0.86,
				},
			},
		},
	}

	ast.PrintAST(expression)
}
