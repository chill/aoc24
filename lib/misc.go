package lib

import "strconv"

func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func Abs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}
