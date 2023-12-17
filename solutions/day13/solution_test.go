package day13

import (
	"strings"
	"testing"
)

func TestVerticalSample(t *testing.T) {
	lines := []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
	}

	got, ok := findReflection(lines)
	wanted := Reflection{VERTICAL, 5}

	if !ok {
		t.Errorf("Expected %s reflection at %d but got no reflection:\n%s", wanted.direction, wanted.index, strings.Join(lines, "\n"))
	}
	if got != wanted {
		t.Errorf("Expected %s reflection at %d but got %s at %d:\n%s", wanted.direction, wanted.index, got.direction, got.index, strings.Join(lines, "\n"))
	}
}

func TestHorizontalSample(t *testing.T) {
	lines := []string{
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}

	got, ok := findReflection(lines)
	wanted := Reflection{HORIZONTAL, 4}

	if !ok {
		t.Errorf("Expected %s reflection at %d but got no reflection:\n%s", wanted.direction, wanted.index, strings.Join(lines, "\n"))
	}
	if got != wanted {
		t.Errorf("Expected %s reflection at %d but got %s at %d:\n%s", wanted.direction, wanted.index, got.direction, got.index, strings.Join(lines, "\n"))
	}
}

func TestProblemInput1(t *testing.T) {
	lines := []string{
		"######.##",
		"##..###..",
		".......#.",
		"#.##.#..#",
		"###.####.",
		"######..#",
		".#..#.#.#",
		"#....###.",
		".......##",
		".####.##.",
		"#....#...",
		"..##....#",
		"..##....#",
		"######...",
		".#..#..##",
		".#..#..##",
		"######...",
	}

	got, ok := findReflection(lines)
	wanted := Reflection{HORIZONTAL, 15}

	if !ok {
		t.Errorf("Expected %s reflection at %d but got no reflection:\n%s", wanted.direction, wanted.index, strings.Join(lines, "\n"))
	}
	if got != wanted {
		t.Errorf("Expected %s reflection at %d but got %s at %d:\n%s", wanted.direction, wanted.index, got.direction, got.index, strings.Join(lines, "\n"))
	}
}
