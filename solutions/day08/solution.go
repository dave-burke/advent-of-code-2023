package day08

import (
	"aoc/utils/aocinput"
	"fmt"
	"log"
	"strings"
)

func Part1() string {
	lines := aocinput.ReadInputAsString(8)
	headerAndBody := strings.Split(lines, "\n\n")
	header := headerAndBody[0]
	body := headerAndBody[1]

	bodyLines := strings.Split(body, "\n")

	networkMap := make(map[string]NetworkNode)
	for _, line := range bodyLines {
		if len(line) > 0 {
			networkNode := NewNetworkNode(line)
			networkMap[networkNode.source] = networkNode
		}
	}

	log.Print(header)
	for source, node := range networkMap {
		log.Printf("%s = (%s, %s)", source, node.left, node.right)
	}
	// fmt.Printf("%s", currentPosition)

	currentPosition := "AAA"
	nSteps := 0
	for currentPosition != "ZZZ" {
		for _, dir := range header {
			if dir == 'L' {
				currentPosition = networkMap[currentPosition].left
			} else if dir == 'R' {
				currentPosition = networkMap[currentPosition].right
			} else {
				log.Fatalf("Unknown direction: %c", dir)
			}
			nSteps++
			// fmt.Printf(" => %s", currentPosition)
		}
	}
	// fmt.Println()

	return fmt.Sprint(nSteps)
}

type NetworkNode struct {
	source string
	left   string
	right  string
}

func NewNetworkNode(line string) NetworkNode {
	return NetworkNode{
		source: line[0:3],
		left:   line[7:10],
		right:  line[12:15],
	}
}

func Part2() string {
	return "todo"
}
