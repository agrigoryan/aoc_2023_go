package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Filterer[T any](filter func(T) bool) func([]T) []T {
	return func(arr []T) []T {
		filtered := []T{}
		for _, item := range arr {
			if filter(item) {
				filtered = append(filtered, item)
			}
		}
		return filtered
	}
}

func Mapper[T any, R any](transform func(T) R) func([]T) []R {
	return func(arr []T) []R {
		transformed := []R{}
		for _, item := range arr {
			transformed = append(transformed, transform(item))
		}
		return transformed
	}
}

func main() {
	content, err := os.ReadFile("d6p1.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)
	lines := strings.Split(input, "\n")

	filterEmpty := Filterer(func(str string) bool { return len(str) > 0 })
	toNumbersMapper := Mapper(func(str string) int {
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
}
