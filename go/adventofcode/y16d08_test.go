package adventofcode

import "testing"

const y16d08TestInput string = `rect 3x2
rotate column x=1 by 1
rotate row y=0 by 4
rotate column x=1 by 1`

const y16d08TestOutputPart2 string = "\n.#..#.#\n#.#....\n.#.....\n\n"

const y16d08TestTall = 3
const y16d08TestWide = 7

func Test_y16d08(t *testing.T) {
	type args struct {
		input string
		part  int
		tall  int
		wide  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer string
	}{
		{"test p1", args{input: y16d08TestInput, part: 1, tall: y16d08TestTall, wide: y16d08TestWide}, "6"},
		{"test p2", args{input: y16d08TestInput, part: 2, tall: y16d08TestTall, wide: y16d08TestWide}, y16d08TestOutputPart2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y16d08(tt.args.input, tt.args.part, tt.args.tall, tt.args.wide); gotAnswer != tt.wantAnswer {
				t.Errorf("y16d08() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
