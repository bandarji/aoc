package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day02(input string, part int) (feet int) {
	input = strings.TrimRight(input, "\n")
	dims := make([]int, 3)
	for _, line := range strings.Split(input, "\n") {
		for i, dim := range strings.Split(line, "x") {
			value, _ := strconv.Atoi(dim)
			dims[i] = value
		}
		sort.Ints(dims)
		if part == 1 {
			feet += 2*dims[0]*dims[1] + 2*dims[1]*dims[2] + 2*dims[0]*dims[2]
			feet += dims[0] * dims[1]
		} else {
			feet += 2*dims[0] + 2*dims[1]
			feet += dims[0] * dims[1] * dims[2]
		}
	}
	return
}

func RunTests() bool {
	tests := map[string]int{
		"2x3x4":           58,
		"1x1x10":          43,
		"2x3x4\n1x1x10\n": 58 + 43,
	}
	for test, expected := range tests {
		received := Day02(test, 1)
		log.Printf("test=\"%s\", expected=%d, received=%d\n", test, expected, received)
		if expected != received {
			return false
		}
	}
	tests = map[string]int{
		"2x3x4":           34,
		"1x1x10":          14,
		"2x3x4\n1x1x10\n": 34 + 14,
	}
	for test, expected := range tests {
		received := Day02(test, 2)
		log.Printf("test=\"%s\", expected=%d, received=%d\n", test, expected, received)
		if expected != received {
			return false
		}
	}
	return true
}

func main() {
	day, part := 2, 1
	RunTests()
	content, _ := os.ReadFile(fmt.Sprintf("input_day%02d.txt", day))
	sqft := Day02(string(content), part)
	log.Printf("[%02d / %02d] %d\n", day, part, sqft)
	part = 2
	feet := Day02(string(content), part)
	log.Printf("[%02d / %02d] %d\n", day, part, feet)
}
