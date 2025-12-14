package helpers

func AssertEqual[T comparable](a, b T) {
	if a != b {
		panic("values are not equal")
	}
}
