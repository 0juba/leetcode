package group_anagrams

import "strconv"

func groupAnagrams(strs []string) [][]string {
	const (
		size       = 26
		shiftAscii = 97
	)

	// 1. calculate hashes
	m := make(map[string][]string, len(strs))
	for _, w := range strs {
		wHash := [size]int{}
		for i := 0; i < len(w); i++ {
			wHash[int(w[i])-shiftAscii]++
		}

		key := []byte{}
		for i := 0; i < len(wHash); i++ {
			if wHash[i] != 0 {
				key = append(key, byte(i+shiftAscii))
				key = strconv.AppendInt(key, int64(wHash[i]), 10)
			}
		}

		sKey := string(key)
		if _, ok := m[sKey]; !ok {
			m[sKey] = make([]string, 0)
		}

		words := m[sKey]
		words = append(words, w)
		m[sKey] = words
	}

	result := [][]string{}
	for _, bag := range m {
		result = append(result, bag)
	}

	return result
}
