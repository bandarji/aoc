package adventofcode

import "testing"

func Test_y16d02(t *testing.T) {
	const y16d02TestInput string = "ULL\nRRDDD\nLURDL\nUUUUD"
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"part 1", args{input: y16d02TestInput, part: 1}, "1985"},
		{"part 2", args{input: y16d02TestInput, part: 2}, "5DB3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d02(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y16d02() = %v, want %v", got, tt.want)
			}
		})
	}
}
