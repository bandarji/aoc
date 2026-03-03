package adventofcode

const y16d19Input string = "3004953"

type y16d19Elf struct {
	number int
	gifts  int
	next   *y16d19Elf
}

func y16d19(input string, part int) (elfNumber int) {
	elfNumber = y16d19FindElfWithGifts(strToInt(input), part)
	return
}

func y16d19FindElfWithGifts(totalElves, part int) int {
	root := &y16d19Elf{number: 1, gifts: 1, next: nil}
	cursor := root
	for i := 2; i <= totalElves; i++ {
		cursor.next = &y16d19Elf{number: i, gifts: 1, next: nil}
		cursor = cursor.next
	}
	cursor.next = root
	if part == 1 {
		for root.next != root {
			root.gifts += root.next.gifts
			root.next = root.next.next
			root = root.next
		}
		return root.number
	}

	odd := totalElves&1 == 1
	across := root
	for i := 0; i < totalElves/2-1; i++ {
		across = across.next
	}
	for root.next != root {
		root.gifts += across.next.gifts
		across.next = across.next.next
		if odd {
			across = across.next
		}
		odd = !odd
		root = root.next
	}
	return root.number
}
