package wildcardmatching

import "strings"

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}

	if strings.Index(p, "?") != -1 {
		return isMatchQuestion(s, p)
	}

	p = compressAsterisks(p)

	if p == "*" {
		return true
	}

	if len(p) == 2 {
		if p[0] == '*' {
			return s[len(s)-1] == p[1]
		}
		if p[1] == '*' {
			return s[0] == p[0]
		}
	}

	return false
}

func compressAsterisks(s string) string {
	buf := strings.Builder{}
	prev := byte(0)
	for i := range s {
		if s[i] == '*' {
			if prev != '*' {
				buf.WriteByte(s[i])
			}
		} else {
			buf.WriteByte(s[i])
		}

		prev = s[i]
	}

	return buf.String()
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
