package day12

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCountGroups(t *testing.T) {
	testLines := map[string][]int{
		"#.#.###":             {1, 1, 3},
		".#...#....###.":      {1, 1, 3},
		".#.###.#.######":     {1, 3, 1, 6},
		"####.#...#...":       {4, 1, 1},
		"#....######..#####.": {1, 6, 5},
		".###.##....#":        {3, 2, 1},
	}
	for line, expected := range testLines {
		result := countGroups(line)

		if !cmp.Equal(expected, result) {
			t.Errorf("%s has %d groups, but got %d", line, expected, result)
		}
	}
}

func TestEnumeratePatterns(t *testing.T) {
	length := 3
	patterns := enumeratePatterns(length)

	expected := 8
	result := len(patterns)

	if result != expected {
		t.Fatalf("Expected %d patterns, but got %d", expected, result)
	}

}

func TestApplyPattern(t *testing.T) {
	line := ".??..??...?##. 1,1,3"
	pattern := "#.#.."
	result := applyPattern(line, pattern)

	expected := ".#...#.....##. 1,1,3"

	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

func TestCountQuestions(t *testing.T) {
	testLines := map[string]int{
		"???.### 1,1,3":             3,
		".??..??...?##. 1,1,3":      5,
		"?#?#?#?#?#?#?#? 1,3,1,6":   8,
		"????.#...#... 4,1,1":       4,
		"????.######..#####. 1,6,5": 4,
		"?###???????? 3,2,1":        9,
	}
	for line, expected := range testLines {
		result := countQuestions(line)

		if result != expected {
			t.Errorf("Expected %d but got %d at line %s", expected, result, line)
		}
	}
}

func TestParseLines(t *testing.T) {
	testLines := map[string][]int{
		"???.### 1,1,3":             {1, 1, 3},
		".??..??...?##. 1,1,3":      {1, 1, 3},
		"?#?#?#?#?#?#?#? 1,3,1,6":   {1, 3, 1, 6},
		"????.#...#... 4,1,1":       {4, 1, 1},
		"????.######..#####. 1,6,5": {1, 6, 5},
		"?###???????? 3,2,1":        {3, 2, 1},
	}
	for line, expected := range testLines {
		result := parseLine(line)

		if !cmp.Equal(expected, result.Groups) {
			t.Errorf("Expected %v but got %v at line %s", expected, result.Groups, line)
		}
	}
}
func TestCountArrangemen(t *testing.T) {
	testLines := map[string]int{
		"???.### 1,1,3":             1,
		".??..??...?##. 1,1,3":      4,
		"?#?#?#?#?#?#?#? 1,3,1,6":   1,
		"????.#...#... 4,1,1":       1,
		"????.######..#####. 1,6,5": 4,
		"?###???????? 3,2,1":        10,
	}
	for line, expected := range testLines {
		result := CountArrangemen(line)

		if expected != result {
			t.Errorf("%s has %d arrangements, but got %d", line, expected, result)
		}
	}
}
