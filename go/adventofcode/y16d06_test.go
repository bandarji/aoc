package adventofcode

import (
	"testing"
)

const y16d06TestInput string = `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`

func Test_y16d06MostAndLeastCommon(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		wantMost  string
		wantLeast string
	}{
		{"test 1", args{"eedadn"}, "e", "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMost, gotLeast := y16d06MostAndLeastCommon(tt.args.s)
			if gotMost != tt.wantMost {
				t.Errorf("y16d06MostAndLeastCommon() gotMost = %v, want %v", gotMost, tt.wantMost)
			}
			if gotLeast != tt.wantLeast {
				t.Errorf("y16d06MostAndLeastCommon() gotLeast = %v, want %v", gotLeast, tt.wantLeast)
			}
		})
	}
}

func Test_y16d06AssembleStripes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 1", args{"abc\ndef\nghi"}, "adg\nbeh\ncfi"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d06AssembleStripes(tt.args.s); got != tt.want {
				t.Errorf("y16d06AssembleStripes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_y16d06(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer string
	}{
		{"test 1", args{input: y16d06TestInput, part: 1}, "easter"},
		{"test 2", args{input: y16d06TestInput, part: 2}, "advent"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y16d06(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y16d06() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
