package aocmath

func Sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func SumChan(nums chan int) int {
	sum := 0
	for n := range nums {
		sum += n
	}
	return sum
}
