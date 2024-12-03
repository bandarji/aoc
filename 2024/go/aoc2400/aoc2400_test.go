package aoc2400

import (
	"testing"
)

func TestAbsInts(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name  string
		args  args
		wantA int
	}{
		{"abs 3, 1", args{3, 1}, 2},
		{"abs 1, 3", args{1, 3}, 2},
		{"abs 0, 0", args{0, 0}, 0},
		{"abs -3, -7", args{-3, -7}, 4},
		{"abs -7, -3", args{-7, -3}, 4},
		{"abs 2, -7", args{2, -7}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotA := AbsInts(tt.args.x, tt.args.y); gotA != tt.wantA {
				t.Errorf("AbsInts() = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}

func TestStrToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		wantV int
	}{
		{"int '123'", args{"123"}, 123},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := StrToInt(tt.args.s); gotV != tt.wantV {
				t.Errorf("StrToInt() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}
