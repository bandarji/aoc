package adventofcode

import "testing"

func Test_y16d05(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name     string
		args     args
		wantCode string
	}{
		{name: "test 1", args: args{input: "abc", part: 1}, wantCode: "18f47a30"},
		{name: "test 2", args: args{input: "abc", part: 2}, wantCode: "05ace8e3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCode := y16d05(tt.args.input, tt.args.part); gotCode != tt.wantCode {
				t.Errorf("y16d05() = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}
