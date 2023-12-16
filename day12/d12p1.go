package day12

import (
	"log"
	"strconv"
	"strings"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func solveLine(springs string, series []int) int {
	stack := []string{springs}
	sum := 0
	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		countDamaged := 0
		seriesIdx := 0
		shouldContinue := false
		for i, r := range s {
			if r == '#' {
				countDamaged += 1
			}
			if r == '.' && countDamaged > 0 {
				if seriesIdx >= len(series) || countDamaged != series[seriesIdx] {
					shouldContinue = true
					break
				}
				seriesIdx += 1
				countDamaged = 0
			}
			if r == '?' {
				stack = append(stack, s[:i]+"."+s[i+1:])
				stack = append(stack, s[:i]+"#"+s[i+1:])
				shouldContinue = true
				break
			}
		}

		if shouldContinue {
			continue
		}

		if countDamaged > 0 {
			if seriesIdx >= len(series) || countDamaged != series[seriesIdx] {
				continue
			}
			seriesIdx += 1
		}

		if seriesIdx != len(series) {
			continue
		}

		sum += 1
	}

	return sum
}

func d12p1(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		springs := parts[0]
		series := aocutils.Map(strings.Split(parts[1], ","), func(s string) int {
			res, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalln(err)
			}
			return res
		})
		sum += solveLine(springs, series)
	}

	return sum
}
