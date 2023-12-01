package day01

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Part1(input string) string {
	lines := strings.Split(input, "\n")
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

func Part2(input string) string {
	return "todo"
}
