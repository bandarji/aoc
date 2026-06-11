package adventofcode

import (
	"fmt"
	"strings"
)

func y17d10(input string, listLength, part int) string {
	sb := strings.Builder{}
	lengths := y17d10ParseInput(input)
	if part == 2 {
		lengths = y17d10ParseInputPart2(input)
	}
	list := y17d10CreateList(listLength)
	if part == 1 {
		y17d10KnotHash(list, lengths, 1)
		sb.WriteString(fmt.Sprintf("%d", list[0]*list[1]))
	} else {
		y17d10KnotHash(list, lengths, 64)
		denseHash := make([]int, len(list)/16)
		for i := 0; i < len(list); i += 16 {
			for j := 0; j < 16; j++ {
				denseHash[i/16] ^= list[i+j]
			}
		}
		sb.WriteString(y17d10ToHex(denseHash))
	}
	return sb.String()
}

func y17d10CreateList(listLength int) (list []int) {
	list = []int{}
	for i := 0; i < listLength; i++ {
		list = append(list, i)
	}
	return
}

func y17d10KnotHash(list []int, lengths []int, cycles int) {
	start, skip, listSize := 0, 0, len(list)
	for i := 0; i < cycles; i++ {
		for _, length := range lengths {
			rStart := start % listSize
			rStop := (start + length - 1) % listSize
			y17d10ReverseSlice(list, rStart, rStop, length)
			start += length + skip
			skip++
		}
	}
}

func y17d10Modulo(number, length int) int {
	if number < 0 {
		return y17d10Modulo(length+number, length)
	} else {
		return number % length
	}
}

func y17d10ParseInput(input string) (lengths []int) {
	for n := range strings.SplitSeq(input, ",") {
		lengths = append(lengths, strToInt(n))
	}
	return
}

func y17d10ParseInputPart2(input string) (lengths []int) {
	for _, c := range input {
		lengths = append(lengths, int(c))
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)
	return
}

func y17d10ReverseSlice(list []int, start, stop, length int) {
	size := len(list)
	dupe := make([]int, size)
	copy(dupe, list)
	for i := 0; i < length; i++ {
		i1 := y17d10Modulo(start+i, size)
		i2 := y17d10Modulo(stop-i, size)
		list[i1] = dupe[i2]
	}
}

func y17d10ToHex(list []int) string {
	sb := strings.Builder{}
	for _, element := range list {
		sb.WriteString(fmt.Sprintf("%02x", element))
	}
	return sb.String()
}
