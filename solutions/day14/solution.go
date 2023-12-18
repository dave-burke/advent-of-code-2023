package day14

import (
	"aoc/utils/aocgrid"
	"aoc/utils/aocinput"
	"fmt"
	"log"
)

func Part1() string {
	// grid := aocgrid.Grid(aocinput.ReadSampleAsGrid(14))
	grid := aocgrid.Grid(aocinput.ReadInputAsGrid(14))

	log.Printf("GRID:\n%s", grid.ToString())
	DropRocks(grid)
	log.Printf("DROPPED:\n%s", grid.ToString())

	result := CalcLoad(grid)

	return fmt.Sprint(result)
}

func DropRocks(g aocgrid.Grid) {
	for i, row := range g {
		for j, char := range row {
			if char == 'O' {
				DropRock(g, g.CursorAt(i, j))
			}
		}
	}
}

func DropRock(g aocgrid.Grid, rock aocgrid.Cursor) {
	if above, err := rock.WalkUp(); err == nil {
		if above.GetValue() == '.' {
			g.SetAt(rock, '.')
			g.SetAt(above, 'O')
			DropRock(g, above)
		}
	}
}

func CalcLoad(g aocgrid.Grid) int {
	result := 0
	for i, row := range g {
		rowLoad := len(g) - i
		rocksInRow := 0
		for _, char := range row {
			if char == 'O' {
				rocksInRow++
			}
		}
		totalRowLoad := rowLoad * rocksInRow
		log.Printf("There are %d rocks in row %d, each contributing %d for a total of %d", rocksInRow, i, rowLoad, totalRowLoad)
		result += totalRowLoad
	}
	return result
}

func Part2() string {
	lines := aocinput.ReadSampleAsLines(14)
	// lines := aocinput.ReadInputAsLines(14)

	log.Printf("Got %d lines", len(lines))

	return fmt.Sprint(len(lines))
}
