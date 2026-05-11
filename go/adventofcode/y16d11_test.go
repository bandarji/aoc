package adventofcode

import "testing"

const y16d11TestInput string = `The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
The second floor contains a hydrogen generator.
The third floor contains a lithium generator.
The fourth floor contains nothing relevant.`

func Test_y16d11(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{input: y16d11TestInput, part: 1}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d11(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y16d11() = %v, want %v", got, tt.want)
			}
		})
	}
}
