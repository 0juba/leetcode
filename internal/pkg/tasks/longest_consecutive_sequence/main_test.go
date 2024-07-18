package longest_consecutive_sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: `empty`,
		},
		{
			name: `example #1`,
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: `example #2`,
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestConsecutive(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
