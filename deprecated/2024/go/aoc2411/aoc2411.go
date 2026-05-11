package aoc2411

import (
	"aoc24/aoc2400"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const AOC2411_TEST string = `125 17`

func Blink182(stones []int) (stones2 []int) {
	for _, stone := range stones {
		digits := strconv.Itoa(stone)
		if stone == 0 {
			stones2 = append(stones2, 1)
		} else if len(digits)%2 == 0 {
			stones2 = append(stones2, aoc2400.StrToInt(digits[:len(digits)/2]))
			stones2 = append(stones2, aoc2400.StrToInt(digits[len(digits)/2:]))
		} else {
			stones2 = append(stones2, stone*2024)
		}
	}
	return
}

func Blink(stones map[int]int, cycles int) (total int) {
	for _ = range cycles {
		nStones := map[int]int{}
		for stone, count := range stones {
			if stone == 0 {
				nStones[1] += count
				continue
			}
			digits := fmt.Sprintf("%d", stone)
			digitsCount := len(digits)
			if digitsCount%2 == 0 {
				nStones[aoc2400.StrToInt(digits[:digitsCount/2])] += count
				nStones[aoc2400.StrToInt(digits[digitsCount/2:])] += count
			} else {
				nStones[stone*2024] += count
			}
			stones = nStones
		}
	}
	for _, count := range stones {
		total += count
	}
	return
}

func readInput(input string) (stones map[int]int) {
	stones = map[int]int{}
	for _, stone := range strings.Fields(input) {
		stones[aoc2400.StrToInt(stone)]++
	}
	return
}

func Aoc2411(input string, cycles int) (count int) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	stones := readInput(input)
	count = Blink(stones, cycles)
	return
}
