package common

import (
	"log"
	"strconv"
)

func StrToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Could not convert '%s' to an integer", s)
	}
	return n
}
