package adventofcode

import (
	"log"
	"strings"
)

var y16d23Registers = map[string]int{"a": 7}

func y16d23(input string, registers map[string]int) int {
	instructions := strings.Split(input, "\n")
	cursor := 0
	for cursor < len(instructions) {
		// if cursor > 24 {
		// 	log.Fatal("cursor", cursor)
		// }
		log.Println("registers", registers, "cursor", cursor, "instruction", instructions[cursor])
		instruction := instructions[cursor]
		fields := strings.Fields(instruction)
		switch fields[0] {
		case "cpy":
			y16d23Copy(registers, fields[1], fields[2])
		case "inc":
			y16d23Inc(registers, fields[1])
		case "dec":
			y16d23Dec(registers, fields[1])
		case "jnz":
			cursor += y16d23JumpNotZero(registers, fields[1], fields[2])
		case "tgl":
			instructions = y16d23Toggle(registers, instructions, cursor)
		}
		cursor++
	}
	return registers["a"]
}

func y16d23Toggle(registers map[string]int, instructions []string, cursor int) []string {
	modIndex := cursor + y16d23ValueOrRegister(registers, instructions[cursor])
	if modIndex < len(instructions) {
		modInstructionFields := strings.Fields(instructions[modIndex])
		if len(modInstructionFields) == 2 {
			if modInstructionFields[0] == "inc" {
				modInstructionFields[0] = "dec"
			} else {
				modInstructionFields[0] = "inc"
			}
		}
		if len(modInstructionFields) == 3 {
			if modInstructionFields[0] == "jnz" {
				modInstructionFields[0] = "cpy"
			} else {
				modInstructionFields[0] = "jnz"
			}
		}
		instructions[modIndex] = strings.Join(modInstructionFields, " ")
	}
	return instructions
}

func y16d23JumpNotZero(registers map[string]int, i1, i2 string) (delta int) {
	switch y16d23ValueOrRegister(registers, i1) {
	case 0:
		delta = 0
	default:
		delta = strToInt(i2) - 1
	}
	return
}

func y16d23ValueOrRegister(registers map[string]int, insPart string) (response int) {
	if strings.ContainsAny(insPart, "abcd") {
		response = registers[insPart]
	} else {
		response = strToInt(insPart)
	}
	return
}

func y16d23Copy(registers map[string]int, i1, i2 string) {
	registers[i2] = y16d23ValueOrRegister(registers, i1)
}

func y16d23Inc(registers map[string]int, i string) {
	registers[i]++
}

func y16d23Dec(registers map[string]int, i string) {
	registers[i]--
}
