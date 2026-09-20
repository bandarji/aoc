package adventofcode

import "testing"

const y17d20TestInput1 string = `p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>
p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>`

const y17d20TestInput2 string = `p=<-6,0,0>, v=<3,0,0>, a=<0,0,0>
p=<-4,0,0>, v=<2,0,0>, a=<0,0,0>
p=<-2,0,0>, v=<1,0,0>, a=<0,0,0>
p=<3,0,0>, v=<-1,0,0>, a=<0,0,0>`

func Test_y17d20(t *testing.T) {
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
			name: "test 1",
			args: args{
				input: y17d20TestInput1,
				part:  1,
			},
			wantAnswer: 0,
		},
		{
			name: "test 2",
			args: args{
				input: y17d20TestInput2,
				part:  2,
			},
			wantAnswer: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y17d20(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y17d20() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
