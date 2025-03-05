package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	ch := make(chan int, 10)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // sends a signal to stop the range
		// close signal range that no more values be sent and it can stop after receiving remaining values
		// once the channel is closed, we can't send more values to it

	}()

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		for v := range ch { // it would run infinitely, channel needs to be closed to stop this range
			fmt.Println(v)
		}
	}()

	wg.Wait()
}
