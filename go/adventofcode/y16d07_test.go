package adventofcode

import "testing"

const y16d07TestInput1 string = `abba[mnop]qrst
abcd[bddb]xyyx
aaaa[qwer]tyui
ioxxoj[asdfgh]zxcvbn`

const y16d07TestInput2 string = `aba[bab]xyz
xyx[xyx]xyx
aaa[kek]eke
zazbz[bzb]cdb`

func Test_y16d07(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{"test part 1", args{input: y16d07TestInput1, part: 1}, 2},
		{"test part 2", args{input: y16d07TestInput2, part: 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y16d07(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y16d07() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
