package day9

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func findNextElement(arr []int) int {
	allZeroes := true
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			allZeroes = false
		}
		if i < len(arr)-1 {
			arr[i] = arr[i+1] - arr[i]
		}
	}
	if allZeroes {
		return 0
	}
	return findNextElement(arr[:len(arr)-1]) + arr[len(arr)-1]
}

func d9p1(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		numStrings := strings.Split(line, " ")
		arr := []int{}
		for _, s := range numStrings {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			arr = append(arr, n)
		}

		sum += findNextElement(arr)
	}

	fmt.Println(sum)
	return sum
}
