package adventofcode

import (
	"fmt"
	"strings"
)

type Y15D05 struct{}

func (d *Y15D05) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D05) Part1(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d", year, day, y15d05(d.GetInput(year, day), 1))
}

func (d *Y15D05) Part2(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d", year, day, y15d05(d.GetInput(year, day), 2))
}

func y15d05(input string, part int) (nice int) {
	input = strings.TrimRight(input, "\n")
	for s := range strings.SplitSeq(input, "\n") {
		if niceString(s, part) {
			nice += 1
		}
	}
	return
}

func niceString(s string, part int) bool {
	if part == 1 {
		if hasThreeVowels(s) && hasCharTwiceTogether(s) && !hasBadCombos(s) {
			return true
		}
	} else {
		if hasTwoTwice(s) && hasSpacedRepeat(s) {
			return true
		}
	}
	return false
}

func hasSpacedRepeat(s string) bool {
	for i := 2; i < len(s); i++ {
		if s[i-2] == s[i] {
			return true
		}
	}
	return false
}

func hasTwoTwice(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		match := s[i : i+2]
		for j := i + 2; j < len(s)-1; j++ {
			if s[j:j+2] == match {
				return true
			}
		}
	}
	return false
}

func hasBadCombos(s string) bool {
	badies := []string{"ab", "cd", "pq", "xy"}
	for i := 1; i < len(s); i++ {
		for _, bad := range badies {
			if fmt.Sprintf("%s%s", string(s[i-1]), string(s[i])) == bad {
				return true
			}
		}
	}
	return false
}

func hasCharTwiceTogether(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return true
		}
	}
	return false
}

func hasThreeVowels(s string) bool {
	count := 0
	for _, char := range s {
		for _, vowel := range VOWELS {
			if char == vowel {
				count++
			}
			if count >= 3 {
				return true
			}
		}
	}
	return false
}
