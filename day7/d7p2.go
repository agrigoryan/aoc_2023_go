package day7

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func d7p2(input string) int {
	lines := strings.Split(input, "\n")
	hands := []*hand{}

	for _, line := range lines {
		var h = parseHand2(line)
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

func parseHand2(str string) *hand {
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

	h.resolveRank2()
	return h
}

func (h *hand) resolveRank2() {
	sortedCounts := make([]int, len(h.counts))
	copy(sortedCounts, h.counts[:])
	numJoker := sortedCounts[cJ]
	if numJoker > 0 {
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
