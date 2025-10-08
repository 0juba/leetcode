package wildcardmatching

import "strings"

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}

	if p == "" {
		return false
	}

	p = compressAsterisks(p)
	var patterns []string
	buf := strings.Builder{}
	patternsLength := 0
	for i := range p {
		switch {
		case p[i] == '*':
			if buf.Len() != 0 {
				patternsLength += buf.Len()
				patterns = append(patterns, buf.String())
				buf.Reset()
			}

			patterns = append(patterns, "")
		default:
			buf.WriteByte(p[i])
		}
	}
	if buf.Len() != 0 {
		patternsLength += buf.Len()
		patterns = append(patterns, buf.String())
	}

	pos := 0
	wasAsterisk := false
	for _, pattern := range patterns {
		if pattern == "" {
			wasAsterisk = true
			continue
		}

		if len(s) < patternsLength || len(s[pos:]) < patternsLength {
			return false
		}

		patternsLength -= len(pattern)
		var idx int
		if !wasAsterisk {
			idx = indexByWildcardFromStart(s[pos:], pattern, 0)
		} else {
			idx = indexByWildcardFromEnd(s[pos:len(s)-patternsLength], pattern, 0)
		}
		if idx == -1 {
			return false
		}

		pos += idx + len(pattern)
		wasAsterisk = false
	}

	if len(patterns) != 0 {
		return patternsLength == 0 && (patterns[len(patterns)-1] == "" || pos == len(s))
	}

	return true
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

func indexByWildcardFromEnd(s string, p string, start int) int {
	if len(s) < len(p) {
		return -1
	}

	var i, j int
	for i, j = len(s)-1, len(p)-1; i >= start && j > -1; {
		if s[i] != p[j] && p[j] != '?' {
			i += len(p) - j - 2
			j = len(p) - 1
		} else {
			j--
			i--
		}
	}

	if i+1 < start {
		return -1
	}

	if j == -1 {
		return i + 1
	}

	return -1
}

func indexByWildcardFromStart(s string, p string, start int) int {
	if len(s) < len(p) {
		return -1
	}

	var i, j int
	i = start
	prevStart := start
	for i < len(s) && j < len(p) {
		if s[i] != p[j] && p[j] != '?' {
			i = prevStart + 1
			prevStart = i
			j = 0
		} else {
			j++
			i++
		}
	}

	if j == len(p) {
		return i - len(p)
	}

	return -1
}
