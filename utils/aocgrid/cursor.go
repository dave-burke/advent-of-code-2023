package aocgrid

import "errors"

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
	value, err := c.Grid.pointAt(c.Position)
	if err != nil {
		// Should never happen
		panic(err)
	}
	return value
}
