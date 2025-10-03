package substringwithconcatenationofallwords

import (
	"sort"
)

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

	total := len(words)

	hash := make(map[string]int, len(words))
	for _, w := range words {
		if _, ok := hash[w]; ok {
			hash[w] += 1
		} else {
			hash[w] = 1
		}
	}

	wordsIndex := make([]string, 0, len(words))
	for w := range hash {
		wordsIndex = append(wordsIndex, w)
	}

	sort.Strings(wordsIndex)

	stat := make([]int, len(wordsIndex))
	for i, w := range wordsIndex {
		stat[i] = hash[w]
	}

	calc := make([]int, len(wordsIndex))
	for i := 0; i < len(s)-substrLen*len(words)+1; i++ {
		copy(calc, stat)
		foundMatches := 0
		for j := i; j < len(s)-substrLen+1; {
			part := s[j : j+substrLen]
			idx := sort.SearchStrings(wordsIndex, part)
			if idx >= len(wordsIndex) {
				break
			}

			if wordsIndex[idx] != part {
				break
			}

			if calc[idx] == 0 {
				break
			}

			foundMatches++
			calc[idx]--
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
