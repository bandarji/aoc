package adventofcode

import "testing"

const y16d12TestInput string = `cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a`

func Test_y16d12(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{"part 1", args{input: y16d12TestInput, part: 1}, 42},
		{"part 2", args{input: y16d12TestInput, part: 2}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y16d12(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y16d12() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
