package wildcardmatching

import (
	"reflect"
	"testing"
)

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
			name: "empty wildcard",
			s:    "a",
		},
		{
			name: "len(s) > len(p)",
			p:    "test",
			s:    "tes",
		},
		{
			name: "example #1",
			p:    "t",
			s:    "tt",
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
		{
			name: "asterisk *aa len = 3",
			p:    "*ab",
			s:    "readyab",
			want: true,
		},
		{
			name: "asterisk aa* len = 3",
			p:    "ab*",
			s:    "abready",
			want: true,
		},
		{
			name: "asterisk a*a len = 3",
			p:    "a*b",
			s:    "areadtb",
			want: true,
		},
		{
			name: "asterisk a*a len = 3",
			p:    "*a*b",
			s:    "tesareadtb",
			want: true,
		},
		{
			name: "asterisk *a* len = 3",
			p:    "*a*",
			s:    "testaaaaaatestaaaaaa",
			want: true,
		},
		{
			name: "asterisk *a?aa len = 3",
			p:    "*a?aa",
			s:    "testaaaaaatestaaaaaa",
			want: true,
		},
		{
			name: "testcase 1489",
			p:    "***a",
			s:    "aaaa",
			want: true,
		},
		{
			name: "testcase 1801",
			s:    "abb",
			p:    "**??",
			want: true,
		},
		{
			name: "testcase 30",
			p:    "*abc???de*",
			s:    "abcabczzzde",
			want: true,
		},
		{
			name: "testcase 1594",
			s:    "mississippi",
			p:    "*miss*iss*",
			want: true,
		},
		{
			name: "testcase 1586",
			s:    "mississippi",
			p:    "m*iss*iss*",
			want: true,
		},
		{
			name: "testcase 1586 #1",
			s:    "missississippi",
			p:    "m*iss*iss*",
			want: true,
		},
		{
			name: "testcase 1586 #1",
			s:    "ississississ",
			p:    "*iss",
			want: true,
		},
		{
			name: "testcase 1530",
			s:    "mississippi",
			p:    "m??*ss*?i*pi",
		},
		{
			name: "testcase 94",
			s:    "abcdef",
			p:    "a?de*",
			want: false,
		},
		{
			name: "testcase 1512",
			s:    "aaa",
			p:    "aa",
			want: false,
		},
		{
			name: "testcase 1522",
			s:    "b",
			p:    "?*?",
			want: false,
		},
		{
			name: "testcase 1523",
			s:    "ab",
			p:    "*a",
			want: false,
		},
		{
			name: "testcase 1623",
			s:    "aaab",
			p:    "b**",
			want: false,
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

func Test_index(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		substr string
		want   int
	}{
		{
			name:   "test #1",
			s:      "aaa",
			substr: "a",
			want:   0,
		},
		{
			name:   "test #2",
			s:      "aaa",
			substr: "?",
			want:   0,
		},
		{
			name:   "test #3",
			s:      "aaa",
			substr: "a?",
			want:   0,
		},
		{
			name:   "test #4",
			s:      "aaa",
			substr: "a??",
			want:   0,
		},
		{
			name:   "test #5",
			s:      "aaa",
			substr: "???",
			want:   0,
		},
		{
			name:   "test #6",
			s:      "caaa",
			substr: "a???",
			want:   -1,
		},
		{
			name:   "test #7",
			s:      "ab",
			substr: "b",
			want:   1,
		},
		{
			name:   "test #8",
			s:      "abc",
			substr: "b?",
			want:   1,
		},
		{
			name:   "test #9",
			s:      "rabc",
			substr: "?b?",
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := index(tt.s, tt.substr); got != tt.want {
				t.Errorf("index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_split(t *testing.T) {
	type args struct {
		s  string
		ch byte
	}
	tests := []struct {
		name string
		s    string
		ch   byte
		want []string
	}{
		{
			name: "test #1",
			s:    "*a*b*c*",
			ch:   '*',
			want: []string{
				"", "a", "", "b", "", "c", "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := split(tt.s, tt.ch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("split() = %v, want %v", got, tt.want)
			}
		})
	}
}
