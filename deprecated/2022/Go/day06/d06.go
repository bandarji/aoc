package main

import (
	"aoc/common"
	"fmt"
	"log"
)

const INPUT_FILENAME string = "input_day06.txt"

func day06(code string, window int) int {
	index := -1
	for i := window; i < len(code); i++ {
		if isUnique(code[i-window:i], window) {
			index = i
			break
		}
	}
	return index
}

func isUnique(chars string, window int) bool {
	set := make(map[string]bool)
	for i := 0; i < len(chars); i++ {
		set[string(chars[i])] = true
	}
	if len(set) == window {
		return true
	}
	return false
}

func test_day0601() {
	tests := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	}
	for testCode, expected := range tests {
		result := day06(testCode, 4)
		if result != expected {
			errMsg := fmt.Sprintf("Test '%s' failed: got %d instead of %d", testCode, result, expected)
			log.Fatal(errMsg)
		}
	}
}

func test_day0602() {
	tests := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
		"nppdvjthqldpwncqszvftbrmjlhg":      23,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
	}
	for testCode, expected := range tests {
		result := day06(testCode, 14)
		if result != expected {
			errMsg := fmt.Sprintf("Test '%s' failed: got %d instead of %d", testCode, result, expected)
			log.Fatal(errMsg)
		}
	}
}

func main() {
	test_day0601()
	input := common.ReadFile(INPUT_FILENAME)
	fmt.Println("Day 06 01:", day06(input, 4))
	fmt.Println("Day 06 02:", day06(input, 14))
}
