package adventofcode

import "testing"

const y17d16Input = `s1,x3/4,pe/b`

func Test_y17d16(t *testing.T) {
	type args struct {
		input        string
		programCount int
		part         int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer string
	}{
		{name: "example", args: args{input: y17d16Input, programCount: 5, part: 1}, wantAnswer: "baedc"},
		{name: "example", args: args{input: y17d16Input, programCount: 5, part: 2}, wantAnswer: "abcde"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y17d16(tt.args.input, tt.args.programCount, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y17d16() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
