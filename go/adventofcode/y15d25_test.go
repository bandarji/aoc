package adventofcode

import "testing"

const y15d25TestInput1 string = "To continue, please consult the code grid in the manual.  Enter the code at row 2947, column 3029."

func Test_y15d25(t *testing.T) {
	type args struct {
		input string
		code  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test p1", args{y15d25TestInput1, 20151125}, 19_980_801},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y15d25(tt.args.input, tt.args.code); got != tt.want {
				t.Errorf("y15d25() = %v, want %v", got, tt.want)
			}
		})
	}
}
