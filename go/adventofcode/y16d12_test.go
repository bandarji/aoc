package adventofcode

import "testing"

const y16d12TestInput string = `cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a`

const y16d12TestInput2 string = `cpy 1 a
cpy 1 b
cpy 26 d
jnz c 2
jnz 1 5
cpy 7 c
inc d
dec c
jnz c -2
cpy a c
inc a
dec b
jnz b -2
cpy c b
dec d
jnz d -6
cpy 18 c
cpy 11 d
inc a
dec d
jnz d -2
dec c
jnz c -5`

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
		{"part 1", args{input: y16d12TestInput2, part: 1}, 318009},
		{"part 2", args{input: y16d12TestInput2, part: 2}, 9227663},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y16d12(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y16d12() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
