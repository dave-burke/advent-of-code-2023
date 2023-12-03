package aocgrid

import "testing"

var exampleCursorGrid = Grid{
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
}

func TestNext(t *testing.T) {
	c := exampleCursorGrid.CursorAt(0, 0)

	chars := []rune{c.GetValue()}
	for c.HasNext() {
		c, _ = c.Next()
		chars = append(chars, c.GetValue())
	}

	got := string(chars)
	want := "abcdefghi"

	if got != want {
		t.Errorf("Wanted '%s' but got '%s'", want, got)
	}

}
