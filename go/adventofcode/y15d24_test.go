package adventofcode

import "testing"

const y15d24TestInput string = "1\n2\n3\n4\n5\n7\n8\n9\n10\n11"

func Test_y15d24(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{"example", args{y15d24TestInput, 1}, 99},
		{"example", args{y15d24TestInput, 2}, 44},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y15d24(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y15d24() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
