package day8

import (
	"slices"
	"strings"
	"sync"

	"github.com/felipe-rochac/advent-of-code-2025/helpers"
)

// Union-Find (Disjoint Set Union) structure
type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return
	}
	// Union by size
	if uf.size[rootX] < uf.size[rootY] {
		uf.parent[rootX] = rootY
		uf.size[rootY] += uf.size[rootX]
	} else {
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
	}
}

type junctionBox struct {
	X, Y, Z  int
	circuits []*circuit
}

func (j *junctionBox) Equal(b junctionBox) bool {
	return j.X == b.X && j.Y == b.Y && j.Z == b.Z
}

type circuit struct {
	From, To *junctionBox
	Distance int
}

func (c *circuit) Find(j junctionBox) {

}

func parse(filename string) []junctionBox {
	lines, err := helpers.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	boxCh := make(chan junctionBox, len(lines))

	for _, line := range lines {
		wg.Go(func() {
			parts := strings.Split(line, ",")

			boxCh <- junctionBox{
				X:        helpers.MustParseInt(parts[0]),
				Y:        helpers.MustParseInt(parts[1]),
				Z:        helpers.MustParseInt(parts[2]),
				circuits: make([]*circuit, 0),
			}
		})
	}

	wg.Wait()
	close(boxCh)

	boxes := make([]junctionBox, 0)
	for box := range boxCh {
		boxes = append(boxes, box)
	}

	return boxes
}

func buildCircuit(jboxes []junctionBox) []*circuit {
	var wg sync.WaitGroup
	circuitCh := make(chan []*circuit, len(jboxes))

	createCircuit := func(a, b *junctionBox) *circuit {
		distance := helpers.EuclidianDistance(a.X, a.Y, a.Z, b.X, b.Y, b.Z)
		c := &circuit{
			From:     a,
			To:       b,
			Distance: distance,
		}
		a.circuits = append(a.circuits, c)
		b.circuits = append(b.circuits, c)
		return c
	}

	for _, box := range jboxes {
		wg.Go(func() {
			circuits := make([]*circuit, 0)
			for _, b := range jboxes {
				if box.Equal(b) {
					continue
				}

				circuit := createCircuit(&box, &b)
				circuits = append(circuits, circuit)
			}

			circuitCh <- circuits
		})
	}

	wg.Wait()
	close(circuitCh)
	circuits := make([]*circuit, 0)

	for _, c := range <-circuitCh {
		circuits = append(circuits, c)
	}

	return circuits
}

func puzzle1(filename string) int64 {
	jbox := parse(filename)

	circuits := buildCircuit(jbox)

	// Sort by distance (ascending)
	slices.SortFunc(circuits, func(a, b *circuit) int {
		if a.Distance < b.Distance {
			return -1
		} else if a.Distance == b.Distance {
			return 0
		}
		return 1
	})

	// Set up union-find for all boxes
	uf := NewUnionFind(len(jbox))

	// Map each box to its index for union-find
	boxIndex := make(map[*junctionBox]int)
	for i, box := range jbox {
		boxIndex[&box] = i
	}

	// Example: connect the closest pairs (you can limit the number of connections as needed)
	for _, c := range circuits {
		aIdx := boxIndex[c.From]
		bIdx := boxIndex[c.To]
		if uf.Find(aIdx) != uf.Find(bIdx) {
			uf.Union(aIdx, bIdx)
			// You can count connections here, or stop after a certain number
		}
	}

	// To find the size of each circuit (connected component):
	componentSizes := make(map[int]int)
	for i := range jbox {
		root := uf.Find(i)
		componentSizes[root]++
	}

	// Now you can get the sizes of all circuits
	sizes := make([]int, 0, len(componentSizes))
	for _, sz := range componentSizes {
		sizes = append(sizes, sz)
	}
	slices.Sort(sizes)
	// For example, multiply the sizes of the three largest circuits
	n := len(sizes)
	if n >= 3 {
		return int64(sizes[n-1] * sizes[n-2] * sizes[n-3])
	}
	return 0
}

func puzzle2(filename string) int64 {
	return 0
}
