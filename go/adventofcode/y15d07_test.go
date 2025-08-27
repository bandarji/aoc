package adventofcode

import "testing"

const y15d07TestInput string = `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`

func Test_y15d07(t *testing.T) {
	type args struct {
		input string
		wire  string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"y15d07 part 1 'd'", args{y15d07TestInput, "d", 1}, 72},
		{"y15d07 part 1 'e'", args{y15d07TestInput, "e", 1}, 507},
		{"y15d07 part 1 'f'", args{y15d07TestInput, "f", 1}, 492},
		{"y15d07 part 1 'g'", args{y15d07TestInput, "g", 1}, 114},
		{"y15d07 part 1 'h'", args{y15d07TestInput, "h", 1}, 65412},
		{"y15d07 part 1 'i'", args{y15d07TestInput, "i", 1}, 65079},
		{"y15d07 part 1 'x'", args{y15d07TestInput, "x", 1}, 123},
		{"y15d07 part 1 'y'", args{y15d07TestInput, "y", 1}, 456},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y15d07(tt.args.input, tt.args.wire, tt.args.part); got != tt.want {
				t.Errorf("y15d07() = %v, want %v", got, tt.want)
			}
		})
	}
}
