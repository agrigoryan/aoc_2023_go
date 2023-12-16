package day10

import (
	"fmt"
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

func d10p1(input string) int {
	lines := strings.Split(input, "\n")
	fmt.Println(lines)

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
		if n.i >= 0 && pipeDirections[c.r].u && pipeDirections[runes[n.i][n.j]].d {
			addToVisit(n, c.d+1)
		}
		n = pos{c.i, c.j - 1}
		if n.j >= 0 && pipeDirections[c.r].l && pipeDirections[runes[n.i][n.j]].r {
			addToVisit(n, c.d+1)
		}
		n = pos{c.i, c.j + 1}
		if n.i < len(runes) && pipeDirections[c.r].r && pipeDirections[runes[n.i][n.j]].l {
			addToVisit(n, c.d+1)
		}
		n = pos{c.i + 1, c.j}
		if n.j < len(runes[0]) && pipeDirections[c.r].d && pipeDirections[runes[n.i][n.j]].u {
			addToVisit(n, c.d+1)
		}
	}

	m := startCell.d

	for _, v := range visited {
		if v.d > m {
			m = v.d
		}
	}

	fmt.Println(m)
	return m
}
