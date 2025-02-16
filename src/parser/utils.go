package parser

import "golox/src/lexer"

// returns next token to be parsed, advancing cur
// if parsing is done, returns the last token (maybe not ideal)
func (parser *Parser) next() lexer.Token {
	if parser.isAtEnd() {
		return parser.tokens[len(parser.tokens)-1]
	}

	parser.cur += 1
	return parser.tokens[parser.cur-1]
}

func (parser *Parser) peek() lexer.Token {
	return parser.tokens[parser.cur]
}

// returns true if cur is out of bounds of tokens
func (parser *Parser) isAtEnd() bool {
	return parser.peek().TokenType() == lexer.EOF
}

// compares the token type of the token at tokens[cur] against a list of tokenTypes
// returns false if none matches, otherwise returns true and iterates cur
func (parser *Parser) matches(tokTypes ...lexer.TokenType) bool {
	cur := parser.peek()
	for _, t := range tokTypes {
		if t == cur.TokenType() {
			parser.cur += 1
			return true
		}
	}

	return false
}

// returns ths previous token without mutating cur
func (parser *Parser) prev() lexer.Token {
	if parser.cur == 0 {
		return parser.tokens[0]
	}

	return parser.tokens[parser.cur-1]
}
