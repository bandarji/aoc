package d04

import (
	"testing"
)

func TestRealRoom(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"day 04 01 01", args{"aaaaa-bbb-z-y-x-123[abxyz]"}, true},
		{"day 04 01 02", args{"a-b-c-d-e-f-g-h-987[abcde]"}, true},
		{"day 04 01 03", args{"not-a-real-room-404[oarel]"}, true},
		{"day 04 01 04", args{"totally-real-room-200[decoy]"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RealRoom(tt.args.s); got != tt.want {
				t.Errorf("RealRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{"Day 04 01", args{"aaaaa-bbb-z-y-x-123[abxyz]\na-b-c-d-e-f-g-h-987[abcde]\nnot-a-real-room-404[oarel]\ntotally-real-room-200[decoy]", 1}, 1514},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := Day(tt.args.input, tt.args.part); gotSum != tt.wantSum {
				t.Errorf("Day() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestCaesar(t *testing.T) {
	type args struct {
		s string
		r int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"caesar cipher 343", args{"qzmtzixmtkozyivhz", 343}, "veryencryptedname"},
		{"caesar cipher 0", args{"qzmtzixmtkozyivhz", 0}, "qzmtzixmtkozyivhz"},
		{"abcdef 1", args{"abcdef", 1}, "bcdefg"},
		{"xyzabc 2", args{"xyzabc", 2}, "zabcde"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Caesar(tt.args.s, tt.args.r); got != tt.want {
				t.Errorf("Caesar() = %v, want %v", got, tt.want)
			}
		})
	}
}
