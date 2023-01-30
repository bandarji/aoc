package main

import (
	"aoc/common"
	"fmt"
	"log"
	"strings"
)

const INPUT_FILENAME string = "input_day04.txt"
const TEST_INPUT string = "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8"
const TEST_EXPECTED1 int = 2
const TEST_EXPECTED2 int = 4

func runTests() {
	var response int
	var errMsg string
	response = partOne(TEST_INPUT)
	if response != TEST_EXPECTED1 {
		errMsg = fmt.Sprintf("Day 04 01 test fails: got %d, expected %d", response, TEST_EXPECTED1)
		log.Fatal(errMsg)
	}
	response = partTwo(TEST_INPUT)
	if response != TEST_EXPECTED2 {
		errMsg = fmt.Sprintf("Day 04 02 test fails: got %d, expected %d", response, TEST_EXPECTED2)
		log.Fatal(errMsg)
	}
}

func partOne(input string) int {
	pairs := 0
	for _, entry := range strings.Split(input, "\n") {
		if isOverlapping(entry) {
			pairs++
		}
	}
	return pairs
}

func isOverlapping(entry string) bool {
	if !strings.Contains(entry, "-") {
		return false
	}
	e1, e2 := splitByString(entry, ",")
	e1_s, e1_e := splitByString(e1, "-")
	e2_s, e2_e := splitByString(e2, "-")
	if (e1_s <= e2_s && e1_e >= e2_e) || (e2_s <= e1_s && e2_e >= e1_e) {
		return true
	}
	return false
}

func partTwo(input string) int {
	pairs := 0
	for _, entry := range strings.Split(input, "\n") {
		if isOverlappingAtAll(entry) {
			pairs++
		}
	}
	return pairs
}

func isOverlappingAtAll(entry string) bool {
	if !strings.Contains(entry, "-") {
		return false
	}
	e1, e2 := splitByString(entry, ",")
	e1_s, e1_e := splitByString(e1, "-")
	e2_s, e2_e := splitByString(e2, "-")
	if (e2_s <= e1_s && e1_s <= e2_e) || (e1_s <= e2_s && e2_s <= e1_e) {
		return true
	}
	return false
}

func splitByString(s string, sp string) (string, string) {
	split := strings.Split(s, sp)
	return split[0], split[1]
}

func main() {
	input := common.ReadFile(INPUT_FILENAME)
	runTests()
	fmt.Println("Day 04 01:", partOne(input))
	fmt.Println("Day 04 02:", partTwo(input))
}
