package adventofcode

import "testing"

func Test_y15d10(t *testing.T) {
	type args struct {
		input  string
		cycles int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test '1' 1 cycle", args{input: "1", cycles: 1}, 2},
		{"test '11' 1 cycle", args{input: "11", cycles: 1}, 2},
		{"test '21' 1 cycle", args{input: "21", cycles: 1}, 4},
		{"test '1211' 1 cycle", args{input: "1211", cycles: 1}, 6},
		{"test '111221' 1 cycle", args{input: "111221", cycles: 1}, 6},
		{"test '1' 5 cycles", args{input: "1", cycles: 5}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y15d10(tt.args.input, tt.args.cycles); got != tt.want {
				t.Errorf("y15d10() = %v, want %v", got, tt.want)
			}
		})
	}
}
