package adventofcode

import "testing"

const y15d14TestInput = `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`

func Test_y15d14(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input   string
		seconds int
		part    int
		want    int
	}{
		{"test p1", y15d14TestInput, 1000, 1, 1120},
		{"test p2", y15d14TestInput, 1000, 2, 689},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := y15d14(tt.input, tt.seconds, tt.part)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("y15d14() = %v, want %v", got, tt.want)
			}
		})
	}
}
