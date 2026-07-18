package adventofcode

import "testing"

const y17d19TestInput string = `     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ 

`

func Test_y17d19(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer string
	}{
		{name: "part 1", args: args{input: y17d19TestInput, part: 1}, wantAnswer: "ABCDEF"},
		{name: "part 2", args: args{input: y17d19TestInput, part: 2}, wantAnswer: "38"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y17d19(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y17d19() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
