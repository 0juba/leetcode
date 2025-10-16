package main

import "testing"

func Test_minMutation(t *testing.T) {
	tests := []struct {
		name      string
		startGene string
		endGene   string
		bank      []string
		want      int
	}{
		{
			name: "empty",
			want: -1,
		},
		{
			name:      "example #1",
			startGene: "AACCGGTT",
			endGene:   "AACCGGTA",
			bank:      []string{"AACCGGTA"},
			want:      1,
		},
		{
			name:      "example #2",
			startGene: "AACCGGTT",
			endGene:   "AAACGGTA",
			bank:      []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			want:      2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMutation(tt.startGene, tt.endGene, tt.bank); got != tt.want {
				t.Errorf("minMutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
