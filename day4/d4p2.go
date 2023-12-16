package day4

import (
	"fmt"
	"slices"
	"strings"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func d4p2(input string) int {
	lines := strings.Split(input, "\n")
	counts := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		counts[i] = 1
	}

	trimSpaces := aocutils.Mapper(func(str string) string { return strings.TrimSpace(str) })
	filterEmpty := aocutils.Filterer(func(str string) bool { return len(str) > 0 })

	for i, line := range strings.Split(input, "\n") {
		ticket := strings.Split(line, ":")[1]
		parts := strings.Split(ticket, "|")
		winningNumbers := filterEmpty(trimSpaces(strings.Split(parts[0], " ")))
		numbers := filterEmpty(trimSpaces(strings.Split(parts[1], " ")))
		matches := 0
		for _, num := range numbers {
			if slices.Contains(winningNumbers, num) {
				matches += 1
			}
		}
		for j := i + 1; j <= i+matches; j++ {
			counts[j] += counts[i]
		}
	}
	numTickets := 0
	for _, num := range counts {
		numTickets += num
	}
	fmt.Println(numTickets)
	return numTickets
}
