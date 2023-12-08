package day08

import (
	"aoc/utils/aocinput"
	"fmt"
	"log"
	"strings"
)

func Part1() string {
	lines := aocinput.ReadSampleAsString(8)
	headerAndBody := strings.Split(lines, "\n\n")
	header := headerAndBody[0]
	body := headerAndBody[1]

	bodyLines := strings.Split(body, "\n")

	log.Print(header)
	log.Printf("%v", bodyLines)
	return fmt.Sprint(len(bodyLines))
}

func Part2() string {
	return "todo"
}
