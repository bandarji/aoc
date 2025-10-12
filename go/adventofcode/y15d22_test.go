package adventofcode

import "testing"

const y15d22TestInput string = "Hit Points: 71\nDamage: 10"

func Test_y15d22(t *testing.T) {
	type args struct {
		input string
		part  int
		hp    int
		mana  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{"part 1", args{input: y15d22TestInput, part: 1, hp: 50, mana: 500}, 1824},
		{"part 2", args{input: y15d22TestInput, part: 2, hp: 50, mana: 500}, 1937},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y15d22(tt.args.input, tt.args.part, tt.args.hp, tt.args.mana); gotAnswer != tt.wantAnswer {
				t.Errorf("y15d22() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
