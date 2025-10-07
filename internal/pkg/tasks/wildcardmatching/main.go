package wildcardmatching

import "strings"

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}

	if strings.Index(p, "?") != -1 {
		return isMatchQuestion(s, p)
	}

	if p == "*" {
		return true
	}

	if len(p) == 2 {
		if p[0] == '*' {
			return s[len(s)-1] == p[1]
		}
	}

	return false
}

func isMatchQuestion(s string, p string) bool {
	if len(s) != len(p) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] != p[i] && p[i] != '?' {
			return false
		}
	}

	return true
}
