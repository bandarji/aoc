package adventofcode

import (
	"crypto/md5"
	"fmt"
	"strings"
)

const y16d14Salt string = "jlmsuwbz"

func y16d14(input string, part int) (answer int) {
	hashes := []string{}
	keys := []int{}
	cycles := 1
	if part == 2 {
		cycles += 2016
	}
	for i := range 1000 {
		hashes = append(hashes, y16d14GetHash(input, i, cycles))
	}
	for i := 0; len(keys) < 64; i++ {
		current := hashes[0]
		hashes = append(hashes, y16d14GetHash(input, i+1000, cycles))
		hashes = hashes[1:]
		if c := y16d14HasThree(current); c != "" {
			if y16d14HasFive(strings.Join(hashes, " "), c) {
				keys = append(keys, i)
			}
		}
	}
	answer = keys[len(keys)-1]
	return
}

func y16d14GetHash(input string, index int, cycles int) string {
	hash := fmt.Sprintf("%s%d", input, index)
	for i := 0; i < cycles; i++ {
		hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	}
	return hash
}

func y16d14HasThree(hash string) string {
	for i := 0; i < len(hash)-2; i++ {
		if hash[i] == hash[i+1] && hash[i] == hash[i+2] {
			return string(hash[i])
		}
	}
	return ""
}

func y16d14HasFive(hash string, char string) bool {
	return strings.Contains(hash, strings.Repeat(char, 5))
}
