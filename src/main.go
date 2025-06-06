package main

import (
	"bufio"
	"fmt"
	"golox/src/executor"
	"golox/src/lexer"
	"golox/src/parser"
	"golox/src/resolver"
	"os"
)

func run(input string, env *executor.Environment) (any, error) {
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

	executor := executor.NewExecutor(parsed, env)
	resolver := resolver.NewResolver(executor)
	_, err = resolver.Resolve(parsed)
	if err != nil {
		fmt.Printf("%s", err)
		return nil, nil
	}

	_, err = executor.Execute()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return nil, nil
}

func HandleFileInput(filePath string) {
	// --- load file into memory
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("[ERROR]:", err)
		return
	}

	// --- execution environment
	env := executor.NewEnvironment(nil)
	_, runtime_err := run(string(file), env)
	if runtime_err != nil {
		os.Exit(70)
	}
}

func HandleReplInput() {
	scanner := bufio.NewScanner(os.Stdin)
	// --- execution environment
	env := executor.NewEnvironment(nil)

	for {
		fmt.Print(">> ")

		// --- handle empty input
		if !scanner.Scan() {
			fmt.Println("\nExiting...")
			return
		}

		input := scanner.Text()
		run(input, env)
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
