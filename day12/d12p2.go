package day12

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func d12p2(input string) int {
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
		newSeries := []int{}
		newSprigns := springs
		for i := 0; i < 5; i++ {
			newSeries = append(newSeries, series...)
			if i < 4 {
				newSprigns += "?" + springs
			}
		}

		sum += solveLine(newSprigns, newSeries)
		fmt.Println(sum)
	}

	return sum
}
