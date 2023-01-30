package main

import (
	"aoc/common"
	"fmt"
	"log"
	"strings"
)

const INPUT_FILENAME string = "input_day03.txt"
const TEST_INPUT string = "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw"
const TEST_EXPECTED1 int = 157
const TEST_EXPECTED2 int = 70
const ALPHABET string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func prioritySum(input string) int {
	total := 0
	for _, entry := range strings.Split(input, "\n") {
		if len(entry) > 1 {
			total += findPriority(entry)
		}
	}
	return total
}

func findPriority(entry string) int {
	return priorityValue(findCommonChar(entry))
}

func findCommonChar(entry string) string {
	charString := ""
	midPoint := len(entry) / 2
	left := entry[:midPoint]
	right := entry[midPoint:]
	for _, char := range left {
		charString = string(char)
		if strings.Contains(right, charString) {
			return charString
		}
	}
	return ""
}

func priorityValue(charString string) int {
	return strings.Index(ALPHABET, charString) + 1
}

func partTwo(input string) int {
	var elves [2]string
	total := 0
	for _, entry := range strings.Split(input, "\n") {
		if len(elves[0]) > 0 {
			if len(elves[1]) > 0 {
				total += findPriorityFromThree(elves[0], elves[1], entry)
				elves[0] = ""
				elves[1] = ""
			} else {
				elves[1] = entry
			}
		} else {
			elves[0] = entry
		}
	}
	return total
}

func findPriorityFromThree(e1, e2, e3 string) int {
	charString := ""
	for _, char := range e1 {
		charString = string(char)
		if strings.Contains(e2, charString) && strings.Contains(e3, charString) {
			return priorityValue(charString)
		}
	}
	return 0
}

func main() {
	input := common.ReadFile(INPUT_FILENAME)
	testResponse := prioritySum(TEST_INPUT)
	if testResponse != TEST_EXPECTED1 {
		errMsg := fmt.Sprintf("Day 03 Part 01 tests fail, Received %d instead of %d", testResponse, TEST_EXPECTED1)
		log.Fatal(errMsg)
	}
	testResponse = partTwo(TEST_INPUT)
	if testResponse != TEST_EXPECTED2 {
		errMsg := fmt.Sprintf("Day 03 Part 02 tests fail, Received %d instead of %d", testResponse, TEST_EXPECTED1)
		log.Fatal(errMsg)
	}
	fmt.Println("Day 03 01:", prioritySum(input))
	fmt.Println("Day 03 02:", partTwo(input))
}
