package _sum

func threeSum(nums []int) [][]int {
	h := make(map[int]int, len(nums))
	for i, n := range nums {
		h[n] = i
	}

	m := make(map[[3]int]struct{}, len(nums))
	for i, target := range nums {
		for j, n := range nums {
			if i != j {
				sum := target + n
				if k, ok := h[-sum]; ok && k != -1 && k != i && k != j {
					m[sortTriplet(target, n, -sum)] = struct{}{}
				}
			}
		}
	}

	var result [][]int
	for triplet := range m {
		sl := make([]int, 3)
		copy(sl, triplet[0:])
		result = append(result, sl)
	}

	return result
}

func sortTriplet(n1, n2, n3 int) [3]int {
	if n1 > n2 {
		n1, n2 = n2, n1
	}

	if n1 > n3 {
		n1, n3 = n3, n1
	}

	if n2 > n3 {
		n2, n3 = n3, n2
	}

	return [3]int{n1, n2, n3}
}
