package main

import (
	"fmt"
	"sync"
)

// use select with unbuffered channels, when working with a fanout pattern
// don't use this with buffered channels
// if you want to use a buffered channel, then use range to receive the remaining value

func main() {
	// select is used when we want to listen or send values to over a multiple channel
	wg := new(sync.WaitGroup)
	wgWorker := new(sync.WaitGroup)
	get := make(chan string)
	post := make(chan string)
	put := make(chan string)
	done := make(chan struct{})

	wgWorker.Add(3)
	go func() {
		defer wgWorker.Done()
		get <- "get"
		fmt.Println("sent get")
	}()

	go func() {
		defer wgWorker.Done()
		post <- "post"

	}()
	go func() {
		defer wgWorker.Done()
		put <- "put"
		put <- "put 1"
	}()
	//
	//fmt.Println(<-get)
	//fmt.Println(<-post)
	//fmt.Println(<-put)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			//whichever case is not blocking exec that first
			//whichever case is ready first, exec that.
			// possible cases are chan recv , send , default
			select {
			case msg := <-get:
				fmt.Println(msg)
			case msg := <-post:
				fmt.Println(msg)
			case msg := <-put:
				fmt.Println(msg)
			case <-done:
				fmt.Println("all the values are processed")
				return

			}
		}
	}()

	// don't put this goroutine before adding counter to wgWorkers
	// there is a chance that worker wait group value is not correct and we send close signal early
	wg.Add(1)
	go func() {
		defer wg.Done()
		wgWorker.Wait()
		close(done)
	}()

	wg.Wait()

}
