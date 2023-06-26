package d12

import (
	"strconv"
	"strings"
)

type Computer struct {
	Instructions   []string
	Cursor, Length int
	Registers      map[string]int
}

func (c *Computer) ReadInstructions(input string) {
	c.Instructions = strings.Split(input, "\n")
	c.Length = len(c.Instructions)
}

func (c *Computer) Cpy(f []string) {
	value, err := strconv.Atoi(f[1])
	if err != nil {
		value = c.Registers[f[1]]
	}
	c.Registers[f[2]] = value
	c.Cursor++
}

func (c *Computer) Inc(f []string) {
	c.Registers[f[1]]++
	c.Cursor++
}

func (c *Computer) Dec(f []string) {
	c.Registers[f[1]]--
	c.Cursor++
}

func (c *Computer) Jnz(f []string) {
	value, err := strconv.Atoi(f[1])
	if err != nil {
		value = c.Registers[f[1]]
	}
	if value != 0 {
		skip, _ := strconv.Atoi(f[2])
		c.Cursor += skip
	} else {
		c.Cursor++
	}
}

func NewComputer() *Computer {
	return &Computer{[]string{}, 0, 0, map[string]int{"a": 0, "b": 0, "c": 0, "d": 0}}
}

func Day(input string, part int) int {
	computer := NewComputer()
	if part == 2 {
		computer.Registers["c"] = 1
	}
	computer.ReadInstructions(input)
	for computer.Cursor < computer.Length {
		instruction := computer.Instructions[computer.Cursor]
		fields := strings.Fields(instruction)
		switch fields[0] {
		case "cpy":
			computer.Cpy(fields)
		case "inc":
			computer.Inc(fields)
		case "dec":
			computer.Dec(fields)
		case "jnz":
			computer.Jnz(fields)
		}
	}
	return computer.Registers["a"]
}
