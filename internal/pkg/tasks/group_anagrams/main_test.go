package group_anagrams

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

//go:embed leetcode_111_test_case.txt
var leetcode111CaseRow string

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: `empty`,
		},
		{
			name: `leetcode example #1`,
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"},
			},
		},
		{
			name: `leetcode example #2`,
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			name: `leetcode example #3`,
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
		{
			name: `leetcode case 98`,
			strs: []string{"ddddddddddg", "dgggggggggg"},
			want: [][]string{
				{"ddddddddddg"}, {"dgggggggggg"},
			},
		},
		{
			name: `leetcode case 111`,
			strs: func() []string {
				strs := []string{}
				for _, s := range strings.Split(leetcode111CaseRow, ",") {
					strs = append(strs, strings.Trim(s, "\" "))
				}

				return strs
			}(),
			want: [][]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t,
				helperJoinSubSlices(tt.want),
				helperJoinSubSlices(groupAnagrams(tt.strs)),
			)
		})
	}
}

func helperJoinSubSlices(strs [][]string) []string {
	result := []string{}
	for _, items := range strs {
		sort.Strings(items)
		result = append(result, strings.Join(items, ","))
	}

	sort.Strings(result)

	return result
}
