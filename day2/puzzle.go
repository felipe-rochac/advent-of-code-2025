package day2

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/felipe-rochac/advent-of-code-2025/helpers"
)

type Range struct {
	First int64
	Last  int64
}

func parse(line string) []Range {
	parts := strings.Split(line, ",")

	ranges := make([]Range, len(parts))
	var wg sync.WaitGroup
	var rw sync.RWMutex

	for i, p := range parts {
		wg.Go(func() {
			r := strings.Split(p, "-")

			rw.Lock()
			ranges[i] = Range{
				First: helpers.MustParseInt64(r[0]),
				Last:  helpers.MustParseInt64(r[1]),
			}
			rw.Unlock()
		})
	}

	wg.Wait()

	return ranges
}

func puzzle1(fileName string) int64 {
	lines, err := helpers.ReadLines(fileName)

	if err != nil {
		panic(err)
	}
	if len(lines) != 1 {
		panic("invalid input file")
	}

	ranges := parse(lines[0])

	isInvalid := func(i int64) bool {
		s := fmt.Sprintf("%d", i)
		l := len(s)
		if l%2 == 1 {
			return false
		}
		mid := int(l / 2)

		return s[:mid] == s[mid:]
	}

	var invalid int64 = 0
	var wg sync.WaitGroup

	for _, r := range ranges {
		for i := r.First; i <= r.Last; i++ {
			wg.Go(func() {
				if isInvalid(i) {
					atomic.AddInt64(&invalid, i)
				}
			})
		}
	}

	wg.Wait()

	return invalid
}

func puzzle2(fileName string) int64 {
	lines, err := helpers.ReadLines(fileName)

	if err != nil {
		panic(err)
	}
	if len(lines) != 1 {
		panic("invalid input file")
	}

	ranges := parse(lines[0])

	isInvalid := func(value int64) bool {
		s := fmt.Sprintf("%d", value)
		l := len(s)
		divisors := helpers.FindDivisors(l)

		if l == 1 {
			return false
		}

		// for each size, tries to find combination
		invalid := false

		for _, d := range divisors {
			if d == l {
				continue // skip the full length
			}
			p := s[0:d]
			r := strings.Repeat(p, l/d)
			if s == r && l/d >= 2 {
				return true
			}
		}

		return invalid
	}

	var invalid int64 = 0
	var wg sync.WaitGroup

	for _, r := range ranges {
		for i := r.First; i <= r.Last; i++ {
			wg.Add(1)
			go func(i int64) {
				defer wg.Done()
				// print("validating range(", r.First, ", ", r.Last, "): ", i, "\n")
				if isInvalid(i) {
					print("stored value was: ", i, "\n")
					atomic.AddInt64(&invalid, i)
				}
			}(i)
		}
	}

	wg.Wait()

	return invalid
}
