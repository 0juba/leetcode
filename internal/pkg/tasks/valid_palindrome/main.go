package valid_palindrome

func isPalindrome(s string) bool {
	const (
		A     = 'A'
		Z     = 'Z'
		a     = 'a'
		z     = 'z'
		shift = 'a' - 'A'
		zero  = '0'
		nine  = '9'
	)

	charsCount := 0
	cmpDoneOnceAtLeast := false
	for i, j := 0, len(s)-1; i <= j; {
		lhs := s[i]
		if !((lhs >= a && lhs <= z) || (lhs >= A && lhs <= Z) || (lhs >= zero && lhs <= nine)) {
			i++

			continue
		}

		charsCount++

		rhs := s[j]
		if !((rhs >= a && rhs <= z) || (rhs >= A && rhs <= Z) || (rhs >= zero && rhs <= nine)) {
			j--

			continue
		}

		charsCount++

		if lhs >= A && lhs <= Z {
			lhs += shift
		}

		if rhs >= A && rhs <= Z {
			rhs += shift
		}

		if rhs != lhs {
			return false
		}

		cmpDoneOnceAtLeast = true
		i++
		j--
	}

	if charsCount == 0 {
		return true
	}

	return cmpDoneOnceAtLeast
}
