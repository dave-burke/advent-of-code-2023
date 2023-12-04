package day04

import (
	"aoc/utils/aocasync"
	"aoc/utils/aocinput"
	"aoc/utils/aocmath"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Part1() string {
	input := aocinput.ReadInputAsChannel(4)

	gamePoints := aocasync.MapLinesAsync[int](input, func(line string) int {
		card := parseCard(line)
		log.Printf("%v => %d", card, card.score)
		return card.score
	})
	return fmt.Sprint(aocmath.SumChan(gamePoints))
}

type Card struct {
	id             int
	cardNumbers    map[int]int
	winningNumbers map[int]int
	nMatches       int
	score          int
}

func parseCard(card string) Card {
	parts := strings.Split(card, ": ")
	title := parts[0]

	id := parseId(title)

	numbers := strings.Split(parts[1], " | ")
	cardNumbers := parseNumbers(numbers[0])
	winningNumbers := parseNumbers(numbers[1])

	nMatches := 0
	for num, count := range cardNumbers {
		if _, ok := winningNumbers[num]; ok {
			nMatches += count
		}
	}
	score := 0
	if nMatches != 0 {
		score = int(math.Pow(float64(2), float64(nMatches-1)))
	}
	return Card{id, cardNumbers, winningNumbers, nMatches, score}
}

func parseId(title string) int {
	id, err := strconv.Atoi(strings.TrimSpace(title[4:]))
	if err != nil {
		log.Fatal(err)
	}
	return id
}

var numbersRegex *regexp.Regexp = regexp.MustCompile(`\d+`)

func parseNumbers(numString string) map[int]int {
	numStrings := numbersRegex.FindAllString(numString, -1)

	result := make(map[int]int, len(numStrings))
	for _, str := range numStrings {
		if len(str) == 0 {
			continue
		}
		if num, err := strconv.Atoi(str); err != nil {
			panic(err)
		} else {
			if _, ok := result[num]; !ok {
				result[num] = 0
			}
			result[num]++
		}
	}
	return result
}

func Part2() string {
	input := aocinput.ReadSampleAsList(4)

	cursor := input.Front()
	for cursor.Next() != nil {
		line := cursor.Value.(string)
		card := parseCard(line)
		score := card.Score()
		log.Printf("%v => %d", card, score)

		lookAhead := cursor
		for i := 0; i < score; i++ {
			lookAhead = lookAhead.Next()
		}
		for score > 0 {
			input.InsertAfter(lookAhead.Value, lookAhead)
			lookAhead = lookAhead.Prev()
			score--
		}
		cursor = cursor.Next()
	}
	return fmt.Sprint(input.Len())
}
