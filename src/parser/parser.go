package parser

import (
	"golox/src/ast"
	"golox/src/lexer"
)

type Parser struct {
	tokens []lexer.Token
	cur    int
}

func NewParser(tokens []lexer.Token) Parser {
	return Parser{
		tokens: tokens,
		cur:    0,
	}
}

func (parser *Parser) Parse() ([]ast.Stmt, error) {
	stmtList := make([]ast.Stmt, 0)

	for !parser.isAtEnd() {
		stmt, err := parser.declaration()
		if err != nil {
			return nil, err
		}

		stmtList = append(stmtList, stmt)
	}

	return stmtList, nil
}
