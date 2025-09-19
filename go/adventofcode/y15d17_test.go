package adventofcode

import "testing"

const y15d17TestInput = "20\n15\n10\n5\n5"

func Test_y15d17(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		liters int
		part   int
		want   int
	}{
		{"test p1", y15d17TestInput, 25, 1, 4},
		{"test p2", y15d17TestInput, 25, 2, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := y15d17(tt.input, tt.liters, tt.part)
			if got != tt.want {
				t.Errorf("y15d17() = %v, want %v", got, tt.want)
			}
		})
	}
}
