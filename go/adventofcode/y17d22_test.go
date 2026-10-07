package adventofcode

import "testing"

const y17d22TestInput string = `..#
#..
...`

func Test_y17d22(t *testing.T) {
	type args struct {
		input  string
		cycles int
		part   int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{name: "Test 1", args: args{input: y17d22TestInput, cycles: 10_000, part: 1}, wantAnswer: 5587},
		{name: "Test 2", args: args{input: y17d22TestInput, cycles: 100, part: 2}, wantAnswer: 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y17d22(tt.args.input, tt.args.cycles, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y17d22() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
