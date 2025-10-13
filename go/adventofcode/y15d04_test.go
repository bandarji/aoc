package adventofcode

import "testing"

func Test_y15d04(t *testing.T) {
	type args struct {
		input  string
		zeroes int
	}
	tests := []struct {
		name       string
		args       args
		wantNumber int
	}{
		{"test p1", args{"abcdef", 5}, 609_043},
		{"test p2", args{"pqrstuv", 5}, 1_048_970},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNumber := y15d04(tt.args.input, tt.args.zeroes); gotNumber != tt.wantNumber {
				t.Errorf("y15d04() = %v, want %v", gotNumber, tt.wantNumber)
			}
		})
	}
}
