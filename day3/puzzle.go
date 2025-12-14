package day3

import (
	"strconv"

	"github.com/felipe-rochac/advent-of-code-2025/helpers"
)

type Bank struct {
	Powers []int
	Number string
}

func parse(filename string) []Bank {
	banks := make([]Bank, 0)
	lines, err := helpers.ReadLines(filename)
	if err != nil {
		panic(err)
	}
	for _, l := range lines {
		powers := make([]int, 0)
		for _, ch := range l {
			powers = append(powers, helpers.RuneToInt(ch))
		}
		banks = append(banks, Bank{Powers: powers, Number: l})
	}

	return banks
}

func puzzle1(filename string) int64 {
	banks := parse(filename)
	result := 0

	for _, b := range banks {
		powers := helpers.Combine(b.Powers, 2)
		max := helpers.Max(powers)
		print("max was ", max, "\n")
		result += max
	}

	return int64(result)
}

func puzzle2(filename string) int64 {
	banks := parse(filename)
	result := 0

	for _, b := range banks {
		max := helpers.MaxNumberWithOrder(b.Number, 12)
		print("max was ", max, "\n")
		n, _ := strconv.Atoi(max)
		result += n
	}

	return int64(result)
}
