package array_and_hashing

import (
	"bytes"
	"fmt"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
	"sort"
)

func containsDuplicateBasedOnRegex(nums []int) bool {
	str := bytes.NewBuffer(make([]byte, len(nums)*4))
	for _, n := range nums {
		str.WriteString(fmt.Sprintf("%d\n", n))
	}

	r, err := pcre.Compile(`^(.*)(\r?\n\1)+$`, pcre.MULTILINE|pcre.CASELESS|pcre.UTF8)
	if err != nil {
		return false
	}

	result := r.ReplaceAll(str.Bytes(), []byte{'_'}, 0)

	return len(result) != str.Len()
}

func containsDuplicateArrayBased(nums []int) bool {
	hash := make([]byte, 2*1e9)
	for _, n := range nums {
		if hash[n] == '0' {
			return true
		}

		hash[n] = '0'
	}

	return false
}

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
