package adventofcode

import (
	"strings"
	"testing"
)

const y16d23TestInput string = `cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a`

func Test_y16d23Processing(t *testing.T) {
	c := y16d23Computer{
		registers:    map[string]int{"a": 0, "b": 0, "c": 0, "d": 0},
		instructions: strings.Split(y16d23TestInput, "\n"),
		cursor:       0,
		size:         len(strings.Split(y16d23TestInput, "\n")),
	}

	// 1: cpy 2 a
	c.executeInstruction()
	if c.registers["a"] != 2 {
		t.Errorf("expected register a to be 2, got %d", c.registers["a"])
	}

	// 2: tgl a
	c.executeInstruction()
	if c.instructions[3] != "inc a" {
		t.Errorf("expected instruction to be 'inc a', got %s", c.instructions[3])
	}

	// 3: tgl a
	c.executeInstruction()
	if c.instructions[4] != "jnz 1 a" {
		t.Errorf("expected instruction to be 'jnz 1 a', got %s", c.instructions[4])
	}

	// 4: inc q
	c.executeInstruction()
	if c.registers["a"] != 3 {
		t.Errorf("expected register a to be 3, got %d", c.registers["a"])
	}

	// 5: jnz 1 a
	c.executeInstruction()
	if c.cursor != 7 {
		t.Errorf("expected cursor to be 8, got %d", c.cursor)
	}

}

func Test_y16d23(t *testing.T) {
	type args struct {
		input     string
		registers map[string]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test p1", args{input: y16d23TestInput, registers: map[string]int{"a": 0, "b": 0, "c": 0, "d": 0}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d23(tt.args.input, tt.args.registers); got != tt.want {
				t.Errorf("y16d23() = %v, want %v", got, tt.want)
			}
		})
	}
}
