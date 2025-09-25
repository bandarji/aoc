package adventofcode

import "testing"

func Test_y15d19(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{"test p1", args{input: "H => HO\nH => OH\nO => HH\n\nHOH", part: 1}, 4},
		{"test p2", args{input: "e => H\ne => O\nH => HO\nH => OH\nO => HH\n\nHOH", part: 2}, 3},
		{"test p2", args{input: "e => H\ne => O\nH => HO\nH => OH\nO => HH\n\nHOHOHO", part: 2}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y15d19(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y15d19() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
