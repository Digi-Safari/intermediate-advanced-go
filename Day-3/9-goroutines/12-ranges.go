package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wgWorker := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			wgWorker.Add(1)
			go func(i int) {
				defer wgWorker.Done()
				ch <- i
			}(i)
		}
		wgWorker.Wait()
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("received:", v)
		}
	}()

	wg.Wait()
}
