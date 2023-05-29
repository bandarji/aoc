package d03

import "testing"

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
		{"Day 03 01 01", args{"5 10 25", 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := Day(tt.args.input, tt.args.part); gotCount != tt.wantCount {
				t.Errorf("Day() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
