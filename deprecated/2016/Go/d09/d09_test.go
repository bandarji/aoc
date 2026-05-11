package d09

import "testing"

func TestDay(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ADVENT", args{"ADVENT", 1}, 6},
		{"A(1x5)BC", args{"A(1x5)BC", 1}, 7},
		{"(3x3)XYZ", args{"XYZXYZXYZ", 1}, 9},
		{"A(2x2)BCD(2x2)EFG", args{"A(2x2)BCD(2x2)EFG", 1}, 11},
		{"(6x1)(1x3)A", args{"(6x1)(1x3)A", 1}, 6},
		{"X(8x2)(3x3)ABCY", args{"X(8x2)(3x3)ABCY", 1}, 18},
		{"(3x3)XYZ", args{"(3x3)XYZ", 2}, 9},
		{"(27x12)(20x12)(13x14)(7x10)(1x12)A", args{"(27x12)(20x12)(13x14)(7x10)(1x12)A", 2}, 241920},
		{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", args{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", 2}, 445},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("Day() = %v, want %v", got, tt.want)
			}
		})
	}
}
