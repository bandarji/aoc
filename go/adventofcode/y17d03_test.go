package adventofcode

import "testing"

func Test_y17d03(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{input: "325489", part: 1}, 552},
		{"test 2", args{input: "325489", part: 2}, 330785},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y17d03(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y17d03() = %v, want %v", got, tt.want)
			}
		})
	}
}
