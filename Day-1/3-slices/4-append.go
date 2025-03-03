package main

import "fmt"

// https://go.dev/ref/spec#Appending_and_copying_slices
func main() {
	// len: number of elements your slice is storing,
	// or number of elems slice is referring to in backing array

	// cap: number of elems your slice can accommodate
	a := []int{10, 20, 30, 40, 50}
	fmt.Printf("before append %p\n", a)
	fmt.Println(len(a), cap(a))

	a = append(a, 60)
	fmt.Printf("after append %p\n", a)
	fmt.Println(len(a), cap(a))
	a = append(a, 70)
	fmt.Printf("after second append %p\n", a)
	fmt.Println(len(a), cap(a))
	fmt.Println(a)

}
