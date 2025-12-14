package day2

import (
	"testing"

	"github.com/felipe-rochac/advent-of-code-2025/helpers"
)

func TestPuzzle1_TestInput(t *testing.T) {
	result := puzzle1("./test.txt")

	helpers.AssertEqual(result, int64(1227775554))
}

func TestPuzzle1_Input(t *testing.T) {
	result := puzzle1("./input.txt")

	helpers.AssertEqual(result, int64(38437576669))
}

func TestPuzzle2_TestInput(t *testing.T) {
	result := puzzle2("./test.txt")

	helpers.AssertEqual(result, int64(4174379265))
}

func TestPuzzle2_Input(t *testing.T) {
	result := puzzle2("./input.txt")

	helpers.AssertEqual(result, int64(49046150754))
}
