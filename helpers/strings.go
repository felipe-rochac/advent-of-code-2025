package helpers

import "strconv"

func MustParseInt64(s string) int64 {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}
	return int64(i)
}
