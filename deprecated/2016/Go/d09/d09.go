package d09

import (
	"fmt"
	"strings"
)

func Day(input string, part int) int {
	uncompressed, distance, repeat, pad := 0, 0, 0, 0
	for i := 0; i < len(input); {
		if input[i] == '(' {
			endMarker := strings.Index(input[i:], ")") + i
			fmt.Sscanf(input[i:endMarker+1], "(%dx%d)", &distance, &repeat)
			sub := input[endMarker+1 : endMarker+distance+1]
			if part == 1 {
				pad = len(sub)
			} else {
				pad = Day(sub, 2)
			}
			uncompressed += pad * repeat
			i = endMarker + pad + 1
		} else {
			uncompressed++
			i++
		}
	}
	return uncompressed
}

func RemoveNewlines(input string) (s string) {
	for i := 0; i < len(input); i++ {
		if input[i] != '\n' {
			s += string(input[i])
		}
	}
	return
}

func Day1(input string, part int) (size int) {
	var count, repeat, sublen int
	input = RemoveNewlines(input)
	for i := 0; i < len(input); {
		if input[i] == '(' {
			closer := strings.Index(input[i:], ")") + i
			fmt.Sscanf(input[i:closer+1], "(%dx%d)", &count, &repeat)
			sub := input[closer+1 : closer+count+1]
			sublen = len(sub)
			if sub[0] == '(' {
				size += sublen + closer
				i = closer + sublen + 1
			} else {
				if part == 2 {
					sublen = Day(sub, 2)
				}
				size += sublen * repeat
				i = closer + sublen + 1
			}
		} else {
			size++
			i++
		}
	}
	return
}
