package d07

import (
	"fmt"
	"strings"
)

type AddressSections struct {
	Inners, Outers []string
}

func AssembleAddress(ipv7 string) AddressSections {
	addr := AddressSections{}
	tmp := ""
	for _, char := range ipv7 + "[" {
		if char == '[' {
			addr.Outers = append(addr.Outers, tmp)
			tmp = ""
		} else if char == ']' {
			addr.Inners = append(addr.Inners, tmp)
			tmp = ""
		} else {
			tmp += string(char)
		}
	}
	// addr.Outers = append(addr.Outers, tmp)
	return addr
}

func HasABBA(strings []string) bool {
	for _, s := range strings {
		for i := 3; i < len(s); i++ {
			if s[i-3] == s[i] && s[i-2] == s[i-1] && s[i] != s[i-1] {
				return true
			}
		}
	}
	return false
}

func IsValidIPv7(a AddressSections) bool {
	if HasABBA(a.Inners) {
		return false
	}
	if HasABBA(a.Outers) {
		return true
	}
	return false
}

type AreaBroadcastAccessor struct {
	ABA, BAB string
}

func FindABAs(outers []string) []AreaBroadcastAccessor {
	ABAs := []AreaBroadcastAccessor{}
	for _, outer := range outers {
		for i := 2; i < len(outer); i++ {
			if outer[i-2] != outer[i-1] && outer[i-2] == outer[i] {
				aba := AreaBroadcastAccessor{}
				aba.ABA = outer[i-2 : i+1]
				aba.BAB = fmt.Sprintf("%s%s%s", string(outer[i-1]), string(outer[i]), string(outer[i-1]))
				ABAs = append(ABAs, aba)
			}
		}
	}
	return ABAs
}

func SupportsSSL(a AddressSections) bool {
	ABAs := FindABAs(a.Outers)
	for _, inner := range a.Inners {
		for _, sections := range ABAs {
			if strings.Contains(inner, sections.BAB) {
				return true
			}
		}
	}
	return false
}

func CountValidIPv7Addresses(input string) (count int) {
	for _, entry := range strings.Split(input, "\n") {
		address := AssembleAddress(entry)
		if IsValidIPv7(address) {
			count++
		}
	}
	return
}

func CountSSL(input string) (count int) {
	for _, entry := range strings.Split(input, "\n") {
		address := AssembleAddress(entry)
		if SupportsSSL(address) {
			count++
		}
	}
	return
}

func Day(input string, part int) (count int) {
	if part == 1 {
		count = CountValidIPv7Addresses(input)
	} else {
		count = CountSSL(input)
	}
	return
}
