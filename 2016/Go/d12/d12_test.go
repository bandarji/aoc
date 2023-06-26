package d12

import "testing"

const TestInstructions string = `cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a`

func TestDay(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{TestInstructions, 1}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("Day() = %v, want %v", got, tt.want)
			}
		})
	}
}
