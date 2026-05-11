package d07

import (
	"reflect"
	"testing"
)

const (
	TESTINPUT1 string = "abba[mnop]qrst\nabcd[bddb]xyyx\naaaa[qwer]tyui\nioxxoj[asdfgh]zxcvbn"
	TESTINPUT2 string = "aba[bab]xyz\nxyx[xyx]xyx\naaa[kek]eke\nzazbz[bzb]cdb"
)

func TestAssembleAddress(t *testing.T) {
	type args struct {
		ipv7 string
	}
	tests := []struct {
		name string
		args args
		want AddressSections
	}{
		{"day 07 01 00", args{"abba[mnop]qrst"}, AddressSections{Inners: []string{"mnop"}, Outers: []string{"abba", "qrst"}}},
		{"day 07 01 01", args{"abcd[bddb]xyyx"}, AddressSections{Inners: []string{"bddb"}, Outers: []string{"abcd", "xyyx"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AssembleAddress(tt.args.ipv7); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssembleAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidIPv7(t *testing.T) {
	type args struct {
		a AddressSections
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"ipv7 01", args{AddressSections{Inners: []string{"mnop"}, Outers: []string{"abbaqrst"}}}, true},
		{"ipv7 02", args{AddressSections{Inners: []string{"bddb"}, Outers: []string{"abcdxyyx"}}}, false},
		{"ipv7 03", args{AddressSections{Inners: []string{"qwer"}, Outers: []string{"aaaatyui"}}}, false},
		{"ipv7 04", args{AddressSections{Inners: []string{"asdfgh"}, Outers: []string{"ioxxojzxcvbn"}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidIPv7(tt.args.a); got != tt.want {
				t.Errorf("IsValidIPv7() = %v, want %v", got, tt.want)
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
		name      string
		args      args
		wantCount int
	}{
		{"day 07 01 00", args{TESTINPUT1, 1}, 2},
		{"day 07 01 00", args{TESTINPUT2, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := Day(tt.args.input, tt.args.part); gotCount != tt.wantCount {
				t.Errorf("Day() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestCountValidIPv7Addresses(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{"count ipv7", args{TESTINPUT1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := CountValidIPv7Addresses(tt.args.input); gotCount != tt.wantCount {
				t.Errorf("CountValidIPv7Addresses() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestFindABAs(t *testing.T) {
	type args struct {
		outers []string
	}
	tests := []struct {
		name string
		args args
		want []AreaBroadcastAccessor
	}{
		{"1", args{[]string{"aba", "xyz"}}, []AreaBroadcastAccessor{{ABA: "aba", BAB: "bab"}}},
		{"2", args{[]string{"zazbz", "cdb"}}, []AreaBroadcastAccessor{{ABA: "zaz", BAB: "aza"}, {"zbz", "bzb"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindABAs(tt.args.outers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindABAs() = %v, want %v", got, tt.want)
			}
		})
	}
}
