package helpers

func AssertEqual[T comparable](a, b T) {
	if a != b {
		panic("values are not equal")
	}
}

func AssertBiggerThan[T ~int | ~int64](a, b T) {
	if a <= b {
		panic("expected a to be bigger than b")
	}
}

func AssertLowerThan[T ~int | ~int64](a, b T) {
	if a >= b {
		panic("expected a to be lower than b")
	}
}
