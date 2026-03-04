package adventofcode

import "testing"

func Test_y16d20(t *testing.T) {
	const y16d20TestInput string = "5-8\n0-2\n4-7" // 3 and 9 free
	type args struct {
		input   string
		ceiling int
		part    int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{"test p1", args{input: y16d20TestInput, ceiling: 9, part: 1}, 3},
		{"test p2", args{input: y16d20TestInput, ceiling: 9, part: 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y16d20(tt.args.input, tt.args.ceiling, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y16d20() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
