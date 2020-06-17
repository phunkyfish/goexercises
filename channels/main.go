package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, work chan int) {
	for job := range work {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d: %d\n", id, job)
	}
	wg.Done()
}

func main() {
	work := make(chan int, 10)

	wg := &sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go worker(i, wg, work)
	}

	for i := 0; i < 100; i++ {
		work <- i
		fmt.Printf("Queued %d\n", i)
	}
	close(work)

	wg.Wait()
	fmt.Printf("done\n")
}
