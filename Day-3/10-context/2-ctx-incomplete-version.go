package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	wg := new(sync.WaitGroup)
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	wg.Add(1)
	go func() {
		// this goroutine would deadlock if the receiver goroutine has already moved on because of time out
		// use select for sending here as well to make sure we send the values in case of receiver is present
		defer wg.Done()
		x := slowFuncV2()
		// in this case slowV2 is still running and it would not stop
		ch <- x
		fmt.Println("sent the value over the channel")

	}()
	func() {
		select {
		case <-ctx.Done():
			// listen over the done channel and if the time is up this case evaluates
			fmt.Println("context done", ctx.Err())
			return
		case x := <-ch:
			// if received value in time, this case evaluates
			fmt.Println("result received from channel", x)
		}
	}()

	fmt.Println("starting doing other things, no longer waiting for other goroutine to finish")
	wg.Wait()
}

func slowFuncV2() int {
	time.Sleep(time.Second * 3)
	fmt.Println("slowFunction: slow fn ran and add 100 records to db")
	fmt.Println()
	return 100
}
