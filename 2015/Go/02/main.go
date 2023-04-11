package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day0201(input string) (sqft int) {
	input = strings.TrimRight(input, "\n")
	dims := make([]int, 3)
	for _, line := range strings.Split(input, "\n") {
		for i, dim := range strings.Split(line, "x") {
			v, _ := strconv.Atoi(dim)
			dims[i] = v
		}
		sort.Ints(dims)
		sqft += 2*dims[0]*dims[1] + 2*dims[1]*dims[2] + 2*dims[0]*dims[2]
		sqft += dims[0] * dims[1]
	}
	return
}

func Day0202(input string) (feet int) {
	input = strings.TrimRight(input, "\n")
	dims := make([]int, 3)
	for _, line := range strings.Split(input, "\n") {
		for i, dim := range strings.Split(line, "x") {
			v, _ := strconv.Atoi(dim)
			dims[i] = v
		}
		sort.Ints(dims)
		feet += 2*dims[0] + 2*dims[1]
		feet += dims[0] * dims[1] * dims[2]
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
		received := Day0201(test)
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
		received := Day0202(test)
		log.Printf("test=\"%s\", expected=%d, received=%d\n", test, expected, received)
		if expected != received {
			return false
		}
	}
	return true
}

func main() {
	RunTests()
	content, _ := os.ReadFile("input_day02.txt")
	sqft := Day0201(string(content))
	log.Println(sqft)
	feet := Day0202(string(content))
	log.Println(feet)
}
