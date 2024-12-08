package aoc2400

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func AbsInts(x, y int) (a int) {
	a = x - y
	if a < 0 {
		a = -a
	}
	return
}

func AbsInt(x int) (a int) {
	a = x
	if x < 0 {
		a = -x
	}
	return
}
