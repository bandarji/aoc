package adventofcode

import "testing"

func Test_y15d12(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test p1 '[1,2,3]'", args{input: "[1,2,3]", part: 1}, 6},
		{"test p1 '[[[3]]]'", args{input: "[[[3]]]", part: 1}, 3},
		{"test p1 '{\"a\":[-1,1]}'", args{input: "{\"a\":[-1,1]}", part: 1}, 0},
		{"test p1 '[]'", args{input: "[]", part: 1}, 0},
		{"test p1 '{}'", args{input: "{}", part: 1}, 0},
		{"test p2 '[1,2,3]'", args{input: "[1,2,3]", part: 2}, 6},
		{"test p2 '[1,{\"c\":\"red\",\"b\":2},3]'", args{input: "[1,{\"c\":\"red\",\"b\":2},3]", part: 2}, 4},
		{"test p2 '{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}'", args{input: "{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}", part: 2}, 0},
		{"test p2 '[1,\"red\",5]'", args{input: "[1,\"red\",5]", part: 2}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y15d12(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y15d12() = %v, want %v", got, tt.want)
			}
		})
	}
}
