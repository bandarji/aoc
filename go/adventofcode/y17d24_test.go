package adventofcode

import "testing"

const y17d24TestInput = `0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`

func Test_y17d24(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{
			name: "part 1",
			args: args{
				input: y17d24TestInput,
				part:  1,
			},
			wantAnswer: 31,
		},
		{
			name: "part 2",
			args: args{
				input: y17d24TestInput,
				part:  2,
			},
			wantAnswer: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y17d24(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y17d24() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
