package day5

import (
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

func findIngredients2(list *List) map[int64]int {
	var wg sync.WaitGroup
	ings := len(list.Ingredients)
	var mu sync.Mutex
	sem := make(chan struct{}, ings)
	fresh := make(map[int64]int)

	for _, r := range list.Ranges {
		wg.Go(func() {
			<-sem

			for j := r.From; j <= r.To; j++ {
				mu.Lock()
				fresh[j]++
				mu.Unlock()
			}
		})
	}

	for range ings {
		sem <- struct{}{}
	}

	wg.Wait()
	close(sem)

	return fresh
}

func puzzle1(filename string) int64 {
	list := parse(filename)

	fresh, _ := findIngredients(list)
	return fresh
}

func puzzle2(filename string) int64 {
	list := parse(filename)

	fresh := findIngredients2(list)
	return int64(len(fresh))
}
