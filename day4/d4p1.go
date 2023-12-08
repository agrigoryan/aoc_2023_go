package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func Filter[T any](arr []T, filter func(T) bool) []T {
	filtered := []T{}
	for _, item := range arr {
		if filter(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func Map[T any, R any](arr []T, transform func(T) R) []R {
	transformed := []R{}
	for _, item := range arr {
		transformed = append(transformed, transform(item))
	}
	return transformed
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

func main() {
	content, err := os.ReadFile("d4p1.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)
	sum := 0

	trimSpaces := Mapper(func(str string) string { return strings.TrimSpace(str) })
	filterEmpty := Filterer(func(str string) bool { return len(str) > 0 })

	for _, line := range strings.Split(input, "\n") {
		ticket := strings.Split(line, ":")[1]
		parts := strings.Split(ticket, "|")
		winningNumbers := filterEmpty(trimSpaces(strings.Split(parts[0], " ")))
		numbers := filterEmpty(trimSpaces(strings.Split(parts[1], " ")))
		matches := 0
		for _, num := range numbers {
			if slices.Contains(winningNumbers, num) {
				matches += 1
			}
		}
		if matches > 0 {
			worth := 1 << (matches - 1)
			sum += worth
		}
	}
	fmt.Println(sum)
}
