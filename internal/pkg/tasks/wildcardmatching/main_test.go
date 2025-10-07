package wildcardmatching

import "testing"

func Test_isMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{
			name: "empty",
			want: true,
		},
		{
			name: "len(s) > len(p)",
			p:    "test",
			s:    "tes",
		},
		{
			name: "?, len = 1",
			p:    "?",
			s:    "t",
			want: true,
		},
		{
			name: "??, len = 2",
			p:    "??",
			s:    "tt",
			want: true,
		},
		{
			name: "a?, len = 2",
			p:    "a?",
			s:    "at",
			want: true,
		},
		{
			name: "?a, len = 2",
			p:    "?a",
			s:    "ta",
			want: true,
		},
		{
			name: "?aa, len = 3",
			p:    "?aa",
			s:    "taa",
			want: true,
		},
		{
			name: "a?a, len = 3",
			p:    "a?a",
			s:    "aba",
			want: true,
		},
		{
			name: "aa?, len = 3",
			p:    "aa?",
			s:    "aab",
			want: true,
		},
		{
			name: "a??, len = 3",
			p:    "a??",
			s:    "atb",
			want: true,
		},
		{
			name: "?a?, len = 3",
			p:    "?a?",
			s:    "ear",
			want: true,
		},
		{
			name: "??a, len = 3",
			p:    "??a",
			s:    "era",
			want: true,
		},
		{
			name: "hello world",
			p:    "h?l?o?w?rl?",
			s:    "hello world",
			want: true,
		},
		{
			name: "asterisk 1",
			p:    "*",
			s:    "a",
			want: true,
		},
		{
			name: "asterisk 0",
			p:    "*",
			s:    "",
			want: true,
		},
		{
			name: "asterisk *a",
			p:    "*a",
			s:    "a",
			want: true,
		},
		{
			name: "asterisk *a len = 3",
			p:    "*a",
			s:    "bca",
			want: true,
		},
		{
			name: "asterisk a* len = 4",
			p:    "a*",
			s:    "abcd",
			want: true,
		},
		{
			name: "asterisk ******** len = 4",
			p:    "*******",
			s:    "abcd",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.s, tt.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
