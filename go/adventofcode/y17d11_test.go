package adventofcode

import "testing"

func Test_y17d11(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{input: "ne,ne,ne", part: 1}, want: 3},
		{name: "test", args: args{input: "ne,ne,sw,sw", part: 1}, want: 0},
		{name: "test", args: args{input: "ne,ne,s,s", part: 1}, want: 2},
		{name: "test", args: args{input: "se,sw,se,sw,sw", part: 1}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y17d11(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y17d11() = %v, want %v", got, tt.want)
			}
		})
	}
}
