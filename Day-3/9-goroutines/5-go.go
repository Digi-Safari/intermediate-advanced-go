package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		helloWorld("Hello")
	}()
	helloWorld("Hello")
	wg.Wait()

}

func helloWorld(s string) {
	fmt.Println(s)
}
