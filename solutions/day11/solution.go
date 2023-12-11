package day09

import (
	"aoc/utils/aocgrid"
	"aoc/utils/aocinput"
	"fmt"
	"log"
)

func Part1() string {
	grid := aocgrid.Grid(aocinput.ReadSampleAsGrid(11))
	fmt.Println(grid.ToString())
	expanded := expandGrid(grid)
	fmt.Println(expanded.ToString())
	log.Printf("%d x %d => %d x %d", len(grid), len(grid[0]), len(expanded), len(expanded[0]))
	return fmt.Sprint(len(grid))
}

func expandGrid(grid aocgrid.Grid) aocgrid.Grid {
	rowsToDuplicate := make([]int, 0)
	for i := range grid {
		if isRowEmpty(grid, i) {
			rowsToDuplicate = append(rowsToDuplicate, i)
		}
	}
	for _, rowNum := range rowsToDuplicate {
		grid = duplicateRow(grid, rowNum)
	}

	colsToDuplicate := make([]int, 0)
	for i := range grid[0] {
		if isColumnEmpty(grid, i) {
			colsToDuplicate = append(colsToDuplicate, i)
		}
	}
	for _, colNum := range colsToDuplicate {
		grid = duplicateCol(grid, colNum)
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
	row := grid[index]
	rowCopy := make([]rune, len(row))
	copy(rowCopy, row)
	grid = append(grid[0:index+1], grid[index:]...)
	return grid
}

func duplicateCol(grid aocgrid.Grid, index int) aocgrid.Grid {
	// won't work on last index
	for i, row := range grid {
		newRow := append(row[0:index+1], row[index:]...)
		grid[i] = newRow
	}
	return grid
}

func Part2() string {
	return ""
}
