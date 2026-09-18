package adventofcode

import "testing"

func Test_y17d14(t *testing.T) {
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
				input: "flqrgnkx",
				part:  1,
			},
			wantAnswer: 8108,
		},
		{
			name: "example 2",
			args: args{
				input: "flqrgnkx",
				part:  2,
			},
			wantAnswer: 1242,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y17d14(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y17d14() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
