package adventofcode

import (
	"strings"
)

const y16d16Input string = "01111001100111011"
const y16d16DiskLengthPart1 int = 272
const y16d16DiskLengthPart2 int = 35651584

func y16d16(input string, diskLength int) (answer string) {
	for len(input) < diskLength {
		input = y16d16ProcessStep(input)
	}
	answer = y16d16Checksum(input[0:diskLength])
	return
}

func y16d16Checksum(input string) string {
	for len(input)%2 == 0 {
		sb := strings.Builder{}
		for i := 0; i < len(input); i += 2 {
			if input[i] == input[i+1] {
				sb.WriteByte('1')
			} else {
				sb.WriteByte('0')
			}
		}
		input = sb.String()
	}
	return input
}

func y16d16ProcessStep(a string) string {
	sb := strings.Builder{}
	b := y16d16Reverse(a)
	b = y16d16ReplaceBits(b)
	sb.WriteString(a)
	sb.WriteByte('0')
	sb.WriteString(b)
	return sb.String()
}

func y16d16Reverse(a string) string {
	sb := strings.Builder{}
	for i := len(a) - 1; i >= 0; i-- {
		sb.WriteByte(a[i])
	}
	return sb.String()
}

func y16d16ReplaceBits(b string) string {
	sb := strings.Builder{}
	for _, char := range b {
		if char == '0' {
			sb.WriteByte('1')
		} else {
			sb.WriteByte('0')
		}
	}
	return sb.String()
}
