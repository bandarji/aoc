package d13

import (
	"slices"
	"strings"
)

const TEST1 string = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func Part1(patterns [][][]string) (sum int) {
	for _, pattern := range patterns {
		vert, hori := Mirrors(pattern)
		sum += hori*100 + vert
	}
	return
}

func Part2(patterns [][][]string) (sum int) {
	for _, pattern := range patterns {
		vert, hori := Smudged(pattern)
		sum += hori*100 + vert
	}
	return
}

func Smudged(pattern [][]string) (vert, hori int) {
	changes := []int{}
	count := 0
	half := len(pattern) / 2
	for i := 0; i < len(pattern[0])-1; i++ {
		count = 0
		for x := i; x >= 0; x-- {
			if i+i+1-x > len(pattern[0])-1 {
				break
			}
			for y := 0; y < len(pattern); y++ {
				if pattern[y][x] != pattern[y][i+i+1-x] {
					count++
				}
			}
		}
		if count == 1 {
			changes = append(changes, i+1)
		}
	}
	vert = MinimumDistance(changes, half)
	changes = []int{}
	for i := 0; i < len(pattern)-1; i++ {
		count = 0
		for y := 1; y >= 0; y-- {
			if i+i+1-y > len(pattern)-1 {
				break
			}
			for x := 0; x < len(pattern[0]); x++ {
				if pattern[y][x] != pattern[i+i+1-y][x] {
					count++
				}
			}
		}
		if count == 1 {
			changes = append(changes, i+1)
		}
	}
	hori = MinimumDistance(changes, half)
	return
}

func MinimumDistance(sources []int, dist int) (minimum int) {
	shortest := 1<<63 - 1
	for _, source := range sources {
		distance := source - dist
		if distance < 0 {
			distance *= -1 // absolute
		}
		if distance < shortest {
			shortest = distance
			minimum = source
		}
	}
	return
}

func Mirrors(pattern [][]string) (vert, hori int) {
	matches := []int{}
	half := len(pattern) / 2
	match := false
	for i := 0; i < len(pattern)-1; i++ {
		match = true
		for y := i; y >= 0; y-- {
			if i+i+1-y > len(pattern)-1 {
				break
			}
			if !slices.Equal(pattern[y], pattern[i+i+1-y]) {
				match = false
				break
			}
		}
		if match {
			matches = append(matches, i+1)
		}
	}
	hori = MinimumDistance(matches, half)
	matches = []int{}
	for i := 0; i < len(pattern[0])-1; i++ {
		match = true
		for x := i; x >= 0; x-- {
			if i+i+1-x > len(pattern[0])-1 {
				break
			}
			for y := 0; y < len(pattern); y++ {
				if pattern[y][x] != pattern[y][i+i+1-x] {
					match = false
					break
				}
			}
		}
		if match {
			matches = append(matches, i+1)
		}
	}
	vert = MinimumDistance(matches, half)
	return
}

func Solve(input string, part int) (answer int) {
	patterns := [][][]string{}
	pattern := [][]string{}
	for _, block := range strings.Split(input, "\n\n") {
		for _, line := range strings.Split(block, "\n") {
			pattern = append(pattern, strings.Split(line, ""))
		}
		patterns = append(patterns, pattern)
		pattern = [][]string{}
	}
	if part == 1 {
		answer = Part1(patterns)
	} else {
		answer = Part2(patterns)
	}
	return
}
