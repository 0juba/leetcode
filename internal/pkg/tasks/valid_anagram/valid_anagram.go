package valid_anagram

func isAnagram(s string, t string) bool {
	const (
		size  = 26
		shift = 97
	)
	if len(s) != len(t) {
		return false
	}

	hash := [size]int{}
	for i, j := 0, 0; i < len(s); i, j = i+1, j+1 {
		hash[int(s[i])-shift]++
		hash[int(t[j])-shift]--
	}

	return hash == [size]int{}
}
