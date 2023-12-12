package day12

import (
	"aoc/utils/aocfuncs"
	"aoc/utils/aocinput"
	"aoc/utils/aocparse"
	"fmt"
	"log"
	"strings"

	"github.com/google/go-cmp/cmp"
)

func Part1() string {
	// lines := aocinput.ReadSampleAsLines(12)
	lines := aocinput.ReadInputAsLines(12)

	log.Printf("Got %d lines", len(lines))

	sum := 0
	for i, line := range lines {
		log.Printf("Counting line %d", i)
		sum += CountArrangemen(line)
	}

	return fmt.Sprint(sum)
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

func enumeratePatterns(nChars int) []string {
	if nChars == 1 {
		return []string{"#", "."}
	} else {
		results := make([]string, 0)
		for _, pattern := range enumeratePatterns(nChars - 1) {
			results = append(results, fmt.Sprintf("%s%s", "#", pattern))
			results = append(results, fmt.Sprintf("%s%s", ".", pattern))
		}
		return results
	}
}

func countQuestions(line string) int {
	result := 0
	for _, char := range line {
		if char == '?' {
			result++
		}
	}
	return result
}

func applyPattern(line string, pattern string) string {
	patternIndex := 0

	result := ""
	for _, char := range line {
		if char == '?' {
			result += string(pattern[patternIndex])
			patternIndex++
		} else {
			result += string(char)
		}
	}
	return result
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

	nQuestions := countQuestions(rec.Springs)
	possiblePatterns := enumeratePatterns(nQuestions)

	result := 0
	for _, possiblePattern := range possiblePatterns {
		possibleArrangement := applyPattern(line, possiblePattern)
		arrangementGroups := countGroups(possibleArrangement)
		if cmp.Equal(arrangementGroups, rec.Groups) {
			result++
		}
	}
	return result
}

func Part2() string {
	lines := aocinput.ReadSampleAsLines(12)
	//lines := aocinput.ReadInputAsLines(12)

	log.Printf("Got %d lines", len(lines))

	return fmt.Sprint(len(lines))
}
