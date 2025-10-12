package adventofcode

import (
	"strings"
)

func y15d08(input string, part int) int {
	if part == 1 {
		return y15d08p1(input)
	}
	return y15d08p2(input)
}

func y15d08p1(input string) int {
	tChars, sChars := 0, 0
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		tChars += len(line)
		for i := 1; i < len(line)-1; i++ {
			switch line[i] {
			case 92: // backslash
				nChar := line[i+1]
				if nChar == 92 || nChar == 34 { // backslash or double quote
					i++
				} else if nChar == 120 { // x
					i += 3
				}
			}
			sChars++
		}
	}
	return tChars - sChars
}

func y15d08p2(input string) int {
	eChars, oChars := 0, 0
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		oChars += len(line)
		eChars += 2
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case 34, 92:
				eChars += 2
			default:
				eChars++
			}
		}
	}
	return eChars - oChars
}
