package adventofcode

import "testing"

const y17d05TestInput = `0
3
0
1
-3`

func Test_y17d05(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{"test p1", args{y17d05TestInput, 1}, 5},
		{"test p2", args{y17d05TestInput, 2}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y17d05(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y17d05() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
