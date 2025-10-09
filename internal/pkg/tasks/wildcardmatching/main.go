package wildcardmatching

import "strings"

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}

	if p == "" {
		return false
	}

	parts := split(p, '*')
	pos := 0
	lastAsterisk := -1
	for i, part := range parts {
		if part == "" {
			lastAsterisk = i
			continue
		}

		idx := index(s[pos:], part)
		if idx == -1 || (lastAsterisk == -1 && idx != 0) {
			return false
		}

		pos += idx + len(part)
	}

	if len(s) == pos || lastAsterisk == len(parts)-1 {
		return true
	}

	if lastAsterisk != -1 {
		lastTpl := parts[len(parts)-1]

		return index(s[len(s)-len(lastTpl):], lastTpl) != -1
	}

	return false
}

func split(s string, ch byte) []string {
	var result []string
	buf := strings.Builder{}
	for i := range s {
		if s[i] == ch {
			if buf.Len() != 0 {
				result = append(result, buf.String())
				buf.Reset()
			}
			result = append(result, "")
		} else {
			buf.WriteByte(s[i])
		}
	}

	if buf.Len() != 0 {
		result = append(result, buf.String())
	}

	return result
}

func index(s, substr string) int {
	for pos := 0; pos < len(s); {
		if len(substr)+pos <= len(s) {
			k := pos
			for ; k < pos+len(substr); k++ {
				if s[k] != substr[k-pos] && substr[k-pos] != '?' {
					break
				}
			}

			if k == pos+len(substr) {
				return pos
			}

			pos++
		} else {
			return -1
		}
	}

	return -1
}
