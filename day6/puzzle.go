package day6

import (
	"fmt"
	"strings"
	"sync"

	"github.com/felipe-rochac/advent-of-code-2025/helpers"
)

type cephalopod struct {
	numbers  []int64
	operator string
}

func parse(filename string) []cephalopod {
	lines, err := helpers.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	line := lines[len(lines)-1]
	parts := strings.Split(line, " ")

	var r []cephalopod

	for _, p := range parts {
		if p == "" {
			continue
		}
		c := cephalopod{
			operator: p,
			numbers:  make([]int64, 0),
		}
		r = append(r, c)
	}

	for _, line := range lines[:len(lines)-1] {
		parts := strings.Split(line, " ")
		i := 0
		for _, p := range parts {
			if p == "" {
				continue
			}
			r[i].numbers = append(r[i].numbers, helpers.MustParseInt64(p))
			i++
		}
	}

	return r
}

func puzzle1(filename string) int64 {
	list := parse(filename)
	var wg sync.WaitGroup
	sumCh := make(chan int64, len(list))

	for _, l := range list {
		wg.Go(func() {
			var sum int64 = 0
			if l.operator == "*" {
				sum = 1
			}
			for _, i := range l.numbers {
				switch l.operator {
				case "+":
					sum += i
				case "*":
					sum *= i
				}
			}
			sumCh <- sum
		})
	}

	wg.Wait()
	close(sumCh)

	var total int64 = 0
	for r := range sumCh {
		total += r
	}

	return total
}

func parse2(filename string) ([][]int64, helpers.Stack[string]) {
	lines, err := helpers.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	operators := helpers.Stack[string]{}

	operatorsLine := lines[len(lines)-1]

	for _, p := range strings.Split(operatorsLine, " ") {
		if p == "" {
			continue
		}

		operators.Push(p)
	}

	matrix := make([][]int64, len(lines)-1)

	for i, l := range lines[:len(lines)-1] {
		col := make([]int64, len(l))
		for j, r := range l {
			i := 0
			if r == ' ' {
				i = -1
			} else {
				i = helpers.RuneToInt(r)
			}
			col[j] = int64(i)
		}
		matrix[i] = col
	}

	return matrix, operators
}

func puzzle2(filename string) int64 {
	matrix, operators := parse2(filename)

	matrix = helpers.RotateMatrix(matrix)
	operators.Reverse()

	isSeparatorLine := func(r int64) bool {
		for c := range matrix[r] {
			if matrix[r][c] != -1 {
				return false
			}
		}
		return true
	}

	// Split matrix into groups so it can be paralellized
	groups := make([][][]int64, 0)
	line := make([][]int64, 0)
	for r := range len(matrix) {
		if isSeparatorLine(int64(r)) {
			groups = append(groups, line)
			line = make([][]int64, 0)
			continue
		}
		line = append(line, matrix[r])
	}
	groups = append(groups, line)

	// parse columns to a number applying * 10
	formNumber := func(value []int64) int64 {
		s := ""
		for col := range value {
			if value[col] == -1 {
				continue
			}
			s += fmt.Sprintf("%d", value[col])
		}
		return helpers.MustParseInt64(s)
	}

	sumCh := make(chan int64, len(groups))
	var wg sync.WaitGroup

	for _, g := range groups {
		op := operators.Pop()
		wg.Go(func() {
			var sum int64 = 0
			if op == "*" {
				sum = 1
			}
			for _, r := range g {
				n := formNumber(r)
				switch op {
				case "+":
					sum += n
				case "*":
					sum *= n
				}
			}

			sumCh <- sum
		})
	}

	wg.Wait()
	close(sumCh)

	var total int64 = 0
	for s := range sumCh {
		total += s
	}

	return total
}
