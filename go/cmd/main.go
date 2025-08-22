package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	firstYear   int = 2015
	currentYear int = 2024
)

func readContent(filename string) string {
	bytes, err := os.ReadFile(filename)
	if err == nil {
		return strings.TrimSpace(string(bytes))
	}
	return ""
}

func main() {
	for year := firstYear; year <= currentYear; year++ {
		for day := 1; day <= 25; day++ {
			formattedDate := fmt.Sprintf("%d-%02d", year, day)
			input := readContent(fmt.Sprintf("%s-input.txt", formattedDate))
			if input == "" {
				fmt.Println(formattedDate, "no input found -- skipping")
			} else {
				fmt.Println(formattedDate, input)
			}
		}
		fmt.Println()
	}
}
