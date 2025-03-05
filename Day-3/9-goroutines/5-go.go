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
		s := helloWorld("Hello")
		fmt.Println(s)
	}()
	helloWorld("Hello")
	wg.Wait()

}

func helloWorld(s string) string {
	fmt.Println(s)
	return s
}
