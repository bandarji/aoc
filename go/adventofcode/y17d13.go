package adventofcode

import (
	"fmt"
	"strings"
)

func y17d13(input string, part int) (answer int) {
	frequencies := [][2]int{}
	for line := range strings.SplitSeq(input, "\n") {
		var depth, range_ int
		fmt.Sscanf(line, "%d: %d", &depth, &range_)
		frequencies = append(frequencies, [2]int{depth, range_})
	}
	if part == 1 {
		answer = y17d13Part1(frequencies)
	} else {
		answer = y17d13Part2(frequencies)
	}
	return
}

func y17d13Part1(frequencies [][2]int) (severity int) {
	for _, frequency := range frequencies {
		depth, range_ := frequency[0], frequency[1]
		if depth%((range_-1)*2) == 0 {
			severity += depth * range_
		}
	}
	return
}

func y17d13Part2(frequencies [][2]int) int {
	for delay := range 1 << 31 {
		caught := false
		for _, frequency := range frequencies {
			depth, range_ := frequency[0], frequency[1]
			if (depth+delay)%((range_-1)*2) == 0 {
				caught = true
				break
			}
		}
		if !caught {
			return delay
		}
	}
	return -1
}
