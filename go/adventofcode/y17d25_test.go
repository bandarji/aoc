package adventofcode

import "testing"

const y17d25TestInput string = `Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.`

func Test_y17d25(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name         string
		args         args
		wantChecksum int
	}{
		{name: "test", args: args{input: y17d25TestInput}, wantChecksum: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotChecksum := y17d25(tt.args.input); gotChecksum != tt.wantChecksum {
				t.Errorf("y17d25() = %v, want %v", gotChecksum, tt.wantChecksum)
			}
		})
	}
}
