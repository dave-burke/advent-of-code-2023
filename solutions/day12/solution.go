package day12

import (
	"aoc/utils/aocasync"
	"aoc/utils/aocfuncs"
	"aoc/utils/aocinput"
	"aoc/utils/aocparse"
	"fmt"
	"log"
	"strings"
	"sync"

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

var patternCache sync.Map

func enumeratePatterns(nChars int) []string {
	if cached, ok := patternCache.Load(nChars); ok {
		return cached.([]string)
	} else if nChars == 1 {
		return []string{"#", "."}
	} else {
		results := make([]string, 0)
		for _, pattern := range enumeratePatterns(nChars - 1) {
			results = append(results, fmt.Sprintf("%s%s", "#", pattern))
			results = append(results, fmt.Sprintf("%s%s", ".", pattern))
		}
		patternCache.Store(nChars, results)
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

	in := make(chan string)
	go func() {
		defer close(in)
		for _, line := range lines {
			in <- line
		}
	}()

	concurrency := 1 //runtime.NumCPU()
	log.Printf("Processing lines with %d workers", concurrency)
	outs := make([]<-chan int, 0, concurrency)
	for i := 0; i < concurrency; i++ {
		out := make(chan int)
		go func(workerNum int) {
			defer close(out)
			for line := range in {
				expandedLine := expandLine(line)
				log.Printf("(%d) Got %s", workerNum, expandedLine)
				result := CountArrangemen2(expandedLine)
				log.Printf("(%d) Finished: %d", i, result)
				out <- result
			}
		}(i)
		outs = append(outs, out)
	}
	results := aocasync.Merge(outs...)

	sum := 0
	for i := range results {
		sum += i
	}

	return fmt.Sprint(sum)
}

func expandLine(line string) string {
	parts := strings.Split(line, " ")
	expandedSprings := strings.Repeat(parts[0], 5)
	expandedGroups := strings.Repeat(parts[1], 5)
	return fmt.Sprintf("%s %s", expandedSprings, expandedGroups)
}

func CountArrangemen2(line string) int {
	rec := parseLine(line)

	nQuestions := countQuestions(rec.Springs)
	possiblePatterns := enumeratePatterns(nQuestions)

	result := 0
	for _, pattern := range possiblePatterns {
		patternIndex := 0

		possibleArrangement := ""
		groupCounts := make([]int, 0)
		currentGroupSize := 0
		for i, char := range rec.Springs {
			if char == '#' {
				currentGroupSize++
				possibleArrangement += "#"
			} else if char == '.' {
				possibleArrangement += "."
				if currentGroupSize > 0 {
					groupCounts = append(groupCounts, currentGroupSize)
					currentGroupSize = 0
					if len(groupCounts) > len(rec.Groups) || !cmp.Equal(groupCounts, rec.Groups[0:len(groupCounts)]) {
						//log.Printf("%s cannot be the beginning of %v", possibleArrangement, rec.Groups)
						break
					}
				}
			} else if char == '?' {
				patternChar := pattern[patternIndex]
				if patternChar == '#' {
					currentGroupSize++
				} else if patternChar == '.' {
					if currentGroupSize > 0 {
						groupCounts = append(groupCounts, currentGroupSize)
						currentGroupSize = 0
						if len(groupCounts) > len(rec.Groups) || !cmp.Equal(groupCounts, rec.Groups[0:len(groupCounts)]) {
							//log.Printf("%s cannot be the beginning of %v", possibleArrangement, rec.Groups)
							break
						}
					}
				}
				possibleArrangement += string(patternChar)
				patternIndex++
			} else {
				log.Fatalf("Invalid char %c at %d in %s", char, i, line)
			}
		}
		arrangementGroups := countGroups(possibleArrangement)
		if cmp.Equal(arrangementGroups, rec.Groups) {
			result++
		}
	}
	//log.Printf("%s => %d", line, result)
	return result
}
