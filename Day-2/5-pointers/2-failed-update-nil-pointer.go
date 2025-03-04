package main

import "fmt"

var x int = 10

func main() {
	// address of p is x80 = and the value it is storing is nil
	var p *int // nil // default value of a pointer is nil
	fmt.Println(&p)

	// after calling the update value p is still nil, as we never updated the pointer
	updateValue(p)
	fmt.Println(*p) // panic // nil pointer dereference
}

func updateValue(p1 *int) {
	if p1 == nil {
		fmt.Println("p1 is nil")
	}
	fmt.Println(&p1)
	p1 = &x
	fmt.Println(*p1)

}
