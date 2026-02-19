package adventofcode

import (
	"strings"
)

func y16d12(input string, part int) (answer int) {
	registers := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0}
	if part == 2 {
		registers["c"] = 1
	}
	var cursor int
	instructions := strings.Split(input, "\n")
	for cursor < len(instructions) {
		// log.Println("cursor", cursor, "instruction", instructions[cursor])
		fields := strings.Fields(instructions[cursor])
		switch fields[0] {
		case "cpy":
			if strings.ContainsAny(fields[1], "abcd") {
				registers[fields[2]] = registers[fields[1]]
			} else {
				registers[fields[2]] = strToInt(fields[1])
			}
			cursor++
		case "inc":
			registers[fields[1]]++
			cursor++
		case "dec":
			registers[fields[1]]--
			cursor++
		case "jnz":
			if strings.ContainsAny(fields[1], "abcd") {
				if registers[fields[1]] != 0 {
					cursor += strToInt(fields[2])
				} else {
					cursor++
				}
			} else {
				if strToInt(fields[1]) != 0 {
					cursor += strToInt(fields[2])
				}
			}
		}
	}
	answer = registers["a"]
	return
}
