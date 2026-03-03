package adventofcode

import "testing"

func Test_y16d19(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name          string
		args          args
		wantElfNumber int
	}{
		{"test p1", args{input: "5", part: 1}, 3},
		{"test p2", args{input: "5", part: 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotElfNumber := y16d19(tt.args.input, tt.args.part); gotElfNumber != tt.wantElfNumber {
				t.Errorf("y16d19() = %v, want %v", gotElfNumber, tt.wantElfNumber)
			}
		})
	}
}
