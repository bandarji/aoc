package main

import (
	"aoc/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const INPUT_FILENAME string = "input_day05.txt"
const TEST_INPUT string = "    [D]    \n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2"
const TEST_EXPECTED_1 string = "CMZ"
const TEST_EXPECTED_2 string = "MCD"

type Instruction struct {
	Count int
	From  int
	To    int
}

func assembleStacks(input string) map[int]string {
	stacks := make(map[int]string)
	for _, line := range strings.Split(input, "\n") {
		if len(line) < 4 || strings.HasPrefix(line, " 1") {
			continue
		}
		for i := 0; i < len(line); i++ {
			if string(line[i]) == "[" {
				index := i/4 + 1
				stacks[index] = string(line[i+1]) + stacks[index]
			}
		}
	}
	return stacks
}

func parseInstruction(input string) Instruction {
	var i Instruction
	// move 3 from 1 to 3
	// 0    1 2    3 4  5
	fields := strings.Fields(input)
	i.Count, _ = strconv.Atoi(fields[1])
	i.From, _ = strconv.Atoi(fields[3])
	i.To, _ = strconv.Atoi(fields[5])
	return i
}

func runInstructions(stacks map[int]string, input string) map[int]string {
	for _, line := range strings.Split(input, "\n") {
		if !strings.HasPrefix(line, "move ") {
			continue
		}
		ins := parseInstruction(line)
		for ins.Count > 0 {
			stacks = processInstruction(stacks, ins.From, ins.To)
			ins.Count--
		}
	}
	return stacks
}

func runInstructions9001(stacks map[int]string, input string) map[int]string {
	for _, line := range strings.Split(input, "\n") {
		if !strings.HasPrefix(line, "move ") {
			continue
		}
		ins := parseInstruction(line)
		stacks = processInstruction9001(stacks, ins)
	}
	return stacks
}

func processInstruction(stacks map[int]string, from int, to int) map[int]string {
	stacks[to] += string(stacks[from][len(stacks[from])-1])
	stacks[from] = stacks[from][:len(stacks[from])-1]
	return stacks
}

func processInstruction9001(stacks map[int]string, i Instruction) map[int]string {
	index := len(stacks[i.From]) - i.Count
	crates := stacks[i.From][index:]
	stacks[i.From] = stacks[i.From][:index]
	stacks[i.To] += crates
	return stacks
}

func day0501(input string) string {
	sections := strings.Split(input, "\n\n")
	stacks := assembleStacks(sections[0])
	stacks = runInstructions(stacks, sections[1])
	return topCrates(stacks)
}

func topCrates(stacks map[int]string) string {
	var s string
	for i := 1; i <= 9; i++ {
		value, exists := stacks[i]
		if !exists {
			break
		}
		s += string(value[len(value)-1])
	}
	return s
}

func day0502(input string) string {
	sections := strings.Split(input, "\n\n")
	stacks := assembleStacks(sections[0])
	stacks = runInstructions9001(stacks, sections[1])
	return topCrates(stacks)
}

func main() {
	var testResponse string
	var errMsg string
	log.SetFlags(log.Ldate | log.Lshortfile)
	testResponse = day0501(TEST_INPUT)
	if testResponse != TEST_EXPECTED_1 {
		errMsg = fmt.Sprintf("Day 05 01 test fails: received '%s' instead of '%s'", testResponse, TEST_EXPECTED_1)
		log.Fatal(errMsg)
	}
	testResponse = day0502(TEST_INPUT)
	if testResponse != TEST_EXPECTED_2 {
		errMsg = fmt.Sprintf("Day 05 02 test fails: received '%s' instead of '%s'", testResponse, TEST_EXPECTED_2)
		log.Fatal(errMsg)
	}
	input := common.ReadFile(INPUT_FILENAME)
	fmt.Println("Day 05 01:", day0501(input))
	fmt.Println("Day 05 02:", day0502(input))
}
