package d04

import (
	"fmt"
	"sort"
	"strings"
)

type Entry struct {
	Name, Checksum string
	SectorID       int
}

func ParseEntry(s string) (e Entry) {
	sections := strings.Split(s, "-")
	fmt.Sscanf(sections[len(sections)-1], "%d[%5s]", &e.SectorID, &e.Checksum)
	e.Name = strings.Join(sections[:len(sections)-1], "")
	return
}

func RealRoom(s string) bool {
	entry := ParseEntry(s)
	frequencies := map[string]int{}
	for _, char := range entry.Name {
		if char >= 97 && char <= 122 {
			frequencies[string(char)]++
		}
	}
	theCounts := []int{}
	for _, count := range frequencies {
		theCounts = append(theCounts, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(theCounts)))
	counts := []int{}
	for i, char := range entry.Checksum {
		_, exists := frequencies[string(char)]
		if !exists {
			return false
		}
		counts = append(counts, frequencies[string(char)])
		if counts[i] != theCounts[i] {
			return false
		}
		if i != 0 {
			if counts[i-1] < counts[i] || (counts[i-1] == counts[i] && string(entry.Checksum[i-1]) > string(char)) {
				return false
			}
		}
	}
	return true
}

func Caesar(s string, r int) (c string) {
	for _, char := range s {
		value := int(char) - int('a')
		value += r
		value %= 26
		value += int('a')
		c += string(rune(value))
	}
	return
}

func Day(input string, part int) int {
	sum := 0
	for _, entry := range strings.Split(input, "\n") {
		if RealRoom(entry) {
			e := ParseEntry(entry)
			if part == 2 {
				e.Name = Caesar(e.Name, e.SectorID)
				if e.Name == "northpoleobjectstorage" {
					return e.SectorID
				}
			}
			sum += e.SectorID
		}
	}
	return sum
}
