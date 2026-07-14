package adventofcode

import (
	"slices"
	"strings"
)

func y17d16(input string, programCount, part int) (answer string) {
	dances := 1
	if part == 2 {
		dances = 1_000_000_000
	}
	answer = y17d16Dances(strings.Split(input, ","), programCount, dances)
	return
}

func y17d16NewPrograms(programCount int) (programs []string) {
	for i := 0; i < programCount; i++ {
		programs = append(programs, string(rune('a'+i)))
	}
	return
}

func y17d16Dances(steps []string, programCount, dances int) string {
	programs := y17d16NewPrograms(programCount)
	cache := map[string]int{}
	for d := 0; d < dances; d++ {
		for _, step := range steps {
			switch step[0] {
			case 's':
				spin := strToInt(step[1:])
				right := programs[len(programs)-spin:]
				left := programs[:len(programs)-spin]
				programs = append(right, left...)
			case 'x':
				parts := strings.Split(step[1:], "/")
				pos1 := strToInt(parts[0])
				pos2 := strToInt(parts[1])
				programs[pos1], programs[pos2] = programs[pos2], programs[pos1]
			case 'p':
				parts := strings.Split(step[1:], "/")
				prog1 := parts[0]
				prog2 := parts[1]
				pos1 := slices.Index(programs, prog1)
				pos2 := slices.Index(programs, prog2)
				programs[pos1], programs[pos2] = programs[pos2], programs[pos1]
			}
		}
		key := y17d16GenerateKey(programs)
		if last, exists := cache[key]; exists {
			delta := d - last
			for d+delta < dances {
				d += delta
			}
		}
		cache[key] = d
	}
	return y17d16GenerateKey(programs)
}

func y17d16GenerateKey(programs []string) string {
	return strings.Join(programs, "")
}
