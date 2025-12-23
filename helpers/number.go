package helpers

// Find dividors of an int
func FindDivisors(num int) []int {
	result := make([]int, 0)
	for i := num / 2; i > 1; i-- {
		if i%num == 1 {
			continue
		}

		result = append(result, i)
	}

	result = append(result, 1)

	return result
}

// Parse a rune to an int
func RuneToInt(r rune) int {
	return int(r - '0')
}

// Combine returns all combinations of the given size from the array, where each combination is represented as a single int (digits merged).
func Combine(array []int, size int) []int {
	if size == 0 || len(array) < size {
		return []int{}
	}
	var result []int
	var comb func(start int, curr []int)
	comb = func(start int, curr []int) {
		if len(curr) == size {
			// Merge digits into a single number
			num := 0
			for _, d := range curr {
				num = num*10 + d
			}
			result = append(result, num)
			return
		}
		for i := start; i < len(array); i++ {
			comb(i+1, append(curr, array[i]))
		}
	}
	comb(0, []int{})
	return result
}

func MaxNumberWithOrder(s string, k int) string {
	n := len(s)
	toRemove := n - k
	stack := make([]byte, 0, k)
	for i := 0; i < n; i++ {
		c := s[i]
		for len(stack) > 0 && toRemove > 0 && stack[len(stack)-1] < c {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, c)
	}
	return string(stack[:k])
}

func Between(v, from, to int64) bool {
	return v >= from && v <= to
}

// Return max element in an int array
func Max(array []int) int {
	max := 0

	for _, a := range array {
		if a > max {
			max = a
		}
	}

	return max
}

func MaxAndPos(array []int) (max, pos int) {
	max = 0
	pos = 0

	for i, a := range array {
		if a > max {
			max = a
			pos = i
		}
	}

	return
}
