package adventofcode

import (
	"strings"
)

var y17d15Factors = [2]int{16807, 48271}

const y17d15Divisor int = 2_147_483_647
const y17d15Mask int = 1 << 16

func y17d15(input string, part int) (total int) {
	mods := [2]int{1, 1}
	if part == 2 {
		mods = [2]int{4, 8}
	}
	cycles := 40_000_000
	if part == 2 {
		cycles = 5_000_000
	}
	gens := y17d15ParseInput(input)
	total = y17d15Duel(gens, mods, cycles)
	return
}

func y17d15ParseInput(input string) (gens [2]int) {
	gens = [2]int{0, 0}
	for i, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		gens[i] = strToInt(fields[len(fields)-1])
	}
	return
}

func y17d15Duel(gens, mods [2]int, cycles int) (total int) {
	for i := range cycles {
		_ = i
		for j := range 2 {
			gens[j] = y17d15CalcNextValue(gens[j], y17d15Factors[j], mods[j])
		}
		cmpv := gens[0] ^ gens[1]
		if (cmpv % y17d15Mask) == 0 {
			total++
		}
	}
	return
}

func y17d15CalcNextValue(value, factor, mod int) int {
	value *= factor
	value %= y17d15Divisor
	for value%mod != 0 {
		value *= factor
		value %= y17d15Divisor
	}
	return value
}
