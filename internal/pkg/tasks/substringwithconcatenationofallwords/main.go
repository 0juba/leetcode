package substringwithconcatenationofallwords

import "maps"

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return nil
	}

	if len(s) < len(words) {
		return nil
	}

	if len(s) < len(words[0]) {
		return nil
	}

	var indexes []int

	wordLen := len(words[0])
	substrLen := wordLen

	total := 0
	hash := make(map[string]int, len(words))
	for _, w := range words {
		total++
		if _, ok := hash[w]; ok {
			hash[w] += 1
		} else {
			hash[w] = 1
		}
	}

	calc := make(map[string]int, len(words))
	for i := 0; i < len(s)-substrLen+1; i++ {
		maps.Copy(calc, hash)
		foundMatches := 0
		for j := i; j < len(s)-substrLen+1; {
			part := s[j : j+substrLen]
			if _, ok := hash[part]; !ok {
				break
			}

			if calc[part] == 0 {
				break
			}

			foundMatches++
			calc[part]--
			j += wordLen
			if foundMatches == total {
				break
			}
		}

		matches := 0
		for _, count := range calc {
			matches += count
		}

		if matches == 0 {
			indexes = append(indexes, i)
		}
	}

	return indexes
}
