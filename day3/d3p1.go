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
	content, err := os.ReadFile("d3p1.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)
	lines := strings.Split(input, "\n")

	sum := 0
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

	digitFound := func(numLine int, line string, start int, end int) int {
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
				return num
			}
		}
		return 0
	}
	for i, line := range lines {
		digitStart := -1
		for j, r := range line {
			if unicode.IsDigit(r) {
				if digitStart == -1 {
					digitStart = j
				}
			} else if digitStart != -1 {
				sum += digitFound(i, line, digitStart, j)
				digitStart = -1
			}
		}
		if digitStart != -1 {
			sum += digitFound(i, line, digitStart, len(line))
		}
	}
	fmt.Println(sum)
}
