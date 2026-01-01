package day9

import (
	"strings"
	"sync"

	"github.com/felipe-rochac/advent-of-code-2025/helpers"
)

type tile struct {
	X, Y int
}

func (t *tile) Equal(t2 tile) bool {
	return t.X == t2.X && t.Y == t2.Y
}

func parse(filename string) []*tile {
	lines, err := helpers.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	tiles := make([]*tile, 0)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		tile := &tile{
			X: helpers.MustParseInt(parts[0]),
			Y: helpers.MustParseInt(parts[1]),
		}
		tiles = append(tiles, tile)
	}

	return tiles
}

func puzzle1(filename string) int64 {
	tiles := parse(filename)

	max := 0
	maxArea := make([]*tile, 2)
	for i, t := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			t2 := tiles[j]

			area := helpers.Area(t.X, t.Y, 0, t2.X, t2.Y, 0)

			if area > max {
				max = area
				maxArea[0] = t
				maxArea[1] = t2
			}
		}
	}

	return int64(max)
}

func puzzle2(filename string) int64 {
	tiles := parse(filename)

	// should find 3 coordinates to form a rectangle
	stack := helpers.Stack[[]*tile]{}

	var wg sync.WaitGroup
	var mu sync.Mutex

	hasEdge := func(x, y int) bool {
		for _, t := range tiles {
			if t.X == x && t.Y == y {
				return true
			}
		}

		return false
	}

	maxDistance := 0
	for i, t := range tiles {
		wg.Go(func() {
			for j := i + 1; j < len(tiles); j++ {
				t2 := tiles[j]

				distance := helpers.EuclidianDistance(t.X, t.Y, 0, t2.X, t2.Y, 0)

				if distance < maxDistance {
					continue
				}

				// only add to stack that have an 3d edge
				if hasEdge(t.X, t2.Y) || hasEdge(t2.X, t.Y) {
					mu.Lock()
					stack.Push([]*tile{t, t2})
					maxDistance = distance
					mu.Unlock()
				}
			}
		})
	}

	wg.Wait()

	maxArea := 0
	var maxRange []*tile
	for !stack.IsEmpty() {
		t := stack.Pop()
		area := helpers.Area(t[0].X, t[0].Y, 0, t[1].X, t[1].Y, 0)

		if area > maxArea {
			maxArea = area
			maxRange = []*tile{t[0], t[1]}
		}
	}

	print("Biggest rectagle found at (", maxRange[0].X, ",", maxRange[0].Y, ") - (", maxRange[1].X, ",", maxRange[1].Y, ")")

	return int64(maxArea)
}
