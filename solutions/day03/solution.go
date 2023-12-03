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

	return fmt.Sprintf("%d", sum)
}

func isPart(point aocgrid.Point, grid aocgrid.Grid) bool {
	c := aocgrid.Cursor{Grid: grid, Position: point}
	for _, neighbor := range c.Neighbors() {
		if isSymbol(neighbor.GetValue()) {
			return true
		}
	}
	return false
}

func isSymbol(r rune) bool {
	return r == '+' ||
		r == '*' ||
		r == '$' ||
		r == '#' ||
		r == '%' ||
		r == '/' ||
		r == '=' ||
		r == '@' ||
		r == '&' ||
		r == '-'
}

func Part2() string {
	grid := aocgrid.Grid(aocinput.ReadInputAsGrid(3))

	re := regexp.MustCompile("[0-9]+")

	starsToNumbers := map[aocgrid.Point]map[int]bool{}
	for i, line := range grid {
		row := string(line)
		numbers := re.FindAllStringSubmatchIndex(row, -1)
		log.Printf("[%d] %s => %v", i, row, numbers)
		for _, indexes := range numbers {
			numString := row[indexes[0]:indexes[1]]
			if number, err := strconv.Atoi(numString); err != nil {
				log.Printf("Oops! %s is not a number", numString)
			} else {
				for j := indexes[0]; j < indexes[1]; j++ {
					p := aocgrid.Point{Row: i, Col: j}
					stars := findAdjacentStars(p, grid)
					if len(stars) > 0 {
						log.Printf("%d is adjacent to stars: %v", number, stars)
						for _, star := range stars {
							if starsToNumbers[star] == nil {
								starsToNumbers[star] = map[int]bool{}
							}
							starsToNumbers[star][number] = true
						}
						break
					}
				}
			}
		}
	}

	sum := 0
	for star, numbers := range starsToNumbers {
		if len(numbers) == 2 {
			ratio := 1
			for number := range numbers {
				ratio *= number
			}
			log.Printf("%v is adjacent to two numbers: %v (ration = %d)", star, numbers, ratio)
			sum += ratio
		}
	}

	return fmt.Sprintf("%d", sum)
}

func findAdjacentStars(p aocgrid.Point, g aocgrid.Grid) []aocgrid.Point {
	cursor := aocgrid.Cursor{Grid: g, Position: p}
	neighbors := cursor.Neighbors()

	neighborStars := make([]aocgrid.Point, 0)
	for _, neighbor := range neighbors {
		if neighbor.GetValue() == '*' {
			neighborStars = append(neighborStars, neighbor.Position)
		}
	}
	return neighborStars
}
