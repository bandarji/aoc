package main

import (
	"aoc2023/common"
	"aoc2023/d01"
	"aoc2023/d02"
	"aoc2023/d03"
	"log"
)

func main() {
	log.Printf("Day  1, part  1 (test): %d\n", d01.Calibrate(d01.TEST0101, 1))
	log.Printf("Day  1, part  2 (test): %d\n", d01.Calibrate(d01.TEST0102, 2))
	log.Printf("Day  1, part  1: %d\n", d01.Calibrate(common.ReadContent(1), 1))
	log.Printf("Day  1, part  2: %d\n\n", d01.Calibrate(common.ReadContent(1), 2))

	log.Printf("Day  2, part  1 (test): %d\n", d02.Solve(d02.TEST1, 1))
	log.Printf("Day  2, part  2 (test): %d\n", d02.Solve(d02.TEST1, 2))
	log.Printf("Day  2, part  1: %d\n", d02.Solve(common.ReadContent(2), 1))
	log.Printf("Day  2, part  2: %d\n\n", d02.Solve(common.ReadContent(2), 2))

	log.Printf("Day  3, part  1 (test): %d\n", d03.Solve(d03.TEST1, 1))
	log.Printf("Day  3, part  2 (test): %d\n", d03.Solve(d03.TEST1, 2))
	log.Printf("Day  3, part  1: %d\n", d03.Solve(common.ReadContent(3), 1))
	log.Printf("Day  3, part  2: %d\n\n", d03.Solve(common.ReadContent(3), 2))

}
