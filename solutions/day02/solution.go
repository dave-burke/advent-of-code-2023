package day02

import (
	"aoc/utils/aocinput"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var minVals = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Part1() string {
	lines := aocinput.ReadInputAsChannel(2)
	result := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		game := parseGame(line)
		isPossible := checkGame(game)
		fmt.Printf("%+v => %t\n", game, isPossible)
		if isPossible {
			result += game.id
		}
	}
	return fmt.Sprintf("%d", result)
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

func checkGame(game Game) bool {
	for color, min := range minVals {
		if game.max[color] > min {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Part2() string {
	lines := aocinput.ReadInputAsLines(2)
	result := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		game := parseGame(line)
		power := power(game)
		fmt.Printf("%+v => %d\n", game, power)
		result += power
	}
	return fmt.Sprintf("%d", result)
}

func power(game Game) int {
	return game.max["red"] * game.max["green"] * game.max["blue"]
}
