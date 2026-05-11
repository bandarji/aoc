package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Solutions struct {
	part1, part2 int
}

const TEST_INPUT string = `1
2
-3
3
-2
0
4`

const (
	INPUT     string = "input_day20.txt"
	TEST_ANS1 int    = 3
	TEST_ANS2 int    = 1623178306
)

func Day(input string, part int) (answer int) {
	numbers := [][2]int{}
	original := GetNumbers(input, part)
	numbers = append(numbers, original...)
	length := len(numbers)
	for t := 0; t < 10; t++ {
		log.Printf("Time %d - %v\n", t, numbers[:5])
		for _, element := range original {
			if element[1] < 0 {
				for v := 0; v < (element[1]*-1)%(length-1); v++ {
					numbers = MoveBack(numbers, length, element)
				}
			} else {
				for v := 0; v < element[1]%(length-1); v++ {
					numbers = MoveAhead(numbers, length, element)
				}
			}
		}
		if part == 1 && t == 0 {
			break
		}
	}
	zero := FindValue(numbers, 0)
	for i := 1; i <= 3; i++ {
		answer += numbers[(i*1000+zero)%length][1]
	}
	return
}

func FindValue(seq [][2]int, val int) int {
	for i, e := range seq {
		if e[1] == val {
			return i
		}
	}
	return -1
}

func MoveBack(seq [][2]int, length int, v [2]int) [][2]int {
	newseq := [][2]int{}
	p := FindElement(seq, v)
	if p == 0 {
		for i := 1; i < length-1; i++ {
			newseq = append(newseq, seq[i])
		}
		newseq = append(newseq, seq[0])
		newseq = append(newseq, seq[length-1])
		return newseq
	} else {
		ov, nv := seq[p], seq[p-1]
		seq[p] = nv
		seq[p-1] = ov
		return seq
	}
}

func MoveAhead(seq [][2]int, length int, v [2]int) [][2]int {
	newseq := [][2]int{}
	p := FindElement(seq, v)
	if p == length-1 {
		newseq = append(newseq, seq[0])
		newseq = append(newseq, v)
		for i := 1; i < length-1; i++ {
			newseq = append(newseq, seq[i])
		}
		return newseq
	} else {
		ov, nv := seq[p], seq[p+1]
		seq[p] = nv
		seq[p+1] = ov
		return seq
	}
}

func FindElement(numbers [][2]int, element [2]int) int {
	for i, e := range numbers {
		if e[0] == element[0] && e[1] == element[1] {
			return i
		}
	}
	return -1
}

func GetNumbers(input string, part int) (numbers [][2]int) {
	multiplier := 1
	if part == 2 {
		multiplier = 811589153
	}
	input = strings.TrimRight(input, "\n")
	for i, n := range strings.Split(input, "\n") {
		numbers = append(numbers, [2]int{i, multiplier * atoi(n)})
	}
	return
}

func atoi(s string) int {
	ns := ""
	for i := 0; i < len(s); i++ {
		if s[i] == 45 || (s[i] <= 57 && s[i] >= 48) {
			ns += string(s[i])
		}
	}
	num, _ := strconv.Atoi(ns)
	return num
}

func RunTest(part int) bool {
	if part == 1 {
		return Day(TEST_INPUT, part) == TEST_ANS1
	} else {
		return Day(TEST_INPUT, part) == TEST_ANS2
	}
}

func main() {
	solutions := Solutions{}
	day, part := 20, 1
	if !RunTest(part) {
		log.Fatalf("Failed day %02d Part %02d\n", day, part)
	}
	log.Printf("Test %02d Part %02d passed\n", day, part)
	part = 2
	if !RunTest(part) {
		log.Fatalf("Failed day %02d Part %02d\n", day, part)
	}
	log.Printf("Test %02d Part %02d passed\n", day, part)
	content, _ := os.ReadFile(fmt.Sprintf("input_day%02d.txt", day))
	log.Printf("%02d Part %02d starting...\n", day, 1)
	solutions.part1 = Day(string(content), 1)
	log.Printf("%02d Part %02d starting...\n", day, 1)
	solutions.part2 = Day(string(content), 2)
	log.Printf("Solutions: %d, %d\n", solutions.part1, solutions.part2)
}
