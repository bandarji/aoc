package adventofcode

import "testing"

const y16d10TestInput string = `value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2`

func Test_y16d10(t *testing.T) {
	type args struct {
		input  string
		part   int
		compLo int
		compHi int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{"test p1", args{input: y16d10TestInput, part: 1, compLo: 2, compHi: 5}, 2},
		{"test p2", args{input: y16d10TestInput, part: 2, compLo: 2, compHi: 5}, 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y16d10(tt.args.input, tt.args.part, tt.args.compLo, tt.args.compHi); gotAnswer != tt.wantAnswer {
				t.Errorf("y16d10() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
