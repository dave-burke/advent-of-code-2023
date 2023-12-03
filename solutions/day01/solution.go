package day01

import (
	"aoc/utils/aocinput"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Part1() string {
	lines := aocinput.ReadInputAsLines(1)
	nums := []int{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		numString := ""
		for i := 0; i < len(line); i++ {
			char := line[i]
			if char >= 48 && char <= 57 {
				numString += string([]byte{char})
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			char := line[i]
			if char >= 48 && char <= 57 {
				numString += string([]byte{char})
				break
			}
		}
		num, err := strconv.Atoi(numString)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
		fmt.Printf("%s => %s\n", line, numString)
	}
	var sum int = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return fmt.Sprint(sum)
}

func Part2() string {
	numMap := map[string]int{
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	nums := []int{}
	lines := aocinput.ReadInputAsLines(1)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if line == "eightone7threenl7mtxbmkpkzqzljrdk" {
			log.Println("bad")
		}
		minNum := 10
		minIndex := len(line)
		for numString, num := range numMap {
			numIndex := strings.Index(line, numString)
			if numIndex != -1 && numIndex < minIndex {
				minNum = num
				minIndex = numIndex
			}
		}

		maxNum := -1
		maxIndex := -1
		for numString, num := range numMap {
			numIndex := strings.LastIndex(line, numString)
			if numIndex != -1 && numIndex > maxIndex {
				maxNum = num
				maxIndex = numIndex
			}
		}

		numString := fmt.Sprintf("%d%d", minNum, maxNum)
		if len(numString) != 2 {
			log.Fatalf("Failed to produce a two digit number: %s => %s", line, numString)
		}

		if minNum == 10 {
			log.Fatalf("Failed to find min number: %s => %s", line, numString)
		}
		if maxNum == -1 {
			log.Fatalf("Failed to find max number: %s => %s", line, numString)
		}

		fmt.Printf("%s => %s\n", line, numString)

		num, err := strconv.Atoi(numString)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, num)
	}
	var sum int = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return fmt.Sprint(sum)
}
