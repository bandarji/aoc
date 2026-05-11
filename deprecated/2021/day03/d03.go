package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "strings"
)

func assembleEpsilon(diags []string) string {
	epsilon := ""

	for position, _ := range(diags[0]) {
		epsilon += findFrequentBit(diags, position, false)
	}

	return epsilon
}

func assembleGamma(diags []string) string {
	gamma := ""

	for position, _ := range(diags[0]) {
		gamma += findFrequentBit(diags, position, true)
	}

	return gamma
}

func findFrequentBit(diags []string, position int, most bool) string {
	var frequentBit string
	count_0 := 0
	count_1 := 0

	frequentBit = "1"

	for _, diag := range diags {
		if diag[position:position+1] == "0" {
			count_0++
		} else {
			count_1++
		}
	}

	if (count_0 > count_1  && most) || (count_1 > count_0 && !most) {
		frequentBit = "0"
	}

	return frequentBit
}

func read_diags(filename string) []string {
	var diags []string

	f, _ := os.Open(filename)
	defer f.Close()

	content := bufio.NewScanner(f)
	for content.Scan() {
		diags = append(diags, content.Text())
	}

	return diags
}


func patternConvert(bits string) int64 {
	if i, err := strconv.ParseInt(bits, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		return i
	}
	return -1
}


func main() {
	var product int64
	diags := read_diags("input.txt")
	gamma := patternConvert(assembleGamma(diags))
	epsilon := patternConvert(assembleEpsilon(diags))
	product = gamma * epsilon
	fmt.Println("answer:", product)
}
