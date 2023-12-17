package day13

import (
	"aoc/utils/aocinput"
	"fmt"
	"log"
	"strings"
)

func Part1() string {
	// content := aocinput.ReadSampleAsString(13)
	content := aocinput.ReadInputAsString(13)
	content = strings.TrimSuffix(content, "\n")

	grids := strings.Split(content, "\n\n")

	horizontalSum := 0
	verticalSum := 0
	for _, gridString := range grids {
		lines := strings.Split(gridString, "\n")
		if reflection, ok := findReflection(lines); ok {
			log.Printf("Found %s reflection at %d", reflection.direction, reflection.index)
			if reflection.direction == HORIZONTAL {
				horizontalSum += reflection.index
			} else {
				verticalSum += reflection.index
			}
		} else {
			log.Print("No reflection found")
		}
	}

	return fmt.Sprint(verticalSum + (100 * horizontalSum))
}

var HORIZONTAL string = "HORIZONTAL"
var VERTICAL string = "VERTICAL"

type Reflection struct {
	direction string
	index     int
}

func findReflection(lines []string) (Reflection, bool) {
	log.Printf("PATTERN:\n%s", ToStringWithIndexes(lines))
	log.Printf("Search horizontal")
	reflectionIndex := findHorizontalReflection(lines)
	if reflectionIndex != -1 {
		return Reflection{HORIZONTAL, reflectionIndex}, true
	} else {
		log.Printf("Search vertical")
		lines = flip(lines)
		reflectionIndex = findHorizontalReflection(lines)
		if reflectionIndex != -1 {
			return Reflection{VERTICAL, reflectionIndex}, true
		} else {
			return Reflection{}, false
		}
	}
}

func findHorizontalReflection(lines []string) int {
	for i := 0; i < len(lines)-1; i++ {
		if lines[i] == lines[i+1] {
			// possible reflection
			doesReflect := true
			log.Printf("CHECK AT %d", i)
			for dist := 1; (i-dist) >= 0 && (i+dist+1) < len(lines); dist++ {
				lower := i - dist
				upper := i + dist + 1
				log.Printf("CHECK: 0 | %d || %d | %d", lower, upper, len(lines)-1)
				if !(lines[lower] == lines[upper]) {
					// found a line that didn't accurately reflect.
					doesReflect = false
					break
				}
			}
			if doesReflect {
				return i + 1
			}
		}
	}
	return -1
}

func flip(lines []string) []string {
	result := make([]string, 0)
	// for each column
	for i := 0; i < len(lines[0]); i++ {
		column := ""
		// for each row
		for _, line := range lines {
			column += string(line[i])
		}
		result = append(result, column)
	}
	return result
}

func ToStringWithIndexes(lines []string) string {
	result := "  "
	for i := range lines[0] {
		result += fmt.Sprint(i)
	}
	result += "\n"
	for i, row := range lines {
		result += fmt.Sprintf("%d %s\n", i, row)
	}
	return result
}

func Part2() string {
	lines := aocinput.ReadSampleAsLines(13)
	// lines := aocinput.ReadInputAsLines(13)

	log.Printf("Got %d lines", len(lines))

	return fmt.Sprint(len(lines))
}
