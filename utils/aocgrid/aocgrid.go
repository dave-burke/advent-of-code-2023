package aocgrid

import (
	"errors"
	"fmt"
)

type Point struct {
	Row int
	Col int
}

func (p Point) topLeft() Point {
	return Point{p.Row - 1, p.Col - 1}
}

func (p Point) topMiddle() Point {
	return Point{p.Row - 1, p.Col}
}

func (p Point) topRight() Point {
	return Point{p.Row - 1, p.Col + 1}
}

func (p Point) left() Point {
	return Point{p.Row, p.Col - 1}
}

func (p Point) right() Point {
	return Point{p.Row, p.Col + 1}
}

func (p Point) bottomLeft() Point {
	return Point{p.Row + 1, p.Col - 1}
}

func (p Point) bottomMiddle() Point {
	return Point{p.Row + 1, p.Col}
}

func (p Point) bottomRight() Point {
	return Point{p.Row + 1, p.Col + 1}
}

type Grid [][]rune

func (g Grid) GetAt(p Point) (rune, error) {
	if g.isInBounds(p) {
		return g[p.Row][p.Col], nil
	} else {
		return ' ', fmt.Errorf("{%d, %d} is not on the grid", p.Row, p.Col)
	}
}

func (g Grid) isInBounds(p Point) bool {
	return (p.Row >= 0 && p.Row < len(g)) && (p.Col >= 0 && p.Col < len(g[p.Row]))
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
	next := Point{pos.Row, pos.Col + 1}
	if c.grid.isInBounds(next) {
		return Cursor{c.grid, next}, nil
	}
	next = Point{pos.Row + 1, 0}
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
