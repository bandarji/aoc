package aoc2404

import (
	"testing"
)

func TestAoc240401(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{"test 2024 Day 4 Part 1", args{AOC2404_TEST}, 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := Aoc240401(tt.args.input); gotCount != tt.wantCount {
				t.Errorf("Aoc240401() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestAoc240402(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{"test 2024 Day 4 Part 2", args{AOC2404_TEST}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := Aoc240402(tt.args.input); gotCount != tt.wantCount {
				t.Errorf("Aoc240402() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
