package adventofcode

import "strings"

func y17d05(input string, part int) (answer int) {
	instructions := y17d05ParseInput(input)
	cursor := 0
	steps := 0
	for cursor < len(instructions) {
		instruction := instructions[cursor]
		if part == 1 {
			instructions[cursor]++
		} else {
			if instruction >= 3 {
				instructions[cursor]--
			} else {
				instructions[cursor]++
			}
		}
		cursor += instruction
		steps++
	}
	answer = steps
	return
}

func y17d05ParseInput(input string) (instructions []int) {
	for line := range strings.SplitSeq(input, "\n") {
		instructions = append(instructions, strToInt(line))
	}
	return
}
