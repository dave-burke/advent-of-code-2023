package aocgrid

import "errors"

type Cursor struct {
	Grid     Grid
	Position Point
}

func newCursor(orig Cursor, newPos Point) (Cursor, error) {
	if orig.Grid.isInBounds(newPos) {
		return Cursor{orig.Grid, newPos}, nil
	} else {
		return Cursor{}, errors.New("out of bounds")
	}
}

func (c Cursor) HasNext() bool {
	_, err := c.Next()
	return err == nil
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
	return newCursor(c, newPos)
}

func (c Cursor) WalkUp() (Cursor, error) {
	newPos := c.Position.topMiddle()
	return newCursor(c, newPos)
}

func (c Cursor) WalkUpRight() (Cursor, error) {
	newPos := c.Position.topRight()
	return newCursor(c, newPos)
}

func (c Cursor) WalkLeft() (Cursor, error) {
	newPos := c.Position.left()
	return newCursor(c, newPos)
}

func (c Cursor) WalkRight() (Cursor, error) {
	newPos := c.Position.right()
	return newCursor(c, newPos)
}

func (c Cursor) WalkDownLeft() (Cursor, error) {
	newPos := c.Position.bottomLeft()
	return newCursor(c, newPos)
}

func (c Cursor) WalkDown() (Cursor, error) {
	newPos := c.Position.bottomMiddle()
	return newCursor(c, newPos)
}

func (c Cursor) WalkDownRight() (Cursor, error) {
	newPos := c.Position.bottomRight()
	return newCursor(c, newPos)
}

func (c Cursor) Neighbors() []Cursor {
	p := c.Position
	potentialNeighbors := []Point{
		p.topLeft(),
		p.topMiddle(),
		p.topRight(),
		p.right(),
		p.bottomRight(),
		p.bottomMiddle(),
		p.bottomLeft(),
		p.left(),
	}

	validNeighbors := make([]Cursor, 0)
	for _, neighbor := range potentialNeighbors {
		if c.Grid.isInBounds(neighbor) {
			validNeighbors = append(validNeighbors, Cursor{c.Grid, neighbor})
		}
	}
	return validNeighbors
}

func (c Cursor) GetValue() rune {
	value, err := c.Grid.pointAt(c.Position)
	if err != nil {
		// Should never happen
		panic(err)
	}
	return value
}
