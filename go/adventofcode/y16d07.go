package adventofcode

import (
	"fmt"
	"strings"
)

type y16d07ipv7Address struct {
	inners, outers []string
}

type y16d07AreaBrodcastAccessor struct {
	aba, bab string
}

func y16d07(input string, part int) (answer int) {
	tlsCount, sslCount := 0, 0
	for _, entry := range strings.Split(input, "\n") {
		address := y16d07ParseAddress(entry)
		if y16d07IsTLS(address) {
			tlsCount++
		}
		if y16d07IsSSL(address) {
			sslCount++
		}
	}
	if part == 1 {
		answer = tlsCount
	} else {
		answer = sslCount
	}
	return
}

func y16d07ParseAddress(entry string) (address y16d07ipv7Address) {
	tmp := ""
	for _, char := range entry {
		if char == '[' {
			address.outers = append(address.outers, tmp)
			tmp = ""
		} else if char == ']' {
			address.inners = append(address.inners, tmp)
			tmp = ""
		} else {
			tmp += string(char)
		}
	}
	if len(tmp) > 0 {
		address.outers = append(address.outers, tmp)
	}
	return
}

func y16d07IsTLS(address y16d07ipv7Address) bool {
	if y16d07HasABBA(address.inners) {
		return false
	}
	if y16d07HasABBA(address.outers) {
		return true
	}
	return false
}

func y16d07HasABBA(strings []string) bool {
	for _, s := range strings {
		for i := 3; i < len(s); i++ {
			if s[i-3] == s[i] && s[i-2] == s[i-1] && s[i] != s[i-1] {
				return true
			}
		}
	}
	return false
}

func y16d07IsSSL(address y16d07ipv7Address) bool {
	abas := y16d07FindABAs(address.outers)
	for _, inner := range address.inners {
		for _, aba := range abas {
			if strings.Contains(inner, aba.bab) {
				return true
			}
		}
	}
	return false
}

func y16d07FindABAs(outers []string) (abas []y16d07AreaBrodcastAccessor) {
	for _, s := range outers {
		for i := 2; i < len(s); i++ {
			if s[i-2] != s[i-1] && s[i-2] == s[i] {
				abas = append(
					abas,
					y16d07AreaBrodcastAccessor{aba: s[i-2 : i+1], bab: fmt.Sprintf("%c%c%c", s[i-1], s[i], s[i-1])},
				)
			}
		}
	}
	return
}
