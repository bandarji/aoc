package main

import (
	"fmt"
	"log"
	"strings"

	"aoc/common"
)

const INPUT_FILENAME string = "input_day02.txt"
const TEST_INPUT_0201 string = "A Y\nB X\nC Z"
const EXPECTED_0201 int = 15
const EXPECTED_0202 int = 12

func sumScoresByStrategy1(input string) int {
	points := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	plays := map[string]int{
		"AX": 3,
		"AY": 6,
		"AZ": 0,
		"BX": 0,
		"BY": 3,
		"BZ": 6,
		"CX": 6,
		"CY": 0,
		"CZ": 3,
	}
	total := 0
	for _, entry := range strings.Split(input, "\n") {
		if len(entry) > 1 {
			moves := strings.Split(entry, " ")
			total += points[moves[1]] + plays[strings.Join(moves, "")]
		}
	}
	return total
}

func sumScoresByStrategy2(input string) int {
	strategies := map[string]int{
		"AX": 3,
		"AY": 1,
		"AZ": 2,
		"BX": 1,
		"BY": 2,
		"BZ": 3,
		"CX": 2,
		"CY": 3,
		"CZ": 1,
	}
	total := 0
	play := 0
	for _, entry := range strings.Split(input, "\n") {
		if len(entry) > 1 {
			moves := strings.Split(entry, " ")
			switch moves[1] {
			case "Z":
				play = 6
			case "Y":
				play = 3
			case "X":
				play = 0
			}
			total += play + strategies[strings.Join(moves, "")]
		}
	}
	return total
}

func main() {
	input := common.ReadFile(INPUT_FILENAME)
	test_response := sumScoresByStrategy1(TEST_INPUT_0201)
	if test_response != EXPECTED_0201 {
		errMsg := fmt.Sprintf("Day 02 Part 01 tests fail, Received %d instead of %d", test_response, EXPECTED_0201)
		log.Fatal(errMsg)
	}
	test_response = sumScoresByStrategy2(TEST_INPUT_0201)
	if test_response != EXPECTED_0202 {
		errMsg := fmt.Sprintf("Day 02 Part 02 tests fail, Received %d instead of %d", test_response, EXPECTED_0202)
		log.Fatal(errMsg)
	}
	fmt.Println("Day 02 01:", sumScoresByStrategy1(input))
	fmt.Println("Day 02 02:", sumScoresByStrategy2(input))
}
