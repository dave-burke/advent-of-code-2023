package day09

import (
	"aoc/utils/aocinput"
	"aoc/utils/aocmath"
	"aoc/utils/aocparse"
	"fmt"
	"strings"
)

func Part1() string {
	// lines := aocinput.ReadSampleAsLines(9)
	lines := aocinput.ReadInputAsLines(9)

	nextValues := make([]int, 0, len(lines))
	for _, line := range lines {
		nums := ParseLine(line)
		nextValue := GetNextIncrement(nums)
		nextValues = append(nextValues, nextValue)
	}
	sum := aocmath.Sum(nextValues)

	// 1939607040 is not the right answer
	return fmt.Sprint(sum)
}

func GetNextIncrement(sequence []int) int {
	diffSequence := DiffSequence(sequence)
	incremented := Increment(diffSequence)
	PrintDiffSequence(incremented)
	return incremented[0][len(incremented[0])-1]
}

func Increment(diffSequence [][]int) [][]int {
	inc := 0
	for i := len(diffSequence) - 1; i > 0; i-- {
		thisSequence := diffSequence[i]
		upSequence := diffSequence[i-1]
		inc = upSequence[len(upSequence)-1] + thisSequence[len(thisSequence)-1]
		diffSequence[i-1] = append(upSequence, inc)
	}
	return diffSequence
}

func PrintDiffSequence(diffSequence [][]int) {
	for i, sequence := range diffSequence {
		fmt.Printf("%s%v\n", strings.Repeat(" ", i), sequence)
	}
}

func ParseLine(line string) []int {
	numStrings := strings.Split(line, " ")
	result := make([]int, 0, len(numStrings))
	for _, numString := range numStrings {
		result = append(result, aocparse.MustAtoi(numString))
	}
	return result
}

func DiffSequence(sequence []int) [][]int {
	result := make([][]int, 0)

	result = append(result, sequence)
	for !allZeroes(sequence) {
		sequence = NextDiff(sequence)
		result = append(result, sequence)
	}
	return result
}

func allZeroes(sequence []int) bool {
	for _, val := range sequence {
		if val != 0 {
			return false
		}
	}
	return true
}

func NextDiff(sequence []int) []int {
	result := make([]int, 0, len(sequence)-1)
	for i := 0; i < len(sequence)-1; i++ {
		result = append(result, sequence[i+1]-sequence[i])
	}
	return result
}

func Part2() string {
	lines := aocinput.ReadInputAsString(9)
	return fmt.Sprint(len(lines))
}
