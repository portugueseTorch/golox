package executor

// FIXME: this deserves some refactoring, should probably expose just on arithmetic and one comparison handler
func anyAsFloat(left, right any) (float64, float64, bool) {
	l, leftOk := left.(float64)
	r, rightOk := right.(float64)

	if !leftOk || !rightOk {
		return 0, 0, false
	}

	return l, r, true
}

// attempts to add lhs to rhs
func plus(left, right any) any {
	// add numbers
	switch l := left.(type) {
	case float64:
		// if right is not a number, error
		if r, ok := right.(float64); ok {
			return l + r
		}

	case string:
		// if right is not a string, error
		if r, ok := right.(string); ok {
			return l + r
		}
	}

	panic("unreachable")
}

func minus(left, right any) any {
	l, r, ok := anyAsFloat(left, right)
	if !ok {
		panic("unreachable")
	}

	return l - r
}

func multiply(left, right any) any {
	l, r, ok := anyAsFloat(left, right)
	if !ok {
		panic("unreachable")
	}

	return l * r
}

func divide(left, right any) any {
	l, r, ok := anyAsFloat(left, right)
	if !ok || r == 0 {
		panic("unreachable")
	}

	return l / r
}

func lt(left, right any) bool {
	l, r, ok := anyAsFloat(left, right)
	if !ok {
		panic("unreachable")
	}

	return l < r
}

func lte(left, right any) bool {
	l, r, ok := anyAsFloat(left, right)
	if !ok {
		panic("unreachable")
	}

	return l <= r
}

func gt(left, right any) bool {
	l, r, ok := anyAsFloat(left, right)
	if !ok {
		panic("unreachable")
	}

	return l > r
}

func gte(left, right any) bool {
	l, r, ok := anyAsFloat(left, right)
	if !ok {
		panic("unreachable")
	}

	return l >= r
}

func equal(left, right any) bool {
	// both null
	if left == nil && right == nil {
		return true
	}
	if left == nil {
		return false
	}

	return left == right
}
