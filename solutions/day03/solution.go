package day03

import (
	"aoc/utils/aocgrid"
	"aoc/utils/aocinput"
	"fmt"
	"log"
	"regexp"
)

func Part1() string {
	grid := aocgrid.Grid(aocinput.ReadSampleAsGrid(3))

	// TODO identify strings of numbers.
	// This line contains 467 and 114:
	// 467..114..

	nParts := 0
	for _, point := range grid.AllPoints() {
		r, err := grid.GetAt(point)
		if err != nil {
			continue
		}
		if isNumber(point, grid) && isPart(point, grid) {
			fmt.Printf("%c at {%d, %d} is a part\n", r, point.Row(), point.Col())
			nParts += int(r - '0')
		} else {
			fmt.Printf("%c at {%d, %d} is NOT a part\n", r, point.Row(), point.Col())
		}
	}
	return fmt.Sprintf("%d", nParts)
}

func isNumber(p aocgrid.Point, g aocgrid.Grid) bool {
	r, err := g.GetAt(p)
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile("[0-9]")
	return re.MatchString(string(r))
}

func isPart(p aocgrid.Point, g aocgrid.Grid) bool {
	neighbors := g.Neighbors(p)
	for _, neighbor := range neighbors {
		if isSymbol(neighbor) {
			return true
		}
	}
	return false
}

func isSymbol(r rune) bool {
	return r == '+' || r == '*' || r == '$' || r == '#'
}

func Part2() string {
	return "todo"
}
