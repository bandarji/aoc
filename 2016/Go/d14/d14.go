package d14

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Day(salt string, part int) int {
	keys := []int{}
	for i := 0; len(keys) < 64; i++ {
		hash := Hash(salt, i, part)
		if triplet := FindTriplet(hash); triplet != "" {
			quint := fmt.Sprintf("%s%s%s%s%s", triplet, triplet, triplet, triplet, triplet)
			if HasQuintuplet(salt, quint, i+1, part) {
				keys = append(keys, i)
			}
		}
	}
	return keys[len(keys)-1]
}

func HasQuintuplet(salt, quint string, index, times int) bool {
	for i := index; i < index+1000; i++ {
		hash := Hash(salt, i, times)
		if strings.Contains(hash, quint) {
			return true
		}
	}
	return false
}

func FindTriplet(s string) (c string) {
	for i := 2; i < len(s); i++ {
		if s[i] == s[i-2] && s[i] == s[i-1] {
			return string(s[i])
		}
	}
	return
}

func Hash(salt string, index, part int) (hash string) {
	hash = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", salt, index))))
	if part == 2 {
		for i := 0; i < 2016; i++ {
			hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
		}
	}
	return
}
