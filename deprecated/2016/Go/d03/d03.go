package d03

import (
	"strconv"
	"strings"
)

func GetSides(input string) (sides [3]int) {
	input = strings.Trim(input, " ")
	index := 0
	ns := ""
	for i, char := range input {
		if i > 0 {
			if input[i] == 32 && input[i-1] == 32 {
				continue
			}
		}
		if char >= 48 && char <= 57 {
			ns += string(char)
		}
		if char == 32 {
			value, _ := strconv.Atoi(ns)
			sides[index] = value
			ns = ""
			index++
		}
	}
	value, _ := strconv.Atoi(ns)
	sides[index] = value
	return
}

func IsValidTriangle(sides [3]int) bool {
	if sides[0]+sides[1] > sides[2] && sides[1]+sides[2] > sides[0] && sides[2]+sides[0] > sides[1] {
		return true
	}
	return false
}

func RebuildSides(allSides [][3]int) [][3]int {
	newSides := [][3]int{}
	for index := 0; index < len(allSides); index += 3 {
		for column := 0; column < 3; column++ {
			edges := [3]int{}
			for offset := 0; offset < 3; offset++ {
				edges[offset] = allSides[offset+index][column]
			}
			newSides = append(newSides, edges)
		}
	}
	return newSides
}

func Day(input string, part int) (count int) {
	allSides := [][3]int{}
	for _, entry := range strings.Split(input, "\n") {
		allSides = append(allSides, GetSides(entry))
	}
	if part == 2 {
		allSides = RebuildSides(allSides)
	}
	for _, sides := range allSides {
		if IsValidTriangle(sides) {
			count++
		}
	}
	return
}
