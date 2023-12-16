package day11

import (
	"strings"
)

type pos struct {
	r, c int
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func isBetween(x, a, b int) bool {
	min := a
	max := b
	if a > b {
		min = b
		max = a
	}
	return min < x && x < max
}

func calcSumDistance(input string, expAmount int) int {
	lines := strings.Split(input, "\n")
	expColumns := []int{}
	expRows := []int{}
	galaxies := []pos{}

	for row, line := range lines {
		shouldExpand := true
		for col, rune := range line {
			if rune != '.' {
				shouldExpand = false
				galaxies = append(galaxies, pos{
					r: row,
					c: col,
				})
			}
		}
		if shouldExpand {
			expRows = append(expRows, row)
		}
	}

	for col := 0; col < len(lines[0]); col++ {
		shouldExpand := true
		for _, line := range lines {
			if line[col] != '.' {
				shouldExpand = false
				break
			}
		}
		if shouldExpand {
			expColumns = append(expColumns, col)
		}
	}

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		ga := galaxies[i]
		for j := i; j < len(galaxies); j++ {
			gb := galaxies[j]
			d := abs(ga.c-gb.c) + abs(ga.r-gb.r)
			for _, row := range expRows {
				if isBetween(row, ga.r, gb.r) {
					d += expAmount
				}
			}
			for _, col := range expColumns {
				if isBetween(col, ga.c, gb.c) {
					d += expAmount
				}
			}
			sum += d
		}
	}

	return sum
}

func d11p1(input string) int {
	return calcSumDistance(input, 1)
}
