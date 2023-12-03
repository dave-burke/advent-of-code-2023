package day03

import (
	"aoc/utils/aocgrid"
	"aoc/utils/aocinput"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func Part1() string {
	grid := aocgrid.Grid(aocinput.ReadInputAsGrid(3))

	re := regexp.MustCompile("[0-9]+")

	partNumbers := make([]int, 0)
	for i, line := range grid {
		row := string(line)
		numbers := re.FindAllStringSubmatchIndex(row, -1)
		log.Printf("[%d] %s => %v", i, row, numbers)
		for _, indexes := range numbers {
			numString := row[indexes[0]:indexes[1]]
			if number, err := strconv.Atoi(numString); err != nil {
				log.Printf("Oops! %s is not a number", numString)
			} else {
				found := false
				for j := indexes[0]; j < indexes[1]; j++ {
					p := aocgrid.Point{Row: i, Col: j}
					if isPart(p, grid) {
						log.Printf("%d IS a part", number)
						partNumbers = append(partNumbers, number)
						found = true
						break
					}
				}
				if !found {
					log.Printf("%d IS NOT a part", number)
				}
			}
		}
	}

	sum := 0
	for _, partNumber := range partNumbers {
		sum += partNumber
	}

	// TODO 400550 is too low
	return fmt.Sprintf("%d", sum)
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
