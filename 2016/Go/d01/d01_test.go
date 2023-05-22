package d01

import (
	"testing"
)

func A(t *testing.T) {}

func TestWalkCity(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"01/01/01", args{"R2, L3", 1}, 5},
		{"01/01/02", args{"R2, R2, R2", 1}, 2},
		{"01/01/03", args{"R5, L5, R5, R3", 1}, 12},
		{"01/02/01", args{"R8, R4, R4, R8", 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WalkCity(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("WalkCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
