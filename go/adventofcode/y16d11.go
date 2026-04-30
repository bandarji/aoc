package adventofcode

import (
	"fmt"
	"sort"
	"strings"
)

type y16d11Item struct {
	name string
	chip bool // false for generator, true for chip
}

type y16d11State struct {
	floors          [][]y16d11Item
	elevator, steps int
}

func (s y16d11State) clone() y16d11State {
	sc := y16d11State{
		floors:   [][]y16d11Item{},
		elevator: s.elevator,
		steps:    s.steps,
	}
	for _, f := range s.floors {
		sc.floors = append(sc.floors, append([]y16d11Item{}, f...))
	}
	return sc
}

func (s y16d11State) done() bool {
	total := 0
	for _, floor := range s.floors[:len(s.floors)-1] {
		total += len(floor)
	}
	return total == 0
}

func (s y16d11State) isValid() bool {
	for f := range s.floors {
		gens := map[string]bool{}
		for _, item := range s.floors[f] {
			if !item.chip {
				gens[item.name] = true
			}
		}
		if len(gens) == 0 {
			continue
		}
		for _, item := range s.floors[f] {
			if item.chip && !gens[item.name] {
				return false
			}
		}
	}
	return true
}

func (s y16d11State) nextStates() (states []y16d11State) {
	states = []y16d11State{}
	moves := [][]int{}
	items := s.floors[s.elevator]
	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			moves = append(moves, []int{i, j})
		}
	}
	for i := range items {
		moves = append(moves, []int{i})
	}
	elevatorDeltas := []int{}
	if s.elevator < len(s.floors)-1 {
		elevatorDeltas = append(elevatorDeltas, 1)
	}
	if s.elevator > 0 {
		elevatorDeltas = append(elevatorDeltas, -1)
	}
	for _, ed := range elevatorDeltas {
		for _, move := range moves {
			sc := s.clone()
			sc.elevator += ed
			sc.steps++
			prevFloor := s.elevator
			newFloor := sc.elevator
			for _, index := range move {
				sc.floors[newFloor] = append(sc.floors[newFloor], sc.floors[prevFloor][index])
			}
			for in := len(move) - 1; in >= 0; in-- {
				sc.floors[prevFloor][move[in]] = sc.floors[prevFloor][len(sc.floors[prevFloor])-1]
				sc.floors[prevFloor] = sc.floors[prevFloor][:len(sc.floors[prevFloor])-1]
			}
			if sc.isValid() {
				states = append(states, sc)
			}
		}
	}
	return
}

func (s y16d11State) hash() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("[%d]", s.elevator))
	for f, floor := range s.floors {
		chips := []string{}
		generators := []string{}
		for _, item := range floor {
			if item.chip {
				chips = append(chips, item.name)
			} else {
				generators = append(generators, item.name)
			}
		}
		sort.Strings(chips)
		sort.Strings(generators)
		sb.WriteString(fmt.Sprintf("%d:%s:%s:", f, strings.Join(chips, ","), strings.Join(generators, ",")))
	}
	return sb.String()
}

func y16d11ParseInput(input string, part int) (state y16d11State) {
	state = y16d11State{floors: [][]y16d11Item{}, elevator: 1, steps: 0}
	for _, line := range strings.Split(input, "\n") {
		items := []y16d11Item{}
		fields := strings.Fields(line)
		for i := range len(fields) {
			if fields[i] == "a" {
				if strings.Contains(fields[i+1], "-") {
					items = append(items, y16d11Item{name: strings.ToUpper(strings.Split(fields[i+1], "-")[0]), chip: false})
				} else {
					items = append(items, y16d11Item{name: strings.ToUpper(strings.Split(fields[i+1], "-")[0]), chip: true})
				}
			}
		}
		state.floors = append(state.floors, items)
	}
	if part == 2 {
		state.floors[0] = append(state.floors[0], y16d11Item{name: "ELERIUM", chip: false})
		state.floors[0] = append(state.floors[0], y16d11Item{name: "ELERIUM", chip: true})
		state.floors[0] = append(state.floors[0], y16d11Item{name: "DILITHIUM", chip: false})
		state.floors[0] = append(state.floors[0], y16d11Item{name: "DILITHIUM", chip: true})
	}
	return
}

func y16d11(input string, part int) int {
	// log.Printf("state: %+v", state)
	q := []y16d11State{y16d11ParseInput(input, part)}
	cache := map[string]bool{}
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		if current.done() {
			return current.steps
		}
		hash := current.hash()
		if cache[hash] {
			continue
		}
		cache[hash] = true
		states := current.nextStates()
		q = append(q, states...)
	}
	return -1
}
