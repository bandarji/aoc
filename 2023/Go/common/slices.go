package common

func SliceContainsString(sequence []string, search string) bool {
	for _, element := range sequence {
		if element == search {
			return true
		}
	}
	return false
}
