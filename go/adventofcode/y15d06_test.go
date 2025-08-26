package adventofcode

import "testing"

func Test_y15d06(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name         string
		args         args
		wantResponse int
	}{
		{"part 1, 0,0 -> 999,999", args{"toggle 0,0 through 999,999", 1}, 1_000_000},
		{"part 1, 0,0 -> 999,0", args{"toggle 0,0 through 999,0", 1}, 1_000},
		{"part 1, 499,499 -> 500,500", args{"toggle 499,499 through 500,500", 1}, 4},
		{"part 2, 0,0 -> 0,0", args{"turn on 0,0 through 0,0", 2}, 1},
		{"part 2, 0,0 -> 999,999", args{"toggle 0,0 through 999,999", 2}, 2_000_000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResponse := y15d06(tt.args.input, tt.args.part); gotResponse != tt.wantResponse {
				t.Errorf("y15d06() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
