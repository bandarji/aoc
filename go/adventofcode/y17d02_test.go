package adventofcode

import (
	"testing"
)

const y17d02TestInput1 string = "5 1 9 5\n7 5 3\n2 4 6 8"
const y17d02TestInput2 string = "5 9 2 8\n9 4 7 3\n3 8 6 5"

func Test_y17d02ProcessRow1(t *testing.T) {
	type args struct {
		row string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test '5 1 9 5' (p1)", args: args{row: "5 1 9 5"}, want: 8},
		{name: "test '7 5 3' (p1)", args: args{row: "7 5 3"}, want: 4},
		{name: "test '2 4 6 8' (p1)", args: args{row: "2 4 6 8"}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y17d02ProcessRow1(tt.args.row); got != tt.want {
				t.Errorf("y17d02ProcessRow1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_y17d02(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name         string
		args         args
		wantChecksum int
	}{
		{name: "test p1", args: args{input: y17d02TestInput1, part: 1}, wantChecksum: 18},
		{name: "test p2", args: args{input: y17d02TestInput2, part: 2}, wantChecksum: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotChecksum := y17d02(tt.args.input, tt.args.part); gotChecksum != tt.wantChecksum {
				t.Errorf("y17d02() = %v, want %v", gotChecksum, tt.wantChecksum)
			}
		})
	}
}

func Test_y17d02ProcessRow2(t *testing.T) {
	type args struct {
		row string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test '5 9 2 8' (p2)", args: args{row: "5 9 2 8"}, want: 4},
		{name: "test '9 4 7 3' (p2)", args: args{row: "9 4 7 3"}, want: 3},
		{name: "test '3 8 6 5' (p2)", args: args{row: "3 8 6 5"}, want: 2},
		{name: "test ''", args: args{row: ""}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y17d02ProcessRow2(tt.args.row); got != tt.want {
				t.Errorf("y17d02ProcessRow2() = %v, want %v", got, tt.want)
			}
		})
	}
}
