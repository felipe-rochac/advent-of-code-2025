package day1

import (
	"sync"
)

type DialPointer struct {
	CurrentPos    int
	LastDirection Direction
	LastPos       int
	InitPos       int
}

type Direction string

const (
	Left  Direction = "L"
	Right Direction = "R"
)

var dial []int

func init() {
	dial = make([]int, 0)
	for i := range 100 {
		dial = append(dial, i)
	}
}

func (d *DialPointer) rotateDial(dir Direction, clicks int) int {
	size := len(dial)
	clicks = clicks % size
	if dir == Left {
		clicks = clicks * -1
	}
	dest := d.CurrentPos + clicks
	curPos := dest
	if dest < 0 || dest > size {
		switch dir {
		case Left:
			curPos = (d.CurrentPos + size) + clicks
		case Right:
			curPos = (clicks + d.CurrentPos) - size
		}
	} else if dest == 100 {
		curPos = 0
	}

	d.LastDirection = dir
	d.LastPos = d.CurrentPos
	d.CurrentPos = curPos

	return dest
}

func (d *DialPointer) rotateDial0x434C49434B(dir Direction, clicks int) int {
	busyCh := make(chan struct{}, 1)
	var wg sync.WaitGroup
	rounds := 0

	for range clicks {
		wg.Go(func() {
			// lock routines
			busyCh <- struct{}{}

			d.rotateDial(dir, 1)

			if d.CurrentPos == 0 {
				rounds++
			}

			// free
			<-busyCh
		})
	}

	wg.Wait()

	return rounds
}
