package aocgrid

import "testing"

var origin = Point{0, 0}

func pointTest(got, wanted Point, t *testing.T) {
	if got != wanted {
		t.Errorf("Wanted %v but got %v", wanted, got)
	}
}

func TestPoint(t *testing.T) {
	pointTest(origin.topLeft(), Point{-1, -1}, t)
	pointTest(origin.topMiddle(), Point{-1, 0}, t)
	pointTest(origin.topRight(), Point{-1, 1}, t)
	pointTest(origin.left(), Point{0, -1}, t)
	pointTest(origin.right(), Point{0, 1}, t)
	pointTest(origin.bottomLeft(), Point{1, -1}, t)
	pointTest(origin.bottomMiddle(), Point{1, 0}, t)
	pointTest(origin.bottomRight(), Point{1, 1}, t)
}
