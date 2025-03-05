package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	// Context is an interface, Background method returns an implementation of that interface
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel() // clean up the resources taken up by the context
	doSomething(ctx, wg)
	wg.Wait()
}

func doSomething(ctx context.Context, wg *sync.WaitGroup) {

	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		rec := slowFunc()
		select {
		case ch <- rec:
			fmt.Println("result sent to channel")
		case <-ctx.Done(): // this case would evaluate if ctx is cancelled or timeout happened
			fmt.Println("context done sending part", ctx.Err())
			fmt.Println("reverse what slow func did")
			return
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("context done", ctx.Err())
		return
	case x := <-ch:
		fmt.Println("result received from channel", x)
	}

	fmt.Println("do something done")

}

func slowFunc() int {
	time.Sleep(time.Second * 3)
	fmt.Println("slowFunction: slow fn ran and add 100 records to db")
	fmt.Println()
	return 100
}
