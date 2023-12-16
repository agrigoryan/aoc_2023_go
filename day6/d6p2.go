package day6

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func d6p2(input string) int {
	lines := strings.Split(input, "\n")

	filterEmpty := aocutils.Filterer(func(str string) bool { return len(str) > 0 })

	t, err := strconv.Atoi(strings.Join(filterEmpty(strings.Split(strings.Split(lines[0], ":")[1], " ")), ""))
	if err != nil {
		log.Fatal(err)
	}
	d, err := strconv.Atoi(strings.Join(filterEmpty(strings.Split(strings.Split(lines[1], ":")[1], " ")), ""))
	if err != nil {
		log.Fatal(err)
	}

	x := t - 1 - 2*int((float64(t)-math.Sqrt(float64(t*t-4*d)))/2)
	fmt.Println(x)

	return x
}
