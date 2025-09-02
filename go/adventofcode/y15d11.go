package adventofcode

import (
	"fmt"
	"strings"
)

const y15d11input string = "hepxcrrq"

type Y15D11 struct{}

func (d *Y15D11) GetInput(year, day int) string {
	return y15d11input
}

func (d *Y15D11) Part1(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %s", year, day, y15d11(d.GetInput(year, day), 1))
}

func (d *Y15D11) Part2(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %s", year, day, y15d11(d.GetInput(year, day), 2))
}

func y15d11(input string, part int) string {
	for part > 0 {
		input = y15d11IncrementPassword(input)
		if y15d11ValidPassword(input) {
			part--
		}
	}
	return input
}

func y15d11IncrementPassword(pw string) string {
	s := strings.Split(pw, "")
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == "z" {
			s[i] = "a"
		} else {
			s[i] = string(rune(s[i][0] + 1))
			break
		}
	}
	return strings.Join(s, "")
}

func y15d11ValidPassword(pw string) bool {
	return y15d11ValidChars(pw) && y15d11HasStraight(pw) && y15d11HasTwoPairs(pw)
}

func y15d11ValidChars(pw string) bool {
	return !strings.ContainsAny(pw, "iol")
}

func y15d11HasStraight(pw string) bool {
	for i := 0; i < len(pw)-2; i++ {
		if pw[i] == pw[i+1]-1 && pw[i] == pw[i+2]-2 {
			return true
		}
	}
	return false
}

func y15d11HasTwoPairs(pw string) bool {
	pairs := map[string]bool{}
	for i := 0; i < len(pw)-1; i++ {
		if pw[i] == pw[i+1] {
			pairs[pw[i:i+2]] = true
		}
	}
	return len(pairs) >= 2
}
