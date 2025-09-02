package adventofcode

import (
	"testing"
)

func Test_y15d11HasTwoPairs(t *testing.T) {
	type args struct {
		pw string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"pass 'abbceffg'", args{pw: "abbceffg"}, true},
		{"fail 'abbcegjk'", args{pw: "abbcegjk"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y15d11HasTwoPairs(tt.args.pw); got != tt.want {
				t.Errorf("y15d11HasTwoPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_y15d11HasStraight(t *testing.T) {
	type args struct {
		pw string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"pass 'hijklmmn'", args{pw: "hijklmmn"}, true},
		{"fail 'zyxwvuts'", args{pw: "zyxwvuts"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y15d11HasStraight(tt.args.pw); got != tt.want {
				t.Errorf("y15d11HasStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_y15d11ValidChars(t *testing.T) {
	type args struct {
		pw string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"fail 'hijklmmn'", args{pw: "hijklmmn"}, false},
		{"pass 'aaaaaaaa'", args{pw: "aaaaaaaa"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y15d11ValidChars(tt.args.pw); got != tt.want {
				t.Errorf("y15d11ValidChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_y15d11(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"next after 'abcdefgh'", args{input: "abcdefgh", part: 1}, "abcdffaa"},
		{"next after 'ghijklmn'", args{input: "ghijklmn", part: 1}, "ghjaabcc"},
		{"next after 'abcdefgh'", args{input: "abcdefgh", part: 2}, "abcdffbb"},
		{"next after 'ghijklmn'", args{input: "ghijklmn", part: 2}, "ghjbbcdd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y15d11(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y15d11() = %v, want %v", got, tt.want)
			}
		})
	}
}
