package main

import (
	"fmt"
	"log"
	"os"

	"go2016/d01"
	"go2016/d02"
)

func GetContent(day int) string {
	content, err := os.ReadFile(fmt.Sprintf("inputs/input_day%02d.txt", day))
	if err != nil {
		log.Fatalln(err)
	}
	return string(content)
}

func main() {
	var content string
	content = GetContent(1)
	log.Println("Day 01/01:", d01.Day(content, 1))
	log.Println("Day 01/02:", d01.Day(content, 2))
	content = GetContent(2)
	log.Println("Day 02/01:", d02.Day(content, 1))
	log.Println("Day 02/02:", d02.Day(content, 2))
}
