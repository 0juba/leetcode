package two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name: `empty`,
		},
		{
			name:   `leetcode example #1`,
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   `leetcode example #2`,
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   `leetcode example #3`,
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, twoSum(tt.nums, tt.target))
		})
	}
}
