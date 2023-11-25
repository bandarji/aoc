package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func JSONValues(input string) (total int) {
	num := 0
	s := strings.Builder{}
	for _, char := range input {
		if char == '\n' {
			continue
		}
		if char == '-' || (char >= '0' && char <= '9') {
			s.WriteRune(char)
		} else {
			if s.Len() > 0 {
				num, _ = strconv.Atoi(s.String())
				total += num
				s.Reset()
			}
		}
	}
	if s.Len() > 0 {
		num, _ := strconv.Atoi(s.String())
		total += num
	}
	return
}

func RunTests() bool {
	tests := map[string]int{
		"[1,2,3]":              6,
		`{"a":2,"b":4}`:        6,
		"[[[3]]]":              3,
		`{"a":{"b":4},"c":-1}`: 3,
		`{"a":[-1,1]}`:         0,
		`[-1,{"a":1}]`:         0,
		"[]":                   0,
		"{}":                   0,
	}
	for test, expected := range tests {
		result := JSONValues(test)
		if result != expected {
			return false
		}
	}
	return true
}

func main() {
	day, part := 12, 1
	_ = day
	_ = part
	if !RunTests() {
		log.Fatal("Tests failed.")
	}
	content, _ := os.ReadFile("input_day12.txt")
	fmt.Println("day", day, "part", part, "ans =", JSONValues(string(content)))
}
