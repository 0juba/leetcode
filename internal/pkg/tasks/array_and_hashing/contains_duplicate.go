package array_and_hashing

import "sort"

func containsDuplicateBruteForce(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

func containsDuplicateSortBased(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}

func containsDuplicate(nums []int) bool {
	hash := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := hash[n]; ok {
			return true
		}

		hash[n] = struct{}{}
	}

	return false
}
