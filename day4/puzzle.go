package day4

import (
	"sync"
	"sync/atomic"

	"github.com/felipe-rochac/advent-of-code-2025/helpers"
)

type pos struct {
	PaperRoll  bool
	Accessible bool
}

type Pair struct {
	X, Y int
}

func parse(filename string) [][]pos {
	lines, err := helpers.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	diagram := make([][]pos, 0)
	for _, l := range lines {
		p := make([]pos, 0)
		for j := range len(l) {
			p = append(p, pos{
				PaperRoll:  string(l[j]) == "@",
				Accessible: false,
			})
		}
		diagram = append(diagram, p)
	}

	return diagram
}

func checkDiagram(diagram [][]pos, remove bool) int {

	safeAdjacentPaperRoll := func(x, y int) bool {
		if x < 0 || x >= len(diagram) {
			return false
		}
		if y < 0 || y >= len(diagram[x]) {
			return false
		}

		return diagram[x][y].PaperRoll
	}

	accessiblePaperRoll := func(x, y int) bool {
		if !diagram[x][y].PaperRoll {
			return false
		}

		pairs := []Pair{
			{X: x - 1, Y: y},
			{X: x - 1, Y: y - 1},
			{X: x - 1, Y: y + 1},
			{X: x + 1, Y: y},
			{X: x + 1, Y: y + 1},
			{X: x + 1, Y: y - 1},
			{X: x, Y: y - 1},
			{X: x, Y: y + 1},
		}

		checkCh := make(chan bool, len(pairs))
		var wg sync.WaitGroup

		for _, p := range pairs {
			wg.Go(func() {
				checkCh <- safeAdjacentPaperRoll(p.X, p.Y)
			})
		}

		wg.Wait()
		close(checkCh)

		paperRollCount := 0
		for found := range checkCh {
			if found {
				paperRollCount++
			}
			if paperRollCount >= 4 {
				return false
			}
		}

		return true
	}

	var wg sync.WaitGroup
	var counter int32 = 0
	for i := range diagram {
		for j := range diagram[i] {
			wg.Go(func() {
				if accessiblePaperRoll(i, j) {
					diagram[i][j].Accessible = true
					if remove {
						diagram[i][j].PaperRoll = false
					}
					atomic.AddInt32(&counter, 1)
				}
			})
		}
	}

	wg.Wait()

	return int(counter)
}

func puzzle1(filename string) int64 {
	diagram := parse(filename)
	counter := checkDiagram(diagram, false)
	return int64(counter)
}

func puzzle2(filename string) int64 {
	diagram := parse(filename)
	total := 0
	for {
		counter := checkDiagram(diagram, true)
		total += counter
		if counter == 0 {
			break
		}
	}
	return int64(total)
}
