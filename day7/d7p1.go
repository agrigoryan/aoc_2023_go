package day7

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

const (
	cA = iota
	cK
	cQ
	cJ
	cT
	c9
	c8
	c7
	c6
	c5
	c4
	c3
	c2
)

const (
	rFiveOfAKind = iota
	rFourOfAKind
	rFullHouse
	rTheeOfAKind
	rTwoPairs
	rPair
	rHighCard
)

var ctoi = map[rune]int{
	'A': cA,
	'K': cK,
	'Q': cQ,
	'J': cJ,
	'T': cT,
	'9': c9,
	'8': c8,
	'7': c7,
	'6': c6,
	'5': c5,
	'4': c4,
	'3': c3,
	'2': c2,
}

type hand struct {
	cards  [5]int
	bid    int
	counts [13]int
	rank   int
}

func d7p1(input string) int {
	lines := strings.Split(input, "\n")
	hands := []*hand{}

	for _, line := range lines {
		var h = parseHand(line)
		hands = append(hands, h)
	}

	slices.SortFunc(hands, func(a, b *hand) int {
		if a.rank != b.rank {
			return b.rank - a.rank
		}
		for i := 0; i < len(a.cards); i++ {
			if a.cards[i] != b.cards[i] {
				return b.cards[i] - a.cards[i]
			}
		}
		return 0
	})

	sum := 0
	for i, h := range hands {
		sum += h.bid * (i + 1)
	}

	fmt.Println(sum)

	return sum
}

func parseHand(str string) *hand {
	h := &hand{}
	for i, r := range str[:5] {
		h.cards[i] = ctoi[r]
		h.counts[ctoi[r]] += 1
		if bid, err := strconv.Atoi(str[6:]); err != nil {
			log.Fatal("error parsing bid", err)
		} else {
			h.bid = bid
		}
	}
	h.resolveRank()
	return h
}

func (h *hand) resolveRank() {
	sortedCounts := make([]int, len(h.counts))
	copy(sortedCounts, h.counts[:])
	slices.SortFunc(sortedCounts, func(a, b int) int { return b - a })
	switch sortedCounts[0] {
	case 5:
		h.rank = rFiveOfAKind
	case 4:
		h.rank = rFourOfAKind
	case 3:
		if sortedCounts[1] == 2 {
			h.rank = rFullHouse
		} else {
			h.rank = rTheeOfAKind
		}
	case 2:
		if sortedCounts[1] == 2 {
			h.rank = rTwoPairs
		} else {
			h.rank = rPair
		}
	default:
		h.rank = rHighCard
	}
}
