package d02

import "testing"

func TestDay(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name             string
		args             args
		wantBathroomCode string
	}{
		{"Day 02 01", args{"ULL\nRRDDD\nLURDL\nUUUUD", 1}, "1985"},
		{"Day 02 01", args{"ULL\nRRDDD\nLURDL\nUUUUD", 2}, "5DB3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBathroomCode := Day(tt.args.input, tt.args.part); gotBathroomCode != tt.wantBathroomCode {
				t.Errorf("Day() = %v, want %v", gotBathroomCode, tt.wantBathroomCode)
			}
		})
	}
}
