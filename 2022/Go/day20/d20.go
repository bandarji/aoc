package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const TEST_INPUT string = `1
2
-3
3
-2
0
4`

type Solutions struct {
	Part1, Part2 int
}

const (
	INPUT     string = "input_day20.txt"
	TEST_ANS1 int    = 3
	TEST_ANS2 int    = 1623178306
)

func Day(input string, part int) (answer int) {
	numbers, key := GetNumbers(input, part)
	length := len(numbers)
	if part == 1 {
		numbers = Mix(numbers, key)
	} else {
		for i := 0; i < 10; i++ {
			numbers = Mix(numbers, key)
			log.Printf("After %d round of mixing...\n", i+1)
			DisplaySequence(numbers)
		}
	}
	index := FindValue(0, numbers)
	for i := 1; i <= 3; i++ {
		answer += numbers[(i*1000+index)%length]
	}
	return
}

func Mix(numbers, key []int) (sequence []int) {
	length := len(numbers)
	for i, number := range key {
		if i == 0 {
			sequence = append(sequence, numbers...)
		}
		if number < 0 {
			for i := number; i < 0; i++ {
				sequence = MoveBackward(number, sequence, length)
			}
		}
		if number > 0 {
			for i := number; i > 0; i-- {
				sequence = MoveForward(number, sequence, length)
			}
		}
	}
	return
}

func MoveForward(value int, sequence []int, length int) []int {
	newseq := make([]int, 0)
	position := FindValue(value, sequence)
	if position == length-1 {
		newseq = append(newseq, sequence[0])
		newseq = append(newseq, value)
		for i := 1; i < length-1; i++ {
			newseq = append(newseq, sequence[i])
		}
		return newseq
	} else {
		oldv, newv := sequence[position], sequence[position+1]
		sequence[position] = newv
		sequence[position+1] = oldv
		return sequence
	}
}

func MoveBackward(value int, sequence []int, length int) []int {
	newseq := make([]int, 0)
	position := FindValue(value, sequence)
	if position == 0 {
		for i := 1; i < length-1; i++ {
			newseq = append(newseq, sequence[i])
		}
		newseq = append(newseq, sequence[0])
		newseq = append(newseq, sequence[length-1])
		return newseq
	} else {
		oldv, newv := sequence[position], sequence[position-1]
		sequence[position] = newv
		sequence[position-1] = oldv
		return sequence
	}
}

func FindValue(value int, sequence []int) int {
	for i := 0; i < len(sequence); i++ {
		if sequence[i] == value {
			return i
		}
	}
	return -1
}

func DisplaySequence(sequence []int) {
	s := ""
	for i := 0; i < len(sequence)-1; i++ {
		s += fmt.Sprintf("%d, ", sequence[i])
	}
	s += fmt.Sprintf("%d\n", sequence[len(sequence)-1])
	fmt.Println(s)
}

func GetNumbers(input string, part int) (numbers, key []int) {
	multiplier := 1
	if part == 2 {
		multiplier = 811589153
	}
	input = strings.TrimRight(input, "\n")
	for _, n := range strings.Split(input, "\n") {
		value := atoi(n)
		numbers = append(numbers, multiplier*value)
		key = append(key, value)
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
	day, part := 20, 1
	if !RunTest(part) {
		log.Fatalf("Failed Day %d Part %02d", day, part)
	}
	content, _ := os.ReadFile(fmt.Sprintf("input_day%d.txt", day))
	log.Println(Day(string(content), 1))
	// part = 2
	// if !RunTest(part) {
	// 	log.Fatalf("Failed Day %d Part %02d", day, part)
	// }
}
