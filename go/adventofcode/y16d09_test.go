package adventofcode

import "testing"

func Test_y16d09(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name                   string
		args                   args
		wantDecompressedLength int
	}{
		{"y16d09p1 'ADVENT'", args{input: "ADVENT", part: 1}, 6},
		{"y16d09p1 'A(1x5)BC'", args{input: "A(1x5)BC", part: 1}, 7},
		{"y16d09p1 '(3x3)XYZ'", args{input: "(3x3)XYZ", part: 1}, 9},
		{"y16d09p1 'A(2x2)BCD(2x2)EFG'", args{input: "A(2x2)BCD(2x2)EFG", part: 1}, 11},
		{"y16d09p1 '(6x1)(1x3)A'", args{input: "(6x1)(1x3)A", part: 1}, 6},
		{"y16d09p1 'X(8x2)(3x3)ABCY'", args{input: "X(8x2)(3x3)ABCY", part: 1}, 18},
		{"y16d09p1 'A(2x2)BCD(2x2)EFG'", args{input: "A(2x2)BCD(2x2)EFG", part: 1}, 11},
		{"y16d09p1 '(6x1)(1x3)A'", args{input: "(6x1)(1x3)A", part: 1}, 6},
		{"y16d09p1 'X(8x2)(3x3)ABCY'", args{input: "X(8x2)(3x3)ABCY", part: 1}, 18},

		{"y16d09p2 '(3x3)XYZ'", args{input: "(3x3)XYZ", part: 2}, len("XYZXYZXYZ")},
		{"y16d09p2 'X(8x2)(3x3)ABCY'", args{input: "X(8x2)(3x3)ABCY", part: 2}, len("XABCABCABCABCABCABCY")},
		{"y16d09p1 '(27x12)(20x12)(13x14)(7x10)(1x12)A'", args{input: "(27x12)(20x12)(13x14)(7x10)(1x12)A", part: 2}, 241920},
		{"y16d09p1 '(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN'", args{input: "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", part: 2}, 445},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDecompressedLength := y16d09(tt.args.input, tt.args.part); gotDecompressedLength != tt.wantDecompressedLength {
				t.Errorf("y16d09() = %v, want %v", gotDecompressedLength, tt.wantDecompressedLength)
			}
		})
	}
}
