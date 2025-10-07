package longestvalidparentheses

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "empty string",
		},
		{
			name: "1 ch (",
			s:    "(",
		},
		{
			name: "1 ch )",
			s:    ")",
		},
		{
			name: "2 ch ()",
			s:    "()",
			want: 2,
		},
		{
			name: "2 ch )(",
			s:    ")(",
		},
		{
			name: "3 ch (()",
			s:    "(()",
			want: 2,
		},
		{
			name: "3 ch ())",
			s:    "())",
			want: 2,
		},
		{
			name: "3 ch )()",
			s:    ")()",
			want: 2,
		},
		{
			name: "3 ch ()(",
			s:    "()(",
			want: 2,
		},
		{
			name: "3 ch (((",
			s:    "(((",
		},
		{
			name: "3 ch )))",
			s:    ")))",
		},
		{
			name: "4 ch ))))",
			s:    "))))",
		},
		{
			name: "4 ch ((((",
			s:    "((((",
		},
		{
			name: "4 ch ()((",
			s:    "()((",
			want: 2,
		},
		{
			name: "4 ch ()()",
			s:    "()()",
			want: 4,
		},
		{
			name: "4 ch (())",
			s:    "(())",
			want: 4,
		},
		{
			name: "5 ch ()())",
			s:    "()())",
			want: 4,
		},
		{
			name: "5 ch (()()",
			s:    "(()()",
			want: 4,
		},
		{
			name: "5 ch ()(()",
			s:    "()(()",
			want: 2,
		},
		{
			name: "5 ch (()))",
			s:    "(()))",
			want: 4,
		},
		{
			name: "6 ch ()()()",
			s:    "()()()",
			want: 6,
		},
		{
			name: "6 ch (()())",
			s:    "(()())",
			want: 6,
		},
		{
			name: "6 ch ((()))",
			s:    "((()))",
			want: 6,
		},
		{
			name: "6 ch (())()",
			s:    "(())()",
			want: 6,
		},
		{
			name: "6 ch ()(())",
			s:    "()(())",
			want: 6,
		},
		{
			name: "6 ch ()()))",
			s:    "()()))",
			want: 4,
		},
		{
			name: "6 ch ())())",
			s:    "())())",
			want: 2,
		},
		{
			name: "7 ch ()()())",
			s:    "()()())",
			want: 6,
		},
		{
			name: "7 ch (()()()",
			s:    "(()()()",
			want: 6,
		},
		{
			name: "test case 171",
			s:    "(()(((()",
			want: 2,
		},
		{
			name: "test case 171 - full valid",
			s:    "(()(((()))))",
			want: 12,
		},
		{
			name: "test case 171 - full valid 1",
			s:    "(()(((()))))(",
			want: 12,
		},
		{
			name: "test case 171 - full valid 2",
			s:    "(()(((())))))",
			want: 12,
		},
		{
			name: "test case 171 - full valid 3",
			s:    ")(()(((())))))",
			want: 12,
		},
		{
			name: "test case 184",
			s:    ")(((((()())()()))()(()))(",
			want: 22,
		},
		{
			name: "timelimit exceeded 118",
			s:    ")(()(()(((())(((((()()))((((()()(()()())())())()))()()()())(())()()(((()))))()((()))(((())()((()()())((())))(())))())((()())()()((()((())))))((()(((((()((()))(()()(())))((()))()))())",
			want: 132,
		},
		{
			name: "`testcase 151",
			s:    "(()))())(",
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
