package common

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func AddCommas(n int) (s string) {
	p := message.NewPrinter(language.English)
	s = p.Sprintf("%d", n)
	return
}

func GCD(n1, n2 int) int {
	for n2 != 0 {
		t := n2
		n2 = n1 % n2
		n1 = t
	}
	return n1
}

func AddUpInts(ints ...int) (total int) {
	for _, number := range ints {
		total += number
	}
	return
}
