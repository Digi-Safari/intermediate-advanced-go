package main

import (
	"fmt"
	"learn-go/slice"
)

func main() {
	x := []int{10, 20, 30}
	x = updateSlice(x, 1, 200)
	slice.Inspect("x", x)
}

func updateSlice(s []int, index, value int) []int {
	// if you want to update the value ,
	//then passing the reference of existing backing array is fine
	s[index] = value
	s = append(s, 10000)
	return s
}

// operatingOnSlice have no intention to update the existing slice x
func operatingOnSlice(s []int) {
	s[0] = 100
	fmt.Println("doing something with s", s)
}
