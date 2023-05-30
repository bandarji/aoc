package d06

import "strings"

type CommonChar struct {
	Char  string
	Count int
}

func LeastCommon(s string) string {
	lc := CommonChar{"", 65536}
	frequencies := map[string]int{}
	for _, char := range s {
		frequencies[string(char)]++
	}
	for char, count := range frequencies {
		if count < lc.Count {
			lc.Char = char
			lc.Count = count
		}
	}
	return lc.Char
}

func MostCommon(s string) string {
	mc := CommonChar{"", 0}
	frequencies := map[string]int{}
	for _, char := range s {
		frequencies[string(char)]++
	}
	for char, count := range frequencies {
		if count > mc.Count {
			mc.Char = char
			mc.Count = count
		}
	}
	return mc.Char
}

func AssembleStripes(input string) []string {
	stripes := []string{}
	entries := map[int]string{}
	for _, entry := range strings.Split(input, "\n") {
		for charNum, char := range entry {
			entries[charNum] += string(char)
		}
	}
	for i := 0; i < len(entries); i++ {
		stripes = append(stripes, entries[i])
	}
	return stripes
}

func Day(input string, part int) (message string) {
	stripes := AssembleStripes(input)
	for _, stripe := range stripes {
		if part == 1 {
			message += MostCommon(stripe)
		} else {
			message += LeastCommon(stripe)
		}
	}
	return
}
