package adventofcode

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strings"
)

const y16d11TestInput string = `The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
The second floor contains a hydrogen generator.
The third floor contains a lithium generator.
The fourth floor contains nothing relevant.`

type y16d11Floor struct {
	chips, generators map[string]bool
}

type y16d11State struct {
	floors                      map[int]y16d11Floor
	count, elevator, steps, top int
}

func (s *y16d11State) dump() string {
	sb := strings.Builder{}
	for f := s.top; f >= 0; f-- {
		sb.WriteString(fmt.Sprintf("[%d]: ", f))
		if f == s.elevator {
			sb.WriteString("E ")
		}
		for _, chip := range y16d11SortKeys(s.floors[f].chips) {
			sb.WriteString(fmt.Sprintf("%s-M ", chip))
		}
		for _, generator := range y16d11SortKeys(s.floors[f].generators) {
			sb.WriteString(fmt.Sprintf("%s-G ", generator))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (s *y16d11State) hash() string {
	sha := sha256.New()
	sha.Write([]byte(s.dump()))
	return fmt.Sprintf("%x", sha.Sum(nil))
}

func y16d11SortKeys(m map[string]bool) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func y16d11ParseInput(input string, part int) (state y16d11State) {
	state = y16d11State{floors: map[int]y16d11Floor{}, count: 0, elevator: 0, steps: 0, top: 0}
	for f, line := range strings.Split(input, "\n") {
		chips := map[string]bool{}
		generators := map[string]bool{}
		fields := strings.Fields(line)
		for i := range len(fields) {
			if fields[i] == "a" {
				state.count++
				target := strings.ToUpper(strings.Split(fields[i+1], "-")[0])
				if strings.Contains(fields[i+1], "-") {
					chips[target] = true
				} else {
					generators[target] = true
				}
			}
		}
		state.floors[f] = y16d11Floor{chips: chips, generators: generators}
		state.top = f
	}
	return
}

func y16d11(input string, part int) int {
	// cache := map[string]bool{}
	// if part == 2 {
	// 	return -2
	// }
	// q := []y16d11State{y16d11ParseInput(input, part)}
	// for len(q) > 0 {
	// 	current := q[0]
	// 	q = q[1:]
	// 	if current.done() {
	// 		return current.steps
	// 	}
	// 	hash := current.hash()
	// 	if cache[hash] {
	// 		continue
	// 	}
	// 	cache[hash] = true
	// 	states := current.nextStates()
	// 	q = append(q, states...)
	// }
	return -1
}
