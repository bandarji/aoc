package adventofcode

import (
	"fmt"
	"strings"
)

var y15d16Tape = map[string]int{
	"children:":    3,
	"cats:":        7,
	"samoyeds:":    2,
	"pomeranians:": 3,
	"akitas:":      0,
	"vizslas:":     0,
	"goldfish:":    5,
	"trees:":       3,
	"cars:":        2,
	"perfumes:":    1,
}

func y15d16(input string, part int) (answer int) {
	answer = y15d16FindSue(input, part)
	return
}

func y15d16FindSue(input string, part int) (rightSue int) {
	// Sue number, item number, item count
	var sue, c1, c2, c3 int
	var i1, i2, i3 string
	for _, line := range strings.Split(input, "\n") {
		fmt.Sscanf(line, "Sue %d: %s %d, %s %d, %s %d", &sue, &i1, &c1, &i2, &c2, &i3, &c3)
		suesThings := map[string]int{i1: c1, i2: c2, i3: c3}
		if y15d16SueFound(sue, suesThings, part) {
			rightSue = sue
			break
		}
	}
	return
}

func y15d16SueFound(sue int, suesThings map[string]int, part int) bool {
	if part == 1 {
		for item, count := range suesThings {
			if y15d16Tape[item] != count {
				return false
			}
		}
		return true
	} else {
		all := true

		// check less and drop wrong Sues
		for _, item := range []string{"cats:", "trees:"} {
			if sueCount, exists := suesThings[item]; exists {
				if sueCount <= y15d16Tape[item] {
					all = false
				}
				delete(suesThings, item)
			}
		}

		// check more and drop wrong Sues
		for _, item := range []string{"pomeranians:", "goldfish:"} {
			if sueCount, exists := suesThings[item]; exists {
				if sueCount >= y15d16Tape[item] {
					all = false
				}
				delete(suesThings, item)
			}
		}

		// check equal and find the right Sue
		for item, count := range suesThings {
			if y15d16Tape[item] != count {
				all = false
			}
		}
		if all {
			return true
		}
	}

	return false
}
