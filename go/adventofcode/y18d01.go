package adventofcode

import (
	"strings"
)

func y18d01(input string, part int) (answer int) {
	numbers := []int{}
	for n := range strings.SplitSeq(strings.ReplaceAll(input, "+", ""), "\n") {
		if n != "" {
			numbers = append(numbers, strToInt(n))
		}
	}
	switch part {
	case 1:
		for _, number := range numbers {
			answer += number
		}
	case 2:
		answer = y18d01Part2(numbers)
	}
	return
}

func y18d01Part2(numbers []int) (total int) {
	seen := map[int]bool{0: true}
	for {
		for _, number := range numbers {
			total += number
			if seen[total] {
				return total
			}
			seen[total] = true
		}
	}
}
