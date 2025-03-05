package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			go func(i int) {
				ch <- i
			}(i)
		}
	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("received:", v)
		}
	}()

	wg.Wait()
}
