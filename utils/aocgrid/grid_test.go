package aocgrid

import "testing"

var exampleGrid = Grid{
	{'.', 'O', '.'},
	{'O', 'X', 'O'},
	{'.', 'O', '.'},
}

func TestAll(t *testing.T) {
	got := len(exampleGrid.All())
	wanted := 9

	if got != wanted {
		t.Errorf("Wanted %v but got %v", wanted, got)
	}
}

type boundsTest struct {
	point      Point
	isInBounds bool
}

var boundsTests = []boundsTest{
	{Point{-1, -1}, false},
	{Point{-1, 0}, false},
	{Point{-1, 1}, false},
	{Point{-1, 2}, false},
	{Point{-1, 3}, false},
	{Point{0, -1}, false},
	{Point{0, 0}, true},
	{Point{0, 1}, true},
	{Point{0, 2}, true},
	{Point{0, 3}, false},
	{Point{1, -1}, false},
	{Point{1, 0}, true},
	{Point{1, 1}, true},
	{Point{1, 2}, true},
	{Point{1, 3}, false},
	{Point{2, -1}, false},
	{Point{2, 0}, true},
	{Point{2, 1}, true},
	{Point{2, 2}, true},
	{Point{2, 3}, false},
	{Point{3, -1}, false},
	{Point{3, 0}, false},
	{Point{3, 1}, false},
	{Point{3, 2}, false},
	{Point{3, 3}, false},
}

func TestBounds(t *testing.T) {
	for _, test := range boundsTests {
		got := exampleGrid.isInBounds(test.point)
		wanted := test.isInBounds

		if got != wanted {
			t.Errorf("Wanted %v but got %v", wanted, got)
		}
	}

}
