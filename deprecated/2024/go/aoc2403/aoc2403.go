package aoc2403

import (
	"aoc24/aoc2400"
	"regexp"
	"strings"
)

const AOC2403_TEST string = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
const AOC2403_TEST_2 string = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func Aoc240301(input string) (total int) {
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, line := range strings.Split(input, "\n") {
		matches := pattern.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			total += aoc2400.StrToInt(match[1]) * aoc2400.StrToInt(match[2])
		}
	}
	return
}

func Aoc240302(input string) (total int) {
	enabled := true
	pattern := regexp.MustCompile(`(?:(?:do(?:n't)?\(\))|(?:mul\((\d+),(\d+)\)))`)
	for _, line := range strings.Split(input, "\n") {
		matches := pattern.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				enabled = true
			} else if match[0] == `don't()` {
				enabled = false
			} else if enabled {
				total += aoc2400.StrToInt(match[1]) * aoc2400.StrToInt(match[2])
			}
		}
	}
	return
}
