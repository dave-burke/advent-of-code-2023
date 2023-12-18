package aocgrid

import (
	"fmt"
)

type Grid [][]rune

func NewGrid(lines []string) Grid {
	if lines[len(lines)-1] == "" {
		// Remove last line if it's empty
		lines = lines[:len(lines)-1]
	}
	result := make([][]rune, len(lines))

	for i, line := range lines {
		result[i] = []rune(line)
	}

	return result
}

func (g Grid) SetAt(c Cursor, newValue rune) {
	g[c.Position.Row][c.Position.Col] = newValue
}

func (g Grid) CursorAt(row, col int) Cursor {
	return Cursor{g, Point{row, col}}
}

func (g Grid) pointAt(p Point) (rune, error) {
	if g.isInBounds(p) {
		return g[p.Row][p.Col], nil
	} else {
		return ' ', fmt.Errorf("{%d, %d} is not on the grid", p.Row, p.Col)
	}
}

func (g Grid) isInBounds(p Point) bool {
	return (p.Row >= 0 && p.Row < len(g)) && (p.Col >= 0 && p.Col < len(g[p.Row]))
}

func (g Grid) All() []Cursor {
	all := make([]Cursor, 0)
	for row, rowSlice := range g {
		for col := range rowSlice {
			all = append(all, Cursor{g, Point{row, col}})
		}
	}
	return all
}

func (g Grid) ToString() string {
	result := ""
	for _, row := range g {
		result += string(row) + "\n"
	}
	return result
}
