package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Tests map[string]int

const VOWELS string = "aeiou"

func Day05(input string, part int) (nice int) {
	input = strings.TrimRight(input, "\n")
	for _, s := range strings.Split(input, "\n") {
		if NiceString(&s, part) {
			nice += 1
		}
	}
	return
}

func NiceString(s *string, part int) bool {
	if part == 1 {
		if HasThreeVowels(s) && HasCharTwiceTogether(s) && !HasBadCombos(s) {
			return true
		}
	} else {
		if HasNonOverlapDoubleDouble(s) && HasSpacedRepeat(s) {
			return true
		}
	}
	return false
}

func HasSpacedRepeat(s *string) bool {
	for i := 2; i < len(*s); i++ {
		if (*s)[i-2] == (*s)[i] {
			return true
		}
	}
	return false
}

func HasNonOverlapDoubleDouble(s *string) bool {
	count := 0
	for i := 2; i < len(*s); i++ {
		if (*s)[i-2] == (*s)[i-1] {
			if (*s)[i-1] != (*s)[i] {
				count++
				i++
				if count > 1 {
					return true
				}
			}
		}
	}
	return false
}

func HasBadCombos(s *string) bool {
	badies := []string{"ab", "cd", "pq", "xy"}
	for i := 1; i < len(*s); i++ {
		for _, bad := range badies {
			if fmt.Sprintf("%s%s", string((*s)[i-1]), string((*s)[i])) == bad {
				return true
			}
		}
	}
	return false
}

func HasCharTwiceTogether(s *string) bool {
	for i := 1; i < len(*s); i++ {
		if (*s)[i] == (*s)[i-1] {
			return true
		}
	}
	return false
}

func HasThreeVowels(s *string) bool {
	count := 0
	for _, char := range *s {
		for _, vowel := range VOWELS {
			if char == vowel {
				count++
			}
			if count >= 3 {
				return true
			}
		}
	}
	return false
}

func RunTests() bool {
	part := 1
	tests := Tests{
		"ugknbfddgicrmopn": 1,
		"aaa":              1,
		"jchzalrnumimnmhp": 0,
		"haegwjzuvuyypxyu": 0,
		"dvszwmarrgswjxmb": 0,
	}
	for test, expected := range tests {
		received := Day05(test, part)
		log.Printf("test=\"%s\", e=%d, r=%d\n", test, expected, received)
		if received != expected {
			return false
		}
	}
	log.Println("First tests pass.")
	part = 2
	tests = Tests{
		"qjhvhtzxzqqjkmpb": 1,
		"xxyxx":            1,
		"uurcxstgmygtbstg": 0,
		"ieodomkazucvgmuy": 0,
	}
	for test, expected := range tests {
		received := Day05(test, part)
		log.Printf("test=\"%s\", e=%d, r=%d\n", test, expected, received)
		if received != expected {
			return false
		}
	}
	log.Println("Second tests pass.")
	return true
}

func main() {
	RunTests()
	content, _ := os.ReadFile("input_day05.txt")
	log.Printf("[%02d / %02d]: %d\n", 5, 1, Day05(string(content), 1))
	log.Printf("[%02d / %02d]: %d\n", 5, 2, Day05(string(content), 2))
}
