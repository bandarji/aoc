package adventofcode

import "testing"

const y17d12TestInput string = `0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`

func Test_y17d12(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{"part 1", args{input: y17d12TestInput, part: 1}, 6},
		{"part 2", args{input: y17d12TestInput, part: 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y17d12(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y17d12() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
