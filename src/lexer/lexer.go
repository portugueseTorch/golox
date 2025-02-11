package lexer

import (
	"fmt"
)

type Lexer struct {
	input string
	start int
	cur   int
	line  int
	//
	hasError bool
	tokens   []Token
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{
		input:  input,
		cur:    0,
		line:   1,
		start:  0,
		tokens: make([]Token, 0),
	}

	return lexer
}

func (lex *Lexer) HasError() bool { return lex.hasError }

func (lex *Lexer) GetTokens() []Token {
	return lex.tokens
}

func (lex *Lexer) ScanTokens() {
	for lex.cur < len(lex.input) {
		lex.start = lex.cur
		lex.scanToken()
	}

	lex.appendToken(EOF, nil)
}

// scans token and appends it to lex.tokens
func (lex *Lexer) scanToken() {
	c := lex.next()

	switch c {
	case '\r', ' ', '\t':
		break
	case '\n':
		lex.line += 1
	case '(':
		lex.appendToken(LEFT_PAREN, nil)
	case ')':
		lex.appendToken(RIGHT_PAREN, nil)
	case '{':
		lex.appendToken(LEFT_BRACE, nil)
	case '}':
		lex.appendToken(RIGHT_BRACE, nil)
	case ',':
		lex.appendToken(COMMA, nil)
	case '.':
		lex.appendToken(DOT, nil)
	case '+':
		lex.appendToken(PLUS, nil)
	case '-':
		lex.appendToken(MINUS, nil)
	case ';':
		lex.appendToken(SEMICOLON, nil)
	case '*':
		lex.appendToken(STAR, nil)
	case '!', '=', '<', '>':
		lex.appendToken(ToTokenType(c, lex.matches('=')), nil)
	case '"':
		lex.buildStringToken()
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		lex.buildNumericToken()
	case '/':
		// --- if next character is also '/', ignore everything until end of the line
		if lex.matches('/') {
			for !lex.isAtEnd() && lex.peek() != '\n' {
				lex.next()
			}
		} else {
			lex.appendToken(SLASH, nil)
		}
	default:
		lex.LogError(fmt.Sprintf("unexpected token '%c'", c))
		lex.hasError = true
	}
}

func (lex *Lexer) buildNumericToken() {
	// --- first half of the number
	for !lex.isAtEnd() && IsDigit(lex.peek()) {
		lex.next()
	}

	// --- if the current char is '.' but the next is not a digit, error
	if lex.peek() == '.' {
		if !IsDigit(lex.peekNext()) {
			lex.LogError("trailing '.'")
			return
		}

		// --- advance from '.'
		lex.next()

		// --- second half of the number, if applicable
		for !lex.isAtEnd() && IsDigit(lex.peek()) {
			lex.next()
		}
	}

	rawNumber := lex.input[lex.start:lex.cur]
	lex.appendToken(NUMBER, &rawNumber)
}

func (lex *Lexer) buildStringToken() {
	for !lex.isAtEnd() && lex.peek() != '"' {
		// --- increment line
		if lex.peek() == '\n' {
			lex.line += 1
		}

		// --- if the current character is '\' and the next is '"', jump twice ahead
		if lex.peek() == '\\' && lex.peekNext() == '"' {
			lex.next()
		}

		lex.next()
	}

	// --- if we're at the end of the file, error with unterminated string
	if lex.isAtEnd() {
		lex.LogError("unterminated string")
		return
	}

	rawString := lex.input[lex.start+1 : lex.cur]
	parsedString := ParseRawString(rawString)

	lex.appendToken(STRING, &parsedString)

	// --- skip last '"'
	lex.next()
}

// returns current byte and advances cur
func (lex *Lexer) next() byte {
	lex.cur += 1
	return lex.input[lex.cur-1]
}

func (lex *Lexer) isAtEnd() bool {
	return lex.cur >= len(lex.input)
}

// returns current byte without advancing
func (lex *Lexer) peek() byte {
	if lex.cur >= len(lex.input) {
		return 0
	}

	return lex.input[lex.cur]
}

// looks ahead to the next byte withoug advancing
func (lex *Lexer) peekNext() byte {
	if lex.cur+1 >= len(lex.input) {
		return 0
	}

	return lex.input[lex.cur+1]
}

// compares the byte at position cur returning false if it does not match
// cmp - if it matches, iterates to next position and returns true
func (lex *Lexer) matches(cmp byte) bool {
	if lex.cur >= len(lex.input) || lex.input[lex.cur] != cmp {
		return false
	}

	lex.cur += 1
	return true
}

// appends token with literal if bool is true
func (lex *Lexer) appendToken(tokenType TokenType, literal *string) {
	tok := Token{
		start:  lex.start,
		length: lex.cur - lex.start,
		line:   lex.line,
		//
		tokenType: tokenType,
		literal:   literal,
	}

	lex.tokens = append(lex.tokens, tok)
}

func (lex *Lexer) LogError(err string) {
	lex.hasError = true
	fmt.Printf("[ERROR]: %s at line %d:%d\n", err, lex.line, lex.cur)
}
