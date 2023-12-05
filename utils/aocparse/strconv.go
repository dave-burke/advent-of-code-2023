package aocparse

import (
	"strconv"
)

func MustAtoi(str string) int {
	if result, err := strconv.Atoi(str); err != nil {
		panic(err)
	} else {
		return result
	}
}
