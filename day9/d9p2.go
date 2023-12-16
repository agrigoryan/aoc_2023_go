package day9

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func findPrevElement(arr []int) int {
	allZeroes := true
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != 0 {
			allZeroes = false
		}
		if i < len(arr)-1 {
			arr[i+1] = arr[i+1] - arr[i]
		}
	}
	if allZeroes {
		return 0
	}
	return arr[0] - findPrevElement(arr[1:])
}

func d9p2(input string) int {
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

		sum += findPrevElement(arr)
	}

	fmt.Println(sum)
	return sum
}
