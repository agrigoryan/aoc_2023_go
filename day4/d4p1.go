package day4

import (
	"fmt"
	"slices"
	"strings"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func d4p1(input string) int {
	sum := 0

	trimSpaces := aocutils.Mapper(func(str string) string { return strings.TrimSpace(str) })
	filterEmpty := aocutils.Filterer(func(str string) bool { return len(str) > 0 })

	for _, line := range strings.Split(input, "\n") {
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
		if matches > 0 {
			worth := 1 << (matches - 1)
			sum += worth
		}
	}
	fmt.Println(sum)
	return sum
}
