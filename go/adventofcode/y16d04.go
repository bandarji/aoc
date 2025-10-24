package adventofcode

import (
	"fmt"
	"sort"
	"strings"
)

type y16d04Room struct {
	name, checksum string
	id             int
}

func y16d04(input string, part int) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		room := y16d04ParseRoom(line)
		if y16d04RealRoom(room) {
			if part == 2 {
				decrypted := y16d04Decrypt(room.name, room.id)
				// fmt.Println(decrypted)
				if decrypted == "northpole object storage" {
					return room.id
				}
			}
			sum += room.id
		}
	}
	return sum
}

func y16d04Decrypt(name string, id int) string {
	decrypted := strings.Builder{}
	for _, char := range name {
		if char >= 'a' && char <= 'z' {
			decrypted.WriteRune(rune((int(char)-int('a')+id)%26 + int('a')))
		} else {
			decrypted.WriteRune(' ')
		}
	}
	return decrypted.String()
}

func y16d04ParseRoom(line string) (room y16d04Room) {
	sections := strings.Split(line, "-")
	fmt.Sscanf(sections[len(sections)-1], "%d[%5s]", &room.id, &room.checksum)
	room.name = strings.Join(sections[:len(sections)-1], "-")
	return
}

func y16d04RealRoom(room y16d04Room) bool {
	frequencies := map[rune]int{}
	charCounts, counts := []int{}, []int{}
	for _, char := range room.name {
		if char >= 'a' && char <= 'z' {
			frequencies[char]++
		}
	}
	for _, count := range frequencies {
		charCounts = append(charCounts, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(charCounts)))
	for i, c := range room.checksum {
		_, exists := frequencies[rune(c)]
		if !exists {
			return false
		}
		counts = append(counts, frequencies[rune(c)])
		if i != 0 {
			if counts[i-1] < counts[i] || (counts[i-1] == counts[i] && string(room.checksum[i-1]) > string(c)) {
				return false
			}
		}
	}
	return true
}
