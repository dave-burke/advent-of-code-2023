package aocmath

import "math"

func Sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func SumChan(nums <-chan int) int {
	sum := 0
	for n := range nums {
		sum += n
	}
	return sum
}

func MinInt(nums []int) int {
	min := math.MaxInt
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func MinIntChan(nums <-chan int) int {
	min := math.MaxInt
	for num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func MultiplyChan(nums <-chan int) int {
	result := 1
	for num := range nums {
		result *= num
	}
	return result
}
