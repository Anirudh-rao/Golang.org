package main

import (
	"fmt"
	"sync"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()

	}
	wg.Wait()

}
