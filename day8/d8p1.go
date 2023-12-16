package day8

import (
	"fmt"
	"strings"
)

type node struct {
	name       string
	left       string
	right      string
	endingNode bool
}

func d8p1(input string) int {
	lines := strings.Split(input, "\n")

	directions := []rune(lines[0])

	nodesByName := map[string]*node{}

	for _, line := range lines[2:] {
		name := strings.Split(line, " ")[0]
		left := line[strings.Index(line, "(")+1 : strings.Index(line, ",")]
		right := line[strings.Index(line, ", ")+2 : strings.Index(line, ")")]
		nodesByName[name] = &node{name, left, right, false}
	}

	steps := 0
	destination := "ZZZ"
	currNode := nodesByName["AAA"]
	for {
		if currNode.name == destination {
			break
		}
		direction := directions[steps%len(directions)]
		if direction == 'R' {
			currNode = nodesByName[currNode.right]
		} else {
			currNode = nodesByName[currNode.left]
		}
		steps += 1
	}

	fmt.Println(steps)
	return steps
}
