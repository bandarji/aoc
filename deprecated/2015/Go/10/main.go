package main

import (
	"fmt"
	"log"
)

const (
	INPUT string = "1113122113"
)

func LookAndSay(input string) (output string) {
	count := 0
	currChar := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] == currChar {
			count++
		} else {
			output += fmt.Sprintf("%d%s", count, string(currChar))
			currChar = input[i]
			count = 1
		}
	}
	output += fmt.Sprintf("%d%s", count, string(currChar))
	return
}

func RunTests() bool {
	tests := map[string]string{
		"1":      "11",
		"11":     "21",
		"21":     "1211",
		"1211":   "111221",
		"111221": "312211",
	}
	for test, expected := range tests {
		result := LookAndSay(test)
		log.Printf("test=%s, expected=%s, result=%s\n", test, expected, result)
		if result != expected {
			return false
		}
	}
	return true
}

func main() {
	day, part := 10, 1
	if !RunTests() {
		log.Println("Tests failed.")
	}
	input := INPUT
	for i := 0; i < 40; i++ {
		input = LookAndSay(input)
		// log.Printf("[%02d / %02d] iter=%d, len=%d\n", day, part, i+1, len(input))
	}
	log.Printf("[%02d / %02d]: %d\n", day, part, len(input))
	for i := 0; i < 10; i++ {
		input = LookAndSay(input)
	}
	part = 2
	log.Printf("[%02d / %02d]: %d\n", day, part, len(input))
	// log.Printf("[%02d / %02d]: %s\n", day, part, input)
}
