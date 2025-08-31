package adventofcode

import (
	"fmt"
)

const y15d10input string = "1113122113"

type Y15D10 struct{}

func (d *Y15D10) GetInput(year, day int) string {
	return y15d10input
}

// Reduced cycles from 40 and 50 to 4 and 5
// Use 'make test' to validate locally
// Edit this file with 40 and 50 but prepare to wait a few minutes for the run
func (d *Y15D10) Part1(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d", year, day, y15d10(d.GetInput(year, day), 4))
}

func (d *Y15D10) Part2(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d", year, day, y15d10(d.GetInput(year, day), 5))
}

func y15d10(input string, cycles int) int {
	for i := 0; i < cycles; i++ {
		input = y15d10LookAndSay(input)
	}
	return len(input)
}

func y15d10LookAndSay(input string) (output string) {
	if len(input) == 0 {
		return ""
	}
	count := 0
	c := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] == c {
			count++
		} else {
			output += fmt.Sprintf("%d%s", count, string(c))
			c = input[i]
			count = 1
		}
	}
	output += fmt.Sprintf("%d%s", count, string(c))
	return
}
