package adventofcode

import (
	"testing"
)

func TestY15D04_Part1(t *testing.T) {
	type args struct {
		year int
		day  int
	}
	tests := []struct {
		name string
		d    *Y15D04
		args args
		want string
	}{
		{"test 1", &Y15D04{}, args{2015, 4}, "Year=2015 Day=04 Part 1: 282749"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Y15D04{}
			if got := d.Part1(tt.args.year, tt.args.day); got != tt.want {
				t.Errorf("Y15D04.Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestY15D04_Part2(t *testing.T) {
	type args struct {
		year int
		day  int
	}
	tests := []struct {
		name string
		d    *Y15D04
		args args
		want string
	}{
		{"test 1", &Y15D04{}, args{2015, 4}, "Year=2015 Day=04 Part 2: 9962624"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Y15D04{}
			if got := d.Part2(tt.args.year, tt.args.day); got != tt.want {
				t.Errorf("Y15D04.Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
