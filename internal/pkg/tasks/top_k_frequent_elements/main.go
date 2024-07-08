package top_k_frequent_elements

import "sort"

type (
	pair     [2]int
	pairList []pair
)

func (p pairList) Len() int {
	return len(p)
}

func (p pairList) Less(i, j int) bool {
	return p[i][1] > p[j][1]
}

func (p pairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	m := make(map[int]int, len(nums))
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = 0
		}

		m[n]++
	}

	pairs := make(pairList, 0, len(nums))
	for n, count := range m {
		pairs = append(pairs, pair{n, count})
	}

	sort.Sort(pairs)

	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, pairs[i][0])
	}

	return result
}
