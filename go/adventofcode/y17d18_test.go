package adventofcode

import "testing"

const y17d18TestInput string = `set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2`

const y17d18TestInput2 string = `snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d`

func Test_y17d18(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name          string
		args          args
		wantFrequency int
	}{
		{"test p1", args{input: y17d18TestInput, part: 1}, 4},
		{"test p2", args{input: y17d18TestInput2, part: 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFrequency := y17d18(tt.args.input, tt.args.part); gotFrequency != tt.wantFrequency {
				t.Errorf("y17d18() = %v, want %v", gotFrequency, tt.wantFrequency)
			}
		})
	}
}
