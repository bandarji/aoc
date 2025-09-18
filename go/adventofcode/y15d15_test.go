package adventofcode

import "testing"

func Test_y15d15(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"test p1", y15d15Input, 1, 13_882_464},
		{"test p2", y15d15Input, 2, 11_171_160},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := y15d15(tt.input, tt.part)
			if got != tt.want {
				t.Errorf("y15d15() = %v, want %v", got, tt.want)
			}
		})
	}
}
