package day1

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func d1p1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		var d1, d2 rune
		for i := 0; i < len(line); i++ {
			if d1 == 0 && unicode.IsDigit(rune(line[i])) {
				d1 = rune(line[i])
			}
			if d2 == 0 && unicode.IsDigit(rune(line[len(line)-1-i])) {
				d2 = rune(line[len(line)-1-i])
			}
			if d1 != 0 && d2 != 0 {
				break
			}
		}
		num, err := strconv.Atoi(string([]rune{d1, d2}))
		if err != nil {
			log.Fatalf("error parsing line %v", line)
		}
		sum += num
	}
	return sum
}
