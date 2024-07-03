package valid_anagram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: `empty set`,
			want: true,
		},
		{
			name: `leetcode example #1`,
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: `leetcode example #2`,
			s:    "cat",
			t:    "rat",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isAnagram(tt.s, tt.t))
		})
	}
}
