package adventofcode

import (
	"sort"
	"strings"
)

func y17d04(input string, part int) (count int) {
	for phrase := range strings.SplitSeq(input, "\n") {
		if part == 1 {
			if y17d04IsValidPassphrase(phrase) {
				count++
			}
		} else {
			if y17d04IsValidPassphraseCheckAnagrams(phrase) {
				count++
			}
		}
	}
	return
}

func y17d04IsValidPassphrase(phrase string) bool {
	words := map[string]bool{}
	for word := range strings.FieldsSeq(phrase) {
		if words[word] {
			return false
		}
		words[word] = true
	}
	return true
}

func y17d04IsValidPassphraseCheckAnagrams(phrase string) bool {
	words := map[string]bool{}
	for word := range strings.FieldsSeq(phrase) {
		letters := strings.Split(word, "")
		sort.Strings(letters)
		anaWord := strings.Join(letters, "")
		if words[anaWord] {
			return false
		}
		words[anaWord] = true
	}
	return true
}
