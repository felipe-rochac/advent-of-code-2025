package day10

import (
	"container/list"
	"math"
	"strings"
	"sync"

	"github.com/felipe-rochac/advent-of-code-2025/helpers"
)

type machine struct {
	indicators   string
	lightDiagram [][]int
	joltage      []int
}

func parse(filename string) []machine {
	lines, err := helpers.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	machines := make([]machine, 0)

	parseIndicators := func(line string) string {
		start, end := strings.Index(line, "["), strings.Index(line, "]")
		indicators := ""

		for i := start + 1; i < end; i++ {
			r := line[i]
			indicators += string(r)
		}

		return indicators
	}

	parseLightDiagram := func(line string) [][]int {
		lights := line[strings.Index(line, "("):strings.LastIndex(line, ")")]
		lights = strings.ReplaceAll(lights, "(", "")
		lights = strings.ReplaceAll(lights, ")", "")
		diagrams := strings.Split(lights, " ")
		diagram := make([][]int, 0)

		for _, d := range diagrams {
			dg := make([]int, 0)
			for _, p := range strings.Split(d, ",") {
				dg = append(dg, helpers.MustParseInt(p))
			}
			diagram = append(diagram, dg)
		}

		return diagram

	}

	parseJoltage := func(line string) []int {
		start, end := strings.Index(line, "{"), strings.Index(line, "}")
		joltage := make([]int, 0)

		s := line[start+1 : end]
		for _, p := range strings.Split(s, ",") {
			joltage = append(joltage, helpers.MustParseInt(p))
		}

		return joltage
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, l := range lines {
		wg.Go(func() {
			joltage := parseJoltage(l)
			diagram := parseLightDiagram(l)
			indicators := parseIndicators(l)

			mu.Lock()
			machines = append(machines, machine{
				indicators:   indicators,
				lightDiagram: diagram,
				joltage:      joltage,
			})
			mu.Unlock()
		})
	}

	wg.Wait()

	return machines
}

type action struct {
	result string
	count  int
}

func applyCombinations(answer string, comb [][][]int) []action {
	var wg sync.WaitGroup
	var mu sync.Mutex
	result := make([]action, 0)

	pressButtons := func(ind string, buttons []int) string {
		for _, b := range buttons {
			input := "."
			s := ind[b]
			if s == '.' {
				input = "#"
			}
			ind = ind[:b] + input + ind[b+1:]
		}

		return ind
	}

	for _, c := range comb {
		wg.Go(func() {

			ind := strings.Repeat(".", len(answer))
			count := 0
			for _, b := range c {
				ind = pressButtons(ind, b)
				count++
			}

			// discard wrong combinations
			if ind != answer {
				return
			}

			mu.Lock()
			result = append(result, action{
				result: ind,
				count:  count,
			})
			mu.Unlock()
		})
	}

	wg.Wait()

	return result
}

func findLowestComb(actions []action) action {
	min := math.MaxInt
	action := action{}

	for _, a := range actions {
		if a.count < min {
			action = a
		}
	}

	return action
}

func minButtonPresses(indicators string, buttons [][]int) int {
	// n := len(indicators)
	target := 0
	for i, c := range indicators {
		if c == '#' {
			target |= 1 << i
		}
	}

	// Precompute button masks
	buttonMasks := make([]int, len(buttons))
	for i, btn := range buttons {
		mask := 0
		for _, pos := range btn {
			mask |= 1 << pos
		}
		buttonMasks[i] = mask
	}

	visited := make(map[int]bool)
	queue := list.New()
	queue.PushBack([2]int{0, 0}) // (state, presses)
	visited[0] = true

	for queue.Len() > 0 {
		elem := queue.Remove(queue.Front()).([2]int)
		state, presses := elem[0], elem[1]
		if state == target {
			return presses
		}
		for _, mask := range buttonMasks {
			next := state ^ mask
			if !visited[next] {
				visited[next] = true
				queue.PushBack([2]int{next, presses + 1})
			}
		}
	}
	return -1 // unreachable
}

func puzzle1(filename string) int64 {
	machines := parse(filename)

	sum := 0
	for _, m := range machines {
		sum += minButtonPresses(m.indicators, m.lightDiagram)
	}

	return int64(sum)
}

func puzzle2(filename string) int64 {
	return 0
}
