package adventofcode

import (
	"testing"
)

const y17d04Tests string = `aa bb cc dd ee
aa bb cc dd aa
aa bb cc dd aaa`

func Test_y17d04(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{"p1", args{input: y17d04Tests, part: 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := y17d04(tt.args.input, tt.args.part); gotCount != tt.wantCount {
				t.Errorf("y17d04() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_y17d04IsValidPassphrase(t *testing.T) {
	type args struct {
		phrase string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test valid passphrase", args{phrase: "aa bb cc dd ee"}, true},
		{"test invalid passphrase", args{phrase: "aa bb cc dd aa"}, false},
		{"test valid passphrase", args{phrase: "aa bb cc dd aaa"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y17d04IsValidPassphrase(tt.args.phrase); got != tt.want {
				t.Errorf("y17d04IsValidPassphrase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_y17d04IsValidPassphraseCheckAnagrams(t *testing.T) {
	type args struct {
		phrase string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test 'abcde fghij' is valid", args{phrase: "abcde fghij"}, true},
		{"test 'abcde xyz ecdab' is not valid", args{phrase: "abcde xyz ecdab"}, false},
		{"test 'a ab abc abd abf abj' is valid", args{phrase: "a ab abc abd abf abj"}, true},
		{"test 'iiii oiii ooii oooi oooo' is valid", args{phrase: "iiii oiii ooii oooi oooo"}, true},
		{"test 'oiii ioii iioi iiio' is not valid", args{phrase: "oiii ioii iioi iiio"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y17d04IsValidPassphraseCheckAnagrams(tt.args.phrase); got != tt.want {
				t.Errorf("y17d04IsValidPassphraseCheckAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
