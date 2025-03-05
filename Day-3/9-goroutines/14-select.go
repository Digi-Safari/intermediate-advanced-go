package main

import (
	"fmt"
	"sync"
	"time"
)

// use select with unbuffered channels, when working with a fanout pattern

func main() {
	// select is used when we want to listen or send values to over a multiple channel
	wg := new(sync.WaitGroup)
	get := make(chan string)
	post := make(chan string)
	put := make(chan string)
	wg.Add(3)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		get <- "get"
	}()

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		post <- "post"
	}()
	go func() {
		defer wg.Done()
		put <- "put"
		put <- "put 1"
	}()
	//
	//fmt.Println(<-get)
	//fmt.Println(<-post)
	//fmt.Println(<-put)

	for i := 0; i < 3; i++ {

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

		}
	}
	wg.Wait()

}
