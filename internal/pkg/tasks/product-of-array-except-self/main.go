package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	h := make([][2]int, len(nums))

	pi, pj := 1, 1
	for i, j := 0, len(nums)-1; i < len(nums) && j > -1; i, j = i+1, j-1 {
		pi *= nums[i]
		h[i][0] = pi

		pj *= nums[j]
		h[j][1] = pj
	}

	result := make([]int, 0, len(nums))
	for i := range nums {
		lhs, rhs := 1, 1
		if i > 0 {
			lhs = h[i-1][0]
		}

		if i < len(nums)-1 {
			rhs = h[i+1][1]
		}

		result = append(result, lhs*rhs)
	}

	return result
}
