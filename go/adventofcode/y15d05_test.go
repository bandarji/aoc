package adventofcode

import "testing"

func Test_y15d05(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name     string
		args     args
		wantNice int
	}{
		{"nice: 'ugknbfddgicrmopn'", args{"ugknbfddgicrmopn", 1}, 1},
		{"nice: 'aaa'", args{"aaa", 1}, 1},
		{"naughty: 'jchzalrnumimnmhp'", args{"jchzalrnumimnmhp", 1}, 0},
		{"naughty: 'haegwjzuvuyypxyu'", args{"haegwjzuvuyypxyu", 1}, 0},
		{"naughty: 'dvszwmarrgswjxmb'", args{"dvszwmarrgswjxmb", 1}, 0},
		{"nice: 'qjhvhtzxzqqjkmpb'", args{"qjhvhtzxzqqjkmpb", 2}, 1},
		{"nice: 'xxyxx'", args{"xxyxx", 2}, 1},
		{"naughty: 'uurcxstgmygtbstg'", args{"uurcxstgmygtbstg", 2}, 0},
		{"naughty: 'ieodomkazucvgmuy'", args{"ieodomkazucvgmuy", 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNice := y15d05(tt.args.input, tt.args.part); gotNice != tt.wantNice {
				t.Errorf("y15d05() = %v, want %v", gotNice, tt.wantNice)
			}
		})
	}
}
