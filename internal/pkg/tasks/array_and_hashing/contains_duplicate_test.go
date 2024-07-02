package array_and_hashing

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"sort"
	"testing"
)

func helperGenSlice() []int {
	const sliceSize = 1e5
	nums := make([]int, 0, sliceSize)
	for i := 0; i < sliceSize; i++ {
		nums = append(nums, i+1)
	}

	return nums
}

func FuzzContainsDuplicates(f *testing.F) {
	testCases := helperGenSlice()

	f.Add(testCases[0], testCases[1], testCases[2], testCases[3], testCases[4], testCases[5], testCases[6],
		testCases[7], testCases[8], testCases[9])

	f.Fuzz(func(t *testing.T, n1 int, n2 int, n3 int, n4 int, n5 int, n6 int, n7 int, n8 int, n9 int, n10 int) {
		nums := []int{n1, n2, n3, n4, n5, n6, n7, n8, n9, n10}
		hasDuplicates := containsDuplicate(nums)
		sort.Ints(nums)
		compacted := slices.Compact(nums)

		assert.Equal(t, !hasDuplicates, len(nums) == len(compacted))
	})
}

// 11    	      27	  43078397
func BenchmarkContainsDuplicatesRegexBased(b *testing.B) {
	for i := 0; i < b.N; i++ {
		containsDuplicateBasedOnRegex(helperGenSlice())
	}
}

// Result: 11    	       1	2017471625 ns/op
func BenchmarkContainsDuplicatesBruteForceBased(b *testing.B) {
	for i := 0; i < b.N; i++ {
		containsDuplicateBruteForce(helperGenSlice())
	}
}

// 11    	   14612	     80911 ns/op
// 11    	    8841	    132812 ns/op
func BenchmarkContainsDuplicateSortBased(b *testing.B) {
	for i := 0; i < b.N; i++ {
		containsDuplicateSortBased(helperGenSlice())
	}
}

// Result: 11    	     433	   2486287 ns/op
func BenchmarkContainsDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		containsDuplicate(helperGenSlice())
	}
}

// 11    	     100	  18383199
func BenchmarkContainsDuplicateArrayBased(b *testing.B) {
	for i := 0; i < b.N; i++ {
		containsDuplicateArrayBased(helperGenSlice())
	}
}

func Test_containsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "empty",
		},
		{
			name: "one element",
			nums: []int{1},
		},
		{
			name: "two unique elements",
			nums: []int{1, 2},
		},
		{
			name: "three items, two duplicates",
			nums: []int{1, 2, 2},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, containsDuplicate(tt.nums))
			assert.Equal(t, tt.want, containsDuplicateSortBased(tt.nums))
			assert.Equal(t, tt.want, containsDuplicateBruteForce(tt.nums))
			assert.Equal(t, tt.want, containsDuplicateArrayBased(tt.nums))
			assert.Equal(t, tt.want, containsDuplicateBasedOnRegex(tt.nums))
		})
	}
}
