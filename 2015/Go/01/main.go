package main

import (
	"log"
	"os"
)

func Day0101(input string) (floor int) {
	for _, char := range input {
		if string(char) == "(" {
			floor += 1
		}
		if string(char) == ")" {
			floor -= 1
		}
	}
	return
}

func Day0102(input string) int {
	floor := 0
	for i, char := range input {
		if string(char) == "(" {
			floor += 1
		}
		if string(char) == ")" {
			floor -= 1
		}
		if floor == -1 {
			return i + 1
		}
	}
	return -1
}

func RunTests() bool {
	tests := map[string]int{
		"(())":    0,
		"()()":    0,
		"(((":     3,
		"(()(()(": 3,
		"))(((((": 3,
		"())":     -1,
		"))(":     -1,
		")))":     -3,
		")())())": -3,
	}
	for test, expected := range tests {
		received := Day0101(test)
		log.Printf("test=\"%s\", expected=%d, received=%d\n", test, expected, received)
		if expected != received {
			return false
		}
	}
	log.Println("Part One tests pass.")
	tests = map[string]int{
		")":     1,
		"()())": 5,
	}
	for test, expected := range tests {
		received := Day0102(test)
		log.Printf("test=\"%s\", expected=%d, received=%d\n", test, expected, received)
		if expected != received {
			return false
		}
	}
	log.Println("Part Two tests pass.")
	return true
}

func main() {
	RunTests()
	day, part := 1, 1
	content, _ := os.ReadFile("input_day01.txt")
	floor := Day0101(string(content))
	log.Printf("%02d / %02d: %d\n", day, part, floor)
	index := Day0102(string(content))
	log.Printf("%02d / %02d: %d\n", day, part, index)
}
