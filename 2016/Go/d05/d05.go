package d05

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func Hash(s string, i int) (string, bool) {
	hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", s, i))))
	if strings.HasPrefix(hash, "00000") {
		return string(hash[5:6]), true
	}
	return "", false
}

func CompoundHash(s string) (password string) {
	i := 0
	var (
		char  string
		yesno bool
	)
	for {
		char, yesno = Hash(s, i)
		if yesno {
			password += char
			if len(password) == 8 {
				break
			}
		}
		i++
	}
	return
}

func BuildPassword(s string) (password string) {
	code := map[int]string{}
	for i := 0; len(code) < 8; i++ {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", s, i))))
		if strings.HasPrefix(hash, "00000") {
			if hash[5] >= 48 && hash[5] <= 55 {
				index, _ := strconv.Atoi(string(hash[5]))
				if _, exists := code[index]; !exists {
					code[index] = hash[6:7]
				}
			}
		}
	}
	for i := 0; i < 8; i++ {
		password += code[i]
	}
	return
}

func Day(input string, part int) string {
	password := ""
	if part == 1 {
		password = CompoundHash(input)
	} else {
		password = BuildPassword(input)
	}
	return password
}
