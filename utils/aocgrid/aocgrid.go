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


func (g Grid) AllPoints() []Point {
	allPoints := make([]Point, 0)
	for row, rowSlice := range g {
		for col := range rowSlice {
			allPoints = append(allPoints, Point{row, col})
		}
	}
	return allPoints
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

type Cursor struct {
	Grid     Grid
	Position Point
}

func InitCursor(grid Grid) Cursor {
	return Cursor{
		grid,
		Point{0, 0},
	}
}

func NewCursor(orig Cursor, newPos Point) (Cursor, error) {
	if orig.Grid.isInBounds(newPos) {
		return Cursor{orig.Grid, newPos}, nil
	} else {
		return Cursor{}, errors.New("out of bounds")
	}
}

func (c Cursor) Next() (Cursor, error) {
	pos := c.Position
	next := Point{pos.Row, pos.Col + 1}
	if c.Grid.isInBounds(next) {
		return Cursor{c.Grid, next}, nil
	}
	next = Point{pos.Row + 1, 0}
	if c.Grid.isInBounds(next) {
		return Cursor{c.Grid, next}, nil
	}
	return Cursor{}, errors.New("no more elements")
}

func (c Cursor) WalkUpLeft() (Cursor, error) {
	newPos := c.Position.topLeft()
	return NewCursor(c, newPos)
}

func (c Cursor) WalkUp() (Cursor, error) {
	newPos := c.Position.topMiddle()
	return NewCursor(c, newPos)
}

func (c Cursor) WalkUpRight() (Cursor, error) {
	newPos := c.Position.topRight()
	return NewCursor(c, newPos)
}

func (c Cursor) WalkLeft() (Cursor, error) {
	newPos := c.Position.left()
	return NewCursor(c, newPos)
}

func (c Cursor) WalkRight() (Cursor, error) {
	newPos := c.Position.right()
	return NewCursor(c, newPos)
}

func (c Cursor) WalkDownLeft() (Cursor, error) {
	newPos := c.Position.bottomLeft()
	return NewCursor(c, newPos)
}

func (c Cursor) WalkDown() (Cursor, error) {
	newPos := c.Position.bottomMiddle()
	return NewCursor(c, newPos)
}

func (c Cursor) WalkDownRight() (Cursor, error) {
	newPos := c.Position.bottomRight()
	return NewCursor(c, newPos)
}

func (c Cursor) GetValue() rune {
	value, err := c.Grid.GetAt(c.Position)
	if err != nil {
		// Should never happen
		panic(err)
	}
	return value
}
