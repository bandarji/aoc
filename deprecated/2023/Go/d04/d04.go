package d04

import (
	"strings"
)

const TEST1 string = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func ScratchCard(card string) (matches int) {
	winners := map[string]bool{}
	sections := strings.Split(card, ":")       // only want after the colon
	sections = strings.Split(sections[1], "|") // separate winners from my numbers
	for _, value := range strings.Fields(sections[0]) {
		winners[value] = true
	}
	for _, value := range strings.Fields(sections[1]) {
		if winners[value] {
			matches++
		}
	}
	return
}

func Part1(input string) (answer int) {
	var points int
	for _, card := range strings.Split(input, "\n") {
		matches := ScratchCard(card)
		points = 0
		if matches > 0 {
			points++
			for i := 1; i < matches; i++ {
				points *= 2
			}
		}
		answer += points
	}
	return
}

func Part2(input string) (answer int) {
	cards := strings.Split(input, "\n")
	counts := make([]int, len(cards))
	for i := 0; i < len(cards); i++ {
		counts[i] = 1
	}
	for i, card := range cards {
		matches := ScratchCard(card)
		if matches > 0 {
			for j := i + 1; j <= matches+i; j++ {
				counts[j] += counts[i]
			}
		}
	}
	for _, cardCount := range counts {
		answer += cardCount
	}
	return
}

func Solve(input string, part int) (answer int) {
	if part == 1 {
		answer = Part1(input)
	} else {
		answer = Part2(input)
	}
	return
}
