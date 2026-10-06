package adventofcode

import "testing"

const y17d21TestInput string = `../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#`

func Test_y17d21(t *testing.T) {
	type args struct {
		input      string
		iterations int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{
			name: "Part 1",
			args: args{
				input:      y17d21TestInput,
				iterations: 2,
			},
			wantAnswer: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y17d21(tt.args.input, tt.args.iterations); gotAnswer != tt.wantAnswer {
				t.Errorf("y17d21() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
