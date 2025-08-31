package adventofcode

import "testing"

const y15d09TestInput string = `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`

func Test_y15d09(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"shortest", args{y15d09TestInput, 1}, 605},
		{"longest", args{y15d09TestInput, 2}, 982},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y15d09(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y15d09() = %v, want %v", got, tt.want)
			}
		})
	}
}
