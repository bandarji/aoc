package d07

import (
	"aoc2023/common"
	"sort"
	"strings"
)

type Hand struct {
	Cards     string
	Bid, Rank int
}

type CardsByRank1 []Hand

func (c CardsByRank1) Len() int      { return len(c) }
func (c CardsByRank1) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CardsByRank1) Less(i, j int) bool {
	if c[i].Rank == c[j].Rank {
		for i, card := range c[i].Cards {
			card2 := rune(c[j].Cards[i])
			if card == card2 {
				continue
			}
			return cardValue1(card) < cardValue1(card2)
		}
		return true
	} else {
		return c[i].Rank < c[j].Rank
	}
}

type CardsByRank2 []Hand

func (c CardsByRank2) Len() int      { return len(c) }
func (c CardsByRank2) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CardsByRank2) Less(i, j int) bool {
	if c[i].Rank == c[j].Rank {
		for i, card := range c[i].Cards {
			card2 := rune(c[j].Cards[i])
			if card == card2 {
				continue
			}
			return cardValue2(card) < cardValue2(card2)
		}
		return true
	} else {
		return c[i].Rank < c[j].Rank
	}
}

const (
	haveHighCard = iota + 1
	haveOnePair
	haveTwoPairs
	haveSet
	haveFullHouse
	haveQuads
	haveFiveOfAKind
)

const TEST1 string = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

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

func HandHighCard(jokers int) (rank int) {
	switch jokers {
	case 5:
		rank = haveFiveOfAKind
	case 4:
		rank = haveFiveOfAKind
	case 3:
		rank = haveQuads
	case 2:
		rank = haveSet
	case 1:
		rank = haveOnePair
	}
	return
}

func HandPair(card rune, counts map[rune]int, jokers int) (rank int) {
	rank = haveOnePair
	if jokers == 3 {
		rank = haveFiveOfAKind
	} else if jokers == 2 {
		rank = haveQuads
	} else {
		for card2, count2 := range counts {
			if card2 == card {
				continue
			}
			if count2 == 3 {
				rank = haveFullHouse
				break
			} else if count2 == 2 {
				if jokers == 1 {
					rank = haveFullHouse
				} else {
					rank = haveTwoPairs
				}
				break
			}
		}
	}
	return
}

func HandQuads(jokers int) (rank int) {
	rank = haveQuads
	if jokers == 1 {
		rank = haveFiveOfAKind
	}
	return
}

func HandSet(card rune, counts map[rune]int, jokers int) (rank int) {
	rank = haveSet
	if jokers == 2 {
		rank = haveFiveOfAKind
	} else if jokers == 1 {
		rank = haveQuads
	} else {
		for card2, count2 := range counts {
			if card2 == card {
				continue
			}
			if count2 == 2 {
				rank = haveFullHouse
				break
			}
		}
	}
	return
}

func RankHand(hand string, part int) (rank int) {
	jokers := 0
	counts := map[rune]int{}
	for _, card := range hand {
		if card == 'J' && part == 2 {
			jokers++
			continue
		}
		counts[card]++
	}
	rank = haveHighCard
	for card, count := range counts {
		switch {
		case count == 5:
			rank = haveFiveOfAKind
		case count == 4:
			rank = HandQuads(jokers)
		case count == 3:
			rank = HandSet(card, counts, jokers)
		case count == 2:
			rank = HandPair(card, counts, jokers)
		}
		if jokers == 1 && rank == haveOnePair {
			rank = haveSet
		}
		// found better than high card
		if rank != haveHighCard {
			break
		}
	}
	if rank == haveHighCard {
		rank = HandHighCard(jokers)
	}
	return
}

func RankHands(input string, part int) (hands []Hand) {
	hands = []Hand{}
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		rank := RankHand(fields[0], part)
		hands = append(hands, Hand{Cards: fields[0], Bid: common.StrToInt(fields[1]), Rank: rank})
	}
	if part == 1 {
		sort.Stable(CardsByRank1(hands))
	} else {
		sort.Stable(CardsByRank2(hands))
	}
	return
}

func Solve(input string, part int) (answer int) {
	hands := RankHands(input, part)
	for rank, hand := range hands {
		answer += (rank + 1) * hand.Bid
	}
	return
}
