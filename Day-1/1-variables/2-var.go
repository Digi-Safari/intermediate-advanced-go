package main

import "fmt"

func main() {

	// every types have a default value
	var a int // int default is 0
	var b string = "ajay"
	var c = "rahul"

	// go compiler would infer the type automatically from the right side value
	d := 100 // shorthand
	{
		// avoid this , adds confusion
		//d = "abc"
		d := "hello" // shadowing
		d = "abc"
		fmt.Println(d)
	}
	fmt.Println(a, b, c, d)
}
