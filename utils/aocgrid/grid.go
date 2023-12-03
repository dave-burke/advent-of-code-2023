package aocgrid

import (
	"fmt"
)

type Grid [][]rune

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

func (g Grid) Neighbors(p Point) []Cursor {
	potentialNeighbors := []Point{
		p.topLeft(),
		p.topMiddle(),
		p.topRight(),
		p.left(),
		p.right(),
		p.bottomLeft(),
		p.bottomMiddle(),
		p.bottomRight(),
	}

	validNeighbors := make([]Cursor, 0)
	for _, neighbor := range potentialNeighbors {
		if g.isInBounds(neighbor) {
			validNeighbors = append(validNeighbors, Cursor{g, neighbor})
		}
	}
	return validNeighbors
}
