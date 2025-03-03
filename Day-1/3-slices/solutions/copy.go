package main

import "learn-go/slice"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}

	// make allocates backing array for slice
	b := make([]int, len(x), 100) // make(type,len,cap)
	copy(b, x)                    // deep copy
	b[0] = 999                    // it will not affect the x slice anymore
	slice.Inspect("x", x)
	slice.Inspect("b", b)
}
