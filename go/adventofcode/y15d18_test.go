package adventofcode

import "testing"

const (
	y15d18TestInput1 string = ".#.#.#\n...##.\n#....#\n..#...\n#.#..#\n####.."
	y15d18TestInput2 string = "##.#.#\n...##.\n#....#\n..#...\n#.#..#\n####.#"
)

func Test_y15d18(t *testing.T) {
	type args struct {
		input string
		steps int
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{"test p1", args{input: y15d18TestInput1, steps: 4, part: 1}, 4},
		{"test p2", args{input: y15d18TestInput2, steps: 5, part: 2}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y15d18(tt.args.input, tt.args.steps, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y15d18() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
