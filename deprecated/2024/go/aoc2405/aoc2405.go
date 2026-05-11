package aoc2405

import (
	"aoc24/aoc2400"
	"fmt"
	"slices"
	"strings"
)

const AOC2405_TEST string = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func checkOrder(rules map[int]map[int]interface{}, line string) int {
	pages := []int{}
	pageInputs := strings.Split(line, ",")
	for _, pageInput := range pageInputs {
		pages = append(pages, aoc2400.StrToInt(pageInput))
	}
	for i, page := range pages {
		for j := i + 1; j < len(pages); j++ {
			if _, ok := rules[pages[j]][page]; ok {
				return 0
			}
		}
	}
	return pages[len(pages)/2]
}

func Aoc240501(input string) (total int) {
	rules := map[int]map[int]interface{}{}
	var (
		pl int
		pr int
	)
	sections := strings.Split(input, "\n\n")
	for _, ruleline := range strings.Split(sections[0], "\n") {
		fmt.Sscanf(ruleline, "%d|%d", &pl, &pr)
		if _, ok := rules[pl]; !ok {
			rules[pl] = map[int]interface{}{}
		}
		rules[pl][pr] = struct{}{}
	}
	for _, updateline := range strings.Split(sections[1], "\n") {
		total += checkOrder(rules, updateline)
	}
	return
}

func Aoc240502(input string) (total int) {
	rules := map[int]map[int]interface{}{}
	var (
		pl int
		pr int
	)
	sections := strings.Split(input, "\n\n")
	for _, ruleline := range strings.Split(sections[0], "\n") {
		fmt.Sscanf(ruleline, "%d|%d", &pl, &pr)
		if _, ok := rules[pl]; !ok {
			rules[pl] = map[int]interface{}{}
		}
		rules[pl][pr] = struct{}{}
	}
	for _, updateline := range strings.Split(sections[1], "\n") {
		mid := checkOrder(rules, updateline)
		if mid != 0 {
			continue
		}
		pagesInput := strings.Split(updateline, ",")
		slices.SortFunc(pagesInput, func(x, y string) int {
			i, j := aoc2400.StrToInt(x), aoc2400.StrToInt(y)
			if _, ok := rules[i][j]; ok {
				return -1
			}
			if _, ok := rules[j][i]; !ok {
				return 1
			}
			return 0
		})
		total += aoc2400.StrToInt(pagesInput[len(pagesInput)/2])
	}
	return
}
