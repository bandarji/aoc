package adventofcode

import "testing"

const y17d08TestInput string = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`

func Test_y17d08(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test p1", args{input: y17d08TestInput, part: 1}, 1},
		{"test p2", args{input: y17d08TestInput, part: 2}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y17d08(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y17d08() = %v, want %v", got, tt.want)
			}
		})
	}
}
