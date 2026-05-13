package adventofcode

import "testing"

func xxxTest_y17d03(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test square 1", args{input: "1", part: 1}, 0},
		{"test square 12", args{input: "12", part: 1}, 3},
		{"test square 23", args{input: "23", part: 1}, 2},
		{"test square 1024", args{input: "1024", part: 1}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y17d03(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y17d03() = %v, want %v", got, tt.want)
			}
		})
	}
}
