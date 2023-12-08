package day07

import "testing"

func TestCompareHands2(t *testing.T) {
	handA := hand{"J345A", 0, "A345A"}
	handB := hand{"2345J", 0, "23455"}

	got := compareHands2(handA, handB)
	wanted := true

	if got != wanted {
		t.Errorf("Wanted %t but got %t", wanted, got)
	}
}
