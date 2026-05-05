package adventofcode

import "testing"

func Test_y17d01(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test '1122' (p1)", args: args{input: "1122", part: 1}, want: 3},
		{name: "test '1111' (p1)", args: args{input: "1111", part: 1}, want: 4},
		{name: "test '1234' (p1)", args: args{input: "1234", part: 1}, want: 0},
		{name: "test '91212129' (p1)", args: args{input: "91212129", part: 1}, want: 9},
		{name: "test '1212' (p2)", args: args{input: "1212", part: 2}, want: 6},
		{name: "test '1221' (p2)", args: args{input: "1221", part: 2}, want: 0},
		{name: "test '123425' (p2)", args: args{input: "123425", part: 2}, want: 4},
		{name: "test '123123' (p2)", args: args{input: "123123", part: 2}, want: 12},
		{name: "test '12131415' (p2)", args: args{input: "12131415", part: 2}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y17d01(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y17d01() = %v, want %v", got, tt.want)
			}
		})
	}
}
