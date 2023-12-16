package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func d2p2(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, ":")
		plays := strings.Split(parts[1], ";")
		requiredCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, play := range plays {
			balls := strings.Split(play, ",")
			for _, ball := range balls {
				ballParts := strings.Split(strings.TrimSpace(ball), " ")
				numBalls, err := strconv.Atoi(ballParts[0])
				if err != nil {
					log.Fatalf("error parsing ball count for %v", ball)
				}
				color := ballParts[1]
				if requiredCubes[color] < numBalls {
					requiredCubes[color] = numBalls
				}
			}
		}
		sum += requiredCubes["red"] * requiredCubes["blue"] * requiredCubes["green"]
	}
	fmt.Println(sum)
	return sum
}
