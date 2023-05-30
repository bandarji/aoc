package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"go2016/d01"
	"go2016/d02"
	"go2016/d03"
	"go2016/d04"
	"go2016/d05"
	"go2016/d06"
)

func GetContent(day int) string {
	content, err := os.ReadFile(fmt.Sprintf("inputs/input_day%02d.txt", day))
	if err != nil {
		log.Fatalln(err)
	}
	return strings.TrimRight(string(content), "\n")
}

func main() {
	var content string
	content = GetContent(1)
	log.Println("Day 01/01:", d01.Day(content, 1))
	log.Println("Day 01/02:", d01.Day(content, 2))
	content = GetContent(2)
	log.Println("Day 02/01:", d02.Day(content, 1))
	log.Println("Day 02/02:", d02.Day(content, 2))
	content = GetContent(3)
	log.Println("Day 03/01:", d03.Day(content, 1))
	log.Println("Day 03/02:", d03.Day(content, 2))
	content = GetContent(4)
	log.Println("Day 04/01:", d04.Day(content, 1))
	log.Println("Day 04/02:", d04.Day(content, 2))
	// no content
	log.Println("Day 05/01:", d05.Day("uqwqemis", 1))
	log.Println("Day 05/02:", d05.Day("uqwqemis", 2))
	content = GetContent(6)
	log.Println("Day 06/01:", d06.Day(content, 1))
	log.Println("Day 06/02:", d06.Day(content, 2))
}
