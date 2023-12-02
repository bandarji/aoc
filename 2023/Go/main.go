package main

import (
	"aoc2023/common"
	"aoc2023/d01"
	"aoc2023/d02"
	"fmt"
)

func main() {
	fmt.Printf("Day  1, part  1 (test): %d\n", d01.Calibrate(d01.TEST0101, 1))
	fmt.Printf("Day  1, part  2 (test): %d\n", d01.Calibrate(d01.TEST0102, 2))
	fmt.Printf("Day  1, part  1: %d\n", d01.Calibrate(common.ReadContent(1), 1))
	fmt.Printf("Day  1, part  2: %d\n\n", d01.Calibrate(common.ReadContent(1), 2))

	fmt.Printf("Day  2, part  1 (test): %d\n", d02.Solve(d02.TEST1, 1))
	fmt.Printf("Day  2, part  2 (test): %d\n", d02.Solve(d02.TEST1, 2))
	fmt.Printf("Day  2, part  1: %d\n", d02.Solve(common.ReadContent(2), 1))
	fmt.Printf("Day  2, part  2: %d\n\n", d02.Solve(common.ReadContent(2), 2))
}
