package main

import (
	"bufio"
	"fmt"
	"golox/src/executor"
	"golox/src/lexer"
	"golox/src/parser"
	"os"
)

func run(input string) (any, error) {
	lexer := lexer.NewLexer(input)
	lexer.ScanTokens()
	if lexer.HasError() {
		return nil, nil
	}

	parser := parser.NewParser(lexer.GetTokens())
	parsed, err := parser.Parse()
	if err != nil {
		fmt.Printf("%s", err)
		return nil, nil
	}

	ret, err := executor.ExecuteAST(parsed)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return ret, nil
}

func HandleFileInput(filePath string) {
	// --- load file into memory
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("[ERROR]:", err)
		return
	}

	_, runtime_err := run(string(file))
	if runtime_err != nil {
		os.Exit(70)
	}
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
