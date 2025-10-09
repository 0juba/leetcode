package wildcardmatching

import (
	"strings"
)

func isMatch(s string, p string) bool {
	if p == "*" || p == "" && s == "" {
		return true
	}

	if p == "" {
		return false
	}

	var patterns []string
	buf := strings.Builder{}
	for i := range p {
		switch {
		case p[i] == '*':
			if buf.Len() != 0 {
				patterns = append(patterns, buf.String())
				buf.Reset()
			}
			patterns = append(patterns, "*")
		default:
			buf.WriteByte(p[i])
		}
	}
	if buf.Len() != 0 {
		patterns = append(patterns, buf.String())
	}

	pos := 0
	hasAsterisk := false
	for i := 0; i < len(patterns); i++ {
		if patterns[i] == "*" {
			hasAsterisk = true
			continue
		}

		matchFound := false
		for ; len(s) >= pos+len(patterns[i]); pos++ {
			if match(s[pos:pos+len(patterns[i])], patterns[i]) {
				matchFound = true
				pos += len(patterns[i])
				break
			} else if !hasAsterisk {
				return false
			}
		}

		if !matchFound {
			return false
		}

		hasAsterisk = false
	}

	return hasAsterisk || len(s) == pos
}

func isMatch2(s string, p string) bool {
	if s == p {
		return true
	}

	if p == "" {
		return false
	}

	var patterns []string
	buf := strings.Builder{}
	for i := range p {
		switch {
		case p[i] == '*':
			if buf.Len() != 0 {
				patterns = append(patterns, buf.String())
				buf.Reset()
			}
		default:
			buf.WriteByte(p[i])
		}
	}
	if buf.Len() != 0 {
		patterns = append(patterns, buf.String())
	}

	if len(patterns) == 0 {
		return true
	}

	if len(patterns) > len(s) {
		return false
	}

	delta := 0
	matches := make([][]int, 0, len(patterns))
	if p[0] == '*' {
		matches = append(matches, []int{0})
		delta++
	}
	for _, pattern := range patterns {
		found := findAllMatches(s, pattern)
		if len(found) == 0 {
			return false
		}

		matches = append(matches, found)
	}

	if p[len(p)-1] == '*' {
		matches = append(matches, []int{len(s) - 1})
		patterns = append(patterns, "")
	}

	if len(matches) == 0 {
		return false
	}

	prevMaxPos := len(s)

	var viewedLen int

	for i := len(matches) - 1; i > -1; i-- {
		positions := matches[i]
		found := false
		for j := len(positions) - 1; j > -1; j-- {
			if positions[j] < prevMaxPos {
				validSubstrLen := true
				if i-delta > -1 {
					validSubstrLen = len(patterns[i-delta]) <= prevMaxPos-positions[j]
				}

				if validSubstrLen {
					found = true
					viewedLen += prevMaxPos - positions[j]
					prevMaxPos = positions[j]

					break
				}
			} else if i == 0 && positions[j] == prevMaxPos {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return viewedLen == len(s)
}

func match(s, p string) bool {
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

func findAllMatches(s, p string) []int {
	var result []int
	for i := 0; i < len(s); {
		j := 0
		for ; j < len(p) && i+j < len(s); j++ {
			if p[j] != s[i+j] && p[j] != '?' {
				break
			}
		}
		if j == len(p) {
			result = append(result, i)
		}
		i++
	}

	return result
}
