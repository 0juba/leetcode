package valid_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: `empty`,
			want: true,
		},
		{
			name: `example #1`,
			s:    "()",
			want: true,
		},
		{
			name: `example #2`,
			s:    "[](){}",
			want: true,
		},
		{
			name: `example #3`,
			s:    "(}",
		},
		{
			name: `testcase #27`,
			s:    "))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isValid(tt.s))
		})
	}
}
