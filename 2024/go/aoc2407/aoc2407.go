package aoc2407

import (
	"aoc24/aoc2400"
	"fmt"
	"strings"
)

const AOC2407_TEST string = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func canSolve(part2 bool, test int, values []int) bool {
	var op func(int, int) bool
	op = func(index, total int) bool {
		if index == len(values) {
			return total == test
		}
		if total > test {
			return false
		}
		catSolve := false
		if part2 {
			catSolve = op(index+1, aoc2400.StrToInt(fmt.Sprintf("%d%d", total, values[index])))
		}
		return op(index+1, total+values[index]) || op(index+1, total*values[index]) || catSolve
	}
	return op(1, values[0])
}

func parseLine(line string) (test int, values []int) {
	sections := strings.Split(line, ": ")
	test = aoc2400.StrToInt(sections[0])
	for _, v := range strings.Fields(sections[1]) {
		values = append(values, aoc2400.StrToInt(v))
	}
	return
}

func Aoc240701(input string) (ans int) {
	for _, line := range strings.Split(input, "\n") {
		test, values := parseLine(line)
		if canSolve(false, test, values) {
			ans += test
		}
	}
	return
}

func Aoc240702(input string) (ans int) {
	for _, line := range strings.Split(input, "\n") {
		test, values := parseLine(line)
		if canSolve(true, test, values) {
			ans += test
		}
	}
	return
}
