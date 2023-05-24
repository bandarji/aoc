package main

import (
	"log"
	"os"

	"go2016/d01"
)

func main() {
	content, err := os.ReadFile("inputs/input_day01.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Day 01/01:", d01.Day(string(content), 1))
	log.Println("Day 01/02:", d01.Day(string(content), 2))
}
