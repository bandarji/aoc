package adventofcode

import "testing"

const y16d15TestInput string = `Disc #1 has 5 positions; at time=0, it is at position 4.
Disc #2 has 2 positions; at time=0, it is at position 1.`

func Test_y16d15(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{"test p1", args{input: y16d15TestInput, part: 1}, 5},
		{"test p2", args{input: y16d15TestInput, part: 2}, 85},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y16d15(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y16d15() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
