package adventofcode

import "testing"

func Test_y17d13(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{
			name: "example 1",
			args: args{
				input: "0: 3\n1: 2\n4: 4\n6: 4",
				part:  1,
			},
			wantAnswer: 24,
		},
		{
			name: "example 2",
			args: args{
				input: "0: 3\n1: 2\n4: 4\n6: 4",
				part:  2,
			},
			wantAnswer: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y17d13(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y17d13() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
