package adventofcode

import (
	"strings"
)

type y17d18Program struct {
	registers                   map[string]int
	cursor, id, frequency, sent int
	q                           []int
}

func y17d18(input string, part int) int {
	programs := [2]y17d18Program{
		{registers: map[string]int{"p": 0}, cursor: 0, id: 0, q: []int{}, sent: 0, frequency: 0},
		{registers: map[string]int{"p": 1}, cursor: 0, id: 1, q: []int{}, sent: 0, frequency: 0},
	}
	instructions := strings.Split(input, "\n")
	insCount := len(instructions)
	answer := 0
	if part == 1 {
		programs[1].cursor = 3 * len(instructions)
	}
	for answer == 0 {
		for id := range 2 {
			other := (id + 1)
			other %= 2
			skip := false
			if programs[id].cursor >= insCount {
				continue
			}
			fields := strings.Fields(instructions[programs[id].cursor])
			switch fields[0] {
			case "set":
				programs[id].registers[fields[1]] = y17d18ValueOrRegister(programs[id].registers, fields[2])
			case "add":
				programs[id].registers[fields[1]] += y17d18ValueOrRegister(programs[id].registers, fields[2])
			case "mul":
				programs[id].registers[fields[1]] *= y17d18ValueOrRegister(programs[id].registers, fields[2])
			case "mod":
				programs[id].registers[fields[1]] %= y17d18ValueOrRegister(programs[id].registers, fields[2])
			case "snd":
				programs[id].frequency = programs[id].registers[fields[1]]
				programs[other].q = append(programs[other].q, programs[id].frequency)
				programs[id].sent++
			case "rcv":
				if part == 1 && programs[id].registers[fields[1]] != 0 {
					answer = programs[id].frequency
				}
				if part == 2 {
					if len(programs[id].q) == 0 && len(programs[other].q) == 0 {
						answer = programs[1].sent
					}
					if len(programs[id].q) > 0 {
						programs[id].registers[fields[1]] = programs[id].q[0]
						programs[id].q = programs[id].q[1:]
					} else {
						skip = true
					}
				}
			case "jgz":
				if y17d18ValueOrRegister(programs[id].registers, fields[1]) > 0 {
					programs[id].cursor += y17d18ValueOrRegister(programs[id].registers, fields[2])
					skip = true
				}
			}
			if !skip {
				programs[id].cursor++
			}
		}
	}
	return answer
}

func y17d18ValueOrRegister(registers map[string]int, possible string) (response int) {
	if len(possible) == 1 && possible[0] >= 'a' && possible[0] <= 'z' {
		response = registers[possible]
	} else {
		response = strToInt(possible)
	}
	return
}
