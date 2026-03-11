package adventofcode

import (
	"fmt"
	"math/rand"
	"strings"
)

func combinationsInt(ints []int, n int) [][]int {
	combos := [][]int{}
	for i := 0; i < len(ints); i++ {
		combos = append(combos, makeCombos(ints[i:], n, []int{})...)
	}
	return combos
}

func makeCombos(ints []int, n int, known []int) [][]int {
	if len(known) == n {
		return [][]int{append([]int{}, known...)}
	}
	combos := [][]int{}
	for i := range ints {
		known = append(known, ints[i])
		newCombos := makeCombos(ints[i+1:], n, known)
		combos = append(combos, newCombos...)
		known = known[:len(known)-1] // backtrack
	}
	return combos
}

func permutations(items []string, n int, result *[][]string) {
	if n == 1 {
		*result = append(*result, append([]string{}, items...))
	} else {
		for i := 0; i < n; i++ {
			permutations(items, n-1, result)
			if n%2 == 0 {
				items[i], items[n-1] = items[n-1], items[i]
			} else {
				items[0], items[n-1] = items[n-1], items[0]
			}
		}
	}
}

// shuffleIndexes returns a shuffled list of indexes and a
// string representation of the indexes separated by commas (good for hashing)
func shuffleIndexes(n int) ([]int, string) {
	sb := strings.Builder{}
	shuffledIndexes := rand.Perm(n)
	for _, index := range shuffledIndexes {
		sb.WriteString(fmt.Sprintf("%d,", index))
	}
	return shuffledIndexes, strings.TrimSuffix(sb.String(), ",")
}
