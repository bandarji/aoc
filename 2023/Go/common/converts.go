package common

import (
	"log"
	"strconv"
)

func StrToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Could not convert '%s' to an integer", s)
	}
	return n
}

func NumbersInString(s string) []int {
	numbers := []int{}
	numS := ""
	for _, char := range s {
		if (char >= '0' && char <= '9') || char == '-' {
			numS += string(char)
		} else {
			if len(numS) > 0 {
				numbers = append(numbers, StrToInt(numS))
			}
			numS = ""
		}
	}
	return numbers
}
