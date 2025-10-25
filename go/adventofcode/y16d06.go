package adventofcode

import "strings"

func y16d06(input string, part int) (answer string) {
	stripes := y16d06AssembleStripes(input)
	msgLeast, msgMost := y16d06AssembleMessage(stripes)
	if part == 1 {
		answer = msgMost
	} else {
		answer = msgLeast
	}
	return
}

func y16d06AssembleStripes(s string) string {
	stripes := []string{}
	entries := map[int]string{}
	for _, entry := range strings.Split(s, "\n") {
		for i, char := range entry {
			entries[i] += string(char)
		}
	}
	for i := 0; i < len(entries); i++ {
		stripes = append(stripes, entries[i])
	}
	return strings.Join(stripes, "\n")
}

func y16d06AssembleMessage(s string) (msgLeast, msgMost string) {
	var most, least string
	for _, line := range strings.Split(s, "\n") {
		most, least = y16d06MostAndLeastCommon(line)
		msgMost += most
		msgLeast += least
	}
	return
}

func y16d06MostAndLeastCommon(s string) (most, least string) {
	frequencies := map[rune]int{}
	lChr, lCnt, mChr, mCnt := ' ', 65535, ' ', 0
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			frequencies[char]++
		}
	}
	for char, count := range frequencies {
		if count > mCnt {
			mChr = char
			mCnt = count
		}
		if count < lCnt {
			lChr = char
			lCnt = count
		}
	}
	most = string(mChr)
	least = string(lChr)
	return
}
