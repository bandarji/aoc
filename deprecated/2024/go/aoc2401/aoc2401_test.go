package aoc2401

import (
	_ "embed"
	"testing"
)

func TestAoc240101(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name      string
		args      args
		wantTotal int
	}{
		{"aoc 2401.1", args{AOC2401_TEST}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := Aoc240101(tt.args.input); gotTotal != tt.wantTotal {
				t.Errorf("Aoc240101() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func TestAoc240102(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name      string
		args      args
		wantTotal int
	}{
		{"aoc 2401.2", args{AOC2401_TEST}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := Aoc240102(tt.args.input); gotTotal != tt.wantTotal {
				t.Errorf("Aoc240102() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
