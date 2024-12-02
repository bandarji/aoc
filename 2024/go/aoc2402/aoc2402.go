package aoc2402

import (
	"aoc24/aoc2400"
	"strings"
)

const AOC2402_TEST string = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func checkDifferent(values []int) bool {
	for i, val := range values {
		if i == 0 {
			continue
		}
		if val == values[i-1] {
			return false
		}
	}
	return true
}

func checkDeltas(values []int) bool {
	for i, val := range values {
		if i == 0 {
			continue
		}
		prev := values[i-1]
		if val > prev {
			if val-prev > 3 {
				return false
			}
		} else {
			if prev-val > 3 {
				return false
			}
		}
	}
	return true
}

func checkLinear(values []int) bool {
	if checkAscending(values) || checkDescending(values) {
		return true
	}
	return false
}

func checkAscending(values []int) bool {
	for i, val := range values {
		if i == 0 {
			continue
		}
		if val < values[i-1] {
			// slog.Info("checkAscending", "val", val, "values[i-1]", values[i-1], "values", values)
			return false
		}
	}
	return true
}

func checkDescending(values []int) bool {
	for i, val := range values {
		if i == 0 {
			continue
		}
		if val > values[i-1] {
			// slog.Info("checkDescending", "val", val, "values[i-1]", values[i-1], "values", values)
			return false
		}
	}
	return true
}

func Aoc240201(input string) (safe int) {
	values := []int{}
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		for _, field := range fields {
			values = append(values, aoc2400.StrToInt(field))
		}
		if checkLinear(values) && checkDifferent(values) && checkDeltas(values) {
			safe++
		}
		values = []int{}
	}
	return
}

func Aoc240202(input string) (safe int) {
	values := []int{}
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		for _, field := range fields {
			values = append(values, aoc2400.StrToInt(field))
		}
		if checkLinear(values) && checkDifferent(values) && checkDeltas(values) {
			safe++
		} else if safeWithDampening(values) {
			safe++
		}
		values = []int{}
	}
	return
}

func safeWithDampening(values []int) bool {
	for removalIndex := range len(values) {
		newValues := []int{}
		for i, val := range values {
			if i == removalIndex {
				continue
			}
			newValues = append(newValues, val)
		}
		if checkLinear(newValues) && checkDifferent(newValues) && checkDeltas(newValues) {
			return true
		}
	}
	return false
}
