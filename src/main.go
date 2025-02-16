package main

import (
	"bufio"
	"fmt"
	"golox/src/ast"
	"golox/src/lexer"
	"golox/src/parser"
	"os"
)

func run(input string) {
	lexer := lexer.NewLexer(input)
	lexer.ScanTokens()
	if lexer.HasError() {
		return
	}

	parser := parser.NewParser(lexer.GetTokens())
	parsed, err := parser.Parse()
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	ast.PrintAST(parsed)
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
	args := os.Args

	if len(args) > 2 {
		panic("[ERROR]: Usage: golox [file_path]")
	} else if len(args) == 2 {
		HandleFileInput(args[1])
	} else {
		HandleReplInput()
	}
}
