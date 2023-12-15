package main

import (
	"aoc2023/common"
	"aoc2023/d01"
	"aoc2023/d02"
	"aoc2023/d03"
	"aoc2023/d04"
	"aoc2023/d05"
	"aoc2023/d06"
	"aoc2023/d07"
	"aoc2023/d08"
	"aoc2023/d09"
	"aoc2023/d10"
	"aoc2023/d11"
	"aoc2023/d12"
	"aoc2023/d13"
	"aoc2023/d14"
	"aoc2023/d15"
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

	log.Printf("Day  8, part  1 (test 1): %d\n", d08.Solve(d08.TEST1, 1))
	log.Printf("Day  8, part  1 (test 2): %d\n", d08.Solve(d08.TEST2, 1))
	log.Printf("Day  8, part  2 (test 1): %d\n\n", d08.Solve(d08.TEST3, 2))
	// log.Printf("Day  8, part  1: %d\n", d08.Solve(common.ReadContent(8), 1))
	// log.Printf("Day  8, part  2: %d\n\n", d08.Solve(common.ReadContent(8), 2))

	log.Printf("Day  9, part  1 (test 1): %d\n", d09.Solve(d09.TEST1, 1))
	log.Printf("Day  9, part  2 (test 1): %d\n\n", d09.Solve(d09.TEST1, 2))
	// log.Printf("Day  9, part  1: %d\n", d09.Solve(common.ReadContent(9), 1))
	// log.Printf("Day  9, part  2: %d\n\n", d09.Solve(common.ReadContent(9), 2))

	log.Printf("Day 10, part  1 (test 1): %d\n", d10.Solve(d10.TEST1, 1))
	log.Printf("Day 10, part  1 (test 2): %d\n", d10.Solve(d10.TEST2, 1))
	log.Printf("Day 10, part  2 (test 1): %d\n\n", d10.Solve(d10.TEST3, 2))
	// log.Printf("Day 10, part  1: %d\n", d10.Solve(common.ReadContent(10), 1))
	// log.Printf("Day 10, part  2: %d\n\n", d10.Solve(common.ReadContent(10), 2))

	log.Printf("Day 11, part  1 (test 1): %d\n", d11.Solve(d11.TEST1, 1, 2))
	log.Printf("Day 11, part  2 (test 10x): %d\n", d11.Solve(d11.TEST1, 2, 10))
	log.Printf("Day 11, part  2 (test 100x): %d\n\n", d11.Solve(d11.TEST1, 2, 100))
	// log.Printf("Day 11, part  1: %d\n", d11.Solve(common.ReadContent(11), 1, 2))
	// log.Printf("Day 11, part  2: %d\n\n", d11.Solve(common.ReadContent(11), 2, 1000000))

	log.Printf("Day 12, part  1 (test 1): %d\n", d12.Solve(d12.TEST1, 1))
	log.Printf("Day 12, part  2 (test 2): %d\n\n", d12.Solve(d12.TEST1, 2))
	// log.Printf("Day 12, part  1: %d\n", d12.Solve(common.ReadContent(12), 1))
	// log.Printf("Day 12, part  2: %d\n\n", d12.Solve(common.ReadContent(12), 2))

	log.Printf("Day 13, part  1 (test 1): %d\n", d13.Solve(d13.TEST1, 1))
	log.Printf("Day 13, part  2 (test 2): %d\n\n", d13.Solve(d13.TEST1, 2))
	// log.Printf("Day 12, part  1: %d\n", d12.Solve(common.ReadContent(12), 1))
	// log.Printf("Day 12, part  2: %d\n\n", d12.Solve(common.ReadContent(12), 2))

	log.Printf("Day 14, part  1 (test 1): %d\n", d14.Solve(d14.TEST1, 1))
	// log.Printf("Day 14, part  2 (test 2): %d\n\n", d14.Solve(d14.TEST1, 2))
	log.Printf("Day 14, part  1: %d\n\n", d14.Solve(common.ReadContent(14), 1))
	// log.Printf("Day 12, part  2: %d\n\n", d12.Solve(common.ReadContent(12), 2))

	log.Printf("Day 15, part  1 (test 1): %d\n", d15.Solve(d15.TEST1, 1))
	log.Printf("Day 15, part  1 (test 2): %d\n", d15.Solve(d15.TEST2, 1))
	log.Printf("Day 15, part  2 (test 1): %d\n\n", d15.Solve(d15.TEST2, 2))
	log.Printf("Day 15, part  1: %d\n", d15.Solve(common.ReadContent(15), 1))
	log.Printf("Day 12, part  2: %d\n\n", d15.Solve(common.ReadContent(15), 2))
}
