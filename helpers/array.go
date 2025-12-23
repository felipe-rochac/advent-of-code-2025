package helpers

// Rotates matrix 90
func RotateMatrix[T any](m [][]T) [][]T {
	r := make([][]T, 0)

	if len(m) == 0 {
		return [][]T{}
	}

	for c := range len(m[0]) {
		row := make([]T, 0)
		for r := range len(m) {
			row = append(row, m[r][c])
		}
		r = append(r, row)
	}

	return r
}
func ReverseArray[T any](arr []T) []T {
	n := len(arr)
	res := make([]T, n)
	for i, v := range arr {
		res[n-1-i] = v
	}
	return res
}

