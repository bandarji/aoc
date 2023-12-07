package d07

import (
	"aoc2023/common"
	"sort"
	"strings"
)

type Hand struct {
	Cards           string
	Bid, Multiplier int
}

const TEST1 string = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

type CardsByRank1 []Hand

func (c CardsByRank1) Len() int      { return len(c) }
func (c CardsByRank1) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CardsByRank1) Less(i, j int) bool {
	if c[i].Multiplier == c[j].Multiplier {
		for i, card := range c[i].Cards {
			card2 := rune(c[j].Cards[i])
			if card == card2 {
				continue
			}
			return cardValue1(card) < cardValue1(card2)
		}
		return true
	} else {
		return c[i].Multiplier < c[j].Multiplier
	}
}

type CardsByRank2 []Hand

func (c CardsByRank2) Len() int      { return len(c) }
func (c CardsByRank2) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CardsByRank2) Less(i, j int) bool {
	if c[i].Multiplier == c[j].Multiplier {
		for i, card := range c[i].Cards {
			card2 := rune(c[j].Cards[i])
			if card == card2 {
				continue
			}
			return cardValue2(card) < cardValue2(card2)
		}
		return true
	} else {
		return c[i].Multiplier < c[j].Multiplier
	}
}

func DetermineMultiplier(hand string, part int) (m int) {
	jokers := 0
	counts := map[rune]int{}
	for _, card := range hand {
		if card == 'J' && part == 2 {
			jokers++
			continue
		}
		counts[card]++
	}
	m = 1
	for card, count := range counts {
		switch {
		case count == 5:
			m = 7
		case count == 4:
			if jokers == 1 {
				m = 7
			} else {
				m = 6
			}
		case count == 3:
			if jokers == 2 {
				m = 7
			} else if jokers == 1 {
				m = 6
			} else {
				m = 4
				for card2, count2 := range counts {
					if card2 == card {
						continue
					}
					if count2 == 2 {
						m = 5
						break
					}
				}
			}
		case count == 2:
			if jokers == 3 {
				m = 7
			} else if jokers == 2 {
				m = 6
			} else {
				m = 2
				for card2, count2 := range counts {
					if card2 == card {
						continue
					}
					if count2 == 3 {
						m = 5
						break
					} else if count2 == 2 {
						if jokers == 1 {
							m = 5
						} else {
							m = 3
						}
						break
					}
				}
			}
		}
		if jokers == 1 && m == 2 {
			m = 4
		}
		// found better than high card
		if m != 1 {
			break
		}

	}
	if m == 1 {
		switch jokers {
		case 5:
			m = 7
		case 4:
			m = 7
		case 3:
			m = 6
		case 2:
			m = 4
		case 1:
			m = 2
		}
	}
	return
}

func RankHands(input string, part int) (hands []Hand) {
	hands = []Hand{}
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		multiplier := DetermineMultiplier(fields[0], part)
		hands = append(hands, Hand{Cards: fields[0], Bid: common.StrToInt(fields[1]), Multiplier: multiplier})
	}
	if part == 1 {
		sort.Stable(CardsByRank1(hands))
	} else {
		sort.Stable(CardsByRank2(hands))
	}
	return
}

func cardValue1(card rune) (v int) {
	v = int(card - '0')
	values := map[rune]int{'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10}
	if value, exists := values[card]; exists {
		v = value
	}
	return
}

func cardValue2(card rune) int {
	if card == 'J' {
		return 1
	}
	values := map[rune]int{'A': 14, 'K': 13, 'Q': 12, 'J': 1, 'T': 10}
	if value, exists := values[card]; exists {
		return value
	}
	return int(card - '0')
}

func Solve(input string, part int) (answer int) {
	hands := RankHands(input, part)
	for rank, hand := range hands {
		answer += (rank + 1) * hand.Bid
	}
	return
}
