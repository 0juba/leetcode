package two_sum

func twoSum(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return nil
	}

	hash := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if j, ok := hash[diff]; ok {
			return []int{i, j}
		}

		hash[nums[i]] = i
	}

	return result
}
