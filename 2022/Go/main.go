package main

import (
	"fmt"
	"strings"

	"aoc/common"
)

func main() {
	x := common.ReadFile("a.txt")
	fmt.Println(x)

	parts := strings.Split(x, "\n\n")
	fmt.Println(strings.Join(parts, "----\n"))
}
