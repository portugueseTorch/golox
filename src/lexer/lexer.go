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

	lex.appendToken(EOF, false)
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
		lex.appendToken(LEFT_PAREN, false)
	case ')':
		lex.appendToken(RIGHT_PAREN, false)
	case '{':
		lex.appendToken(LEFT_BRACE, false)
	case '}':
		lex.appendToken(RIGHT_BRACE, false)
	case ',':
		lex.appendToken(COMMA, false)
	case '.':
		lex.appendToken(DOT, false)
	case '+':
		lex.appendToken(PLUS, false)
	case '-':
		lex.appendToken(MINUS, false)
	case ';':
		lex.appendToken(SEMICOLON, false)
	case '*':
		lex.appendToken(STAR, false)
	case '!', '=', '<', '>':
		if lex.matches('=') {
			lex.appendToken(ToTokenTypeWithEqual(c), false)
		} else {
			lex.appendToken(ToTokenType(c), false)
		}
	case '/':
		// --- if next character is also '/', ignore everything until end of the line
		if lex.matches('/') {
			for !lex.isAtEnd() && lex.peek() != '\n' {
				lex.next()
			}
		} else {
			lex.appendToken(SLASH, false)
		}
	default:
		lex.LogError(fmt.Sprintf("unexpected token '%c'", c))
		lex.hasError = true
	}
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
func (lex *Lexer) appendToken(tokenType TokenType, literal bool) {
	var raw *string = nil
	if literal {
		tmpRaw := lex.input[lex.start : lex.cur+1]
		raw = &tmpRaw
	}

	tok := Token{
		start:  lex.start,
		length: lex.cur - lex.start,
		line:   lex.line,
		//
		tokenType: tokenType,
		raw:       raw,
	}

	lex.tokens = append(lex.tokens, tok)
}

func (lex *Lexer) LogError(err string) {
	// lineAsString := strconv.Itoa(lex.line)
	// var lineDigits int = len(lineAsString)
	// spaces := strings.Repeat(" ", 12+lineDigits+lex.cur)

	fmt.Printf("[ERROR]: %s at line %d:%d\n", err, lex.line, lex.cur)
	// fmt.Printf("   └───> %d | %s", lex.line, lex.input[lex.start:lex.cur])
	// fmt.Printf("%s^-- error\n", spaces)
}
