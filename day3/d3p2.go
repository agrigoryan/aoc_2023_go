package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	content, err := os.ReadFile("d3p2.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)
	lines := strings.Split(input, "\n")

	const empty = '.'

	hasSymbol := func(parts []string) bool {
		for _, p := range parts {
			for _, r := range p {
				if !unicode.IsDigit(r) && r != empty {
					return true
				}
			}
		}
		return false
	}

	type partNumber struct {
		numLine int
		value   int
		start   int
		end     int
	}

	parts := []partNumber{}

	digitFound := func(numLine int, line string, start int, end int) {
		pStart := max(start-1, 0)
		pEnd := min(end+1, len(line))
		partsToCheck := []string{line[pStart:pEnd]}

		if numLine > 0 {
			partsToCheck = append(partsToCheck, lines[numLine-1][pStart:pEnd])
		}
		if numLine < len(lines)-1 {
			partsToCheck = append(partsToCheck, lines[numLine+1][pStart:pEnd])
		}

		if hasSymbol(partsToCheck) {
			if num, err := strconv.Atoi(line[start:end]); err == nil {
				parts = append(parts, partNumber{
					numLine: numLine,
					value:   num,
					start:   start,
					end:     end,
				})
			}
		}
	}

	for i, line := range lines {
		digitStart := -1
		for j, r := range line {
			if unicode.IsDigit(r) {
				if digitStart == -1 {
					digitStart = j
				}
			} else if digitStart != -1 {
				digitFound(i, line, digitStart, j)
				digitStart = -1
			}
		}
		if digitStart != -1 {
			digitFound(i, line, digitStart, len(line))
		}
	}

	sum := 0
	for i, line := range lines {
		for j, r := range line {
			if r == '*' {
				adjacentPartNumbers := []int{}
				for _, p := range parts {
					if p.numLine < i+2 && p.numLine > i-2 && j > p.start-2 && j < p.end+1 {
						adjacentPartNumbers = append(adjacentPartNumbers, p.value)
					}
				}
				if len(adjacentPartNumbers) == 2 {
					sum += adjacentPartNumbers[0] * adjacentPartNumbers[1]
				}
			}
		}
	}
	fmt.Println(sum)
}
