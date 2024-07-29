package valid_parentheses

func isValid(s string) bool {
	const (
		roundOpen     = '('
		roundClosing  = ')'
		squareOpen    = '['
		squareClosing = ']'
		curvedOpen    = '{'
		curvedClosing = '}'
	)

	isOpening := func(pr byte) bool {
		return pr == roundOpen || pr == squareOpen || pr == curvedOpen
	}

	isClosing := func(pr byte) bool {
		return pr == roundClosing || pr == squareClosing || pr == curvedClosing
	}

	isValidPair := func(pr1, pr2 byte) bool {
		switch pr1 {
		case roundOpen:
			return roundClosing == pr2
		case squareOpen:
			return squareClosing == pr2
		case curvedOpen:
			return curvedClosing == pr2
		default:
			panic("unsupported char")
		}
	}

	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch {
		case len(stack) == 0 && isClosing(s[i]):
			return false
		case len(stack) == 0:
			stack = append(stack, s[i])
		case isOpening(s[i]):
			stack = append(stack, s[i])
		case !isValidPair(stack[len(stack)-1], s[i]):
			return false
		default:
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
