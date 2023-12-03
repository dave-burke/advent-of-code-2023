package aocasync

import (
	"log"
	"sync"
)

func MapLinesAsync[T any](lines chan string, lineHandler func(line string) T) chan T {
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
