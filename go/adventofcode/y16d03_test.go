package adventofcode

import "testing"

const y16d03TestInput string = "5 10 25\n10 20 30\n15 25 35"

func Test_y16d03(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{"test part 1", args{input: y16d03TestInput, part: 1}, 1},
		{"test part 2", args{input: y16d03TestInput, part: 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y16d03(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y16d03() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
