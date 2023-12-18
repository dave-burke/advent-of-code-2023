package day15

import (
	"aoc/utils/aocinput"
	"fmt"
	"log"
	"strings"
)

func Part1() string {
	// input := aocinput.ReadSampleAsString(15)
	input := aocinput.ReadInputAsString(15)

	input = strings.TrimSpace(input) // remove newline at the end

	steps := strings.Split(input, ",")

	sum := 0
	for _, step := range steps {
		result := hash(step)
		log.Printf("%s => %d", step, result)
		sum += result
	}

	return fmt.Sprint(sum)
}

func hash(s string) int {
	currentValue := 0

	for _, char := range s {
		code := int(char)
		currentValue += code
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

func Part2() string {
	lines := aocinput.ReadSampleAsLines(15)
	// lines := aocinput.ReadInputAsLines(15)

	log.Printf("Got %d lines", len(lines))

	return fmt.Sprint(len(lines))
}
