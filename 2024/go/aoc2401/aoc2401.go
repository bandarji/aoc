package aoc2401

import (
	"aoc24/aoc2400"
	_ "embed"
	"fmt"
	"log/slog"
	"sort"
	"strings"
)

const AOC2401 bool = true

const AOC2401_TEST string = `3   4
4   3
2   5
1   3
3   9
3   3`

func Aoc240101(input string) (total int) {
	left := []int{}
	right := []int{}
	var distance int
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		lv := aoc2400.StrToInt(fields[0])
		rv := aoc2400.StrToInt(fields[1])
		left = append(left, lv)
		right = append(right, rv)
	}
	sort.Ints(left)
	sort.Ints(right)
	fmt.Println(left, right)
	for i, lv := range left {
		rv := right[i]
		if lv > rv {
			distance = lv - rv
		} else {
			distance = rv - lv
		}
		total += distance
		slog.Info("2401.1_test", "left", lv, "right", rv, "distance", distance, "total", total)
	}
	return
}

func Aoc240102(input string) (total int) {
	left := []int{}
	right := []int{}
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		lv := aoc2400.StrToInt(fields[0])
		rv := aoc2400.StrToInt(fields[1])
		left = append(left, lv)
		right = append(right, rv)
	}
	for _, lv := range left {
		score := 0
		for _, rv := range right {
			if rv == lv {
				score++
			}
		}
		total += lv * score
	}
	return
}
