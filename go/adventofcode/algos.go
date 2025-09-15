package adventofcode

func permutations(items []string, n int, result *[][]string) {
	if n == 1 {
		*result = append(*result, append([]string{}, items...))
	} else {
		for i := 0; i < n; i++ {
			permutations(items, n-1, result)
			if n%2 == 0 {
				items[i], items[n-1] = items[n-1], items[i]
			} else {
				items[0], items[n-1] = items[n-1], items[0]
			}
		}
	}
}
