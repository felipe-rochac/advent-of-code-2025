package template

import (
	"fmt"

	"github.com/felipe-rochac/advent-of-code-2025/helpers"
)

type manifold struct {
	diagram   [][]rune
	visited   [][]bool
	splitters []splitter
	start     *position
}

type position struct {
	X, Y int
}

type splitter struct {
	X, Y    int
	Visited bool
}

func parse(filename string) *manifold {
	lines, err := helpers.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	manifold := &manifold{
		diagram:   make([][]rune, 0),
		visited:   make([][]bool, 0),
		splitters: make([]splitter, 0),
	}
	for i, line := range lines {
		row := make([]rune, len(line))
		visited := make([]bool, len(line))
		for j, r := range line {
			switch r {
			case 'S':
				manifold.start = &position{
					X: i,
					Y: j,
				}
			case '^':
				splitter := splitter{
					X:       i,
					Y:       j,
					Visited: false,
				}
				manifold.splitters = append(manifold.splitters, splitter)
			}
			row[j] = r
			visited[j] = false
		}
		manifold.diagram = append(manifold.diagram, row)
		manifold.visited = append(manifold.visited, visited)
	}

	return manifold
}

func isSplitter(splitters []splitter, pos position) (found, visited bool) {
	for _, s := range splitters {
		if s.X == pos.X && s.Y == pos.Y {
			found = true
			visited = s.Visited
			return
		}
	}
	found = false
	visited = false
	return
}

func puzzle1(filename string) int64 {
	manifold := parse(filename)

	pos := &position{
		X: manifold.start.X,
		Y: manifold.start.Y,
	}

	total := 0
	memo := make(map[string]int)
	var propagate func(pos *position, depth int) int
	propagate = func(pos *position, depth int) int {
		// Create a unique key for the current state (position and depth)
		key := fmt.Sprintf("%d,%d,%d", pos.X, pos.Y, depth)
		if val, ok := memo[key]; ok {
			return val
		}

		p := &position{X: pos.X, Y: pos.Y}
		p.X++
		// size overflow
		if p.X > len(manifold.diagram)-1 {
			// print("beam reach to the end at position (", p.X, ",", p.Y, ") depth ", depth, " \n")
			memo[key] = depth
			return depth
		}
		// side overflow
		if p.Y < 0 || p.Y > len(manifold.diagram[p.X])-1 {
			memo[key] = 0
			return 0
		}

		// if visited position, then skips
		if manifold.visited[p.X][p.Y] {
			memo[key] = 0
			return 0
		}

		manifold.visited[p.X-1][p.Y] = true

		found, visited := isSplitter(manifold.splitters, *p)

		// if splitter visited, stop propagation
		if visited {
			memo[key] = 0
			return 0
		}

		res := 0
		// if splitter propagate next row left and right, otherwise just next row
		if found {
			res += propagate(&position{
				X: p.X,
				Y: p.Y - 1,
			}, depth+1)
			res += propagate(&position{
				X: p.X,
				Y: p.Y + 1,
			}, depth+1)
		} else {
			res += propagate(p, depth+1)
		}
		memo[key] = res
		return res
	}

	total = propagate(pos, 1)
	return int64(total)
}

func puzzle2(filename string) int64 {
	manifold := parse(filename)

	var dfs func(x, y int) int64
	dfs = func(x, y int) int64 {
		// Out of bounds
		if x < 0 || x >= len(manifold.diagram) || y < 0 || y >= len(manifold.diagram[x]) {
			return 0
		}
		// Already visited in this timeline
		if manifold.visited[x][y] {
			return 0
		}

		// Mark as visited for this timeline
		manifold.visited[x][y] = true

		// If at the bottom row, this is a valid timeline
		if x == len(manifold.diagram)-1 {
			manifold.visited[x][y] = false
			return 1
		}

		found, splitterVisited := isSplitter(manifold.splitters, position{X: x, Y: y})
		if splitterVisited {
			manifold.visited[x][y] = false
			return 0
		}

		var timelines int64 = 0
		if found {
			// Splitter: branch left and right
			timelines += dfs(x+1, y-1)
			timelines += dfs(x+1, y+1)
		} else {
			// Normal: go straight down
			timelines += dfs(x+1, y)
		}

		// Backtrack
		manifold.visited[x][y] = false
		return timelines
	}

	return dfs(manifold.start.X, manifold.start.Y)
}
