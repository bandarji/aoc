package adventofcode

import (
	"testing"
)

const y16d24TestInput string = `###########
#0.1.....2#
#.#######.#
#4.......3#
###########`

func Test_y16d24(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test p1", args{input: y16d24TestInput, part: 1}, 14},
		{"test p2", args{input: y16d24TestInput, part: 2}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d24(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y16d24() = %v, want %v", got, tt.want)
			}
		})
	}
}
