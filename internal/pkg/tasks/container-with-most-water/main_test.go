package container_with_most_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name: `empty`,
		},
		{
			name:   `example #1`,
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, maxArea(tt.height), tt.want)
		})
	}
}
