package encode_and_decode_strings

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: `empty`,
		},
		{
			name: `one item`,
			strs: []string{"str1"},
			want: "4;str1",
		},
		{
			name: `two items`,
			strs: []string{"str1", "str2"},
			want: "4;str14;str2",
		},
		{
			name: `three items`,
			strs: []string{"str1", `4;str24;`, `4;str4;str4;`},
			want: `4;str18;4;str24;12;4;str4;str4;`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.strs); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want []string
	}{
		{
			name: `empty`,
		},
		{
			name: `one simple item`,
			str:  "4;str1",
			want: []string{"str1"},
		},
		{
			name: `two simple items`,
			str:  "4;str14;str2",
			want: []string{"str1", "str2"},
		},
		{
			name: `with escaping`,
			str:  `4;str15;str206;;str30`,
			want: []string{"str1", "str20", ";str30"},
		},
		{
			name: `empty strings`,
			str:  `0;0;0;0;0;`,
			want: []string{"", "", "", "", ""},
		},
		{
			name: `complex`,
			str:  `4;str18;4;str24;12;4;str4;str4;`,
			want: []string{"str1", `4;str24;`, `4;str4;str4;`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
