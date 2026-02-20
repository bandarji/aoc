package adventofcode

import "testing"

func Test_y16d14(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{"test p1", args{input: "abc", part: 1}, 22728},
		{"test p2", args{input: "abc", part: 2}, 22551},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y16d14(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y16d14() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
