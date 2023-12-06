package aocasync

import (
	"log"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

func MapLinesAsync[T any](lines <-chan string, lineHandler func(line string) T) <-chan T {
	var wg sync.WaitGroup
	results := make(chan T)

	for line := range lines {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			log.Printf("RECEIVE: %s\n", line)
			results <- lineHandler(line)
		}(line)
	}

	go func() {
		wg.Wait()
		close(results) // Close the result channel after all workers are done
	}()

	return results
}

func Merge(workerOutputs ...<-chan int) (out chan int) {
	var wg sync.WaitGroup
	out = make(chan int)

	wg.Add(len(workerOutputs))
	for _, c := range workerOutputs {
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				out <- n
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func ProgressTracker(in <-chan int, total int) <-chan int {
	out := make(chan int)

	hundredth := int(math.Max(math.Floor(float64(total)/100), 1))
	log.Printf("1%% is about %d items", hundredth)

	var count atomic.Uint32
	go func() {
		defer close(out)
		start := time.Now()
		for x := range in {
			out <- x

			count.Add(1)
			currentValue := count.Load()
			if int(currentValue)%hundredth == 0 {
				percent := (float32(currentValue) / float32(total)) * 100
				duration := time.Since(start)
				log.Printf("Completed %d of %d (%f%%) in %v", currentValue, total, percent, duration)
				start = time.Now()
			}
		}
	}()
	return out
}
