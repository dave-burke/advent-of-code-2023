package day04

import (
	"aoc/utils/aocasync"
	"aoc/utils/aocinput"
	"aoc/utils/aocmath"
	"fmt"
)

func Part1() string {
	input := aocinput.ReadSampleAsChannel(4)

	gamePoints := aocasync.MapLinesAsync[int](input, func(line string) int {
		// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		return 0
	})
	return fmt.Sprint(aocmath.SumChan(gamePoints))
}

func Part2() string {
	return "todo"
}
