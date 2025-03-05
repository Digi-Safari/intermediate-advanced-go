package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go work(i, wg)
	}

	wg.Wait()
	fmt.Println("done")
}

func work(workId int, wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		fmt.Println("work", workId, "anon function")
	}()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("work", workId, "main function")
}
