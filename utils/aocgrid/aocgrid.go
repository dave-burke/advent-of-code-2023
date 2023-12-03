package aocgrid

import (
	"errors"
	"fmt"
)

type Point struct {
	row int
	col int
}

func (p Point) Row() int {
	return p.row
}

func (p Point) Col() int {
	return p.col
}

func (p Point) topLeft() Point {
	return Point{p.row - 1, p.col - 1}
}

func (p Point) topMiddle() Point {
	return Point{p.row - 1, p.col}
}

func (p Point) topRight() Point {
	return Point{p.row - 1, p.col + 1}
}

func (p Point) left() Point {
	return Point{p.row, p.col - 1}
}

func (p Point) right() Point {
	return Point{p.row, p.col + 1}
}

func (p Point) bottomLeft() Point {
	return Point{p.row + 1, p.col - 1}
}

func (p Point) bottomMiddle() Point {
	return Point{p.row + 1, p.col}
}

func (p Point) bottomRight() Point {
	return Point{p.row + 1, p.col + 1}
}

type Grid [][]rune

func (g Grid) GetAt(p Point) (rune, error) {
	if g.isInBounds(p) {
		return g[p.row][p.col], nil
	} else {
		return ' ', fmt.Errorf("{%d, %d} is not on the grid", p.row, p.col)
	}
}

func (g Grid) isInBounds(p Point) bool {
	return (p.row >= 0 && p.row < len(g)) && (p.col >= 0 && p.col < len(g[p.row]))
}

func (g Grid) size() int {
	return len(g.AllPoints())
}

func (g Grid) AllPoints() []Point {
	allPoints := make([]Point, 0)
	for row, rowSlice := range g {
		for col := range rowSlice {
			allPoints = append(allPoints, Point{row, col})
		}
	}
	return allPoints
}

func (g Grid) Neighbors(p Point) []rune {
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

	validNeighbors := make([]rune, 3)
	for _, neighbor := range potentialNeighbors {
		if g.isInBounds(neighbor) {
			neighborRune, _ := g.GetAt(neighbor)
			validNeighbors = append(validNeighbors, neighborRune)
		}
	}
	return validNeighbors
}

type Cursor struct {
	grid     Grid
	position Point
}

func NewCursor(grid Grid) Cursor {
	return Cursor{
		grid,
		Point{0, 0},
	}
}

func (c Cursor) Next() (Cursor, error) {
	pos := c.position
	next := Point{pos.row, pos.col + 1}
	if c.grid.isInBounds(next) {
		return Cursor{c.grid, next}, nil
	}
	next = Point{pos.row + 1, 0}
	if c.grid.isInBounds(next) {
		return Cursor{c.grid, next}, nil
	}
	return Cursor{}, errors.New("no more elements")
}

func (c Cursor) GetValue() rune {
	value, err := c.grid.GetAt(c.position)
	if err != nil {
		// Should never happen
		panic(err)
	}
	return value
}
