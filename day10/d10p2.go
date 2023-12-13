package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var pipeDirections = map[rune]struct {
	u, r, d, l bool
}{
	'-': {
		u: false,
		r: true,
		d: false,
		l: true,
	},
	'|': {
		u: true,
		r: false,
		d: true,
		l: false,
	},
	'7': {
		u: false,
		r: false,
		d: true,
		l: true,
	},
	'L': {
		u: true,
		r: true,
		d: false,
		l: false,
	},
	'J': {
		u: true,
		r: false,
		d: false,
		l: true,
	},
	'F': {
		u: false,
		r: true,
		d: true,
		l: false,
	},
	'S': {
		u: true,
		r: true,
		d: true,
		l: true,
	},
}

func main() {
	content, err := os.ReadFile("d10p2.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)
	lines := strings.Split(input, "\n")

	type pos struct {
		i, j int
	}
	type cell struct {
		pos
		d int
		r rune
	}

	runes := make([][]rune, len(lines))

	var startCell cell
	for i, line := range lines {
		runes[i] = []rune(line)
		for j, r := range runes[i] {
			if r == 'S' {
				startCell = cell{pos{i, j}, 0, r}
			}
		}
	}

	visited := map[pos]*cell{}
	toVisit := []cell{startCell}

	addToVisit := func(n pos, d int) {
		if visited[n] != nil {
			visited[n].d = min(visited[n].d, d)
		} else {
			toVisit = append(toVisit, cell{
				n, d, runes[n.i][n.j],
			})
		}
	}

	for len(toVisit) > 0 {
		c := toVisit[0]
		toVisit = toVisit[1:]
		visited[c.pos] = &c
		n := pos{c.i - 1, c.j}
		var nu, nd, nl, nr bool
		if n.i >= 0 && pipeDirections[c.r].u && pipeDirections[runes[n.i][n.j]].d {
			nu = true
			addToVisit(n, c.d+1)
		}
		n = pos{c.i, c.j - 1}
		if n.j >= 0 && pipeDirections[c.r].l && pipeDirections[runes[n.i][n.j]].r {
			nl = true
			addToVisit(n, c.d+1)
		}
		n = pos{c.i, c.j + 1}
		if n.i < len(runes) && pipeDirections[c.r].r && pipeDirections[runes[n.i][n.j]].l {
			nr = true
			addToVisit(n, c.d+1)
		}
		n = pos{c.i + 1, c.j}
		if n.j < len(runes[0]) && pipeDirections[c.r].d && pipeDirections[runes[n.i][n.j]].u {
			nd = true
			addToVisit(n, c.d+1)
		}

		// replace S with actual pipe type
		if c == startCell {
			var actual rune
			if nd && nl {
				actual = '7'
			} else if nr && nl {
				actual = '-'
			} else if nd && nu {
				actual = '|'
			} else if nl && nu {
				actual = 'J'
			} else if nr && nd {
				actual = 'F'
			} else if nr && nu {
				actual = 'L'
			}
			fmt.Println(actual)
			runes[startCell.i][startCell.j] = actual
		}
	}

	numCellsInside := 0

	for i, row := range runes {
		inside := false
		marker := 0
		for j, r := range row {
			if (visited[pos{i, j}] != nil) {
				if pipeDirections[r].u {
					marker -= 1
				}
				if pipeDirections[r].d {
					marker += 1
				}
				if marker == 0 {
					inside = !inside
				} else if marker < -1 || marker > 1 {
					marker = 0
				}
			} else if inside {
				fmt.Println(i, j)
				numCellsInside += 1
			}
		}
	}

	fmt.Println(numCellsInside)
}
