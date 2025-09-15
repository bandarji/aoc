package adventofcode

import "testing"

const y15d13TestInput = `Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.`

func Test_y15d13(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"test p1", y15d13TestInput, 1, 330},
		{"test p2", y15d13TestInput, 2, 286},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := y15d13(tt.input, tt.part)
			if got != tt.want {
				t.Errorf("y15d13() = %v, want %v", got, tt.want)
			}
		})
	}
}
