package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const INPUT_FILENAME string = "input_day13.txt"
const TEST_ANS1 int = 13
const TEST_ANS2 int = 140
const TEST_INPUT string = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

func Day1301(input string) (answer int) {
	for i, pair := range assemblePairs(input) {
		left, right := pair[0], pair[1]
		if inOrder(left, right) {
			answer += i + 1
		}
	}
	return
}

func Day1302(input string) (answer int) {
	packets := [][]interface{}{
		{[]interface{}{(float64(2))}},
		{[]interface{}{(float64(6))}},
	}
	for _, pair := range assemblePairs(input) {
		packets = append(packets, pair[0])
		packets = append(packets, pair[1])
	}
	sort.Slice(packets, func(a, b int) bool {
		return inOrder(packets[a], packets[b])
	})
	answer = 1
	for i, packet := range packets {
		if fmt.Sprint(packet) == "[[2]]" || fmt.Sprint(packet) == "[[6]]" {
			answer *= i + 1
		}
	}
	return
}

func inOrder(left, right []interface{}) bool {
	for e := 0; e < len(left); e++ {
		if e > len(right)-1 {
			return false
		}
		numLeft, isNumLeft := left[e].(float64)
		numRight, isNumRight := right[e].(float64)
		sliceLeft, isSliceLeft := left[e].([]interface{})
		sliceRight, isSliceRight := right[e].([]interface{})
		if isNumLeft && isNumRight {
			if numLeft != numRight {
				return numLeft < numRight
			}
		} else if isNumLeft || isNumRight {
			if isNumLeft {
				sliceLeft = []interface{}{numLeft}
			} else if isNumRight {
				sliceRight = []interface{}{numRight}
			} else {
				log.Println("Expected a number here: ", left[e], right[e])
				log.Fatal("Exiting.")
			}
			return inOrder(sliceLeft, sliceRight)
		} else {
			if !isSliceLeft || !isSliceRight {
				log.Println("Expected two sequences: ", left[e], right[e])
				log.Fatal("Exiting.")
			}
			return inOrder(sliceLeft, sliceRight)
		}
	}
	return true
}

func assemblePairs(input string) (pairs [][2][]interface{}) {
	input = strings.TrimRight(input, "\n")
	for _, lines := range strings.Split(input, "\n\n") {
		pairings := strings.Split(lines, "\n")
		pairs = append(pairs, [2][]interface{}{
			structBuild(pairings[0]),
			structBuild(pairings[1]),
		})
	}
	return
}

func structBuild(s string) []interface{} {
	structure := []interface{}{}
	json.Unmarshal([]byte(s), &structure)
	return structure
}

func RunTests(day int) {
	result := 0
	if day == 1 {
		result = Day1301(TEST_INPUT)
		if result != TEST_ANS1 {
			log.Fatal(fmt.Sprintf("Day 13 01 test failed: %d instead of %d.", result, TEST_ANS1))
		} else {
			fmt.Println("Day 13 01 test passed.")
		}
	}
	if day == 2 {
		result = Day1302(TEST_INPUT)
		if result != TEST_ANS2 {
			log.Fatal(fmt.Sprintf("Day 13 02 test failed: %d instead of %d.", result, TEST_ANS2))
		} else {
			fmt.Println("Day 13 02 test passed.")
		}
	}
}

func main() {
	log.Println("Day 13.")
	RunTests(1)
	RunTests(2)
	input, _ := os.ReadFile(INPUT_FILENAME)
	fmt.Println("Day 13 01:", Day1301(string(input)))
	fmt.Println("Day 13 02:", Day1302(string(input)))
}
