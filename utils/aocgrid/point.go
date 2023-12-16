package aocgrid

type Point struct {
	Row int
	Col int
}

func pointsEqual(a, b Point) bool {
	return a.Row == b.Row && a.Col == b.Col
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
