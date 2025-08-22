package adventofcode

import (
	"testing"
)

func Test_y15d02p1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name     string
		args     args
		wantSqft int
	}{
		{
			name:     "2x3x4",
			args:     args{input: "2x3x4"},
			wantSqft: 58,
		},
		{
			name:     "1x1x10",
			args:     args{input: "1x1x10"},
			wantSqft: 43,
		},
		{
			name:     "2x3x4\n1x1x10",
			args:     args{input: "2x3x4\n1x1x10"},
			wantSqft: 58 + 43,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSqft := y15d02p1(tt.args.input); gotSqft != tt.wantSqft {
				t.Errorf("y15d02p1() = %v, want %v", gotSqft, tt.wantSqft)
			}
		})
	}
}

func Test_y15d02p2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		args   args
		wantFt int
	}{
		{
			name:   "2x3x4",
			args:   args{input: "2x3x4"},
			wantFt: 34,
		},
		{
			name:   "1x1x10",
			args:   args{input: "1x1x10"},
			wantFt: 14,
		},
		{
			name:   "2x3x4\n1x1x10",
			args:   args{input: "2x3x4\n1x1x10"},
			wantFt: 34 + 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFt := y15d02p2(tt.args.input); gotFt != tt.wantFt {
				t.Errorf("y15d02p2() = %v, want %v", gotFt, tt.wantFt)
			}
		})
	}
}

func Test_y15d02solve(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name:       "2x3x4",
			args:       args{input: "2x3x4", part: 1},
			wantResult: 58,
		},
		{
			name:       "1x1x10",
			args:       args{input: "1x1x10", part: 1},
			wantResult: 43,
		},
		{
			name:       "2x3x4\n1x1x10",
			args:       args{input: "2x3x4\n1x1x10", part: 1},
			wantResult: 58 + 43,
		},
		{
			name:       "2x3x4",
			args:       args{input: "2x3x4", part: 2},
			wantResult: 34,
		},
		{
			name:       "1x1x10",
			args:       args{input: "1x1x10", part: 2},
			wantResult: 14,
		},
		{
			name:       "2x3x4\n1x1x10",
			args:       args{input: "2x3x4\n1x1x10", part: 2},
			wantResult: 34 + 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := y15d02solve(tt.args.input, tt.args.part); gotResult != tt.wantResult {
				t.Errorf("y15d02solve() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
