package day12

import (
	"aoc/utils/aocinput"
	"fmt"
	"log"
)

func Part1() string {
	lines := aocinput.ReadSampleAsLines(12)
	//lines := aocinput.ReadInputAsLines(12)

	log.Printf("Got %d lines", len(lines))

	return fmt.Sprint(len(lines))
}

func Part2() string {
	lines := aocinput.ReadSampleAsLines(12)
	//lines := aocinput.ReadInputAsLines(12)

	log.Printf("Got %d lines", len(lines))

	return fmt.Sprint(len(lines))
}

