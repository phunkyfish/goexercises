package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		//fmt.Printf("%d: %d\n", id, i)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 200000; i++ {
		wg.Add(1)
		go worker(i, wg)
	}
	wg.Wait()
	fmt.Printf("done\n")
}
