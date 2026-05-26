package adventofcode

func y17d09(input string, part int) int {
	score, openers, garChars := 0, 0, 0
	garbaged := false
	for i := 0; i < len(input); i++ {
		ch := input[i]
		switch garbaged {
		case true:
			switch ch {
			case '!':
				i++
			case '>':
				garbaged = false
			default:
				garChars++
			}
		case false:
			switch ch {
			case '{':
				openers++
			case '<':
				garbaged = true
			case '}':
				score += openers
				openers--
			}
		}
	}
	if part == 1 {
		return score
	}
	return garChars
}
