package day3

import (
	"testing"

	"github.com/felipe-rochac/advent-of-code-2025/helpers"
)

func TestPuzzle1_Test(t *testing.T) {
	result := puzzle1("test.txt")

	helpers.AssertEqual(result, 357)
}

func TestPuzzle1_Input(t *testing.T) {
	result := puzzle1("input.txt")

	helpers.AssertEqual(result, 17403)
}

func TestPuzzle2_Test(t *testing.T) {
	result := puzzle2("test.txt")

	helpers.AssertEqual(result, 3121910778619)
}

func TestPuzzle2_Input(t *testing.T) {
	result := puzzle2("input.txt")

	helpers.AssertEqual(result, 173416889848394)
}
