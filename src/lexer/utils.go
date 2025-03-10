package lexer

// strips all '\' from the raw string
func ParseRawString(raw string) string {
	output := make([]rune, len(raw))
	inBackslash := false

	for _, c := range raw {
		if c == '\\' && !inBackslash {
			inBackslash = true
			continue
		}

		inBackslash = false
		output = append(output, c)
	}

	return string(output)
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func IsAlpha(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || c == '_'
}

func IsAlphaNumeric(c byte) bool {
	return IsAlpha(c) || IsDigit(c)
}
