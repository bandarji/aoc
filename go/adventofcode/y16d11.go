package adventofcode

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

type y16d11Floor struct {
	items map[string]bool
}

type y16d11State struct {
	elevator, steps, floorCount int
	floors                      map[int]y16d11Floor
}

func (s *y16d11State) moves() []y16d11State {
	states := []y16d11State{}
	newState := s.clone()
	for item := range s.floors[s.elevator].items {
		for delta := range []int{-1, 1} {
			newState = s.clone()
			newState.moveItem(delta, item, s.elevator)
			if newState.isValid() {
				newState.steps++
				newState.elevator += delta
				states = append(states, newState)
			}
			newState = s.clone()
			if strings.HasSuffix(item, "C") && newState.floors[s.elevator].items[fmt.Sprintf("%sG", item[0:1])] {
				newState.moveItems(delta, item, fmt.Sprintf("%sG", item[0:1]), s.elevator)
			}
			if newState.isValid() {
				newState.steps++
				newState.elevator += delta
				states = append(states, newState)
			}
		}
	}
	return states
}

func (s *y16d11State) moveItem(delta int, item string, elevator int) {
	log.Println("MOVING ITEM", item, "FROM", elevator, "TO", elevator+delta)
	delete(s.floors[elevator].items, item)
	s.floors[elevator+delta].items[item] = true
	s.elevator += delta
}

func (s *y16d11State) moveItems(delta int, item string, partner string, elevator int) {
	log.Println("MOVING ITEMS", item, "AND", partner, "FROM", elevator, "TO", elevator+delta)
	delete(s.floors[elevator].items, item)
	delete(s.floors[elevator].items, partner)
	s.floors[elevator+delta].items[item] = true
	s.floors[elevator+delta].items[partner] = true
	s.elevator += delta
}

func (s *y16d11State) predict() []y16d11State {
	states := []y16d11State{}
	states = append(states, s.moves()...)
	return states
}

func (s *y16d11State) clone() y16d11State {
	newState := y16d11State{elevator: s.elevator, steps: s.steps, floorCount: s.floorCount, floors: map[int]y16d11Floor{}}
	for f, floor := range s.floors {
		newState.floors[f] = y16d11Floor{items: map[string]bool{}}
		for item := range floor.items {
			newState.floors[f].items[item] = true
		}
	}
	return newState
}

func (s y16d11State) isValid() bool {
	for f := 1; f < s.floorCount; f++ {
		for item := range s.floors[f].items {
			if strings.HasSuffix(item, "C") && !s.floors[f].items[fmt.Sprintf("%sG", item[0:1])] {
				return false
			}
		}
	}
	return true
}

func (s y16d11State) encode() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("e%d", s.elevator))
	for _, k := range y16d11SortFloorKeys(s.floors) {
		sb.WriteString(fmt.Sprintf("|f%d|", k))
		sb.WriteString(strings.Join(y16d11SortFloorItemKeys(s.floors[k]), ""))
	}
	return sb.String()
}

func y16d11SortFloorKeys(floors map[int]y16d11Floor) []int {
	keys := []int{}
	for k := range floors {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func y16d11SortFloorItemKeys(floor y16d11Floor) []string {
	keys := []string{}
	for k := range floor.items {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func y16d11ParseInput(input string) (state y16d11State) {
	state = y16d11State{elevator: 1, steps: 0, floors: map[int]y16d11Floor{}}
	for f, line := range strings.Split(input, "\n") {
		items := map[string]bool{}
		fields := strings.Fields(line)
		for i := range len(fields) {
			if fields[i] == "a" {
				if strings.Contains(fields[i+1], "-") {
					items[fmt.Sprintf("%sC", strings.ToUpper(fields[i+1][0:1]))] = true
				} else {
					items[fmt.Sprintf("%sG", strings.ToUpper(fields[i+1][0:1]))] = true
				}
			}
		}
		state.floors[f] = y16d11Floor{items: items}
	}
	state.floorCount = len(state.floors)
	return
}

func y16d11(input string, part int) int {
	_ = part
	cache := map[string]bool{}
	state := y16d11ParseInput(input)
	q := []y16d11State{state}
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		if current.isDone() {
			return current.steps
		}
		hash := current.encode()
		log.Println("HASH", hash)
		if cache[hash] {
			continue
		}
		cache[hash] = true
		states := current.predict()
		q = append(q, states...)
	}
	return -1
}

func (s y16d11State) isDone() bool {
	for f := 1; f < s.floorCount-1; f++ {
		if len(s.floors[f].items) > 0 {
			return false
		}
	}
	return true
}
