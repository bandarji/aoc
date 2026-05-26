package adventofcode

import "testing"

func Test_y17d09(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"y17d09p1 '{}'", args{input: "{}", part: 1}, 1},
		{"y17d09p1 '{{{}}}", args{input: "{{{}}}", part: 1}, 6},
		{"y17d09p1 '{{},{}}", args{input: "{{},{}}", part: 1}, 5},
		{"y17d09p1 '{{{},{},{{}}}}", args{input: "{{{},{},{{}}}}", part: 1}, 16},
		{"y17d09p1 '{<a>,<a>,<a>,<a>}", args{input: "{<a>,<a>,<a>,<a>}", part: 1}, 1},
		{"y17d09p1 '{{<ab>},{<ab>},{<ab>},{<ab>}}", args{input: "{{<ab>},{<ab>},{<ab>},{<ab>}}", part: 1}, 9},
		{"y17d09p1 '{{<!!>},{<!!>},{<!!>},{<!!>}}", args{input: "{{<!!>},{<!!>},{<!!>},{<!!>}}", part: 1}, 9},
		{"y17d09p1 '{{<a!>},{<a!>},{<a!>},{<ab>}}", args{input: "{{<a!>},{<a!>},{<a!>},{<ab>}}", part: 1}, 3},
		{"y17d09p1 '<>'", args{input: "<>", part: 2}, 0},
		{"y17d09p1 '<random characters>'", args{input: "<random characters>", part: 2}, 17},
		{"y17d09p1 '<<<<>'", args{input: "<<<<>", part: 2}, 3},
		{"y17d09p1 '<{!>}>'", args{input: "<{!>}>", part: 2}, 2},
		{"y17d09p1 '<!!>'", args{input: "<!!>", part: 2}, 0},
		{"y17d09p1 '<!!!>>'", args{input: "<!!!>>", part: 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y17d09(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y17d09() = %v, want %v", got, tt.want)
			}
		})
	}
}
