package adventofcode

import (
	"strings"
)

func y17d10Modulo(number int, length int) int {
	if number < 0 {
		return y17d10Modulo(length+number, length)
	} else {
		return number % length
	}
}

func y17d10ReverseSlice(list []int, start int, stop int, length int) {
	size := len(list)
	dupe := make([]int, size)
	copy(dupe, list)
	for i := 0; i < length; i++ {
		i1 := y17d10Modulo(start+i, size)
		i2 := y17d10Modulo(stop-i, size)
		list[i1] = dupe[i2]
	}
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

func y17d10(input string, listLength, part int) (product int) {
	lengths := y17d10ParseInput(input)
	list := y17d10CreateList(listLength)
	y17d10KnotHash(list, lengths, 1)
	product = list[0] * list[1]
	return
}

func y17d10CreateList(listLength int) (list []int) {
	list = []int{}
	for i := 0; i < listLength; i++ {
		list = append(list, i)
	}
	return
}

func y17d10ParseInput(input string) (lengths []int) {
	for n := range strings.SplitSeq(input, ",") {
		lengths = append(lengths, strToInt(n))
	}
	return
}
