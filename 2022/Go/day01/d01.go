package main

import (
	"aoc/common"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

const INPUT_FILENAME string = "input_day01.txt"
const TEST_CALORIES string = "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000"
const TEST_EXPECTED int = 24000
const TEST_EXPECTED2 int = 45000

func mostCarriedCalories(input string) int {
	elves := strings.Split(input, "\n\n")
	mostCalores := 0
	for _, elf := range elves {
		totalCalories := caloriesOnOneElf(elf)
		if totalCalories > mostCalores {
			mostCalores = totalCalories
		}
	}
	return mostCalores
}

func caloriesCarriedByElves(input string) []int {
	var caloriesCarried []int
	for _, elf := range strings.Split(input, "\n\n") {
		totalCalories := caloriesOnOneElf(elf)
		caloriesCarried = append(caloriesCarried, totalCalories)
	}
	// sort in descending order
	sort.Slice(caloriesCarried, func(a, b int) bool { return caloriesCarried[a] > caloriesCarried[b] })
	return caloriesCarried
}

func caloriesOnOneElf(input string) int {
	total := 0
	for _, calorieCount := range strings.Split(input, "\n") {
		if len(calorieCount) > 1 {
			calories, _ := strconv.Atoi(calorieCount)
			total += calories
		}
	}
	return total
}

func sumTopElves(input string, count int) int {
	total := 0
	caloriesCarried := caloriesCarriedByElves(input)
	for i, calories := range caloriesCarried {
		if i == count {
			break
		}
		total += calories
	}
	return total
}

func main() {
	input := common.ReadFile(INPUT_FILENAME)
	test_response := mostCarriedCalories(TEST_CALORIES)
	if test_response != TEST_EXPECTED {
		errMsg := fmt.Sprintf("Day 01 Part 01 tests fail, Received %d instead of %d", test_response, TEST_EXPECTED)
		log.Fatal(errMsg)
	}
	test_response = sumTopElves(TEST_CALORIES, 3)
	if test_response != TEST_EXPECTED2 {
		errMsg := fmt.Sprintf("Day 01 Part 02 tests fail, Received %d instead of %d", test_response, TEST_EXPECTED2)
		log.Fatal(errMsg)
	}

	fmt.Println("Day 01 01:", mostCarriedCalories(input))
	fmt.Println("Day 01 02:", sumTopElves(input, 3))
}
