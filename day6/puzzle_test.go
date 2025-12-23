package day6

import (
	"testing"

	"github.com/felipe-rochac/advent-of-code-2025/helpers"
)

func TestPuzzle1_Test(t *testing.T) {
	result := puzzle1("test.txt")

	helpers.AssertEqual(result, 4277556)
}

func TestPuzzle1_Input(t *testing.T) {
	result := puzzle1("input.txt")

	helpers.AssertEqual(result, 6343365546996)
}

func TestPuzzle2_Test(t *testing.T) {
	result := puzzle2("test.txt")

	helpers.AssertEqual(result, 3263827)
}

func TestPuzzle2_Input(t *testing.T) {
	result := puzzle2("input.txt")

	helpers.AssertEqual(result, 11136895955912)
}
