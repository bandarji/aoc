package adventofcode

import (
	"fmt"
	"sort"
	"strings"
)

type y16d10Instruction struct {
	bot    int
	lo, hi string
}

func y16d10(input string, part, compLo, compHi int) (answer int) {
	bots, instructions := y16d10ParseInput(input)
	answer = y16d10Process(bots, instructions, part, compLo, compHi)
	return
}

func y16d10Process(bots map[int][]int, instructions []y16d10Instruction, part, compLo, compHi int) int {
	outputs := map[int]int{}
	for outputs[0] == 0 || outputs[1] == 0 || outputs[2] == 0 {
		for _, instruction := range instructions {
			if len(bots[instruction.bot]) == 2 {
				sort.Ints(bots[instruction.bot])
				lo, hi := bots[instruction.bot][0], bots[instruction.bot][1]
				if part == 1 && lo == compLo && hi == compHi {
					return instruction.bot
				}
				var out, receive int
				if strings.HasPrefix(instruction.lo, "output ") {
					fmt.Sscanf(instruction.lo, "output %d", &out)
					outputs[out] = lo
				} else {
					fmt.Sscanf(instruction.lo, "bot %d", &receive)
					bots[receive] = append(bots[receive], lo)
				}
				if strings.HasPrefix(instruction.hi, "output ") {
					fmt.Sscanf(instruction.hi, "output %d", &out)
					outputs[out] = hi
				} else {
					fmt.Sscanf(instruction.hi, "bot %d", &receive)
					bots[receive] = append(bots[receive], hi)
				}
				bots[instruction.bot] = []int{}
			}
		}
	}
	return outputs[0] * outputs[1] * outputs[2]
}

func y16d10ParseInput(input string) (bots map[int][]int, instructions []y16d10Instruction) {
	bots = map[int][]int{}
	instructions = []y16d10Instruction{}
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		if fields[0] == "value" {
			var chip, bot int
			chip = strToInt(fields[1])
			bot = strToInt(fields[5])
			bots[bot] = append(bots[bot], chip)
		} else if strings.HasPrefix(line, "bot ") {
			var bot int
			var lo, hi string
			bot = strToInt(fields[1])
			lo = strings.Join(fields[5:7], " ")
			hi = strings.Join(fields[len(fields)-2:], " ")
			instructions = append(instructions, y16d10Instruction{bot, lo, hi})
		}
	}
	return
}
