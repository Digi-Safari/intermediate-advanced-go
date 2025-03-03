package main

import "fmt"

func main() {
	var s []int // default value of slice is nil
	if s == nil {
		fmt.Println("slice is nil")
		return
	}
	fmt.Println(s)
	fmt.Printf("%#v\n", s)
	//s[0] = 100 // this is meant to update values

	fmt.Println("end of the main")

}
