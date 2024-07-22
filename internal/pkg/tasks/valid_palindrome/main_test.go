package valid_palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
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
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: `example #2`,
			s:    "race a car",
		},
		{
			name: `example #3`,
			s:    " ",
			want: true,
		},
		{
			name: `test case #462`,
			s:    "0P",
		},
		{
			name: `test case #457`,
			s:    "a.",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isPalindrome(tt.s))
		})
	}
}
