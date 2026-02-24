package adventofcode

import (
	"fmt"
	"strings"
)

func y16d15(input string, part int) (answer int) {
	discs := y16d15ParseInput(input)
	if part == 2 {
		discs = append(discs, [3]int{len(discs) + 1, 11, 0})
	}
	t := 0
	for {
		collision := false
		for _, disc := range discs {
			p := (disc[2] + disc[0] + t) % disc[1]
			if p != 0 {
				collision = true
			}
		}
		if !collision {
			answer = t
			break
		}
		t++
	}
	return
}

func y16d15ParseInput(input string) (discs [][3]int) {
	for _, line := range strings.Split(input, "\n") {
		var id, positions, position int
		fmt.Sscanf(
			line,
			"Disc #%d has %d positions; at time=0, it is at position %d.",
			&id,
			&positions,
			&position,
		)
		discs = append(discs, [3]int{id, positions, position})
	}
	return
}
