package d12

import (
	"aoc2023/common"
	"fmt"
	"strings"
	"sync"
)

var Cache = make(map[string]int)
var mutex = &sync.RWMutex{}

const TEST1 string = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

// operational (.) or damaged (#) or unknown (?)

func UnfoldPattern(s string) (unfolded string) {
	sections := []string{}
	for i := 0; i < 5; i++ {
		sections = append(sections, s)
	}
	unfolded = strings.Join(sections, "?")
	return
}

func UnfoldGroups(g []int) (unfolded []int) {
	unfolded = []int{}
	for i := 0; i < 5; i++ {
		unfolded = append(unfolded, g...)
	}
	return
}

func ProcessRow(ch chan int, row string, index, part int) {
	groups := []int{}
	sections := strings.Fields(row)
	pattern := sections[0]
	for _, group := range strings.Split(sections[1], ",") {
		groups = append(groups, common.StrToInt(group))
	}
	if part == 2 {
		pattern = UnfoldPattern(pattern)
		groups = UnfoldGroups(groups)
	}
	ch <- CountArrangements(pattern, groups, false)
}

func DecrementedNum0(nums []int) (decremented []int) {
	decremented = []int{}
	for i, num := range nums {
		if i == 0 {
			decremented = append(decremented, num-1)
		} else {
			decremented = append(decremented, num)
		}
	}
	return
}

func CacheSet(k string, v int) int {
	mutex.Lock()
	Cache[k] = v
	mutex.Unlock()
	return v
}

func CountArrangements(pattern string, groups []int, toggle bool) int {
	var count int
	key := fmt.Sprintf("%s-%+v-%v", pattern, groups, toggle)
	mutex.Lock()
	if value, exists := Cache[key]; exists {
		mutex.Unlock()
		return value
	} else {
		mutex.Unlock()
	}
	switch {
	case len(pattern) == 0:
		if common.AddUpInts(groups...) == 0 {
			return 1
		}
		return 0
	case common.AddUpInts(groups...) == 0:
		if strings.Contains(pattern, "#") {
			return 0
		}
		return 1
	case pattern[0] == '#':
		if groups[0] == 0 && toggle {
			return 0
		}
		decGroups := DecrementedNum0(groups)
		count = CountArrangements(pattern[1:], decGroups, true)
		return CacheSet(fmt.Sprintf("%s-%+v-%v", pattern[1:], decGroups, true), count)
	case pattern[0] == '.':
		if groups[0] > 0 && toggle {
			return 0
		}
		if groups[0] == 0 {
			count = CountArrangements(pattern[1:], groups[1:], false)
			return CacheSet(fmt.Sprintf("%s-%+v-%v", pattern[1:], groups[1:], false), count)
		}
		count = CountArrangements(pattern[1:], groups, false)
		return CacheSet(fmt.Sprintf("%s-%+v-%v", pattern[1:], groups, false), count)
	case toggle:
		if groups[0] == 0 {
			count = CountArrangements(pattern[1:], groups[1:], false)
			return CacheSet(fmt.Sprintf("%s-%+v-%v", pattern[1:], groups[1:], false), count)
		}
		decGroups := DecrementedNum0(groups)
		count = CountArrangements(pattern[1:], DecrementedNum0(groups), true)
		return CacheSet(fmt.Sprintf("%s-%+v-%v", pattern[1:], decGroups, true), count)
	} // esac
	return CountArrangements(pattern[1:], groups, false) + CountArrangements(pattern[1:], DecrementedNum0(groups), true)
}

func Solve(input string, part int) (answer int) {
	ch := make(chan int)
	rows := strings.Split(input, "\n")
	counts := make([]int, len(rows))
	for i, row := range rows {
		go ProcessRow(ch, row, i, part)

	}
	for i := range counts {
		counts[i] = <-ch
	}
	answer = common.AddUpInts(counts...)
	return
}
