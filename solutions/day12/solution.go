package day12

import (
	"aoc/utils/aocfuncs"
	"aoc/utils/aocinput"
	"aoc/utils/aocparse"
	"fmt"
	"log"
	"strings"
)

func Part1() string {
	lines := aocinput.ReadSampleAsLines(12)
	//lines := aocinput.ReadInputAsLines(12)

	log.Printf("Got %d lines", len(lines))

	return fmt.Sprint(len(lines))
}

type SpringRecord struct {
	Springs string
	Groups  []int
}

func parseLine(line string) SpringRecord {
	parts := strings.Split(line, " ")
	groups := strings.Split(parts[1], ",")
	groupNums := aocfuncs.Map[string, int](groups, aocparse.MustAtoi)
	return SpringRecord{parts[0], groupNums}
}

func countGroups(line string) []int {
	counts := make([]int, 0)
	for i := 0; i < len(line); i++ {
		count := 0
		for i < len(line) && line[i] == '#' {
			count++
			i++
		}
		if count > 0 {
			counts = append(counts, count)
		}
	}
	return counts
}

func CountArrangemen(line string) int {
	rec := parseLine(line)
	return len(rec.Groups)
}

func Part2() string {
	lines := aocinput.ReadSampleAsLines(12)
	//lines := aocinput.ReadInputAsLines(12)

	log.Printf("Got %d lines", len(lines))

	return fmt.Sprint(len(lines))
}
