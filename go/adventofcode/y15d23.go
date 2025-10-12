package adventofcode

import (
	"strconv"
	"strings"
)

func y15d23(input string, register string, part int) (answer int) {
	instructions := strings.Split(input, "\n")
	answer = y15d23ProcessInstructions(instructions, register, part)
	return
}

func y15d23ProcessInstructions(instructions []string, register string, part int) (answer int) {
	registers := map[string]int{"a": 0, "b": 0}
	index := 0
	if part == 2 {
		registers["a"] = 1
	}
	for index < len(instructions) {
		fields := strings.Fields(strings.ReplaceAll(instructions[index], ",", ""))
		// log.Println(fields)
		switch fields[0] {
		case "hlf":
			registers[fields[1]] /= 2
			index++
		case "tpl":
			registers[fields[1]] *= 3
			index++
		case "inc":
			registers[fields[1]]++
			index++
		case "jmp":
			index += y15d23AtoI(fields[1])
		case "jie":
			if registers[fields[1]]%2 == 0 {
				index += y15d23AtoI(fields[2])
			} else {
				index++
			}
		case "jio":
			if registers[fields[1]] == 1 {
				index += y15d23AtoI(fields[2])
			} else {
				index++
			}
		}
	}
	answer = registers[register]
	return
}

func y15d23AtoI(s string) int {
	s = strings.ReplaceAll(s, "+", "")
	i, _ := strconv.Atoi(s)
	return i
}
