package day4

import (
	"testing"

	"github.com/felipe-rochac/advent-of-code-2025/helpers"
)

func TestPuzzle1_Test(t *testing.T) {
	result := puzzle1("test.txt")

	helpers.AssertEqual(result, 13)
}

func TestPuzzle1_Input(t *testing.T) {
	result := puzzle1("input.txt")

	helpers.AssertEqual(result, 1564)
}

func TestPuzzle2_Test(t *testing.T) {
	result := puzzle2("test.txt")

	helpers.AssertEqual(result, 43)
}

func TestPuzzle2_Input(t *testing.T) {
	result := puzzle2("input.txt")

	helpers.AssertEqual(result, 9401)
}
