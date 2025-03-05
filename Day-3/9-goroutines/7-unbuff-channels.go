package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// we never send and recv in the same goroutine

	ch := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(2)

	//this goroutine would be blocked while sending the values to the another goroutine
	// if recv is not ready
	// if this goroutine blocks, go scheduler would pick other goroutine which would
	// start the recvs
	go func() {
		defer wg.Done()
		for i := 1; i < 3; i++ {
			ch <- i
			fmt.Println("sent", i)
		}
		// blocked until there is a recv

	}()

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		for i := 1; i < 3; i++ {
			fmt.Println(<-ch)
		}

	}()
	wg.Wait()
}
