package d10

import (
	"reflect"
	"testing"
)

func TestIntersectStringSlices(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"empties", args{s1: []string{}, s2: []string{}}, []string{}},
		{"empty result", args{s1: []string{"foo"}, s2: []string{"bar"}}, []string{}},
		{"foo and bar", args{s1: []string{"foo", "bar", "dog"}, s2: []string{"bar", "foo", "cat"}}, []string{"bar", "foo"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectStringSlices(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectStringSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
