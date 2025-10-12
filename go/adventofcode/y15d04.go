package adventofcode

import (
	"crypto/md5"
	"fmt"
	"strings"
)

const y15d04input string = "yzbqklnj"

func y15d04(input string, zeroes int) (number int) {
	for {
		number++
		data := []byte(fmt.Sprintf("%s%d", input, number))
		s := fmt.Sprintf("%x", md5.Sum(data))
		if strings.Repeat("0", zeroes) == s[:zeroes] {
			break
		}
	}
	return
}
