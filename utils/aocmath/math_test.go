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

type LcmPair struct {
	input    []int
	expected int
}

func TestLcm(t *testing.T) {
	testCases := []LcmPair{
		{[]int{12, 15, 75}, 300},
	}

	for _, testCase := range testCases {
		result := Lcm(testCase.input...)
		if result != testCase.expected {
			t.Errorf("Wanted %d but got %d", result, testCase.expected)
		}
	}
}
