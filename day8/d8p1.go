package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type node struct {
	name  string
	left  string
	right string
}

func main() {
	content, err := os.ReadFile("d8p1.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)
	lines := strings.Split(input, "\n")

	directions := []rune(lines[0])

	nodesByName := map[string]*node{}

	for _, line := range lines[2:] {
		name := strings.Split(line, " ")[0]
		left := line[strings.Index(line, "(")+1 : strings.Index(line, ",")]
		right := line[strings.Index(line, ", ")+2 : strings.Index(line, ")")]
		nodesByName[name] = &node{name, left, right}
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
}
