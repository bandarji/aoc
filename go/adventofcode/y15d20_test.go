package adventofcode

import "testing"

func Test_y15d20(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test p1 - 1 present", args{input: "10", part: 1}, 1},
		{"test p1 - 70 presents", args{input: "70", part: 1}, 4},
		{"test p2 - 34_000_000 presents", args{input: y15d20Input, part: 2}, 831_600},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y15d20(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y15d20() = %v, want %v", got, tt.want)
			}
		})
	}
}
