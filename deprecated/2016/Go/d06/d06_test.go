package d06

import "testing"

const TESTINPUT string = `eedadn
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

func TestDay(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name        string
		args        args
		wantMessage string
	}{
		{"Day 06 01 00", args{TESTINPUT, 1}, "easter"},
		{"Day 06 02 00", args{TESTINPUT, 2}, "advent"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMessage := Day(tt.args.input, tt.args.part); gotMessage != tt.wantMessage {
				t.Errorf("Day() = %v, want %v", gotMessage, tt.wantMessage)
			}
		})
	}
}
