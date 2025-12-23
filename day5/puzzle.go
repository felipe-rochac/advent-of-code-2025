package day5

import (
	"slices"
	"sort"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/felipe-rochac/advent-of-code-2025/helpers"
)

type Range struct {
	From, To    int64
	Ingredients []int64
}

type List struct {
	Ranges      []Range
	Ingredients []int64
}

type Pair struct {
	X, Y int64
}

func parse(filename string) *List {
	lines, err := helpers.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	list := &List{
		Ranges:      make([]Range, 0),
		Ingredients: make([]int64, 0),
	}
	emptyLine := 0
	for i := range lines {
		if lines[i] == "" {
			emptyLine = i
			break
		}
	}
	var wg sync.WaitGroup
	var mu sync.RWMutex
	// force all routines to start together
	// must be buffered as unbuffered can only sync 1 reader and receiver
	startCh := make(chan struct{}, len(lines))

	// parse ranges
	for _, l := range lines[:emptyLine] {
		wg.Go(func() {
			<-startCh
			parts := strings.Split(l, "-")
			mu.Lock()
			list.Ranges = append(list.Ranges, Range{
				From:        helpers.MustParseInt64(parts[0]),
				To:          helpers.MustParseInt64(parts[1]),
				Ingredients: make([]int64, 0),
			})
			mu.Unlock()
		})
	}

	// parse ingredients
	for _, l := range lines[emptyLine+1:] {
		wg.Go(func() {
			<-startCh
			mu.Lock()
			list.Ingredients = append(list.Ingredients, helpers.MustParseInt64(l))
			mu.Unlock()
		})
	}

	// waking up workers
	for range lines {
		startCh <- struct{}{}
	}
	wg.Wait()
	close(startCh)

	return list
}

func findIngredients(list *List) (fresh, spoiled int64) {
	var wg sync.WaitGroup
	ings := len(list.Ingredients)
	var mu sync.Mutex
	sem := make(chan struct{}, ings)

	for _, ing := range list.Ingredients {
		wg.Go(func() {
			<-sem
			for i, r := range list.Ranges {
				if ing >= r.From && ing <= r.To {
					mu.Lock()
					list.Ranges[i].Ingredients = append(list.Ranges[i].Ingredients, ing)
					atomic.AddInt64(&fresh, 1)
					mu.Unlock()
					return
				}
			}

			atomic.AddInt64(&spoiled, 1)
		})
	}

	for range ings {
		sem <- struct{}{}
	}

	wg.Wait()
	close(sem)

	return
}

func MergeRanges(r1, r2 *Range) (bool, *Range) {
	r := r1
	merged := false
	if helpers.Between(r1.From, r2.From, r2.To) ||
		helpers.Between(r1.To, r1.From, r1.To) {
		r.From = min(r1.From, r2.From)
		r.To = max(r1.To, r2.To)
		merged = true
	}

	return merged, r
}

func findIngredients2(list *List) int64 {
	var wg sync.WaitGroup
	var mu sync.Mutex
	events := make([]Pair, 0)

	for i := 0; i < len(list.Ranges); i++ {
		wg.Go(func() {
			r := list.Ranges[i]
			mu.Lock()
			events = append(events, Pair{
				X: r.From,
				Y: +1,
			})
			events = append(events, Pair{
				X: r.To + 1,
				Y: -1,
			})
			mu.Unlock()
		})
	}

	wg.Wait()

	sort.Slice(events, func(i, j int) bool {
		return events[i].X < events[j].X
	})

	ranges := make([]int64, 0)
	active := int32(0)
	i := 0
	// TODO if i < 1 should fail
	for {
		from := events[i]
		to := events[i+1]
		atomic.AddInt32(&active, int32(from.Y))
		print("(", from.X, ", ", to.X, ") active = ", active, "\n")
		for j := from.X; j < to.X; j++ {
			mu.Lock()
			if !slices.Contains(ranges, j) {
				ranges = append(ranges, j)
			}
			mu.Unlock()
		}
		atomic.AddInt32(&active, int32(to.Y))
	}

	// wg.Wait()

	return int64(len(ranges))
}

func puzzle1(filename string) int64 {
	list := parse(filename)

	fresh, _ := findIngredients(list)
	return fresh
}

func puzzle2(filename string) int64 {
	list := parse(filename)

	total := findIngredients2(list)
	return total
}
