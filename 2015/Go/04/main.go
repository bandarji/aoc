package main

import (
	"crypto/md5"
	"fmt"
	"log"
)

const KEY string = "yzbqklnj"

func Day04(zeroes int) (number int) {
	for {
		number++
		data := []byte(fmt.Sprintf("%s%d", KEY, number))
		s := fmt.Sprintf("%x", md5.Sum(data))
		if CountLeadZeroes(&s, zeroes) {
			break
		}
	}
	return
}

func CountLeadZeroes(s *string, z int) bool {
	zs := ""
	for i := 0; i < z; i++ {
		zs += "0"
	}
	return zs == string((*s)[0:z])
}

func main() {
	day, part, zeroes := 4, 1, 5
	first := Day04(zeroes)
	log.Printf("[%02d / %02d]: %d\n", day, part, first)
	part, zeroes = 2, 6
	first = Day04(zeroes)
	log.Printf("[%02d / %02d]: %d\n", day, part, first)
}
