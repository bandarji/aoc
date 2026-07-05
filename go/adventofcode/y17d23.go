package adventofcode

import (
	"strconv"
	"strings"
)

func y17d23(input string, part int) (answer int) {
	registers := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0}
	instructions := strings.Split(input, "\n")
	answer = y17d23Run(registers, instructions)
	// log.Printf("%v\n", registers)
	// map[a:0 b:67 c:67 d:67 e:67 f:1 g:0 h:0]
	if part == 2 {
		answer = y17d23DebugMode(registers)
	}
	return
}

func y17d23Run(registers map[string]int, instructions []string) (answer int) {
	length := len(instructions)
	cursor := 0
	for cursor < length {
		instruction := strings.Fields(instructions[cursor])
		opcode := instruction[0]
		switch opcode {
		case "set":
			registers[instruction[1]] = y17d23ValueOrRegister(registers, instruction[2])
		case "sub":
			registers[instruction[1]] -= y17d23ValueOrRegister(registers, instruction[2])
		case "mul":
			registers[instruction[1]] *= y17d23ValueOrRegister(registers, instruction[2])
			answer++
		case "jnz":
			if y17d23ValueOrRegister(registers, instruction[1]) != 0 {
				// log.Printf("%s %v %d\n", instruction, registers, answer)
				cursor += y17d23ValueOrRegister(registers, instruction[2])
				continue
			}
		}
		cursor++
	}
	return
}

func y17d23ValueOrRegister(registers map[string]int, insPart string) (response int) {
	if v, err := strconv.Atoi(insPart); err != nil {
		response = registers[insPart]
	} else {
		response = v
	}
	return
}

func y17d23DebugMode(registers map[string]int) int {
	b := registers["b"]
	c := registers["c"]
	b *= 100
	b += 100000
	var h int
	for c = b + 17000; ; b += 17 {
		if !isPrime(b) {
			h++
		}
		if b == c {
			break
		}
	}
	return h
}
