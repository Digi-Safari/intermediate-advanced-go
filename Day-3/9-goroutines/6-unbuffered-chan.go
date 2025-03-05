package main

import (
	"fmt"
	"sync"
)

// https://go.dev/ref/spec#Send_statements
// A send on an unbuffered channel can proceed if a receiver is ready.
// send will block until there is no recv
// channels are only meant to be used in concurrent programming

func main() {
	wg := &sync.WaitGroup{}

	ch := make(chan int) // unbuffered channel has size of 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1 // send
	}()
	x := <-ch // recv // this would block if no sender is present,
	//and another goroutine from the queue would be picked up
	//which is sender goroutine in this case
	//communication completes, and we get 1 on the screen
	fmt.Println(x)

	// if a goroutine is running, go would not report deadlock with a hope that
	// the goroutine would unblock the other goroutine
	//go func() {
	//	for {
	//		time.Sleep(5 * time.Second)
	//	}
	//}()
	wg.Wait()
}
