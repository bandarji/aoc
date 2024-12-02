package aoc2400

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const AOC2400 bool = true
const contentDir string = "inputfiles/"

func StrToInt(s string) (v int) {
	v, _ = strconv.Atoi(s)
	return
}

func ReadContent(day int) (content string) {
	filename := fmt.Sprintf("%sinput_%02d.txt", contentDir, day)
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to read file: %s", filename)
	}
	content = strings.Trim(string(bytes), "\n")
	return
}
