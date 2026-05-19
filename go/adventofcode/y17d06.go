package adventofcode

import (
	"fmt"
	"strings"
)

type y17d06State struct {
	memSize, bigBankIndex, bigBankValue int
	banks                               []int
}

func (s *y17d06State) findLargestBank() {
	hi := s.banks[0]
	for i := 0; i < s.memSize; i++ {
		if s.banks[i] > hi {
			hi = s.banks[i]
		}
	}
	for i := 0; i < s.memSize; i++ {
		if s.banks[i] == hi {
			s.bigBankIndex = i
			s.bigBankValue = hi
			break
		}
	}
}

func (s y17d06State) hash() string {
	hash := strings.Builder{}
	for _, bank := range s.banks {
		hash.WriteString(fmt.Sprintf("%d,", bank))
	}
	return hash.String()
}

func (s *y17d06State) reallocate() {
	s.findLargestBank()
	s.banks[s.bigBankIndex] = 0
	for i := (s.bigBankIndex + 1) % s.memSize; s.bigBankValue > 0; i = (i + 1) % s.memSize {
		s.banks[i]++
		s.bigBankValue--
	}
}

func y17d06(input string, part int) int {
	state := y17d06ParseInput(input)
	seen := map[string]int{}
	i := 0
	for {
		if seen[state.hash()] != 0 {
			if part == 1 {
				return i
			} else {
				return i - seen[state.hash()]
			}
		}
		seen[state.hash()] = i
		state.reallocate()
		i++
	}
}

func y17d06ParseInput(input string) *y17d06State {
	state := &y17d06State{
		memSize: 0,
		banks:   []int{},
	}
	for n := range strings.FieldsSeq(input) {
		state.banks = append(state.banks, strToInt(n))
		state.memSize++
	}
	state.memSize = len(state.banks)
	return state
}
