package adventofcode

import "testing"

const y17d06TestInput string = `0 2 7 0`

func Test_y17d06(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantCycles int
	}{
		{"test 1", args{input: y17d06TestInput, part: 1}, 5},
		{"test 2", args{input: y17d06TestInput, part: 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCycles := y17d06(tt.args.input, tt.args.part); gotCycles != tt.wantCycles {
				t.Errorf("y17d06() = %v, want %v", gotCycles, tt.wantCycles)
			}
		})
	}
}
