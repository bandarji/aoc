package adventofcode

import "testing"

const y17d07TestInput string = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

func Test_y17d07(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name     string
		args     args
		wantName string
	}{
		{"test part 1", args{y17d07TestInput, 1}, "tknk"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotName := y17d07(tt.args.input, tt.args.part); gotName != tt.wantName {
				t.Errorf("y17d07() = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}
