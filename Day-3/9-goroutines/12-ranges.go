package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			go func(i int) {
				ch <- i
			}(i)
		}
	}()

	go func() {
		for v := range ch {
			fmt.Println("received:", v)
		}
	}()
}
