package adventofcode

type y17d17Node struct {
	value int
	next  *y17d17Node
}

const y17d17StartValue string = "328"

func y17d17(input string, valueToInsert, valueToFind int) (answer int) {
	steps := strToInt(input)
	curr := &y17d17Node{value: 0}
	curr.next = curr
	for i := 1; i <= valueToInsert; i++ {
		for range steps {
			curr = curr.next
		}
		next := curr.next
		curr.next = &y17d17Node{value: i, next: next}
		curr = curr.next
	}
	for curr.value != valueToFind {
		curr = curr.next
	}
	return curr.next.value
}
