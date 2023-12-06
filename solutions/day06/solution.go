package day06

import (
	"aoc/utils/aocinput"
	"aoc/utils/aocmath"
	"aoc/utils/aocparse"
	"fmt"
	"log"
	"strings"
)

func Part1() string {
	lines := aocinput.ReadInputAsLines(6)

	times := parseLine(lines[0])

	dist := parseLine(lines[1])

	races := genRaces(times, dist)
	nWins := countWins(races)
	result := aocmath.MultiplyChan(nWins)

	return fmt.Sprint(result)
}

type race struct {
	time int
	dist int
}

func genRaces(times []int, maxDistances []int) <-chan race {
	if len(times) != len(maxDistances) {
		log.Fatalf("Expected equal lengths, but got %d times and %d distances", len(times), len(maxDistances))
	}

	out := make(chan race)

	go func() {
		defer close(out)
		for i := 0; i < len(times); i++ {
			out <- race{times[i], maxDistances[i]}
		}
	}()
	return out
}

func countWins(in <-chan race) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for race := range in {
			wins := findWins(race)
			out <- len(wins)
		}
	}()

	return out
}

func findWins(r race) []int {
	// naive impl
	wins := make([]int, 0)
	for i := 1; i < r.time-1; i++ {
		speed := i
		moveTimeMs := r.time - i
		dist := speed * moveTimeMs
		if dist > r.dist {
			wins = append(wins, dist)
		}
	}
	return wins
}

func parseLine(line string) []int {
	fullValue := strings.Split(line, ":")[1]
	valueStrings := strings.Split(fullValue, " ")

	result := make([]int, 0)
	for _, valueString := range valueStrings {
		trimmed := strings.TrimSpace(valueString)
		if len(trimmed) == 0 {
			continue
		}
		valueInt := aocparse.MustAtoi(trimmed)
		result = append(result, valueInt)
	}
	return result
}

func Part2() string {
	return "todo"
}
