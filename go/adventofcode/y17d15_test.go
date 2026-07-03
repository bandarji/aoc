package adventofcode

import "testing"

const y17d15TestInput string = `Generator A starts with 65
Generator B starts with 8921`

func Test_y17d15(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name      string
		args      args
		wantTotal int
	}{
		{name: "Part 1", args: args{input: y17d15TestInput, part: 1}, wantTotal: 588},
		{name: "Part 2", args: args{input: y17d15TestInput, part: 2}, wantTotal: 309},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := y17d15(tt.args.input, tt.args.part); gotTotal != tt.wantTotal {
				t.Errorf("y17d15() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
