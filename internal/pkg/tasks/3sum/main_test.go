package _sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: `empty`,
		},
		{
			name: `example #1`,
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: `example #2`,
			nums: []int{0, 1, 1},
		},
		{
			name: `example #3`,
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, threeSum(tt.nums))
		})
	}
}
