package day04

import (
	"aoc/utils/aocinput"
	"fmt"
)

func Part1() string {
	lines := aocinput.ReadSampleAsLines(5)
	return fmt.Sprint(len(lines))
}

type farmMap []farmMapping

func (m farmMap) get(key int) int {
	for _,mapping := range m {
		if val, ok := mapping.getWithin(key); ok {
			return val
		}
	}
	return key
}

type farmMapping struct {
	destinationRangeStart int
	sourceRangeStart      int
	length                int
}

func (m farmMapping) contains(key int) bool {
	return key >= m.sourceRangeStart && key < m.sourceRangeStart+m.length
}

func (m farmMapping) getWithin(key int) int, bool {
	if m.contains(key) {
		dist := key - m.sourceRangeStart
		return m.destinationRangeStart + dist, true
	}
	return _, false
}

func Part2() string {
	return "todo"
}
