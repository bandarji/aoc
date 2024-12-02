package main

import (
	"aoc24/aoc2401"
	"aoc24/aoc2402"
	"fmt"
)

func main() {
	fmt.Printf("AOC 2401.1_test: %d\n\n", aoc2401.Aoc240101(aoc2401.AOC2401_TEST))
	// fmt.Printf("AOC 2401.1: %d\n", aoc2401.Aoc240101(aoc2400.ReadContent(1)))
	fmt.Printf("AOC 2401.2_test: %d\n\n", aoc2401.Aoc240102(aoc2401.AOC2401_TEST))
	// fmt.Printf("AOC 2401.1: %d\n", aoc2401.Aoc240102(aoc2400.ReadContent(1)))
	fmt.Printf("AOC 2402.1_test: %d\n\n", aoc2402.Aoc240201(aoc2402.AOC2402_TEST))
	// fmt.Printf("AOC 2402.1: %d\n", aoc2402.Aoc240201(aoc2400.ReadContent(2)))
	fmt.Printf("AOC 2402.2_test: %d\n\n", aoc2402.Aoc240202(aoc2402.AOC2402_TEST))
	// fmt.Printf("AOC 2402.2: %d\n", aoc2402.Aoc240202(aoc2400.ReadContent(2)))
}
