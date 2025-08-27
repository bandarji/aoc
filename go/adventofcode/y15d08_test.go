package adventofcode

import "testing"

var y25d08TestInput = string([]byte{34, 34, 10, 34, 97, 98, 99, 34, 10, 34, 97, 97, 97, 92, 34, 97, 97, 97, 34, 10, 34, 92, 120, 50, 55, 34})

func Test_y15d08(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"y15d08p1", args{y25d08TestInput, 1}, 12},
		{"y15d08p2", args{y25d08TestInput, 2}, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y15d08(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y15d08() = %v, want %v", got, tt.want)
			}
		})
	}
}
