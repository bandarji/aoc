package common

func Min(values ...int) (minimum int) {
	if len(values) < 1 {
		return 0
	}
	minimum = values[0]
	for _, value := range values {
		if value < minimum {
			minimum = value
		}
	}
	return
}
