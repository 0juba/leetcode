package two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		sum := numbers[i] + numbers[j]
		switch {
		case sum == target:
			return []int{i + 1, j + 1}
		case sum > target:
			j--
		case sum < target:
			i++
		}
	}

	panic("pair not exists")
}
