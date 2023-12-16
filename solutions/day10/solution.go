package day10

import (
	"aoc/utils/aocgrid"
	"aoc/utils/aocinput"
	"fmt"
	"log"
	"slices"
)

const (
	north = 0
	south = 1
	east  = 2
	west  = 3
)

func directionToString(dir int) string {
	switch dir {
	case 0:
		return "north"
	case 1:
		return "south"
	case 2:
		return "east"
	case 3:
		return "west"
	default:
		log.Fatalf("Invalid direction: %d", dir)
	}
	return ""
}

func Part1() string {
	// grid := aocgrid.Grid(aocinput.ReadSampleAsGrid(10))
	grid := aocgrid.Grid(aocinput.ReadInputAsGrid(10))

	start := grid.CursorAt(0, 0)
	for start.GetValue() != 'S' {
		var err error
		if start, err = start.Next(); err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("%c is at (%v)", start.GetValue(), start.Position)

	loop := startLoop(start)

	return fmt.Sprint((len(loop) + 1) / 2)
}

func isPipe(cursor aocgrid.Cursor) bool {
	return slices.Contains([]rune{'J', 'F', 'L', '7', '-', '|'}, cursor.GetValue())

}

func startLoop(start aocgrid.Cursor) []aocgrid.Cursor {
	if left, err := start.WalkLeft(); err == nil && isPipe(left) {
		return walkLoop(start, east, left)
	}
	if up, err := start.WalkUp(); err == nil && isPipe(up) {
		return walkLoop(start, south, up)
	}
	if right, err := start.WalkRight(); err == nil && isPipe(right) {
		return walkLoop(start, west, right)
	}
	if down, err := start.WalkDown(); err == nil && isPipe(down) {
		return walkLoop(start, north, down)
	}
	log.Fatalf("No pipe is adjacent to %v", start.Position)
	return []aocgrid.Cursor{}
}

func walkLoop(start aocgrid.Cursor, from int, first aocgrid.Cursor) []aocgrid.Cursor {
	current := first
	loop := make([]aocgrid.Cursor, 0)
	for !aocgrid.CursorsEqual(start, current) {
		log.Printf("Moving from %s to %v", directionToString(from), current.Position)
		loop = append(loop, current)
		from, current = followPipse(from, current)
	}
	return loop
}

func followPipse(from int, cursor aocgrid.Cursor) (int, aocgrid.Cursor) {
	pipe := cursor.GetValue()
	if pipe == 'J' {
		if from == west {
			if next, err := cursor.WalkUp(); err != nil {
				log.Fatalf("Could not walk north through J pipe from (%v): %v", cursor.Position, err)
			} else {
				return south, next
			}
		} else if from == north {
			if next, err := cursor.WalkLeft(); err != nil {
				log.Fatalf("Could not walk west through J pipefrom (%v): %v", cursor.Position, err)
			} else {
				return east, next
			}
		} else {
			log.Fatalf("Cannot go through pipe J from %s", directionToString(from))
		}
	} else if pipe == 'L' {
		if from == north {
			if next, err := cursor.WalkRight(); err != nil {
				log.Fatalf("Could not walk east through L pipe from (%v): %v", cursor.Position, err)
			} else {
				return west, next
			}
		} else if from == east {
			if next, err := cursor.WalkUp(); err != nil {
				log.Fatalf("Could not walk north through L pipe from (%v): %v", cursor.Position, err)
			} else {
				return south, next
			}
		} else {
			log.Fatalf("Cannot go through pipe L from %s", directionToString(from))
		}
	} else if pipe == '7' {
		if from == west {
			if next, err := cursor.WalkDown(); err != nil {
				log.Fatalf("Could not walk south through 7 pipe from (%v): %v", cursor.Position, err)
			} else {
				return north, next
			}
		} else if from == south {
			if next, err := cursor.WalkLeft(); err != nil {
				log.Fatalf("Could not walk west through 7 pipe from (%v): %v", cursor.Position, err)
			} else {
				return east, next
			}
		} else {
			log.Fatalf("Cannot go through pipe 7 from %s", directionToString(from))
		}
	} else if pipe == 'F' {
		if from == south {
			if next, err := cursor.WalkRight(); err != nil {
				log.Fatalf("Could not walk east through F pipe from (%v): %v", cursor.Position, err)
			} else {
				return west, next
			}
		} else if from == east {
			if next, err := cursor.WalkDown(); err != nil {
				log.Fatalf("Could not walk south through F pipe from (%v): %v", cursor.Position, err)
			} else {
				return north, next
			}
		} else {
			log.Fatalf("Cannot go through pipe F from %s", directionToString(from))
		}
	} else if pipe == '|' {
		if from == south {
			if next, err := cursor.WalkUp(); err != nil {
				log.Fatalf("Could not walk north through | pipe from (%v): %v", cursor.Position, err)
			} else {
				return south, next
			}
		} else if from == north {
			if next, err := cursor.WalkDown(); err != nil {
				log.Fatalf("Could not walk south through | pipe from (%v): %v", cursor.Position, err)
			} else {
				return north, next
			}
		} else {
			log.Fatalf("Cannot go through pipe | from %s", directionToString(from))
		}
	} else if pipe == '-' {
		if from == east {
			if next, err := cursor.WalkLeft(); err != nil {
				log.Fatalf("Could not walk west through - pipe from (%v): %v", cursor.Position, err)
			} else {
				return east, next
			}
		} else if from == west {
			if next, err := cursor.WalkRight(); err != nil {
				log.Fatalf("Could not walk east through - pipe from (%v): %v", cursor.Position, err)
			} else {
				return west, next
			}
		} else {
			log.Fatalf("Cannot go through pipe - from %s", directionToString(from))
		}
	} else {
		log.Fatalf("Unknown pipe %c", pipe)
	}
	log.Fatalf("Failed to move ahead at %v coming from %s", cursor, directionToString(from))
	return -1, cursor
}

func Part2() string {
	lines := aocinput.ReadSampleAsLines(10)
	// lines := aocinput.ReadInputAsLines(10)

	return fmt.Sprint(len(lines))
}
