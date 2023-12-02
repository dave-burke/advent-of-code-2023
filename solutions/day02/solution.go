package day02

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Part1(input string) string {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		game := parseGame(line)
		fmt.Printf("%+v\n", game)
	}
	return "todo"
}

type Game struct {
	id  int
	max map[string]int
}

func parseGame(line string) Game {
	result := Game{-1, map[string]int{}}

	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	parts := strings.Split(line, ": ")
	title := parts[0]
	game := parts[1]

	id, err := strconv.Atoi(title[5:])
	if err != nil {
		log.Fatal(err)
	}
	result.id = id

	turns := strings.Split(game, "; ")

	for _, turn := range turns {
		colors := strings.Split(turn, ", ")
		for _, color := range colors {
			colorParts := strings.Split(color, " ")
			numString := colorParts[0]
			num, err := strconv.Atoi(numString)
			if err != nil {
				log.Fatal(err)
			}
			color := colorParts[1]
			currentMax := result.max[color]
			result.max[color] = max(currentMax, num)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Part2(input string) string {
	return "todo"
}
