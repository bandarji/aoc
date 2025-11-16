package adventofcode

import (
	"fmt"
	"strings"
)

func y16d09(input string, part int) (decompressedLength int) {
	copyLength, copyCount, contentLength := 0, 0, 0
	for i := 0; i < len(input); {
		if input[i] == '(' {
			nextCloser := strings.Index(input[i:], ")")
			nextStart := nextCloser + i
			fmt.Sscanf(input[i:nextStart+1], "(%dx%d)", &copyLength, &copyCount)
			nextContent := input[nextStart+1 : nextStart+1+copyLength]
			contentLength = len(nextContent)
			if part == 2 {
				contentLength = y16d09(nextContent, 2)
			}
			decompressedLength += contentLength * copyCount
			i = nextStart + 1 + copyLength
		} else {
			decompressedLength++
			i++
		}
	}
	return
}
