package d01

import (
	"aoc2023/common"
	"fmt"
	"strings"
)

const (
	TEST0101 string = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	TEST0102 string = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
)

func BuildCalibration1(s string) (value int) {
	values := []rune{}
	for _, char := range s {
		if char >= '0' && char <= '9' {
			values = append(values, char)
		}
	}
	value = common.StrToInt(fmt.Sprintf("%c%c", values[0], values[len(values)-1]))
	return
}

func BuildCalibration2(s string) (total int) {
	values := []int{}
	wordNums := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	value := 0
	length := len(s)
	for cursor := 0; cursor < length; cursor++ {
		value = 0
		char := rune(s[cursor])
		if char >= '0' && char <= '9' {
			value = common.StrToInt(string(char))
		}
		for i, wordNum := range wordNums {
			if cursor+len(wordNum) <= length {
				if s[cursor:cursor+len(wordNum)] == wordNum {
					value = i + 1
					break
				}
			}
		}
		if value > 0 {
			values = append(values, value)
		}
	}
	total = common.StrToInt(fmt.Sprintf("%d%d", values[0], values[len(values)-1]))
	return
}

func Calibrate(s string, part int) (total int) {
	for _, line := range strings.Split(s, "\n") {
		if part == 1 {
			total += BuildCalibration1(line)
		} else {
			total += BuildCalibration2(line)
		}
	}
	return
}
