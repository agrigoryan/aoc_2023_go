package day6

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func d6p1(input string) int {
	lines := strings.Split(input, "\n")

	filterEmpty := aocutils.Filterer(func(str string) bool { return len(str) > 0 })
	toNumbersMapper := aocutils.Mapper(func(str string) int {
		res, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal("failed to map to numbers")
		}
		return res
	})

	times := toNumbersMapper(filterEmpty(strings.Split(strings.Split(lines[0], ":")[1], " ")))
	distances := toNumbersMapper(filterEmpty(strings.Split(strings.Split(lines[1], ":")[1], " ")))

	result := 1
	for i := 0; i < len(times); i++ {
		t := times[i]
		d := distances[i]
		x := t - 1 - 2*int((float64(t)-math.Sqrt(float64(t*t-4*d)))/2)
		result *= x
	}
	fmt.Println(result)
	return result
}
