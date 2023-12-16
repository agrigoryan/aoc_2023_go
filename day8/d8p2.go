package day8

import (
	"fmt"
	"strings"
)

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func d8p2(input string) int {
	lines := strings.Split(input, "\n")

	directions := []rune(lines[0])

	nodesByName := map[string]*node{}

	startNodes := []*node{}

	for _, line := range lines[2:] {
		name := strings.Split(line, " ")[0]
		left := line[strings.Index(line, "(")+1 : strings.Index(line, ",")]
		right := line[strings.Index(line, ", ")+2 : strings.Index(line, ")")]
		nodesByName[name] = &node{name, left, right, name[len(name)-1] == 'Z'}
		if name[len(name)-1] == 'A' {
			startNodes = append(startNodes, nodesByName[name])
		}
	}

	cycleLength := func(n *node) int {
		steps := 0
		for {
			direction := directions[steps%len(directions)]
			if n.endingNode {
				break
			}
			if direction == 'R' {
				n = nodesByName[n.right]
			} else {
				n = nodesByName[n.left]
			}
			steps += 1
		}
		return steps
	}

	cycleLengths := make([]int, len(startNodes))

	for i, n := range startNodes {
		cycleLengths[i] = cycleLength(n)
	}

	res := lcm(cycleLengths[0], cycleLengths[1], cycleLengths[2:]...)
	fmt.Println(res)
	return res
}
