package adventofcode

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

const y16d11TestInput1 string = `The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
The second floor contains a hydrogen generator.
The third floor contains a lithium generator.
The fourth floor contains nothing relevant.`

type y16d11State struct {
	count, elevator, steps, top int
	floors                      []map[string]bool
}

func y16d11ParseInput(input string, part int) (state y16d11State) {
	state = y16d11State{floors: []map[string]bool{}, count: 0, elevator: 0, steps: 0, top: 0}
	for f, line := range strings.Split(input, "\n") {
		floorItems := map[string]bool{}
		fields := strings.Fields(line)
		for i := range len(fields) {
			if fields[i] == "a" {
				state.count++
				save := fields[i+1]
				if strings.Contains(save, "-") {
					parts := strings.Split(save, "-")
					save = fmt.Sprintf("%s-%s", parts[0][:3], parts[1][:1])
				} else {
					save = save[:3]
				}
				floorItems[save] = true
			}
		}
		state.floors = append(state.floors, floorItems)
		state.top = f
	}
	if part == 2 {
		state.floors[0]["ele"] = true
		state.floors[0]["ele-c"] = true
		state.floors[0]["dil"] = true
		state.floors[0]["dil-c"] = true
		state.count += 4
	}
	return
}

func y16d11(input string, part int) int {
	// if part == 2 {
	// 	return -2
	// }
	cache := map[string]bool{}
	state := y16d11ParseInput(input, part)
	q := []y16d11State{state}
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		dump := current.dump()
		// if len(cache)%1_000_000 == 0 || len(q)%10_000_000 == 0 {
		// 	log.Println("STEPS:", current.steps, "CACHE LEN:", len(cache), "Q LEN:", len(q), "DUMP:", dump)
		// }
		if cache[dump] {
			continue
		}
		cache[dump] = true
		if !current.isValid() {
			continue
		}
		if current.isDone() {
			log.Println("DONE:", dump)
			return current.steps
		}
		q = append(q, current.nextStates()...)
	}
	return -1
}

func (s y16d11State) isValid() bool {
	for i := s.top; i >= 0; i-- {
		if isInvalidFloor(s.floors[i]) {
			return false
		}
	}
	return true
}

func isInvalidFloor(floor map[string]bool) bool {
	chips, gens := 0, 0 // count of unpaired
	for i := range floor {
		base := strings.Split(i, "-")[0]
		if strings.Contains(i, "-") {
			if !floor[base] {
				chips++
			}
		} else {
			if !floor[fmt.Sprintf("%s-compatible", base)] {
				gens++
			}
		}
	}
	return chips > 0 && gens > 0
}

func (s y16d11State) dump() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("[E%d]", s.elevator))
	for f := range s.top + 1 {
		sb.WriteString(fmt.Sprintf("[%d%s]", f, strings.Join(sortedKeys(s.floors[f]), "|")))
	}
	return sb.String()
}

func sortedKeys(m map[string]bool) (keys []string) {
	keys = []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}

func (s y16d11State) isDone() bool {
	return len(s.floors[s.top]) == s.count
}

func (s y16d11State) nextStates() (states []y16d11State) {
	states = []y16d11State{}
	for _, dy := range [2]int{1, -1} {
		ne := s.elevator + dy
		if ne < 0 || ne > s.top {
			continue
		}
		combos := y16d11Combos(sortedKeys(s.floors[s.elevator]))
		for _, combo := range combos {
			states = append(states, s.move(ne, combo...))
		}
	}
	return
}

func (s y16d11State) move(ne int, items ...string) y16d11State {
	ns := y16d11State{
		count:    s.count,
		elevator: ne,
		steps:    s.steps + 1,
		top:      s.top,
		floors:   make([]map[string]bool, s.top+1),
	}
	for f := range s.top + 1 {
		ns.floors[f] = make(map[string]bool)
		for k := range s.floors[f] {
			ns.floors[f][k] = true
		}
	}
	for _, item := range items {
		delete(ns.floors[s.elevator], item)
		ns.floors[ne][item] = true
	}
	return ns
}

func y16d11Combos(items []string) (combos [][]string) {
	combos = [][]string{}
	for i := range len(items) {
		combos = append(combos, []string{items[i]})
	}
	for i := range len(items) {
		for j := i + 1; j < len(items); j++ {
			combos = append(combos, []string{items[i], items[j]})
		}
	}
	return
}
