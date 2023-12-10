package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	cA = iota
	cK
	cQ
	cT
	c9
	c8
	c7
	c6
	c5
	c4
	c3
	c2
	cJ
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
	'T': cT,
	'9': c9,
	'8': c8,
	'7': c7,
	'6': c6,
	'5': c5,
	'4': c4,
	'3': c3,
	'2': c2,
	'J': cJ,
}

type hand struct {
	cards  [5]int
	bid    int
	counts [13]int
	rank   int
}

func main() {
	content, err := os.ReadFile("d7p2.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)
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
		fmt.Printf("hand : %v\n", h)
		sum += h.bid * (i + 1)
	}

	fmt.Println(sum)
}

func parseHand(str string) *hand {
	h := &hand{}
	numJokers := 0
	for i, r := range str[:5] {
		card := ctoi[r]
		h.cards[i] = card
		h.counts[card] += 1
		if card == cJ {
			numJokers += 1
		}
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
	numJoker := sortedCounts[cJ]
	if numJoker > 0 {
		fmt.Println(numJoker)
		maxCard := 0
		maxCount := 0
		for i, c := range sortedCounts {
			if c > maxCount && i != cJ {
				maxCard = i
				maxCount = c
			}
		}
		sortedCounts[maxCard] += numJoker
		sortedCounts[cJ] = 0
	}
	fmt.Println(sortedCounts)
	slices.SortFunc(sortedCounts, func(a, b int) int { return b - a })
	switch sortedCounts[0] {
	case 5:
		h.rank = rFiveOfAKind
	case 4:
		h.rank = rFourOfAKind
	case 3:
		if sortedCounts[1] >= 2 {
			h.rank = rFullHouse
		} else {
			h.rank = rTheeOfAKind
		}
	case 2:
		if sortedCounts[1] >= 2 {
			h.rank = rTwoPairs
		} else {
			h.rank = rPair
		}
	default:
		h.rank = rHighCard
	}
}
