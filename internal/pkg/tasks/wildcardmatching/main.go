package wildcardmatching

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}

	if p == "?" && len(s) == 1 {
		return true
	}

	if len(s) == 2 {
		if p == "??" {
			return true
		}

		if p[1] == '?' {
			return p[0] == s[0]
		}

		if p[0] == '?' {
			return p[1] == s[1]
		}
	}

	return false
}
