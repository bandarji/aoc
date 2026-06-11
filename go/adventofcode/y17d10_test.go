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
		wantProduct string
	}{
		{"test p1", args{input: y17d10TestInput, listLength: 5, part: 1}, "12"},
		{"test p2", args{input: "", listLength: 256, part: 2}, "a2582a3a0e66e6e86e3812dcb672a272"},
		{"test p2", args{input: "AoC 2017", listLength: 256, part: 2}, "33efeb34ea91902bb2f59c9920caa6cd"},
		{"test p2", args{input: "1,2,3", listLength: 256, part: 2}, "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"test p2", args{input: "1,2,4", listLength: 256, part: 2}, "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotProduct := y17d10(tt.args.input, tt.args.listLength, tt.args.part); gotProduct != tt.wantProduct {
				t.Errorf("y17d10() = %v, want %v", gotProduct, tt.wantProduct)
			}
		})
	}
}
