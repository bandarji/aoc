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
	"aoc2023/d16"
	"aoc2023/d17"
	"aoc2023/d18"
	"aoc2023/d19"
	"aoc2023/d21"
	"aoc2023/d22"
	"aoc2023/d23"
	"aoc2023/d24"
	"aoc2023/d25"
	"log"
)

func main() {
	log.Printf("Day  1, part  1 (test): %d", d01.Calibrate(d01.TEST0101, 1))
	log.Printf("Day  1, part  2 (test): %d", d01.Calibrate(d01.TEST0102, 2))
	// log.Printf("Day  1, part  1: %d\n", d01.Calibrate(common.ReadContent(1), 1))
	// log.Printf("Day  1, part  2: %d\n\n", d01.Calibrate(common.ReadContent(1), 2))

	log.Print("")
	log.Printf("Day  2, part  1 (test): %d", d02.Solve(d02.TEST1, 1))
	log.Printf("Day  2, part  2 (test): %d", d02.Solve(d02.TEST1, 2))
	// log.Printf("Day  2, part  1: %d\n", d02.Solve(common.ReadContent(2), 1))
	// log.Printf("Day  2, part  2: %d\n\n", d02.Solve(common.ReadContent(2), 2))

	log.Print("")
	log.Printf("Day  3, part  1 (test): %d", d03.Solve(d03.TEST1, 1))
	log.Printf("Day  3, part  2 (test): %d", d03.Solve(d03.TEST1, 2))
	// log.Printf("Day  3, part  1: %d\n", d03.Solve(common.ReadContent(3), 1))
	// log.Printf("Day  3, part  2: %d\n\n", d03.Solve(common.ReadContent(3), 2))

	log.Print("")
	log.Printf("Day  4, part  1 (test): %d", d04.Solve(d04.TEST1, 1))
	log.Printf("Day  4, part  2 (test): %d", d04.Solve(d04.TEST1, 2))
	// log.Printf("Day  4, part  1: %d\n", d04.Solve(common.ReadContent(4), 1))
	// log.Printf("Day  4, part  2: %d\n\n", d04.Solve(common.ReadContent(4), 2))

	log.Print("")
	log.Printf("Day  5, part  1 (test): %d", d05.Solve(d05.TEST1, 1))
	log.Printf("Day  5, part  2 (test): %d", d05.Solve(d05.TEST1, 2))
	// log.Printf("Day  5, part  1: %d\n", d05.Solve(common.ReadContent(5), 1))
	// log.Printf("Day  5, part  2: %d\n\n", d05.Solve(common.ReadContent(5), 2))

	log.Print("")
	log.Printf("Day  6, part  1 (test): %d", d06.Solve(d06.TEST1, 1))
	log.Printf("Day  6, part  2 (test): %d", d06.Solve(d06.TEST1, 2))
	// log.Printf("Day  6, part  1: %d\n", d06.Solve(common.ReadContent(6), 1))
	// log.Printf("Day  6, part  2: %d\n\n", d06.Solve(common.ReadContent(6), 2))

	log.Print("")
	log.Printf("Day  7, part  1 (test): %d", d07.Solve(d07.TEST1, 1))
	log.Printf("Day  7, part  2 (test): %d", d07.Solve(d07.TEST1, 2))
	// log.Printf("Day  7, part  1: %d\n", d07.Solve(common.ReadContent(7), 1))
	// log.Printf("Day  7, part  2: %d\n\n", d07.Solve(common.ReadContent(7), 2))

	log.Print("")
	log.Printf("Day  8, part  1 (test 1): %d", d08.Solve(d08.TEST1, 1))
	log.Printf("Day  8, part  1 (test 2): %d", d08.Solve(d08.TEST2, 1))
	log.Printf("Day  8, part  2 (test 1): %d", d08.Solve(d08.TEST3, 2))
	// log.Printf("Day  8, part  1: %d\n", d08.Solve(common.ReadContent(8), 1))
	// log.Printf("Day  8, part  2: %d\n\n", d08.Solve(common.ReadContent(8), 2))

	log.Print("")
	log.Printf("Day  9, part  1 (test 1): %d", d09.Solve(d09.TEST1, 1))
	log.Printf("Day  9, part  2 (test 1): %d", d09.Solve(d09.TEST1, 2))
	// log.Printf("Day  9, part  1: %d\n", d09.Solve(common.ReadContent(9), 1))
	// log.Printf("Day  9, part  2: %d\n\n", d09.Solve(common.ReadContent(9), 2))

	log.Print("")
	log.Printf("Day 10, part  1 (test 1): %d", d10.Solve(d10.TEST1, 1))
	log.Printf("Day 10, part  1 (test 2): %d", d10.Solve(d10.TEST2, 1))
	log.Printf("Day 10, part  2 (test 1): %d", d10.Solve(d10.TEST3, 2))
	// log.Printf("Day 10, part  1: %d\n", d10.Solve(common.ReadContent(10), 1))
	// log.Printf("Day 10, part  2: %d\n\n", d10.Solve(common.ReadContent(10), 2))

	log.Print("")
	log.Printf("Day 11, part  1 (test 1): %d", d11.Solve(d11.TEST1, 1, 2))
	log.Printf("Day 11, part  2 (test 10x): %d", d11.Solve(d11.TEST1, 2, 10))
	log.Printf("Day 11, part  2 (test 100x): %d", d11.Solve(d11.TEST1, 2, 100))
	// log.Printf("Day 11, part  1: %d\n", d11.Solve(common.ReadContent(11), 1, 2))
	// log.Printf("Day 11, part  2: %d\n\n", d11.Solve(common.ReadContent(11), 2, 1000000))

	log.Print("")
	log.Printf("Day 12, part  1 (test 1): %d", d12.Solve(d12.TEST1, 1))
	log.Printf("Day 12, part  2 (test 2): %d", d12.Solve(d12.TEST1, 2))
	// log.Printf("Day 12, part  1: %d\n", d12.Solve(common.ReadContent(12), 1))
	// log.Printf("Day 12, part  2: %d\n\n", d12.Solve(common.ReadContent(12), 2))

	log.Print("")
	log.Printf("Day 13, part  1 (test 1): %d", d13.Solve(d13.TEST1, 1))
	log.Printf("Day 13, part  2 (test 2): %d", d13.Solve(d13.TEST1, 2))
	// log.Printf("Day 13, part  1: %d\n", d13.Solve(common.ReadContent(13), 1))
	// log.Printf("Day 13, part  2: %d\n\n", d13.Solve(common.ReadContent(13), 2))

	log.Print("")
	log.Printf("Day 14, part  1 (test 1): %d", d14.Solve(d14.TEST1, 1))
	// log.Printf("Day 14, part  2 (test 2): %d\n\n", d14.Solve(d14.TEST1, 2))
	// log.Printf("Day 14, part  1: %d\n\n", d14.Solve(common.ReadContent(14), 1))
	// log.Printf("Day 12, part  2: %d\n\n", d12.Solve(common.ReadContent(12), 2))

	log.Print("")
	log.Printf("Day 15, part  1 (test 1): %d", d15.Solve(d15.TEST1, 1))
	log.Printf("Day 15, part  1 (test 2): %d", d15.Solve(d15.TEST2, 1))
	log.Printf("Day 15, part  2 (test 1): %d", d15.Solve(d15.TEST2, 2))
	// log.Printf("Day 15, part  1: %d\n", d15.Solve(common.ReadContent(15), 1))
	// log.Printf("Day 15, part  2: %d\n\n", d15.Solve(common.ReadContent(15), 2))

	log.Print("")
	log.Printf("Day 16, part  1 (test 1): %d", d16.Solve(d16.TEST1, 1))
	log.Printf("Day 16, part  2 (test 1): %d", d16.Solve(d16.TEST1, 2))
	// log.Printf("Day 16, part  1: %d\n", d16.Solve(common.ReadContent(16), 1))
	// log.Printf("Day 16, part  2: %d\n\n", d16.Solve(common.ReadContent(16), 2))

	log.Print("")
	log.Printf("Day 17, part  1 (test 1): %d", d17.Solve(d17.TEST1, 1))
	log.Printf("Day 17, part  2 (test 1): %d", d17.Solve(d17.TEST1, 2))
	log.Printf("Day 17, part  2 (test 2): %d", d17.Solve(d17.TEST2, 2))
	// log.Printf("Day 17, part  1: %d", d17.Solve(common.ReadContent(17), 1))
	// log.Printf("Day 17, part  2: %d", d17.Solve(common.ReadContent(17), 2))

	log.Print("")
	log.Printf("Day 18, part  1 (test 1): %d", d18.Solve(d18.TEST1, 1))
	log.Printf("Day 18, part  2 (test 1): %d", d18.Solve(d18.TEST1, 2))
	// log.Printf("Day 18, part  1: %d", d18.Solve(common.ReadContent(18), 1))
	// log.Printf("Day 18, part  2: %d", d18.Solve(common.ReadContent(18), 2))

	log.Print("")
	log.Printf("Day 19, part  1 (test 1): %d", d19.Solve(d19.TEST1, 1))
	log.Printf("Day 19, part  2 (test 1): %d", d19.Solve(d19.TEST1, 2))
	// log.Printf("Day 19, part  1: %d\n", d19.Solve(common.ReadContent(19), 1))
	// log.Printf("Day 19, part  2: %d\n\n", d19.Solve(common.ReadContent(19), 2))

	log.Print("")
	log.Printf("Day 21, part  1 (test 1): %d", d21.Solve(d21.TEST1, 6, 1))
	// log.Printf("Day 21, part  2 (test 1): %d\n\n", d21.Solve(d21.TEST1, 2))
	// In exactly 6 steps, he can still reach 16 garden plots.
	// In exactly 10 steps, he can reach any of 50 garden plots.
	// In exactly 50 steps, he can reach 1594 garden plots.
	// In exactly 100 steps, he can reach 6536 garden plots.
	// In exactly 500 steps, he can reach 167004 garden plots.
	// In exactly 1000 steps, he can reach 668697 garden plots.
	// In exactly 5000 steps, he can reach 16733044 garden plots.
	// for _, test := range []int{6, 10, 50, 100, 500, 1000, 5000} {
	// 	log.Printf("In exactly %d steps, he can reach %d garden plots.", test, d21.Solve(d21.TEST1, test, 2))
	// }
	log.Print("")
	// log.Printf("Day 21, part  1: %d\n", d21.Solve(common.ReadContent(21), 64, 1))
	// log.Printf("Day 21, part  2: %d\n\n", d21.Solve(common.ReadContent(21), 26501365, 2))

	log.Print("")
	log.Printf("Day 22, part  1 (test 1): %d", d22.Solve(d22.TEST1, 1))
	log.Printf("Day 22, part  2 (test 1): %d", d22.Solve(d22.TEST1, 2))
	// log.Printf("Day 22, part  1: %d\n", d22.Solve(common.ReadContent(22), 1))
	// log.Printf("Day 22, part  2: %d\n\n", d22.Solve(common.ReadContent(22), 2))

	log.Print("")
	log.Printf("Day 23, part  1 (test 1): %d", d23.Solve(d23.TEST1, 1))
	// log.Printf("Day 23, part  2 (test 1): %d\n\n", d23.Solve(d23.TEST1, 2))
	// log.Printf("Day 23, part  1: %d\n", d23.Solve(common.ReadContent(23), 1))
	// log.Printf("Day 23, part  2: %d\n\n", d23.Solve(common.ReadContent(23), 2))

	log.Print("")
	log.Printf("Day 24, part  1 (test 1): %d", d24.Solve(d24.TEST1, 1, 7, 27))
	// log.Printf("Day 24, part  2 (test 1): %d\n\n", d24.Solve(d24.TEST1, 2))
	log.Printf("Day 24, part  1: %d", d24.Solve(common.ReadContent(24), 1, 200_000_000_000_000, 400_000_000_000_000))
	// log.Printf("Day 24, part  2: %d\n\n", d24.Solve(common.ReadContent(24), 2))

	log.Print("")
	log.Printf("Day 25, part  1 (test 1): %d", d25.Solve(d25.TEST1, 1))
	// log.Printf("Day 24, part  2 (test 1): %d\n\n", d24.Solve(d24.TEST1, 2))
	// log.Printf("Day 25, part  1: %d\n", d25.Solve(common.ReadContent(25), 1))
	// log.Printf("Day 24, part  2: %d\n\n", d24.Solve(common.ReadContent(24), 2))
}
