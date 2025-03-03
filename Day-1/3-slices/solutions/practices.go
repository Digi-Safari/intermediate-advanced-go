package main

import (
	"fmt"
	"learn-go/slice"
)

func main() {
	x := []int{10, 20, 30}
	x = updateSlice(x, 1, 200)
	operatingOnSlice(x)
	slice.Inspect("x", x)
}

func updateSlice(s []int, index, value int) []int {
	// if you want to update the value ,
	//then passing the reference of existing backing array is fine
	s[index] = value
	s = append(s, 10000)
	return s // make sure to return the updated slice if using append
}

// ideally a function should never update a slice if its job is to work on the slice,
// but in this case just for assignment we are doing it
// operatingOnSlice have no intention to update the existing slice x
func operatingOnSlice(s []int) {
	copiedSlice := make([]int, len(s))
	copy(copiedSlice, s) // Use the `copy` function to duplicate the slice.
	copiedSlice[0] = 100
	fmt.Println("doing something with s", s)
}
