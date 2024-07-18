package longest_consecutive_sequence

import "math"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]int, len(nums))
	for _, n := range nums {
		m[n] = 1
	}

	mx := math.MinInt
	for _, n := range nums {
		v := walk(m, n)
		if v > mx {
			mx = v
		}
	}

	return mx
}

func walk(m map[int]int, target int) int {
	n, ok := m[target]
	if !ok {
		return 0
	}

	if n == -1 {
		return 0
	}

	m[target] = -1

	return 1 + walk(m, target+1) + walk(m, target-1)
}
