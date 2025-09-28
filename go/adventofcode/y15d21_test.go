package adventofcode

import "testing"

const y15d21TestInput string = "Hit Points: 104\nDamage: 8\nArmor: 1"

func Test_y15d21(t *testing.T) {
	type args struct {
		input      string
		startingHP int
		part       int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{"part 1", args{input: y15d21TestInput, startingHP: y15d21StartingHP, part: 1}, 78},
		{"part 2", args{input: y15d21TestInput, startingHP: y15d21StartingHP, part: 2}, 148},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y15d21(tt.args.input, tt.args.startingHP, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y15d21() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
