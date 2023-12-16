package day1

import (
	"fmt"
	"regexp"
	"strings"
)

var digits = map[string]int{
	"one":   1,
	"1":     1,
	"two":   2,
	"2":     2,
	"three": 3,
	"3":     3,
	"four":  4,
	"4":     4,
	"five":  5,
	"5":     5,
	"six":   6,
	"6":     6,
	"seven": 7,
	"7":     7,
	"eight": 8,
	"8":     8,
	"nine":  9,
	"9":     9,
}

func findSumWithRegexp(lines []string) int {
	sum := 0
	re := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	for _, line := range lines {
		matches := re.FindAllString(line, len(line))
		fmt.Println(line, matches)
		d1, d2 := matches[0], matches[len(matches)-1]
		fmt.Println(d1, d2)
		sum += digits[d1]*10 + digits[d2]
	}
	return sum
}

func findSumBrutally(lines []string) int {
	sum := 0
	for _, line := range lines {
		first_digit, last_digit := 0, 0
		first_idx, last_idx := -1, -1
		for k, v := range digits {
			if i := strings.Index(line, k); i != -1 && (i < first_idx || first_idx == -1) {
				first_idx = i
				first_digit = v
			}
			if i := strings.LastIndex(line, k); i > last_idx || last_idx == -1 {
				last_idx = i
				last_digit = v
			}
		}
		fmt.Println(line, first_digit, last_digit)
		sum += first_digit*10 + last_digit
	}
	return sum
}

func d1p2(input string) int {
	sum := findSumBrutally(strings.Split(input, "\n"))
	fmt.Println(sum)
	return sum
}
