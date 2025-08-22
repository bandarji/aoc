package adventofcode

import (
	"testing"
)

func Test_y15d03p1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantHouses int
	}{
		{"test '>'", args{input: ">"}, 2},
		{"test '^>v<'", args{input: "^>v<"}, 4},
		{"test '^v^v^v^v^v'", args{input: "^v^v^v^v^v"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHouses := y15d03p1(tt.args.input); gotHouses != tt.wantHouses {
				t.Errorf("y15d03p1() = %v, want %v", gotHouses, tt.wantHouses)
			}
		})
	}
}

func Test_y15d03p2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantHouses int
	}{
		{"test '^v'", args{input: "^v"}, 3},
		{"test '^>v<'", args{input: "^>v<"}, 3},
		{"test '^v^v^v^v^v'", args{input: "^v^v^v^v^v"}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHouses := y15d03p2(tt.args.input); gotHouses != tt.wantHouses {
				t.Errorf("y15d03p2() = %v, want %v", gotHouses, tt.wantHouses)
			}
		})
	}
}
