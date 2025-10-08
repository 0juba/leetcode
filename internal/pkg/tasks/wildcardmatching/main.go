package wildcardmatching

import "strings"

func isMatch(s string, p string) bool {
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
		matches = append(matches, []int{len(s)})
		patterns = append(patterns, "")
	}

	if len(matches) == 0 {
		return false
	}

	prevMaxPos := -1

	var viewedLen int
	if p[len(p)-1] == '*' {
		viewedLen = 0
	} else {
		last := matches[len(matches)-1]
		viewedLen = len(s) - last[len(last)-1]
	}

	for i := len(matches) - 1; i > -1; i-- {
		positions := matches[i]
		if prevMaxPos == -1 {
			prevMaxPos = positions[len(positions)-1]
		} else {
			found := false
			for j := len(positions) - 1; j > -1; j-- {
				if positions[j] < prevMaxPos {
					validSubstrLen := true
					if i-delta > -1 {
						validSubstrLen = len(patterns[i-delta]) <= prevMaxPos-positions[j]
					}

					if validSubstrLen {
						found = true
						if i == 0 {
							viewedLen += prevMaxPos - positions[0]
							prevMaxPos = positions[0]
						} else {
							viewedLen += prevMaxPos - positions[j]
							prevMaxPos = positions[j]
						}
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
	}

	return viewedLen == len(s)
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
			i += len(p)
		} else {
			i++
		}
	}

	return result
}
