package adventofcode

import "testing"

const y15d23TestInput string = "inc a\njio a, +2\ntpl a\ninc a"

func Test_y15d23(t *testing.T) {
	type args struct {
		input    string
		register string
		part     int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{"example", args{y15d23TestInput, "a", 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y15d23(tt.args.input, tt.args.register, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y15d23() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
