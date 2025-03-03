package main

import "fmt"

func main() {
	//var a [2]int // fixed size
	//i := [2]int{10, 20, 30}
	// slices can grow

	var i []int = []int{10, 20, 30, 40}
	//i[5] = 100 // updating or accessing the value
	fmt.Println(i)

	var b []int // default is nil
	fmt.Println(b)
	if b == nil {
		fmt.Println("it is nil slice....")
	}
	fmt.Printf("%#v", b)
}
