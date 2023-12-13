package d12

import (
	"reflect"
	"testing"
)

func TestUnfoldPattern(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name         string
		args         args
		wantUnfolded string
	}{
		{".# -> .#?.#?.#?.#?.#", args{s: ".#"}, ".#?.#?.#?.#?.#"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUnfolded := UnfoldPattern(tt.args.s); gotUnfolded != tt.wantUnfolded {
				t.Errorf("UnfoldPattern() = %v, want %v", gotUnfolded, tt.wantUnfolded)
			}
		})
	}
}

func TestUnfoldGroups(t *testing.T) {
	type args struct {
		g []int
	}
	tests := []struct {
		name         string
		args         args
		wantUnfolded []int
	}{
		{"1 -> 1,1,1,1,1", args{g: []int{1}}, []int{1, 1, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUnfolded := UnfoldGroups(tt.args.g); !reflect.DeepEqual(gotUnfolded, tt.wantUnfolded) {
				t.Errorf("UnfoldGroups() = %v, want %v", gotUnfolded, tt.wantUnfolded)
			}
		})
	}
}
