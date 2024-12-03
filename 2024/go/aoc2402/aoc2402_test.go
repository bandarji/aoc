package aoc2402

import (
	"testing"
)

func Test_checkDifferent(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"checkDifferent true [1, 2, 3, 4, 5]", args{[]int{1, 2, 3, 4, 5}}, true},
		{"checkDifferent false [1, 1, 3, 4, 5]", args{[]int{1, 1, 3, 4, 5}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkDifferent(tt.args.values); got != tt.want {
				t.Errorf("checkDifferent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkDeltas(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"checkDeltas true [1, 2, 3, 4, 5]", args{[]int{1, 2, 3, 4, 5}}, true},
		{"checkDeltas false [1, 2, 30, 31, 32]", args{[]int{1, 2, 30, 31, 32}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkDeltas(tt.args.values); got != tt.want {
				t.Errorf("checkDeltas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkLinear(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"checkLinear true [1, 2, 3, 4, 5]", args{[]int{1, 2, 3, 4, 5}}, true},
		{"checkLinear true [5, 4, 3, 2, 1]", args{[]int{5, 4, 3, 2, 1}}, true},
		{"checkLinear false [1, 2, 3, 5, 4]", args{[]int{1, 2, 3, 5, 4}}, false},
		{"checkLinear false [1, 2, 3, 4, 3]", args{[]int{1, 2, 3, 4, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkLinear(tt.args.values); got != tt.want {
				t.Errorf("checkLinear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkAscending(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"checkAscending true [1, 2, 3, 4, 5]", args{[]int{1, 2, 3, 4, 5}}, true},
		{"checkAscending false [1, 2, 3, 5, 4]", args{[]int{1, 2, 3, 5, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkAscending(tt.args.values); got != tt.want {
				t.Errorf("checkAscending() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkDescending(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"checkDescending true [5, 4, 3, 2, 1]", args{[]int{5, 4, 3, 2, 1}}, true},
		{"checkDescending false [1, 2, 3, 5, 4]", args{[]int{1, 2, 3, 5, 4}}, false},
		{"checkDescending false [5, 4, 3, 2, 3]", args{[]int{5, 4, 3, 2, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkDescending(tt.args.values); got != tt.want {
				t.Errorf("checkDescending() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAoc240201(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name     string
		args     args
		wantSafe int
	}{
		{"AOC 2402.1", args{AOC2402_TEST}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSafe := Aoc240201(tt.args.input); gotSafe != tt.wantSafe {
				t.Errorf("Aoc240201() = %v, want %v", gotSafe, tt.wantSafe)
			}
		})
	}
}

func TestAoc240202(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name     string
		args     args
		wantSafe int
	}{
		{"AOC 2402.2", args{AOC2402_TEST}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSafe := Aoc240202(tt.args.input); gotSafe != tt.wantSafe {
				t.Errorf("Aoc240202() = %v, want %v", gotSafe, tt.wantSafe)
			}
		})
	}
}
