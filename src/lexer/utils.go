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
