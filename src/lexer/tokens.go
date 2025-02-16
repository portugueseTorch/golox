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
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case LEFT_BRACE:
		return "LEFT_BRACE"
	case RIGHT_BRACE:
		return "RIGHT_BRACE"
	case COMMA:
		return "COMMA"
	case DOT:
		return "DOT"
	case MINUS:
		return "MINUS"
	case PLUS:
		return "PLUS"
	case SEMICOLON:
		return "SEMICOLON"
	case SLASH:
		return "SLASH"
	case STAR:
		return "STAR"
	case BANG:
		return "BANG"
	case BANG_EQUAL:
		return "BANG_EQUAL"
	case EQUAL:
		return "EQUAL"
	case EQUAL_EQUAL:
		return "EQUAL_EQUAL"
	case GREATER:
		return "GREATER"
	case GREATER_EQUAL:
		return "GREATER_EQUAL"
	case LESS:
		return "LESS"
	case LESS_EQUAL:
		return "LESS_EQUAL"
	case AND:
		return "AND"
	case OR:
		return "OR"
	case IF:
		return "IF"
	case ELSE:
		return "ELSE"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case NIL:
		return "NIL"
	case PRINT:
		return "PRINT"
	case RETURN:
		return "RETURN"
	case SUPER:
		return "SUPER"
	case THIS:
		return "THIS"
	case VAR:
		return "VAR"
	case FOR:
		return "FOR"
	case WHILE:
		return "WHILE"
	case FUN:
		return "FUN"
	case CLASS:
		return "CLASS"
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

func (t Token) Literal() *string {
	return t.literal
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
