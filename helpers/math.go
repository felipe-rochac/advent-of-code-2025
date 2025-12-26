package helpers

import "math"

func Abs[T ~int | ~float64](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func Pow[T ~int | ~float64](a T, n int) T {
	result := T(1)
	for i := 0; i < n; i++ {
		result *= a
	}
	return result
}

func Sqrt[T ~int | ~float64](a T, n int) T {
	return T(math.Pow(float64(a), 1.0/float64(n)))
}

// d= sqrt{(x_{2}-x_{1})^{2}+(y_{2}-y_{1})^{2}+(z_{2}-z_{1})^{2}}
func EuclidianDistance[T ~int | ~float64](x1, y1, z1, x2, y2, z2 T) T {
	dx := T(x1 - x2)
	dy := T(y1 - y2)
	dz := T(z1 - z2)
	return Sqrt(Pow(dx, 2)+Pow(dy, 2)+Pow(dz, 2), 2)
}

func ManhattanDistance[T ~int | ~float64](x1, y1, z1, x2, y2, z2 T) T {
	dx := T(x1 - x2)
	dy := T(y1 - y2)
	dz := T(z1 - z2)
	return dx + dy + dz
}

func Area[T ~int | ~float64](x1, y1, z1, x2, y2, z2 T) T {
	return Abs(x1-x2+1) * Abs(y1-y2+1) * Abs(z1-z2+1)
}
