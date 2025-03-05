package main

import (
	"fmt"
	"sync"
	"time"
)

// https://go.dev/ref/spec#Send_statements
// A send on a buffered channel can proceed if there is room in the buffer.
func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int, 5)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			// it would only see if there is room in the buffer, if yes it would send the value
			// buf chan doesn't care about recv
			ch <- i
			fmt.Println("sent", i)
		}
	}()

	go func() {
		defer wg.Done()
		// when we recv value we make one slot empty in the buffer, and more value could be sent over it
		// make sure to recv all the values from the sender, no guarantees given by buffered chan
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("received", <-ch)
		}
	}()

	wg.Wait()
}
