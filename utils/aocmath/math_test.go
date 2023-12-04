package aocmath

import "testing"

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3}

	want := 6
	got := Sum(nums)

	if want != got {
		t.Errorf("Wanted %d but got %d", want, got)
	}
}
