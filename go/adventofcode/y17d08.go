package adventofcode

import (
	"strings"
)

type y17d08Instruction struct {
	regChange, regCompare string
	inc, compare          int
	operation             string
}

func y17d08(input string, part int) int {
	highestSeen := 0
	computer := map[string]int{}
	instructions := y17d08ParseInput(input)
	for _, ins := range instructions {
		var flag bool
		val := computer[ins.regCompare]
		switch ins.operation {
		case ">":
			flag = val > ins.compare
		case "<":
			flag = val < ins.compare
		case ">=":
			flag = val >= ins.compare
		case "<=":
			flag = val <= ins.compare
		case "==":
			flag = val == ins.compare
		case "!=":
			flag = val != ins.compare
		}
		if flag {
			computer[ins.regChange] += ins.inc
		}
		highestSeen = max(highestSeen, y17d08HighestRegisterValue(computer))
	}
	if part == 1 {
		return y17d08HighestRegisterValue(computer)
	} else {
		return highestSeen
	}
}

func y17d08HighestRegisterValue(computer map[string]int) int {
	hi := 0
	for _, v := range computer {
		hi = max(hi, v)
	}
	return hi
}

func y17d08ParseInput(input string) (instructions []y17d08Instruction) {
	// b inc 5 if a > 1
	// 0 1   2 3  4 5 6
	for line := range strings.SplitSeq(input, "\n") {
		fields := strings.Fields(line)
		ins := y17d08Instruction{
			regChange:  fields[0],
			inc:        strToInt(fields[2]),
			regCompare: fields[4],
			compare:    strToInt(fields[6]),
			operation:  fields[5],
		}
		if fields[1] == "dec" {
			ins.inc *= -1
		}
		instructions = append(instructions, ins)
	}
	return
}
