package adventofcode

import (
	"strings"
)

func y15d17(input string, liters int, part int) (answer int) {
	containers := y15d17GetContainers(input)
	combinations := y15d17GetCombinations(containers, 0, liters, []int{})
	if part == 1 {
		answer = len(combinations)
	} else {
		minimum := 0xffff0000
		for _, combination := range combinations {
			minimum = min(minimum, len(combination))
		}
		count := 0
		for _, combination := range combinations {
			if len(combination) == minimum {
				count++
			}
		}
		answer = count
	}
	return
}

func y15d17GetContainers(input string) (containers []int) {
	for _, line := range strings.Split(input, "\n") {
		containers = append(containers, strToInt(line))
	}
	return
}

func y15d17GetCombinations(containers []int, start int, remaining int, known []int) [][]int {
	if remaining == 0 {
		return [][]int{append([]int{}, known...)}
	}
	if remaining < 0 {
		return nil
	}
	var goodCombinations [][]int
	for i := start; i < len(containers); i++ {
		known = append(known, i)
		goodCombinations = append(goodCombinations, y15d17GetCombinations(containers, i+1, remaining-containers[i], known)...)
		known = known[:len(known)-1]
	}
	return goodCombinations

}
