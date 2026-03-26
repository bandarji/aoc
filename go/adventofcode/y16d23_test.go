package adventofcode

import "testing"

const y16d23TestInput string = `cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a`

func Test_y16d23(t *testing.T) {
	type args struct {
		input     string
		registers map[string]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test p1", args{input: y16d23TestInput, registers: map[string]int{"a": 0, "b": 0, "c": 0, "d": 0}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d23(tt.args.input, tt.args.registers); got != tt.want {
				t.Errorf("y16d23() = %v, want %v", got, tt.want)
			}
		})
	}
}
