package day1

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/felipe-rochac/advent-of-code-2025/helpers"
	"github.com/stretchr/testify/assert"
)

func TestDial(t *testing.T) {
	dialPoints := DialPointer{
		CurrentPos: 50,
		InitPos:    50,
	}
	clicks := dialPoints.rotateDial(Left, 55)

	assert.Equal(t, clicks, 95)
	assert.Equal(t, dialPoints.CurrentPos, 95)
	assert.Equal(t, dialPoints.LastDirection, Left)
	assert.Equal(t, dialPoints.LastPos, 50)

	clicks = dialPoints.rotateDial(Right, 10)

	assert.Equal(t, clicks, 15)
	assert.Equal(t, dialPoints.CurrentPos, 10)
	assert.Equal(t, dialPoints.LastDirection, Right)
	assert.Equal(t, dialPoints.LastPos, 95)
}

func TestPuzzle1(t *testing.T) {
	lines, err := helpers.ReadLines("./input1.txt")
	if err != nil {
		panic(err)
	}

	pointer := DialPointer{
		CurrentPos: 50,
		InitPos:    50,
	}

	parse := func(s string) (Direction, int) {
		pos, err := strconv.Atoi(s[1:])
		if err != nil {
			panic(fmt.Sprintf("error parsing %s to int", s[1:]))
		}
		dir := Left
		if strings.EqualFold(string(s[0]), "R") {
			dir = Right
		}
		return dir, pos
	}

	password := 0
	for _, l := range lines {
		dir, pos := parse(l)
		pointer.rotateDial(dir, pos)

		fmt.Printf("pointer (%s, %d) cur: %d last: %d \n", dir, pos, pointer.CurrentPos, pointer.LastPos)

		if pointer.CurrentPos == 0 {
			password++
		}
	}

	fmt.Printf("the current password is %d\n", password)
}

func TestPuzzle2(t *testing.T) {
	lines, err := helpers.ReadLines("./input2.txt")
	if err != nil {
		panic(err)
	}

	pointer := DialPointer{
		CurrentPos: 50,
		InitPos:    50,
	}

	parse := func(s string) (Direction, int) {
		pos, err := strconv.Atoi(s[1:])
		if err != nil {
			panic(fmt.Sprintf("error parsing %s to int", s[1:]))
		}
		dir := Left
		if strings.EqualFold(string(s[0]), "R") {
			dir = Right
		}
		return dir, pos
	}

	password := 0
	for _, l := range lines {
		dir, pos := parse(l)
		rounds := pointer.rotateDial0x434C49434B(dir, pos)

		fmt.Printf("pointer (%s, %d) cur: %d last: %d \n", dir, pos, pointer.CurrentPos, pointer.LastPos)

		password += rounds
	}

	fmt.Printf("the current password is %d\n", password)
}
