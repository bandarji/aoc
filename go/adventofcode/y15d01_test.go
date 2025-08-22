package adventofcode

import (
	"testing"
)

func Test_y15d01p1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name      string
		args      args
		wantFloor int
	}{
		{"test '(())'", args{"(())"}, 0},
		{"test '()()'", args{"()()"}, 0},
		{"test '((('", args{"((("}, 3},
		{"test '(()(()('", args{"(()(()("}, 3},
		{"test '))((((('", args{"))((((("}, 3},
		{"test '())'", args{"())"}, -1},
		{"test '))('", args{"))("}, -1},
		{"test ')))'", args{")))"}, -3},
		{"test ')())())'", args{")())())"}, -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFloor := y15d01p1(tt.args.input); gotFloor != tt.wantFloor {
				t.Errorf("y15d01p1() = %v, want %v", gotFloor, tt.wantFloor)
			}
		})
	}
}

func Test_y15d01p2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test ')'", args{")"}, 1},
		{"test '()())'", args{"()())"}, 5},
		{"test '((('", args{"((("}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y15d01p2(tt.args.input); got != tt.want {
				t.Errorf("y15d01p2() = %v, want %v", got, tt.want)
			}
		})
	}
}
