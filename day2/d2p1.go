package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func d2p1(input string) int {
	maxBalls := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, ":")
		id, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
		if err != nil {
			log.Fatalf("failed to parse id for line %v", line)
		}
		plays := strings.Split(parts[1], ";")
		gamePossbile := true
		for _, play := range plays {
			balls := strings.Split(play, ",")
			for _, ball := range balls {
				ballParts := strings.Split(strings.TrimSpace(ball), " ")
				numBalls, err := strconv.Atoi(ballParts[0])
				if err != nil {
					log.Fatalf("error parsing ball count for %v", ball)
				}
				color := ballParts[1]
				if numBalls > maxBalls[color] {
					gamePossbile = false
					break
				}
			}
			if !gamePossbile {
				break
			}
		}
		if gamePossbile {
			sum += id
		}
	}
	fmt.Println(sum)
	return sum
}
