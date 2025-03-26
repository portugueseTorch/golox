package lexer

import "fmt"

type TokenType int

const (
	INVALID TokenType = iota
	EOF
	//
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	//
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	//
	AND
	OR
	IF
	ELSE
	TRUE
	FALSE
	NIL
	PRINT
	RETURN
	SUPER
	THIS
	VAR
	FOR
	WHILE
	FUN
	CLASS
	IDENTIFIER
	STRING
	NUMBER
)

func (t TokenType) String() string {
	switch t {
	case INVALID:
		return "INVALID"
	case EOF:
		return "EOF"
	case LEFT_PAREN:
		return "("
	case RIGHT_PAREN:
		return ")"
	case LEFT_BRACE:
		return "{"
	case RIGHT_BRACE:
		return "}"
	case COMMA:
		return ","
	case DOT:
		return "."
	case MINUS:
		return "-"
	case PLUS:
		return "+"
	case SEMICOLON:
		return ";"
	case SLASH:
		return "/"
	case STAR:
		return "*"
	case BANG:
		return "!"
	case BANG_EQUAL:
		return "!="
	case EQUAL:
		return "="
	case EQUAL_EQUAL:
		return "=="
	case GREATER:
		return ">"
	case GREATER_EQUAL:
		return ">="
	case LESS:
		return "<"
	case LESS_EQUAL:
		return "<="
	case AND:
		return "and"
	case OR:
		return "or"
	case IF:
		return "if"
	case ELSE:
		return "else"
	case TRUE:
		return "true"
	case FALSE:
		return "false"
	case NIL:
		return "nil"
	case PRINT:
		return "print"
	case RETURN:
		return "return"
	case SUPER:
		return "super"
	case THIS:
		return "this"
	case VAR:
		return "var"
	case FOR:
		return "for"
	case WHILE:
		return "while"
	case FUN:
		return "fun"
	case CLASS:
		return "class"
	case IDENTIFIER:
		return "IDENTIFIER"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"
	default:
		return "UNKNOWN"
	}
}

type Token struct {
	// offset from the start of the file
	start int
	// length of the token literal
	length int
	// line number
	line int

	tokenType TokenType
	literal   *string
}

func ToTokenType(char byte, withEqual bool) TokenType {
	switch char {
	case '!':
		if withEqual {
			return BANG_EQUAL
		}
		return BANG
	case '=':
		if withEqual {
			return EQUAL_EQUAL
		}
		return EQUAL
	case '<':
		if withEqual {
			return LESS_EQUAL
		}
		return LESS
	case '>':
		if withEqual {
			return GREATER_EQUAL
		}
		return GREATER
	default:
		return INVALID
	}
}

func (t Token) String() string {
	if t.literal == nil {
		return fmt.Sprintf("[%s] at %d:%d", t.tokenType, t.line, t.start+1)
	}
	return fmt.Sprintf("[%s]: %s at %d:%d", t.tokenType, *t.literal, t.line, t.start+1)
}

func (t Token) Type() string {
	return t.tokenType.String()
}

func (t Token) TokenType() TokenType {
	return t.tokenType
}

func (t Token) Literal() string {
	return *t.literal
}

func (t Token) Line() int {
	return t.line
}

func NewToken(tokenType TokenType) Token {
	return Token{
		start:     0,
		length:    0,
		line:      0,
		tokenType: tokenType,
		literal:   nil,
	}
}
