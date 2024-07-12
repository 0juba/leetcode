package product_of_array_except_self

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed test_case_19.txt
var testCase19In string

//go:embed test_case_19_expected.txt
var testCase19Expected string

func getTestCase19In() []int {
	var result []int
	err := json.Unmarshal([]byte(testCase19In), &result)
	if err != nil {
		panic(err)
	}

	return result
}

func getTestCase19Expected() []int {
	var result []int
	err := json.Unmarshal([]byte(testCase19Expected), &result)
	if err != nil {
		panic(err)
	}

	return result
}

func Test_productExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: `example 1`,
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: `example 2`,
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
		{
			name: `test case 19`,
			nums: getTestCase19In(),
			want: getTestCase19Expected(),
		},
		{
			name: `example 3`,
			nums: []int{1, 2},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := productExceptSelf(tt.nums)

			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
