package adventofcode

import (
	"strings"
)

const XXX string = `
set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2
The first four instructions set a to 1, add 2 to it, square it, and then set it to itself modulo 5, resulting in a value of 4.
Then, a sound with frequency 4 (the value of a) is played.
After that, a is set to 0, causing the subsequent rcv and jgz instructions to both be skipped (rcv because a is 0, and jgz because a is not greater than 0).
Finally, a is set to 1, causing the next jgz instruction to activate, jumping back two instructions to another jump, which jumps again to the rcv, which ultimately triggers the recover operation.
At the time the recover operation is executed, the frequency of the last sound played is 4.
`

func y17d18ValueOrRegister(registers map[string]int, possible string) (response int) {
	if len(possible) == 1 && possible[0] >= 'a' && possible[0] <= 'z' {
		response = registers[possible]
	} else {
		response = strToInt(possible)
	}
	return
}

func y17d18(input string, part int) (frequency int) {
	registers := map[string]int{}
	instructions := strings.Split(input, "\n")
	insCount := len(instructions)
	cursor, frequency := 0, 0
	for cursor < insCount {
		instruction := instructions[cursor]
		fields := strings.Fields(instruction)
		firstRegister := fields[1]
		switch fields[0] {
		case "set":
			registers[firstRegister] = y17d18ValueOrRegister(registers, fields[2])
		case "add":
			registers[firstRegister] += y17d18ValueOrRegister(registers, fields[2])
		case "mul":
			registers[firstRegister] *= y17d18ValueOrRegister(registers, fields[2])
		case "mod":
			registers[firstRegister] %= y17d18ValueOrRegister(registers, fields[2])
		case "snd":
			// log.Printf("%s - %+v", instruction, registers)
			frequency = registers[firstRegister]
		case "rcv":
			if registers[firstRegister] != 0 {
				return frequency
			}
		case "jgz":
			if y17d18ValueOrRegister(registers, firstRegister) > 0 {
				cursor += y17d18ValueOrRegister(registers, fields[2])
				continue
			}
		}
		// log.Printf("cursor: %d, frequency: %d, registers: %v", cursor, frequency, registers)
		cursor++
	}
	return frequency
}
