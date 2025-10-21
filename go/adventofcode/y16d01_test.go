package adventofcode

import "testing"

func Test_y16d01(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Part 1 - R2, L3", args: args{input: "R2, L3", part: 1}, want: 5},
		{name: "Part 1 - R2, R2, R2", args: args{input: "R2, R2, R2", part: 1}, want: 2},
		{name: "Part 1 - R5, L5, R5, R3", args: args{input: "R5, L5, R5, R3", part: 1}, want: 12},
		{name: "Part 2 - R8, R4, R4, R8", args: args{input: "R8, R4, R4, R8", part: 2}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d01(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y16d01() = %v, want %v", got, tt.want)
			}
		})
	}
}
