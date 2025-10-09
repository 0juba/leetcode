package substringmatchingpattern

import (
	"testing"
)

func Test_hasMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{
			name: "testcase #870",
			s:    "l",
			p:    "*",
			want: true,
		},
		{
			name: "empty",
			want: true,
		},
		{
			name: "empty pattern",
			s:    "a",
		},
		{
			name: "example #1",
			s:    "leetcode",
			p:    "ee*e",
			want: true,
		},
		{
			name: "example #2",
			s:    "car",
			p:    "c*v",
		},
		{
			name: "test case 93",
			s:    "kvb",
			p:    "k*v",
			want: true,
		},
		{
			name: "example #3",
			s:    "luck",
			p:    "u*",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasMatch(tt.s, tt.p); got != tt.want {
				t.Errorf("hasMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_index(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name   string
		s      string
		substr string
		want   int
	}{
		{
			name:   "test #1",
			s:      "a",
			substr: "a",
			want:   0,
		},
		{
			name:   "test #1",
			s:      "a",
			substr: "aa",
			want:   -1,
		},
		{
			name:   "test #1",
			s:      "foobar",
			substr: "foo",
			want:   0,
		},
		{
			name:   "test #1",
			s:      "foobar",
			substr: "bar",
			want:   3,
		},
		{
			name:   "test #1",
			s:      "",
			substr: "",
			want:   0,
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
