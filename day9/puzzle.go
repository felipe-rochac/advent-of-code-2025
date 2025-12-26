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

func greenTiles(redTiles []tile) []tile {
	var wg sync.WaitGroup
	syncCh := make(chan []tile, len(redTiles))

	for i, t := range redTiles {
		wg.Go(func() {
			greenTiles := make([]tile, 0)
			for j := i + 1; j < len(redTiles); j++ {

			}
		})
	}
}

func puzzle2(filename string) int64 {
	tiles := parse(filename)

	// should frm a rectagle from objects that shares same coordinate
	stack := helpers.Stack[[]*tile]{}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i, t := range tiles {
		wg.Go(func() {
			group := []*tile{t}
			for j := i + 1; j < len(tiles); j++ {
				if t.X == tiles[j].X || t.Y == tiles[j].Y {
					group = append(group, tiles[j])
				}

				if len(group) == 2 {
					mu.Lock()
					stack.Push(group)
					group = []*tile{t}
					mu.Unlock()

				}
			}

		})
	}

	wg.Wait()

	maxArea := 0
	for !stack.IsEmpty() {
		t := stack.Pop()
		area := helpers.Area(t[0].X, t[0].Y, 0, t[1].X, t[1].Y, 0)

		if area > maxArea {
			maxArea = area
		}
	}

	return 0
}
