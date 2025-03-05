package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//wg := &sync.WaitGroup{}
	// using waitgroup keep track of goroutines
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("hello world")
		wg.Done() // decrement the counter
	}()

	fmt.Println("some kind of work going in the main")

	wg.Wait() // wait until the counter is not 0
	fmt.Println("done")

}
