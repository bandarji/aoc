package adventofcode

import (
	"testing"
)

func Test_y16d13(t *testing.T) {
	type args struct {
		input        string
		destinationX int
		destinationY int
		part         int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{"test p1", args{input: "1352", destinationX: 4, destinationY: 7, part: 1}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y16d13(tt.args.input, tt.args.destinationX, tt.args.destinationY, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y16d13() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
