package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func obtain_readings(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var readings []int

	content := bufio.NewScanner(f)
	for content.Scan() {
		reading, _ := strconv.Atoi(content.Text())
		readings = append(readings, reading)
	}

	return readings
}

func increases_single(readings []int) int {
	increases := 0
	for index, entry := range readings {
		if index == 0 {
			continue
		}
		if entry > readings[index - 1] {
			increases++
		}
	}
	return increases
}

func increases_window(readings []int) int {
	var windowed_readings []int
	var total int
	for index, entry := range readings {
		if index < 3 {
			continue
		}
		total = entry + readings[index - 1] + readings[index - 2]
		windowed_readings = append(windowed_readings, total)
	}
	return increases_single(windowed_readings)
}

func main() {
	readings := obtain_readings("input.txt")
	// readings = convert_readings(readings)
	increases := increases_single(readings)
	fmt.Println(increases)
	increases = increases_window(readings)
	fmt.Println(increases)
}


// number, err := strconv.Atoi(str)
