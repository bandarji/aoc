package adventofcode

import (
	"crypto/md5"
	"fmt"
	"strings"
)

const y16d05Input = "uqwqemis"

func y16d05(input string, part int) (code string) {
	if part == 1 {
		code = y16d05Part1(input)
	} else {
		code = y16d05Part2(input)
	}
	return
}

func y16d05Part1(input string) (pw string) {
	hash := ""
	for i := 0; len(pw) < 8; i++ {
		hash = y16d05GetHash(input, i)
		if strings.HasPrefix(hash, "00000") {
			pw += string(hash[5])
		}
	}
	return
}

func y16d05Part2(input string) (pw string) {
	hash := ""
	code := map[int]string{}
	for i := 0; len(code) < 8; i++ {
		hash = y16d05GetHash(input, i)
		if strings.HasPrefix(hash, "00000") && hash[5] >= '0' && hash[5] <= '7' {
			pos := strToInt(string(hash[5]))
			if _, ok := code[pos]; !ok {
				code[pos] = string(hash[6])
			}
		}
	}
	for i := range 8 {
		pw += code[i]
	}
	return
}

func y16d05GetHash(s string, i int) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", s, i))))
}
