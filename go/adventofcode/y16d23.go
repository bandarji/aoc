package adventofcode

import (
	"log"
	"strings"
)

var y16d23RegistersPart1 = map[string]int{"a": 7, "b": 0, "c": 0, "d": 0}
var y16d23RegistersPart2 = map[string]int{"a": 12, "b": 0, "c": 0, "d": 0}

func y16d23(input string, registers map[string]int) int {
	instructions := strings.Split(input, "\n")
	cursor := 0
	for cursor < len(instructions) {
		instruction := instructions[cursor]
		fields := strings.Fields(instruction)
		switch fields[0] {
		case "cpy":
			registers[fields[2]] = y16d23ValueOrRegister(registers, fields[1])
		case "inc":
			registers[fields[1]]++
		case "dec":
			registers[fields[1]]--
		case "jnz":
			if y16d23ValueOrRegister(registers, fields[1]) != 0 {
				cursor += y16d23ValueOrRegister(registers, fields[2])
			} else {
				cursor++
			}
			continue
		case "tgl":
			target := cursor + y16d23ValueOrRegister(registers, fields[1])
			if target < len(instructions) {
				targetFields := strings.Fields(instructions[target])
				switch len(targetFields) {
				case 2:
					if targetFields[0] == "inc" {
						targetFields[0] = "dec"
					} else {
						targetFields[0] = "inc"
					}
				case 3:
					if targetFields[0] == "jnz" {
						targetFields[0] = "cpy"
					} else {
						targetFields[0] = "jnz"
					}
				}
				instructions[target] = strings.Join(targetFields, " ")
			}
		default:
			log.Println("unknown instruction", instruction)
		}
		cursor++
	}
	return registers["a"]
}

func y16d23ValueOrRegister(registers map[string]int, insPart string) (response int) {
	if strings.ContainsAny(insPart, "abcd") {
		response = registers[insPart]
	} else {
		response = strToInt(insPart)
	}
	return
}
