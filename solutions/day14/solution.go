package day14

import (
	"aoc/utils/aocgrid"
	"aoc/utils/aocinput"
	"fmt"
	"log"
)

func Part1() string {
	grid := aocgrid.Grid(aocinput.ReadSampleAsGrid(14))
	//grid := aocgrid.Grid(aocinput.ReadInputAsGrid(14))

	log.Printf("GRID:\n%s", grid.ToString())
	for i, row := range grid {
		for j, char := range row {
			if char == 'O' {
				DropRock(grid, grid.CursorAt(i, j))
			}
		}
	}
	log.Printf("DROPPED:\n%s", grid.ToString())

	return fmt.Sprint(len(grid))
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

func Part2() string {
	lines := aocinput.ReadSampleAsLines(14)
	// lines := aocinput.ReadInputAsLines(14)

	log.Printf("Got %d lines", len(lines))

	return fmt.Sprint(len(lines))
}
