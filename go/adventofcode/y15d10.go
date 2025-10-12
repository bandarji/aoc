package adventofcode

import (
	"fmt"
)

const y15d10input string = "1113122113"

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
