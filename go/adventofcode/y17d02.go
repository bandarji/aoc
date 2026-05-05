package adventofcode

import "strings"

func y17d02(input string, part int) (checksum int) {
	for _, row := range strings.Split(input, "\n") {
		if part == 1 {
			checksum += y17d02ProcessRow1(row)
		} else {
			checksum += y17d02ProcessRow2(row)
		}
	}
	return
}

func y17d02ProcessRow1(row string) int {
	lo, hi := 0, 0
	fields := strings.Fields(row)
	for i, field := range fields {
		n := strToInt(field)
		if i == 0 {
			lo = n
			hi = n
		} else {
			if n < lo {
				lo = n
			}
			if n > hi {
				hi = n
			}
		}
	}
	return hi - lo
}

func y17d02ProcessRow2(row string) int {
	fields := strings.Fields(row)
	nums := []int{}
	for _, field := range fields {
		nums = append(nums, strToInt(field))
	}
	for _, n1 := range nums {
		for _, n2 := range nums {
			if n1 != n2 && n1%n2 == 0 {
				return n1 / n2
			}
		}
	}
	return 0
}
