package day09

import (
	"aoc/utils/aocfuncs"
	"aoc/utils/aocgrid"
	"aoc/utils/aocinput"
	"fmt"
	"log"
	"math"
)

func Part1() string {
	grid := aocgrid.Grid(aocinput.ReadSampleAsGrid(11))
	fmt.Println(grid.ToString())
	expanded := expandGrid(grid)
	fmt.Println(expanded.ToString())
	log.Printf("Expanded %d x %d => %d x %d", len(grid), len(grid[0]), len(expanded), len(expanded[0]))
	galaxies := findGalaxies(expanded)

	coords := aocfuncs.Map[aocgrid.Cursor, string](galaxies, func(item aocgrid.Cursor) string {
		return fmt.Sprint(item.Position)
	})
	log.Printf("Galaxies at: %v", coords)

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			dist := calcDist(galaxies[i], galaxies[j])
			log.Printf("%d <-> %d | %v <-> %v = %d", i, j, galaxies[i].Position, galaxies[j].Position, dist)
			sum += dist
		}
	}

	return fmt.Sprint(sum)
}

func calcDist(a, b aocgrid.Cursor) int {
	colDist := math.Abs(float64(a.Position.Col) - float64(b.Position.Col))
	rowDist := math.Abs(float64(a.Position.Row) - float64(b.Position.Row))
	return int(colDist) + int(rowDist)
}

func findGalaxies(grid aocgrid.Grid) []aocgrid.Cursor {
	result := make([]aocgrid.Cursor, 0)
	for _, cursor := range grid.All() {
		if cursor.GetValue() == '#' {
			result = append(result, cursor)
		}
	}
	return result
}

func expandGrid(grid aocgrid.Grid) aocgrid.Grid {
	for i := len(grid) - 1; i >= 0; i-- {
		if isRowEmpty(grid, i) {
			grid = duplicateRow(grid, i)
		}
	}
	for i := len(grid[0]) - 1; i >= 0; i-- {
		if isColumnEmpty(grid, i) {
			grid = duplicateCol(grid, i)
		}
	}
	return grid
}

func isRowEmpty(grid aocgrid.Grid, rowNum int) bool {
	row := grid[rowNum]
	for _, char := range row {
		if char != '.' {
			return false
		}
	}
	return true
}

func isColumnEmpty(grid aocgrid.Grid, colNum int) bool {
	for _, row := range grid {
		if row[colNum] != '.' {
			return false
		}
	}
	return true
}

func duplicateRow(grid aocgrid.Grid, index int) aocgrid.Grid {
	// won't work on last index
	newGrid := append(grid[0:index+1], grid[index:]...)
	return newGrid
}

func duplicateCol(grid aocgrid.Grid, index int) aocgrid.Grid {
	// won't work on last index
	newGrid := make([][]rune, 0)
	for _, row := range grid {
		newRow := append(row[0:index+1], row[index:]...)
		newGrid = append(newGrid, newRow)
	}
	return newGrid
}

func Part2() string {
	return ""
}
