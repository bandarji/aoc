package main

import (
	"aoc2023/common"
	"aoc2023/d01"
	"fmt"
)

func main() {
	fmt.Printf("Day  1, part  1 (test): %d\n", d01.Calibrate(d01.TEST0101, 1))
	fmt.Printf("Day  1, part  2 (test): %d\n", d01.Calibrate(d01.TEST0102, 2))
	fmt.Printf("Day  1, part  1: %d\n", d01.Calibrate(common.ReadContent(1), 1))
	fmt.Printf("Day  1, part  2: %d\n", d01.Calibrate(common.ReadContent(1), 2))
}
