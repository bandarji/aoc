package main

import (
	"aoc2023/d01"
	"aoc2023/d02"
	"aoc2023/d03"
	"aoc2023/d04"
	"aoc2023/d05"
	"aoc2023/d06"
	"aoc2023/d07"
	"log"
)

func main() {
	log.Printf("Day  1, part  1 (test): %d\n", d01.Calibrate(d01.TEST0101, 1))
	log.Printf("Day  1, part  2 (test): %d\n\n", d01.Calibrate(d01.TEST0102, 2))
	// log.Printf("Day  1, part  1: %d\n", d01.Calibrate(common.ReadContent(1), 1))
	// log.Printf("Day  1, part  2: %d\n\n", d01.Calibrate(common.ReadContent(1), 2))

	log.Printf("Day  2, part  1 (test): %d\n", d02.Solve(d02.TEST1, 1))
	log.Printf("Day  2, part  2 (test): %d\n\n", d02.Solve(d02.TEST1, 2))
	// log.Printf("Day  2, part  1: %d\n", d02.Solve(common.ReadContent(2), 1))
	// log.Printf("Day  2, part  2: %d\n\n", d02.Solve(common.ReadContent(2), 2))

	log.Printf("Day  3, part  1 (test): %d\n", d03.Solve(d03.TEST1, 1))
	log.Printf("Day  3, part  2 (test): %d\n\n", d03.Solve(d03.TEST1, 2))
	// log.Printf("Day  3, part  1: %d\n", d03.Solve(common.ReadContent(3), 1))
	// log.Printf("Day  3, part  2: %d\n\n", d03.Solve(common.ReadContent(3), 2))

	log.Printf("Day  4, part  1 (test): %d\n", d04.Solve(d04.TEST1, 1))
	log.Printf("Day  4, part  2 (test): %d\n\n", d04.Solve(d04.TEST1, 2))
	// log.Printf("Day  4, part  1: %d\n", d04.Solve(common.ReadContent(4), 1))
	// log.Printf("Day  4, part  2: %d\n\n", d04.Solve(common.ReadContent(4), 2))

	log.Printf("Day  5, part  1 (test): %d\n", d05.Solve(d05.TEST1, 1))
	log.Printf("Day  5, part  2 (test): %d\n\n", d05.Solve(d05.TEST1, 2))
	// log.Printf("Day  5, part  1: %d\n", d05.Solve(common.ReadContent(5), 1))
	// log.Printf("Day  5, part  2: %d\n\n", d05.Solve(common.ReadContent(5), 2))

	log.Printf("Day  6, part  1 (test): %d\n", d06.Solve(d06.TEST1, 1))
	log.Printf("Day  6, part  2 (test): %d\n\n", d06.Solve(d06.TEST1, 2))
	// log.Printf("Day  6, part  1: %d\n", d06.Solve(common.ReadContent(6), 1))
	// log.Printf("Day  6, part  2: %d\n\n", d06.Solve(common.ReadContent(6), 2))

	log.Printf("Day  7, part  1 (test): %d\n", d07.Solve(d07.TEST1, 1))
	log.Printf("Day  7, part  2 (test): %d\n\n", d07.Solve(d07.TEST1, 2))
	// log.Printf("Day  7, part  1: %d\n", d07.Solve(common.ReadContent(7), 1))
	// log.Printf("Day  7, part  2: %d\n\n", d07.Solve(common.ReadContent(7), 2))
}
