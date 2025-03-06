package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	go func() {
		x := slowFuncV2()
		ch <- x

	}()
	func() {
		select {
		case <-ctx.Done():
			// listen over the done channel and if the time is up this case evaluates
			fmt.Println("context done", ctx.Err())
			return
		case x := <-ch:
			fmt.Println("result received from channel", x)
		}
	}()

	fmt.Println("starting doing other things, no longer waiting for other goroutine to finish")
}

func slowFuncV2() int {
	time.Sleep(time.Second * 3)
	fmt.Println("slowFunction: slow fn ran and add 100 records to db")
	fmt.Println()
	return 100
}
