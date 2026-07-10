package adventofcode

import (
	"fmt"
	"strings"
)

type y17d25Rule struct {
	value           int
	direction, next string
}

func y17d25(input string) (checksum int) {
	state, steps, rules := parseInput(input)
	tape := map[int]bool{}
	position := 0

	for i := 0; i < steps; i++ {
		var cv int
		if tape[position] {
			cv = 1
		} else {
			cv = 0
		}
		rule := rules[state][cv]
		// write
		if cv == 0 && rule.value == 1 {
			tape[position] = true
		} else if cv == 1 && rule.value == 0 {
			delete(tape, position)
		}
		// move
		if rule.direction == "left" {
			position--
		} else {
			position++
		}
		// state
		state = rule.next
	}
	checksum = len(tape)
	return
}

func parseInput(input string) (begin string, steps int, rules map[string]map[int]y17d25Rule) {
	blocks := strings.Split(input, "\n\n")
	preamble := strings.Split(blocks[0], "\n")
	begin = y17d25GetBegin(preamble[0])
	steps = y17d25GetSteps(preamble[1])
	rules = y17d25GetRules(blocks[1:])
	return
}

func y17d25GetBegin(line string) (begin string) {
	f := strings.Fields(strings.Trim(line, "."))
	begin = f[3]
	return
}

func y17d25GetSteps(line string) (steps int) {
	fmt.Sscanf(line, "Perform a diagnostic checksum after %d steps.", &steps)
	return
}

func y17d25GetRules(blocks []string) (rules map[string]map[int]y17d25Rule) {
	rules = map[string]map[int]y17d25Rule{}
	for _, block := range blocks {
		lines := strings.Split(block, "\n")
		state := y17d25GetState(lines[0])
		v0, d0, n0 := y17d25ParseRule(lines[2:5])
		v1, d1, n1 := y17d25ParseRule(lines[6:])
		rules[state] = map[int]y17d25Rule{0: {v0, d0, n0}, 1: {v1, d1, n1}}
	}
	return
}

func y17d25ParseRule(lines []string) (v int, d, n string) {
	fv := strings.Fields(strings.Trim(lines[0], "."))
	vd := strings.Fields(strings.Trim(lines[1], "."))
	fn := strings.Fields(strings.Trim(lines[2], "."))
	fmt.Sscanf(fv[4], "%d", &v)
	d = vd[6]
	n = fn[4]
	return
}

func y17d25GetState(line string) (state string) {
	f := strings.Fields(strings.Trim(line, ":"))
	state = f[2]
	return
}
