package adventofcode

import (
	"testing"
)

func Test_y16d16ProcessStep(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{a: "1"}, want: "100"},
		{name: "0", args: args{a: "0"}, want: "001"},
		{name: "11111", args: args{a: "11111"}, want: "11111000000"},
		{name: "111100001010", args: args{a: "111100001010"}, want: "1111000010100101011110000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d16ProcessStep(tt.args.a); got != tt.want {
				t.Errorf("y16d16ProcessStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_y16d16Checksum(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "110010110100", args: args{input: "110010110100"}, want: "100"},
		{name: "10000011110010000111", args: args{input: "10000011110010000111"}, want: "01100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d16Checksum(tt.args.input); got != tt.want {
				t.Errorf("y16d16Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_y16d16(t *testing.T) {
	type args struct {
		input      string
		diskLength int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer string
	}{
		{"test p1", args{input: y16d16Input, diskLength: y16d16DiskLengthPart1}, "11111000111110000"},
		{"test p2", args{input: y16d16Input, diskLength: y16d16DiskLengthPart2}, "10111100110110100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y16d16(tt.args.input, tt.args.diskLength); gotAnswer != tt.wantAnswer {
				t.Errorf("y16d16() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
