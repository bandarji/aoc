package adventofcode

import "testing"

const y17d10TestInput string = `3,4,1,5`

func Test_y17d10(t *testing.T) {
	type args struct {
		input      string
		listLength int
		part       int
	}
	tests := []struct {
		name        string
		args        args
		wantProduct int
	}{
		{"test p1", args{input: y17d10TestInput, listLength: 5, part: 1}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotProduct := y17d10(tt.args.input, tt.args.listLength, tt.args.part); gotProduct != tt.wantProduct {
				t.Errorf("y17d10() = %v, want %v", gotProduct, tt.wantProduct)
			}
		})
	}
}
