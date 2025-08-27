package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

// const y15d07InterestingWire string = "a"

type Y15D07 struct{}

func (d *Y15D07) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D07) Part1(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d", year, day, y15d07(d.GetInput(year, day), "a", 1))
}

func (d *Y15D07) Part2(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d", year, day, y15d07(d.GetInput(year, day), "a", 2))
}

func y15d07(input string, wire string, part int) int {
	circuit := map[string]string{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")
		circuit[parts[1]] = parts[0]
	}
	answer := y15d07Process(circuit, wire, map[string]int{})
	if part == 2 {
		circuit["b"] = strconv.Itoa(answer)
		answer = y15d07Process(circuit, wire, map[string]int{})
	}
	return answer
}

func y15d07Process(circuit map[string]string, info string, cache map[string]int) int {
	if memoized, ok := cache[info]; ok {
		return memoized
	}
	if info[0] >= '0' && info[0] <= '9' {
		v, _ := strconv.Atoi(info)
		cache[info] = v
		return v
	}

	parts := strings.Split(circuit[info], " ")
	var result int
	switch {
	case len(parts) == 1:
		result = y15d07Process(circuit, parts[0], cache)
	case parts[0] == "NOT":
		result = 0xffff ^ y15d07Process(circuit, parts[1], cache)
	case parts[1] == "AND":
		result = y15d07Process(circuit, parts[0], cache) & y15d07Process(circuit, parts[2], cache)
	case parts[1] == "OR":
		result = y15d07Process(circuit, parts[0], cache) | y15d07Process(circuit, parts[2], cache)
	case parts[1] == "LSHIFT":
		ls, _ := strconv.Atoi(parts[2])
		result = y15d07Process(circuit, parts[0], cache) << ls
	case parts[1] == "RSHIFT":
		rs, _ := strconv.Atoi(parts[2])
		result = y15d07Process(circuit, parts[0], cache) >> rs
	}
	cache[info] = result
	return result
}
