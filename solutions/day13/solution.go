package day13

import (
	"aoc/utils/aocinput"
	"fmt"
	"log"
)

func Part1() string {
	lines := aocinput.ReadSampleAsLines(13)
	// lines := aocinput.ReadInputAsLines(13)

	log.Printf("Got %d lines", len(lines))

	return fmt.Sprint(len(lines))
}

func Part2() string {
	lines := aocinput.ReadSampleAsLines(13)
	// lines := aocinput.ReadInputAsLines(13)

	log.Printf("Got %d lines", len(lines))

	return fmt.Sprint(len(lines))
}
